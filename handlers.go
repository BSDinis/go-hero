package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "strconv"
  "io/ioutil"
  "io"
  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome to the hero handler!")
}

func ListAll(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(heroes); err != nil {
    panic(err)
  }
}

func AddNew(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  var hero Hero
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 0xfffff))

  if err != nil {
    panic(err)
  }
  if err := r.Body.Close(); err != nil {
    panic(err)
  }
  if err := json.Unmarshal(body, &hero); err != nil {
    w.WriteHeader(422) // cannot process
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  h, err := DBCreateHero(hero)
  if err != nil {
    w.WriteHeader(409) // id taken
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(h); err != nil {
    panic(err)
  }
}

func DeleteHero(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  vars := mux.Vars(r)
  idStr := vars["id"]
  id, err := strconv.Atoi(idStr)
  if err != nil {
    w.WriteHeader(400) // wrong format
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  hero, err := DBDeleteHero(id)
  if err != nil {
    w.WriteHeader(404) // hero not found
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(hero); err != nil {
    panic(err)
  }
}

func GetHero(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  vars := mux.Vars(r)
  idStr := vars["id"]
  id, err := strconv.Atoi(idStr)
  if err != nil {
    w.WriteHeader(400) // wrong format
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  hero, err := DBFindHero(id)
  if err != nil {
    w.WriteHeader(404) // hero not found
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(hero); err != nil {
    panic(err)
  }
}

func ChangeHero(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  decoder := json.NewDecoder(r.Body)
  var h Hero
  err := decoder.Decode(&h)
  if err != nil {
    w.WriteHeader(400) // wrong format
  }

  hero, err := DBUpdateHero(h)
  if err != nil {
    w.WriteHeader(404) // hero not found
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(hero); err != nil {
    panic(err)
  }
}
