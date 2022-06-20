let NavLink : Type =
  { Title : Text
  , Target : Text
  }

let nav : List NavLink =
  [ { Title = "home", Target = "/" }
  ,
  ]
in { Port     = 3000
   , Domain   = ""
   , NavLinks = nav
   }
