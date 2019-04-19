package main

import "net/http"

type Route struct {
  Name string
  Method string
  Pattern string
  Handler http.HandlerFunc
}

var routes = []Route{
  Route{
    "Index",
    "GET",
    "/",
    Index,
  },
  Route{
    "ListAll",
    "GET",
    "/hero",
    ListAll,
  },
  Route{
    "AddNew",
    "POST",
    "/hero",
    AddNew,
  },
  Route{
    "DeleteHero",
    "DELETE",
    "/hero/{id}",
    DeleteHero,
  },
  Route{
    "GetHero",
    "GET",
    "/hero/{id}",
    GetHero,
  },
  Route{
    "ChangeHero",
    "PUT",
    "/hero",
    ChangeHero,
  },
}
