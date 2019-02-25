package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Title    string `json:"Tittle"`
	Subtitle string `json:"Subtitle"`
	Other    string `json:"Other"`
	Desc     string `json:"Desc"`
	Year     string `json:"Year"`
	Area     string `json:"Area"`
	Tag      string `json:"Tag"`
	Star     string `json:"Star"`
	Comment  string `json:"Comment"`
	//Quote    string `json:"Quote"`
	//Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
//type Pokemon struct {
//	EntryNo int            `json:"entry_number"`
//	Species PokemonSpecies `json:"pokemon_species"`
//}

// A struct to map our Pokemon's Species which includes it's name
//type PokemonSpecies struct {
//	Name string `json:"name"`
//}
//var v map[string]interface{}

func main() {

	filePtr, err := os.Open("Douban.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
	}
	defer filePtr.Close()





	//response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	//if err != nil {
	//	fmt.Print(err.Error())
	//	os.Exit(1)
	//}

	responseData, err := ioutil.ReadAll(filePtr)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s",responseData)

	//var responseObject Response
	//json.Unmarshal(responseData, &responseObject)

	//decoder := json.NewDecoder()
	//err = decoder.Decode(&responseData)
	//if err != nil {
	//	fmt.Println("Decoder failed", err.Error())
	//
	//} else {
	//	fmt.Println("Decoder success")
	//	//fmt.Println(movies)
	//}

	fmt.Println(string(responseData))
	//fmt.Println(len(responseObject.Pokemon))



}