package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
)


type Response struct {
	Name 	string	`json:"name"`
	Pokemon	[]Pokemon	`json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo	int	`json:"entry_number"`
	Species	PokemonSpecies	`json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name	string	`json:"name"`
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

	// fmt.Println(string(responseData))

	var responseObject Response

	json.Unmarshal(responseData, &responseObject)

	// fmt.Println(responseObject.Name)
	// fmt.Println(len(responseObject.Pokemon))

	for i:= 0; i < len(responseObject.Pokemon); i++ {
		fmt.Printf("My pokemon name is: %s \n", responseObject.Pokemon[i].Species.Name)
	}

}