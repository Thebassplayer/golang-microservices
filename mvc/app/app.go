package app

import (
	"net/http"
)

func StartApp() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!"))
	})

	err := http.ListenAndServe(":8080", nil) // Change the port if needed
	if err != nil {
		panic(err)
	}
}
