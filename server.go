package main
import (
	"net/http"
	"time"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", nil).Methods("GET").HandlerFunc(func(w http.ResponseWriter, r * http.Request) {
		w.Write([]byte("Hello World!"))
	});

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
