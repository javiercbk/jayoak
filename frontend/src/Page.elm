module Page exposing (Page(..), view, viewErrors)

import Asset
import Browser exposing (Document)
import Bulma.Classes as Bulma
import Bulma.Helpers as BulmaHelpers
import Html exposing (Html, a, article, button, div, footer, i, img, li, nav, p, span, text, ul)
import Html.Attributes exposing (alt, attribute, class, classList, href, style)
import Html.Attributes.Aria as Aria
import Html.Events exposing (onClick)
import Route exposing (Route)
import Username exposing (Username)
import Viewer exposing (Viewer)


{-| Determines which navbar link (if any) will be rendered as active.

Note that we don't enumerate every page here, because the navbar doesn't
have links for every page. Anything that's not part of the navbar falls
under Other.

-}
type Page
    = Other
    | Home
    | Login
    | Register
    | Settings
    | Profile Username
    | NewArticle


{-| Take a page's Html and frames it with a header and footer.

The caller provides the current user, so we can display in either
"signed in" (rendering username) or "signed out" mode.

isLoading is for determining whether we should show a loading spinner
in the header. (This comes up during slow page transitions.)

-}
view : Maybe Viewer -> Page -> { title : String, content : Html msg } -> Document msg
view maybeViewer page { title, content } =
    { title = title ++ " - jayoak"
    , body = viewHeader page maybeViewer :: content :: []
    }


viewHeader : Page -> Maybe Viewer -> Html msg
viewHeader page maybeViewer =
    nav [ attribute "aria-label" "main navigation", attribute "role" "navigation", BulmaHelpers.classList [ Bulma.navbar, Bulma.isFixedTop, Bulma.hasShadow ] ]
        [ div [ class Bulma.container ]
            [ div [ class Bulma.navbarBrand ]
                [ a [ class Bulma.navbarItem, Route.href Route.Home ]
                    [ img [ alt "jayoak", Asset.src Asset.logoLog ]
                        []
                    ]
                ]
            ]
        ]


{-| Render dismissable errors. We use this all over the place!
-}
viewErrors : msg -> List String -> Html msg
viewErrors dismissErrors errors =
    if List.isEmpty errors then
        Html.text ""

    else
        div [] <|
            List.map
                (\error ->
                    article
                        [ BulmaHelpers.classList [ Bulma.message, Bulma.isDanger ] ]
                        [ div [ class Bulma.messageHeader ]
                            [ p [] [ text "Error" ]
                            , button [ class Bulma.delete, Aria.ariaLabel "delete", onClick dismissErrors ] []
                            ]
                        , div [ class Bulma.messageBody ]
                            [ text error
                            ]
                        ]
                )
                errors
