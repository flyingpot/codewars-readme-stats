package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type UserInfo struct {
	Username            string      `json:"username"`
	Name                string      `json:"name"`
	Honor               int         `json:"honor"`
	Clan                interface{} `json:"clan"`
	Leaderboardposition int         `json:"leaderboardPosition"`
	Skills              interface{} `json:"skills"`
	Ranks               struct {
		Overall struct {
			Rank  int    `json:"rank"`
			Name  string `json:"name"`
			Color string `json:"color"`
			Score int    `json:"score"`
		} `json:"overall"`
		Languages struct {
			Java struct {
				Rank  int    `json:"rank"`
				Name  string `json:"name"`
				Color string `json:"color"`
				Score int    `json:"score"`
			} `json:"java"`
		} `json:"languages"`
	} `json:"ranks"`
	Codechallenges struct {
		Totalauthored  int `json:"totalAuthored"`
		Totalcompleted int `json:"totalCompleted"`
	} `json:"codeChallenges"`
}

func fetch(username string) UserInfo {
	resp, err := http.Get("https://www.codewars.com/api/v1/users/" + username)
	if err != nil {
		log.Fatal("get codewars api error")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var info UserInfo
	json.Unmarshal(body, &info)
	return info
}
