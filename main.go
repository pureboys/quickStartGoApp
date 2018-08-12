package main

import (
	"net/http"
	"quickStartGoApp/routes"
	"quickStartGoApp/models"
	"quickStartGoApp/utils"
)

func main() {
	models.Init()
	utils.LoadTemplates("templates/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8888", nil)
}
