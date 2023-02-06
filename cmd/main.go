package main

import (
	"clerks/app/services"
	"clerks/persistence/arango"
	"clerks/persistence/random_user"
	"clerks/presentation/users"
	"fmt"
	"github.com/arangodb/go-driver"
	arangogttp "github.com/arangodb/go-driver/http"
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
	arangoUrl := getEnv("ARANGO_URL", "http://localhost:8572")
	arangoUser := "root"
	arangoPass := "rootpassword"
	arangoDatabase := "example"
	serveOnPort := getEnv("SERVE_ON_PORT", "8000")
	randomUserURL := "https://randomuser.me/api/?results=5000&inc=,name,email,phone,picture,registered&noinfo"

	conn, err := arangogttp.NewConnection(arangogttp.ConnectionConfig{
		Endpoints: []string{arangoUrl},
	})
	if err != nil {
		log.Fatal("cannot connect to arangodb")
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(arangoUser, arangoPass),
	})
	if err != nil {
		log.Fatal("Cannot login to arangodb")
	}

	databaseExists, err := client.DatabaseExists(nil, arangoDatabase)
	if err != nil {
		log.Fatalf("failed to check databsase existance, err: %+v", err)
	}

	if !databaseExists {
		_, createDatabaseErr := client.CreateDatabase(nil, arangoDatabase, nil)
		if createDatabaseErr != nil {
			log.Fatalf("failed to create databsase, err: %+v", createDatabaseErr)
		}
	}

	database, err := client.Database(nil, arangoDatabase)
	if err != nil {
		log.Fatalf("Could not find database %s. %v", arangoDatabase, err)
	}

	userRepository, err := arango.NewUserRepository(database, "users")
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
	fmt.Printf("Started clerks service at port %s\n", serveOnPort)
	log.Fatal(http.ListenAndServe(":"+serveOnPort, r))

}
