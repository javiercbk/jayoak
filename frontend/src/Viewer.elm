module Viewer exposing (Viewer, cred, decoder, minPasswordChars, store, username)

{-| The logged-in user currently viewing this page. It stores enough data to
be able to render the menu bar (username and avatar), along with Api.Cred so it's
impossible to have a Viewer if you aren't logged in.
-}

import Api
import Json.Decode as Decode exposing (Decoder)
import Json.Decode.Pipeline exposing (custom, required)
import Json.Encode as Encode exposing (Value)
import Username exposing (Username)



-- TYPES


type Viewer
    = Viewer Api.Cred



-- INFO


cred : Viewer -> Api.Cred
cred (Viewer val) =
    val


username : Viewer -> Username
username (Viewer val) =
    Api.username val


{-| Passwords must be at least this many characters long!
-}
minPasswordChars : Int
minPasswordChars =
    6



-- SERIALIZATION


decoder : Decoder (Api.Cred -> Viewer)
decoder =
    Decode.succeed Viewer


store : Viewer -> Cmd msg
store (Viewer credVal) =
    Api.storeCredWith
        credVal
