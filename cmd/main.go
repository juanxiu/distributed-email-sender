package main

import (
	"distributed-email-sender/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := api.SetupRoutes()

	fmt.Println("서버를 시작합니다. 포트: 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
