package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request){
	todo := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
}

func main() {
	
}
