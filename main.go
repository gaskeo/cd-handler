package main

import (
	"cd-handler/files"
	"cd-handler/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	secret, existSecret := os.LookupEnv("CD_SECRET")

	if !existSecret {
		log.Fatal("CD_SECRET env not provided")
	}

	err := files.InitPath()
	if err != nil {
		log.Fatal(err)
	}

	handlers.RegisterSecretHandler(
		secret,
		func() error {
			err := files.ExecSecret()
			return err
		},
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
