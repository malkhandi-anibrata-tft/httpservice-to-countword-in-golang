package requests

import (
	"fmt"
	"log"
	"net/http"
	"wcservice/service"
)

// func input(
// 	handler.inputHandler()
// )

func Request() {
	// handler.hii()
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/input", service.PostRequest) // user input page

	err := http.ListenAndServe(":8000", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "WelCome to My API")
}
