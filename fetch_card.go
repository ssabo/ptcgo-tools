package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GetCard(url string, wg *sync.WaitGroup) Card {
	defer wg.Done()
	doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
	}

	card := ParseCard(doc)

	fmt.Printf("Name: %s\t\tType: %s, ID: %s\n", card.Name(), card.Type(), card.Id())

	return card
}

func main() {
	var url string
	//var card Card
	var wg sync.WaitGroup

	for _, expansion := range Expansions() {
		fmt.Println("Processing: ", expansion.Name())

		wg.Add(expansion.Count())
		for i := 1; i <= expansion.Count(); i++ {

			url = fmt.Sprintf("http://www.pokemon.com/us/pokemon-tcg/pokemon-cards/%s/%s/%d/",
				expansion.Series(), expansion.Id(), i)
			go GetCard(url, &wg)
		}

		wg.Wait()
		// Pokemon.com fails if you hit it too agressively so backoff between sets
		time.Sleep(15 * time.Second)
	}
}
