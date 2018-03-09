package main
import (
	"net/http"
	"time"
	"log"

	"github.com/ziken/Register-of-Expenses-REST-API-go/routes"
)

func main() {
	r := routes.NewRouter()

	server := &http.Server{
		Addr: "127.0.0.1:3000",
		Handler: r,

		WriteTimeout: time.Second * 15,
		ReadTimeout: time.Second * 15,
		IdleTimeout: time.Second * 60,
	}
	c := make(chan int)

	go func(){
		log.Println("Server is running on port 3000");
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	<-c
}
