package main

import (
	"clerks/app/services"
	"clerks/persistence/arango"
	"clerks/persistence/random_user"
	"clerks/presentation/users"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	serveOnPort := getEnv("SERVE_ON_PORT", "8000")
	randomUserURL := "https://randomuser.me/api/?results=5000&inc=,name,email,phone,picture,registered&noinfo"

	userRepository, err := arango.NewUserRepository()
	if err != nil {
		log.Fatalf("failed to instantiate user repository, err: %+v", err)
	}
	randomUserRepository, err := random_user.NewRandomUserRepository(randomUserURL)
	if err != nil {
		log.Fatalf("failed to instantiate random user repository, err: %+v", err)
	}

	uService := services.NewUserService(userRepository, randomUserRepository)

	userRoutes := users.CreateRoutes(uService)

	r := mux.NewRouter()
	s := r.PathPrefix("").Subrouter()

	for routePath, routeMethods := range userRoutes {
		fmt.Printf("adding %s route with methods %v\n", routePath, routeMethods.Methods)
		if routeMethods.Handler != nil {
			s.Handle(routePath, routeMethods.Handler).Methods(routeMethods.Methods...)
		} else {
			s.HandleFunc(routePath, routeMethods.HandlerFunc).Methods(routeMethods.Methods...)
		}
	}

	http.Handle("/", r)
	// Bind to a port and pass our router in
	fmt.Printf("Started rndm-usr at port %s\n", serveOnPort)
	log.Fatal(http.ListenAndServe(":"+serveOnPort, r))

}
