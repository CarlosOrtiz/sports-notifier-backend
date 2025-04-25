package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CarlosOrtiz/sports-notifier-backend/config"
	"github.com/CarlosOrtiz/sports-notifier-backend/database"
)

func main() {
	config.Load()

	//database.ConnectionPostgres()
	database.ConnectionMongoDb()

	//database.DB.AutoMigrate(models.Person{}, models.User{})

	appEnv := config.EnvConfig.AppEnv
	appHostServer := config.EnvConfig.AppHostServer
	appPort := config.EnvConfig.AppPort

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Servidor Go en el puerto %s!\n", appPort)
	})

	log.Printf("Server running in url: %s:%s on env %s", appPort, appHostServer, appEnv)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}
