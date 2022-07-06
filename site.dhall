let NavLink : Type =
  { Title : Text
  , Target : Text
  }

let nav : List NavLink =
  [ { Title = "git", Target = "https://git.sr.ht/~lazyload/" }
  , { Title = "twit", Target = "https://twitter.com/lazyIoad" }
  , { Title = "mail", Target = "mailto:hi@lazyloading.net" }
  , { Title = "rss", Target = "/feeds/blog/rss" }
  ]
in { Port        = 3000
   , Origin      = "https://lazyloading.net"
   , NavLinks    = nav
   , Title       = "lazyloading"
   , Description = "words by areeb khan"
   }
