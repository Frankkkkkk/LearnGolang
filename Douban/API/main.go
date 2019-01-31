package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct{
	Tittle string `jason:Tittle`
	Desc string `jason:Desc`
	Content string `jason:Content`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request){
	articles :=Articles{
		Article{Tittle:"Test Tittle",Desc:"Test Desc", Content:"Test Content"},
	}

	fmt.Println("Endpoint: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w,"Homepage Endpoint Hit")
}

func handleRequest(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/articles",allArticles)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func main(){
	handleRequest()
}
