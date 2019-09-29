port module Api exposing (Cred, ErrorDetailed, addServerError, application, decodeErrors, delete, get, login, logout, post, put, register, storeCredWith, username, viewerChanges)

{-| This module is responsible for communicating to the Conduit API.

It exposes an opaque Endpoint type which is guaranteed to point to the correct URL.

-}

import Api.Endpoint as Endpoint exposing (Endpoint)
import Browser
import Browser.Navigation as Nav
import Http exposing (Body, Expect)
import Json.Decode as Decode exposing (Decoder, Value, decodeString, field, string)
import Json.Decode.Pipeline as Pipeline exposing (optional, required)
import Json.Encode as Encode
import Url exposing (Url)
import Username exposing (Username)



-- CRED


{-| The authentication credentials for the Viewer (that is, the currently logged-in user.)

This includes:

  - The cred's Username
  - The cred's authentication token

By design, there is no way to access the token directly as a String.
It can be encoded for persistence, and it can be added to a header
to a HttpBuilder for a request, but that's it.

This token should never be rendered to the end user, and with this API, it
can't be!

-}
type Cred
    = Cred Username String


username : Cred -> Username
username (Cred val _) =
    val


credHeader : Cred -> Http.Header
credHeader (Cred _ str) =
    Http.header "authorization" ("Token " ++ str)


{-| It's important that this is never exposed!

We epxose `login` and `application` instead, so we can be certain that if anyone
ever has access to a `Cred` value, it came from either the login API endpoint
or was passed in via flags.

-}
credDecoder : Decoder Cred
credDecoder =
    Decode.succeed Cred
        |> required "username" Username.decoder
        |> required "token" Decode.string



-- HTTP HELPERS


type ErrorDetailed body
    = BadUrl String
    | Timeout
    | NetworkError
    | BadStatus Http.Metadata body
    | BadBody Http.Metadata body String



-- convertResponseString : Http.Response String -> Result (ErrorDetailed String) ( Http.Metadata, String )
-- convertResponseString httpResponse =
--     case httpResponse of
--         Http.BadUrl_ url ->
--             Err (BadUrl url)
--         Http.Timeout_ ->
--             Err Timeout
--         Http.NetworkError_ ->
--             Err NetworkError
--         Http.BadStatus_ metadata body ->
--             Err (BadStatus metadata body)
--         Http.GoodStatus_ metadata body ->
--             Ok ( metadata, body )


convertResponseStringToJson : Decoder a -> Http.Response String -> Result (ErrorDetailed String) ( Http.Metadata, a )
convertResponseStringToJson decoder httpResponse =
    case httpResponse of
        Http.BadUrl_ url ->
            Err (BadUrl url)

        Http.Timeout_ ->
            Err Timeout

        Http.NetworkError_ ->
            Err NetworkError

        Http.BadStatus_ metadata body ->
            Err (BadStatus metadata body)

        Http.GoodStatus_ metadata body ->
            Result.mapError (BadBody metadata body) <|
                Result.mapError Decode.errorToString
                    (Decode.decodeString (Decode.map (\res -> ( metadata, res )) decoder) body)



-- convertResponseBytes : Bytes.Decode.Decoder a -> Http.Response Bytes -> Result (ErrorDetailed Bytes) ( Http.Metadata, a )
-- convertResponseBytes decoder httpResponse =
--     case httpResponse of
--         Http.BadUrl_ url ->
--             Err (BadUrl url)
--         Http.Timeout_ ->
--             Err Timeout
--         Http.NetworkError_ ->
--             Err NetworkError
--         Http.BadStatus_ metadata body ->
--             Err (BadStatus metadata body)
--         Http.GoodStatus_ metadata body ->
--             Result.mapError (BadBody metadata body) <|
--                 Result.fromMaybe "Error decoding bytes" <|
--                     Bytes.Decode.decode (Bytes.Decode.map (\res -> ( metadata, res )) decoder) body


expectJsonDetailed : (Result (ErrorDetailed String) ( Http.Metadata, a ) -> msg) -> Decoder a -> Http.Expect msg
expectJsonDetailed msg decoder =
    Http.expectStringResponse msg (convertResponseStringToJson decoder)



-- expectStringDetailed : (Result (ErrorDetailed String) ( Http.Metadata, String ) -> msg) -> Http.Expect msg
-- expectStringDetailed msg =
--     Http.expectStringResponse msg convertResponseString
-- expectBytesDetailed : (Result (ErrorDetailed Bytes) ( Http.Metadata, a ) -> msg) -> Bytes.Decode.Decoder a -> Http.Expect msg
-- expectBytesDetailed msg decoder =
--     Http.expectBytesResponse msg (convertResponseBytes decoder)
-- PERSISTENCE


decode : Decoder (Cred -> a) -> Value -> Result Decode.Error a
decode decoder value =
    -- It's stored in localStorage as a JSON String;
    -- first decode the Value as a String, then
    -- decode that String as JSON.
    Decode.decodeValue Decode.string value
        |> Result.andThen (\str -> Decode.decodeString (Decode.field "user" (decoderFromCred decoder)) str)


port onStoreChange : (Value -> msg) -> Sub msg


viewerChanges : (Maybe viewer -> msg) -> Decoder (Cred -> viewer) -> Sub msg
viewerChanges toMsg decoder =
    onStoreChange (\value -> toMsg (decodeFromChange decoder value))


