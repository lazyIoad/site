let NavLink : Type =
  { Title : Text
  , Target : Text
  }

let nav : List NavLink =
  [ { Title = "home", Target = "/" }
  ,
  ]
in { Port        = 3000
   , Origin      = ""
   , NavLinks    = nav
   , Title       = "lazyloading"
   , Description = "words by areeb khan"
   }
