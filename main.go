package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Batman",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World!"
	w.Write([]byte(message))
}

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var port = ":9000"
	fmt.Printf("🚀 [SERVER] is running on port http://localhost%s\n", port)
	server := new(http.Server)
	server.Addr = port
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
