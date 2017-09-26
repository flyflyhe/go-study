package clientGame

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://log.lonlife.org/game/list"

type game struct {
	Id   int
	Name string
	Area string
}

func getGameJson() []byte {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func GetGame() map[int]game {
	strByte := getGameJson()
	m1 := make(map[int]game)
	var games []game
	err := json.Unmarshal(strByte, &games)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, val := range games {
		m1[val.Id] = val
	}
	return m1
}
