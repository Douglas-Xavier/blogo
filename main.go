package main

import (
	"net/http"
	"yakki/blogo/providers"
)

func main() {

	userApi := providers.InitUserApi()
	server := http.NewServeMux()
	server.HandleFunc("/users", userApi.RouteByMethod)
	err := http.ListenAndServe(":8080", server)

	if err != nil {
		panic("Impossivel criar api")
	}

}
