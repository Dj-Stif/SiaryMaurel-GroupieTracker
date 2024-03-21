package main

import (
	"GroupieTracker/routeur"
	initTemplate "GroupieTracker/temps"
)

func main() {
	initTemplate.InitTemplate()
	routeur.InitServe()
}
