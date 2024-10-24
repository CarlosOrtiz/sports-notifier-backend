package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CarlosOrtiz/sports-notifier-backend/config"
)

func main() {
	config.Load()

	appEnv := config.EnvConfig.AppEnv
	appHostServer := config.EnvConfig.AppHostServer
	appPort := config.EnvConfig.AppPort

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Servidor Go en el puerto %s!\n", appPort)
	})

	log.Printf("Server running in %s url: %s with env %s", appPort, appHostServer, appEnv)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}
