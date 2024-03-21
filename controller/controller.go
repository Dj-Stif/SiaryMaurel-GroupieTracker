package controller

import (
	"GroupieTracker/backend"
	initTemplate "GroupieTracker/temps"
	"net/http"
)

func DisplayHome(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "index", nil)
}

func DisplayArmures(w http.ResponseWriter, r *http.Request) {
	var data backend.Armures
	armures := Armures(data)
	initTemplate.Temp.ExecuteTemplate(w, "armures", armures)
}

func DisplayWeapons(w http.ResponseWriter, r *http.Request) {
	var data backend.Weapons
	weapons := Weapons(data)
	initTemplate.Temp.ExecuteTemplate(w, "weapons", weapons)
}

func DisplayMonsters(w http.ResponseWriter, r *http.Request) {
	var data backend.Monsters
	monsters := Monsters(data)
	initTemplate.Temp.ExecuteTemplate(w, "monsters", monsters)
}
