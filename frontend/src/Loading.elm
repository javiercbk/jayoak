module Loading exposing (error, icon, slowThreshold)

{-| A loading spinner icon.
-}

import Asset
import Html exposing (Attribute, Html)
import Html.Attributes exposing (class)
import Process
import Task exposing (Task)


icon : Html msg
icon =
    Html.i
        [ class "fas fa-compact-disc fa-spin" ]
        []


error : String -> Html msg
error str =
    Html.text ("Error loading " ++ str ++ ".")


slowThreshold : Task x ()
slowThreshold =
    Process.sleep 500