decodeFromChange : Decoder (Cred -> viewer) -> Value -> Maybe viewer
decodeFromChange viewerDecoder val =
    -- It's stored in localStorage as a JSON String;
    -- first decode the Value as a String, then
    -- decode that String as JSON.
    Decode.decodeValue (storageDecoder viewerDecoder) val
        |> Result.toMaybe


storeCredWith : Cred -> Cmd msg
storeCredWith (Cred uname token) =
    let
        json =
            Encode.object
                [ ( "user"
                  , Encode.object
                        [ ( "username", Username.encode uname )
                        , ( "token", Encode.string token )
                        ]
                  )
                ]
    in
    storeCache (Just json)


logout : Cmd msg
logout =
    storeCache Nothing


port storeCache : Maybe Value -> Cmd msg



-- SERIALIZATION
-- APPLICATION


application :
    Decoder (Cred -> viewer)
    ->
        { init : Maybe viewer -> Url -> Nav.Key -> ( model, Cmd msg )
        , onUrlChange : Url -> msg
        , onUrlRequest : Browser.UrlRequest -> msg
        , subscriptions : model -> Sub msg
        , update : msg -> model -> ( model, Cmd msg )
        , view : model -> Browser.Document msg
        }
    -> Program Value model msg
application viewerDecoder config =
    let
        init flags url navKey =
            let
                maybeViewer =
                    Decode.decodeValue Decode.string flags
                        |> Result.andThen (Decode.decodeString (storageDecoder viewerDecoder))
                        |> Result.toMaybe
            in
            config.init maybeViewer url navKey
    in
    Browser.application
        { init = init
        , onUrlChange = config.onUrlChange
        , onUrlRequest = config.onUrlRequest
        , subscriptions = config.subscriptions
        , update = config.update
        , view = config.view
        }


storageDecoder : Decoder (Cred -> viewer) -> Decoder viewer
storageDecoder viewerDecoder =
    Decode.field "user" (decoderFromCred viewerDecoder)



-- HTTP


get : Endpoint -> Maybe Cred -> (Result (ErrorDetailed String) ( Http.Metadata, a ) -> a) -> Decoder a -> Cmd a
get url maybeCred toMsg decoder =
    Endpoint.request
        { method = "GET"
        , url = url
        , expect = expectJsonDetailed toMsg decoder
        , headers =
            case maybeCred of
                Just cred ->
                    [ credHeader cred ]

                Nothing ->
                    []
        , body = Http.emptyBody
        , timeout = Nothing
        , tracker = Nothing
        }


put : Endpoint -> Cred -> Body -> (Result (ErrorDetailed String) ( Http.Metadata, a ) -> a) -> Decoder a -> Cmd a
put url cred body toMsg decoder =
    Endpoint.request
        { method = "PUT"
        , url = url
        , expect = expectJsonDetailed toMsg decoder
        , headers = [ credHeader cred ]
        , body = body
        , timeout = Nothing
        , tracker = Nothing
        }


post : Endpoint -> Maybe Cred -> Body -> (Result (ErrorDetailed String) ( Http.Metadata, a ) -> b) -> Decoder a -> Cmd b
post url maybeCred body toMsg decoder =
    Endpoint.request
        { method = "POST"
        , url = url
        , expect = expectJsonDetailed toMsg decoder
        , headers =
            case maybeCred of
                Just cred ->
                    [ credHeader cred ]

                Nothing ->
                    []
        , body = body
        , timeout = Nothing
        , tracker = Nothing
        }


delete : Endpoint -> Cred -> Body -> (Result (ErrorDetailed String) ( Http.Metadata, a ) -> a) -> Decoder a -> Cmd a
delete url cred body toMsg decoder =
    Endpoint.request
        { method = "DELETE"
        , url = url
        , expect = expectJsonDetailed toMsg decoder
        , headers = [ credHeader cred ]
        , body = body
        , timeout = Nothing
        , tracker = Nothing
        }


login : Http.Body -> (Result (ErrorDetailed String) ( Http.Metadata, a ) -> a) -> Decoder (Cred -> a) -> Cmd a
login body msg decoder =
    post Endpoint.login Nothing body msg (Decode.field "user" (decoderFromCred decoder))


register : Http.Body -> (Result (ErrorDetailed String) ( Http.Metadata, a ) -> a) -> Decoder (Cred -> a) -> Cmd a
register body msg decoder =
    post Endpoint.users Nothing body msg (Decode.field "user" (decoderFromCred decoder))


decoderFromCred : Decoder (Cred -> a) -> Decoder a
decoderFromCred decoder =
    Decode.map2 (\fromCred cred -> fromCred cred)
        decoder
        credDecoder



-- ERRORS


addServerError : List String -> List String
addServerError list =
    "Server error" :: list


{-| Many API endpoints include an "errors" field in their BadStatus responses.
-}
decodeErrors : ErrorDetailed String -> List String
decodeErrors error =
    case error of
        BadStatus _ body ->
            body
                |> decodeString (field "errors" errorsDecoder)
                |> Result.withDefault [ "Server error" ]

        err ->
            [ "Server error" ]


errorsDecoder : Decoder (List String)
errorsDecoder =
    Decode.keyValuePairs (Decode.list Decode.string)
        |> Decode.map (List.concatMap fromPair)


fromPair : ( String, List String ) -> List String
fromPair ( field, errors ) =
    List.map (\error -> field ++ " " ++ error) errors



-- LOCALSTORAGE KEYS


cacheStorageKey : String
cacheStorageKey =
    "cache"


credStorageKey : String
credStorageKey =
    "cred"
