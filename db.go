package main

import "fmt"

var heroes []Hero

// init data
func init() {
  DBCreateHero(Hero{id: 1, name: "chico"})
  DBCreateHero(Hero{id: 2, name: "buarque"})
}

func DBCreateHero(h Hero) (Hero, error) {
  if _, err := DBFindHero(h.id) ; err != nil {
    return Hero{}, fmt.Errorf("Id %d is taken", h.id)
  }

  heroes = append(heroes, h)
  return h, nil
}

func DBFindHero(id int) (Hero, error) {
  for _, h := range heroes {
    if h.id == id {
      return h, nil
    }
  }

  return Hero{}, fmt.Errorf("Id %d not found", id)
}

func DBDeleteHero(id int) (Hero, error) {
  for i, h := range heroes {
    if h.id == id {
      heroes = append(heroes[:i], heroes[i+1:]...)
      return h, nil
    }
  }

  return Hero{}, fmt.Errorf("Id %d not found ", id)
}

func DBUpdateHero(hero Hero) (Hero, error) {
  for i, h := range heroes {
    if hero.id == h.id {
      heroes[i] = hero
      return hero, nil
    }
  }

  return hero, fmt.Errorf("Id %d not found", hero.id)
}
