package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func isTrainer(c_type string) bool {
	trainer_types := []string{
		"Trainer - Supporter", "Trainer - Item", "Trainer [Pokémon Tool]",
		"Trainer - Stadium",
	}

	for _, i := range trainer_types {
		if i == c_type {
			return true
		}
	}
	return false
}

func isPokemon(c_type string) bool {
	pokemon_types := []string{
		"Basic Pokémon", "Stage 1 Pokémon", "Stage 2 Pokémon",
		"Pokémon-EX", "MEGA",
		"BREAK", "Restored Pokémon",
	}

	for _, i := range pokemon_types {
		if i == c_type {
			return true
		}
	}
	return false
}

func isEnergy(c_type string) bool {
	energy_types := []string{
		"Special Energy",
		"Basic Energy",
	}

	for _, i := range energy_types {
		if i == c_type {
			return true
		}
	}
	return false
}

func ParseCard(doc *goquery.Document) Card {
	// Create the card object
	var card Card
	// Create a map to store arbitrary data in before we create the card
	card_map := make(map[string]string)

	// Card description parsing
	// Find the name and type
	doc.Find("div.card-description").Each(func(i int, s *goquery.Selection) {
		// Card Name
		c_name := s.Find("div h1").Text()
		card_map["name"] = c_name

		// Card Type
		c_type := s.Find("div.card-type h2").Text()
		card_map["type"] = c_type
	})

	// Card Footer parsing
	// Find the ID, set, and rarity
	doc.Find("div.stats-footer").Each(func(i int, s *goquery.Selection) {
		// Set card is in
		set := s.Find("h3 a").Text()
		card_map["set"] = set

		// Card's ID # and Rarity
		id_rarity_str := string(s.Find("span").Text())
		id_rarity_arr := strings.Split(id_rarity_str, " ")

		id_full := id_rarity_arr[0]
		id := strings.Split(id_full, "/")[0]
		card_map["id"] = id

		rarity := id_rarity_arr[1]
		card_map["rarity"] = rarity
	})

	if isPokemon(card_map["type"]) {
		// Parse pokemon stats
		doc.Find("span.card-hp").Each(func(i int, s *goquery.Selection) {
			hp := s.Text()[2:]
			card_map["hp"] = hp
		})
		hp_int, _ := strconv.Atoi(card_map["hp"])

		card = NewPokemon(
			card_map["name"], card_map["type"], card_map["set"],
			card_map["id"], card_map["rarity"], hp_int,
		)
	} else if isTrainer(card_map["type"]) {
		// Parse Trainer text
		doc.Find("div.pokemon-abilities div.ability p").Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			card_map["text"] = text
		})
		card = NewTrainer(
			card_map["name"], card_map["type"], card_map["set"],
			card_map["id"], card_map["rarity"], card_map["text"],
		)
	} else if isEnergy(card_map["type"]) {
		// Parse Energy
		// TODO deal with energy better
		card = NewCard(card_map["name"], card_map["type"], card_map["set"],
			card_map["id"], card_map["rarity"])
	} else {
		// Could not detect card type, create a generic card with only the basics
		// card = NewCard(card_map["name"], card_map["type"], card_map["set"],
		// 	card_map["id"], card_map["rarity"])
	}

	if card == nil {
		fmt.Println(card_map)
		log.Fatal("Card not initialized correctly")
	}

	return card
}
