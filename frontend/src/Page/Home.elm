module Page.Home exposing (Model, Msg, init, subscriptions, toSession, update, view)

{-| The homepage. You can get here via either the / or /#/ routes.
-}

import Api
import Api.Endpoint as Endpoint
import Browser.Dom as Dom
import Html exposing (..)
import Html.Attributes exposing (attribute, class, classList, href, id, placeholder)
import Html.Events exposing (onClick)
import Http
import Loading
import Log
import Page
import Session exposing (Session)
import Task exposing (Task)
import Time
import Url.Builder
import Username exposing (Username)



-- MODEL


type alias Model =
    { session : Session
    , timeZone : Time.Zone
    }


type Status a
    = Loading
    | LoadingSlowly
    | Loaded a
    | Failed


init : Session -> ( Model, Cmd Msg )
init session =
    ( { session = session
      , timeZone = Time.utc
      }
    , Cmd.none
    )



-- VIEW


view : Model -> { title : String, content : Html Msg }
view model =
    { title = "Conduit"
    , content =
        div []
            [ text "home"
            ]
    }



-- UPDATE


type Msg
    = GotTimeZone Time.Zone
    | GotSession Session


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    ( model, Cmd.none )



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
    Session.changes GotSession (Session.navKey model.session)



-- EXPORT


toSession : Model -> Session
toSession model =
    model.session
