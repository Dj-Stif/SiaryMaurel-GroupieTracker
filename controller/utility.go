package controller

import (
	"GroupieTracker/backend"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Armures(data backend.Armures) backend.Armures {

	apiUrl := "https://mhw-db.com/armor/sets"

	req, err := http.NewRequest("GET", apiUrl, nil) //création requete
	if err != nil {
		fmt.Println("Error creatoin request", err)
		return data
	}

	client := &http.Client{}
	resp, err := client.Do(req) //envoie requete dans resp
	if err != nil {
		fmt.Println("error envoie request", err)
		return data
	}
	defer resp.Body.Close() //ferme la requete http

	body, err := io.ReadAll(resp.Body) //lecture de la requete
	if err != nil {
		fmt.Println("error lecture request", err)
		return data
	}

	json.Unmarshal(body, &data)
	for _, i := range data {
		fmt.Println("Nom : ", i.Name)
		for _, j := range i.Pieces {
			fmt.Println("Type : ", j.Type)
			fmt.Println("Def Base : ", j.Defense.Base)
		}
	}
	return data
}

func Weapons(data backend.Weapons) backend.Weapons {
	apiUrl := "https://mhw-db.com/weapons"

	req, err := http.NewRequest("GET", apiUrl, nil) //création requete
	if err != nil {
		fmt.Println("Error creatoin request", err)
		return data
	}

	client := &http.Client{}
	resp, err := client.Do(req) //envoie requete dans resp
	if err != nil {
		fmt.Println("error envoie request", err)
		return data
	}
	defer resp.Body.Close() //ferme la requete http

	body, err := io.ReadAll(resp.Body) //lecture de la requete
	if err != nil {
		fmt.Println("error lecture request", err)
		return data
	}

	json.Unmarshal(body, &data)
	for _, i := range data {
		fmt.Println("Nom : ", i.Name)
		fmt.Println("Type : ", i.Type)
		fmt.Println("Rarity : ", i.Rarity)
		fmt.Println("DamageType", i.DamageType)
	}
	return data
}

func Monsters(data backend.Monsters) backend.Monsters {
	apiUrl := "https://mhw-db.com/monsters"

	req, err := http.NewRequest("GET", apiUrl, nil) //création requete
	if err != nil {
		fmt.Println("Error creatoin request", err)
		return data
	}

	client := &http.Client{}
	resp, err := client.Do(req) //envoie requete dans resp
	if err != nil {
		fmt.Println("error envoie request", err)
		return data
	}
	defer resp.Body.Close() //ferme la requete http

	body, err := io.ReadAll(resp.Body) //lecture de la requete
	if err != nil {
		fmt.Println("error lecture request", err)
		return data
	}

	json.Unmarshal(body, &data)
	for _, i := range data {
		fmt.Println("Nom : ", i.Name)
		fmt.Println("Type : ", i.Type)
		fmt.Println("Species : ", i.Species)
		for _, j := range i.Weaknesses {
			fmt.Println("Type : ", j.Element)
		}
	}
	return data
}
