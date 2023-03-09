package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// Pokemon A pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// PokemonSpecies A struct to map our Pokemon's Species which includes its name
type PokemonSpecies struct {
	Name string `json:"name"`
}

func getPokemons() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].Species.Name)
	}
}
