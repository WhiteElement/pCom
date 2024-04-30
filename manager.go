package main

import (
	"encoding/xml"
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

func readXml(players *[]Player, npc *[]Npc) {
	data, err := os.Open("data/entities.xml")
	if err != nil {
		log.Fatal(err)
	}

	var parsedPlayers XMLPlayers
	var parsedNpcs XMLNpcs
	err = xml.NewDecoder(data).Decode(&parsedPlayers)
	_, err = data.Seek(0, 0)

	if err != nil {
		return
	}

	err = xml.NewDecoder(data).Decode(&parsedNpcs)

	if err != nil {
		return
	}
}

func main() {
	npcs := []Npc{}
	players := []Player{}

	readXml(&players, &npcs)

}
