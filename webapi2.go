package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type API struct {
	Message string "json:message"
}

type User struct {
	ID    int
	FName string
	LName string
	Age   int
}

func main() {

	apiRoot := "/api"

	// .../api
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home"}
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	// .../api/users
	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{ID: 1, FName: "remy0", LName: "pezo0", Age: 310},
			User{ID: 2, FName: "remy1", LName: "pezo1", Age: 311},
			User{ID: 3, FName: "remy2", LName: "pezo2", Age: 312},
		}
		message := users
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))
	})

	// .../api/me

	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := User{6, "dant", "rahmi", 995}
		message := user
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))

	})

	http.ListenAndServe(":8080", nil)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error: ", err.Error())
	}
}
