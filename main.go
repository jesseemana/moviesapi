package main

import (
	"fmt"
	"log"
	"net/http"

	connectdb "github.com/jesseemana/moviesapi/connect"
	"github.com/jesseemana/moviesapi/router"
)

func main() {
	connectdb.ConnectDB()
	r := router.Router()
	
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":3030", r))
	fmt.Println("Server listening on port 3030...")
}
