module Page.Login exposing (Model, Msg, init, subscriptions, toSession, update, view)

{-| The login page.
-}

import Api
import Browser.Navigation as Nav
import Bulma.Classes as Bulma
import Bulma.Helpers as BulmaHelpers
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http
import Json.Decode as Decode exposing (Decoder, decodeString, field, string)
import Json.Decode.Pipeline exposing (optional)
import Json.Encode as Encode
import Route exposing (Route)
import Session exposing (Session)
import Viewer exposing (Viewer)



-- MODEL


type alias Model =
    { session : Session
    , problems : List Problem
    , form : Form
    }


{-| Recording validation problems on a per-field basis facilitates displaying
them inline next to the field where the error occurred.
I implemented it this way out of habit, then realized the spec called for
displaying all the errors at the top. I thought about simplifying it, but then
figured it'd be useful to show how I would normally model this data - assuming
the intended UX was to render errors per field.
(The other part of this is having a view function like this:
viewFieldErrors : ValidatedField -> List Problem -> Html msg
...and it filters the list of problems to render only InvalidEntry ones for the
given ValidatedField. That way you can call this:
viewFieldErrors Email problems
...next to the `email` field, and call `viewFieldErrors Password problems`
next to the `password` field, and so on.
The `LoginError` should be displayed elsewhere, since it doesn't correspond to
a particular field.
-}
type Problem
    = InvalidEntry ValidatedField String
    | ServerError String


type alias Form =
    { email : String
    , password : String
    }


init : Session -> ( Model, Cmd msg )
init session =
    ( { session = session
      , problems = []
      , form =
            { email = ""
            , password = ""
            }
      }
    , Cmd.none
    )



-- VIEW


view : Model -> { title : String, content : Html Msg }
view model =
    { title = "Login"
    , content =
        section
            [ BulmaHelpers.classList [ Bulma.section, Bulma.hasBackgroundWhiteBis ] ]
            [ div [ BulmaHelpers.classList [ Bulma.columns, Bulma.isCentered ] ]
                [ div [ BulmaHelpers.classList [ Bulma.columns, Bulma.is8Tablet, Bulma.is6Desktop, Bulma.is4Widescreen ] ]
                    [ div [ class Bulma.box ]
                        [ div [ class Bulma.cardContent ]
                            [ viewForm model.form ]
                        ]
                    ]
                ]
            ]
    }


viewForm : Form -> Html Msg
viewForm form =
    Html.form [ onSubmit SubmittedForm ]
        [ h2 [ BulmaHelpers.classList [ Bulma.title, Bulma.hasTextCentered, Bulma.isSize3 ] ]
            [ text "Please log in" ]
        , List.map viewProblem model.problems
        , div [ class Bulma.field ]
            [ label [ class Bulma.label, for "email" ]
                [ text "Email address" ]
            , div [ class Bulma.control ]
                [ input [ class Bulma.input, id "email", name "_username", placeholder "E-mail", attribute "required" "required", type_ "email", value "" ]
                    []
                ]
            ]
        , div [ class Bulma.field ]
            [ label [ class Bulma.label, for "password" ]
                [ text "Password" ]
            , input [ class Bulma.input, id "password", name "_password", placeholder "Password", attribute "required" "required", type_ "password" ]
                []
            ]
        , div [ class Bulma.field ]
            [ button [ BulmaHelpers.classList [ Bulma.button, Bulma.isMedium, Bulma.isPrimary, Bulma, isFullwidth ], type_ "submit" ]
                [ text "Log in" ]
            ]
        ]


viewProblem : Problem -> Html msg
viewProblem problem =
    let
        errorMessage =
            case problem of
                InvalidEntry _ str ->
                    str

                ServerError str ->
                    str
    in
    div [ BulmaHelpers.classList [ Bulma.notification, Bulma.isDanger ] ]
        [ text errorMessage ]



-- UPDATE


type Msg
    = SubmittedForm
    | EnteredEmail String
    | EnteredPassword String
    | CompletedLogin (Result Http.Error Viewer)
    | GotSession Session


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        SubmittedForm ->
            case validate model.form of
                Ok validForm ->
                    ( { model | problems = [] }
                    , Http.send CompletedLogin (login validForm)
                    )

                Err problems ->
                    ( { model | problems = problems }
                    , Cmd.none
                    )

        EnteredEmail email ->
            updateForm (\form -> { form | email = email }) model

        EnteredPassword password ->
            updateForm (\form -> { form | password = password }) model

        CompletedLogin (Api.ErrorDetailed error) ->
            let
                serverErrors =
                    Api.decodeErrors error
                        |> List.map ServerError
            in
            ( { model | problems = List.append model.problems serverErrors }
            , Cmd.none
            )

        CompletedLogin (Ok viewer) ->
            ( model
            , Viewer.store viewer
            )

        GotSession session ->
            ( { model | session = session }
            , Route.replaceUrl (Session.navKey session) Route.Home
            )


{-| Helper function for `update`. Updates the form and returns Cmd.none.
Useful for recording form fields!
-}
updateForm : (Form -> Form) -> Model -> ( Model, Cmd Msg )
updateForm transform model =
    ( { model | form = transform model.form }, Cmd.none )



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Session.changes GotSession (Session.navKey model.session)



-- FORM


{-| Marks that we've trimmed the form's fields, so we don't accidentally send
it to the server without having trimmed it!
-}
type TrimmedForm
    = Trimmed Form


{-| When adding a variant here, add it to `fieldsToValidate` too!
-}
type ValidatedField
    = Email
    | Password


fieldsToValidate : List ValidatedField
fieldsToValidate =
    [ Email
    , Password
    ]


{-| Trim the form and validate its fields. If there are problems, report them!
-}
validate : Form -> Result (List Problem) TrimmedForm
validate form =
    let
        trimmedForm =
            trimFields form
    in
    case List.concatMap (validateField trimmedForm) fieldsToValidate of
        [] ->
            Ok trimmedForm

        problems ->
            Err problems


validateField : TrimmedForm -> ValidatedField -> List Problem
validateField (Trimmed form) field =
    List.map (InvalidEntry field) <|
        case field of
            Email ->
                if String.isEmpty form.email then
                    [ "email can't be blank." ]

                else
                    []

            Password ->
                if String.isEmpty form.password then
                    [ "password can't be blank." ]

                else
                    []


{-| Don't trim while the user is typing! That would be super annoying.
Instead, trim only on submit.
-}
trimFields : Form -> TrimmedForm
trimFields form =
    Trimmed
        { email = String.trim form.email
        , password = String.trim form.password
        }



-- HTTP


login : TrimmedForm -> Http.Request Viewer
login (Trimmed form) =
    let
        user =
            Encode.object
                [ ( "email", Encode.string form.email )
                , ( "password", Encode.string form.password )
                ]

        body =
            Encode.object [ ( "user", user ) ]
                |> Http.jsonBody
    in
    Api.login body Viewer.decoder



-- EXPORT


toSession : Model -> Session
toSession model =
    model.session
