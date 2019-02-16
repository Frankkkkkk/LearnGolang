//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"os"
//	"strconv"
//)
//
//type Users struct{
//	Users[]User `json:"users"`
//}
//
//type User struct{
//	Name string `json:"name"`
//	Type string `json:"type"`
//	Age int `json:"age"`
//	Social Social `json:"social"`
//}
//
//type Social struct{
//	Facebook string `json:"facebook"`
//	Twitter string `json:"twitter"`
//}
//
//
//func main() {
//	jsonFile, err := os.Open("users.json")
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("Successfully Opened users.json")
//	defer jsonFile.Close()
//
//	byteValue, _ := ioutil.ReadAll(jsonFile)
//
//	var users Users
//
//	// we unmarshal our byteArray which contains our
//	// jsonFile's content into 'users' which we defined above
//	json.Unmarshal(byteValue, &users)
//
//	// we iterate through every user within our users array and
//	// print out the user Type, their name, and their facebook url
//	// as just an example
//	for i := 0; i < len(users.Users); i++ {
//		fmt.Println("User Type: " + users.Users[i].Type)
//		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
//		fmt.Println("User Name: " + users.Users[i].Name)
//		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
//
//	}
//}
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].Species.Name)
	}

}