package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

func main() {
	handler := rest.ResourceHandler{
		EnableStatusService: true,
	}
	auth := &rest.AuthBasicMiddleware{
		Realm: "test zone",
		Authenticator: func(userId string, password string) bool {
			if userId == "admin" && password == "admin" {
				return true
			}
			return false
		},
	}
	err := handler.SetRoutes(
		&rest.Route{"GET", "/countries", GetAllCountries},
		&rest.Route{"GET", "/.status",
			auth.MiddlewareFunc(
				func(w rest.ResponseWriter, r *rest.Request) {
					w.WriteJson(handler.GetStatus())
				},
			),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", &handler))
}

type Country struct {
	Code string
	Name string
}

func GetAllCountries(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(
		[]Country{
			Country{
				Code: "FR",
				Name: "France",
			},
			Country{
				Code: "US",
				Name: "United States",
			},
		},
	)
}
