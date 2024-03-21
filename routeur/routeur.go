package routeur

import (
	"GroupieTracker/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServe() {
	http.HandleFunc("/", controller.DisplayHome)
	http.HandleFunc("/armures", controller.DisplayArmures)
	http.HandleFunc("/weapons", controller.DisplayWeapons)
	http.HandleFunc("/monsters", controller.DisplayMonsters)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("\nLien vers le site : http://localhost:8080 (CTRL+CLICK)")
	http.ListenAndServe("localhost:8080", nil)
}
