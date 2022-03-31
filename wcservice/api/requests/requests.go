package requests

import (
	"fmt"
	"log"
	"net/http"
	"wcservice/api/handler"
)

func homepage(m http.ResponseWriter, r *http.Request) {
	fmt.Fprint(m, "WelCome to My API")
}

// func input(
// 	handler.inputHandler()
// )

func Request() {
	// handler.hii()
	http.HandleFunc("/", homepage)
	http.HandleFunc("/input", handler.InputHandler) // user input page

	err := http.ListenAndServe(":8000", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
