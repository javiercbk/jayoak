module Page.NotFound exposing (view)

import Bulma.Classes as Bulma
import Bulma.Helpers as BulmaHelpers
import Html exposing (Html, div, h1, h2, img, main_, section, text)
import Html.Attributes exposing (alt, class, id, src, tabindex)



-- VIEW


view : { title : String, content : Html msg }
view =
    { title = "Page Not Found"
    , content =
        section
            [ BulmaHelpers.classList [ Bulma.section, Bulma.hasBackgroundWhiteBis ] ]
            [ div [ BulmaHelpers.classList [ Bulma.columns, Bulma.isCentered ] ]
                [ div [ BulmaHelpers.classList [ Bulma.columns, Bulma.is8Tablet, Bulma.is6Desktop, Bulma.is4Widescreen ] ]
                    [ div [ class Bulma.box ]
                        [ div [ class Bulma.cardContent ]
                            [ h2 [ BulmaHelpers.classList [ Bulma.title, Bulma.hasTextCentered, Bulma.isSize3 ] ]
                                [ text "Not found" ]
                            ]
                        ]
                    ]
                ]
            ]
    }
