package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type XMLNpcs struct {
	Npcs []Npc `xml:"npcs>entry"`
}

type XMLPlayers struct {
	Players []Player `xml:"players>entry"`
}

type Npc struct {
	Init uint8  `xml:"ini"`
	Name string `xml:"name"`
}

type Player struct {
	Init uint8  `xml:"ini"`
	Name string `xml:"name"`
}

func readXml(players *XMLPlayers, npcs *XMLNpcs) {
	data, err := os.Open("data/entities.xml")
	if err != nil {
		log.Fatal(err)
	}

	err = xml.NewDecoder(data).Decode(&players)
	_, err = data.Seek(0, 0)

	if err != nil {
		return
	}

	err = xml.NewDecoder(data).Decode(npcs)

	if err != nil {
		return
	}
}

func printEntities(players XMLPlayers, npcs XMLNpcs) {
	fmt.Println("Players\n===============")
	for i, p := range players.Players {
		fmt.Printf("%v: %s (Ini: %v)\n", i+1, p.Name, p.Init)
	}

	fmt.Println("'\n\nNPCs\n===============")
	for i, p := range npcs.Npcs {
		fmt.Printf("%v: %s (Ini: %v)\n", i+1, p.Name, p.Init)
	}
}

func choseEntities(players *XMLPlayers, npcs *XMLNpcs) {
	fmt.Println("Welche Spieler sollen am Kampf teilnehmen?")
	fmt.Println("1. Alle Spieler")
	fmt.Println("2. Auswählen")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	switch scanner.Text() {
	case "1":
		//TODO npcs auswählen
		fmt.Println("NPC auswählen screen")
	case "2":
		//TODO Subset der Player auswählen
		fmt.Println("Subset der Player auswählen")
	}
}

func main() {
	npcs := XMLNpcs{}
	players := XMLPlayers{}

	readXml(&players, &npcs)
	printEntities(players, npcs)

	fmt.Println("\n1. Alten Kampf fortsetzen\n2. Neuen Kampf starten")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	switch scanner.Text() {
	case "1":
		fmt.Println("Continuing fight not yet implemented")
	case "2":
		choseEntities(&players, &npcs)
	}
}
