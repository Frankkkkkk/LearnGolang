package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"welcome")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	filePtr, err := os.Open("Douban.json")
	if err != nil {
		fmt.Println("Open file failed ", err.Error())
	}
	defer filePtr.Close()

	responseData, err := ioutil.ReadAll(filePtr)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBufferString(string(responseData))
	fmt.Fprint(w,buf)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusNotFound)
	//if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
	//	panic(err)
	//}

	value :=r.Header.Get("Authorization")
	s := strings.Split(value, " ")
	fmt.Println(s[1])

	//http.Redirect(w, r, "")
	fmt.Printf("host=%s\n",r.URL.Host)
	fmt.Printf("path=%s\n",r.URL.Path)
	fmt.Printf("raw path=%s\n",r.URL.RawPath)
	fmt.Printf("raw query=%s\n",r.URL.RawQuery)
	data := map[string]interface{}{
		"expires_in": time.Hour * 2,
		"client_id":222222,
		"user_id":2222222,
	}
	bytesRepresentation, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post("http://localhost:9096/login?"+r.URL.RawQuery, "application/json; charset=UTF-8", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
	log.Println(result["data"])
	fmt.Fprint(w,result)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}