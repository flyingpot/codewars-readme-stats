package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang/freetype"

	"github.com/hqbobo/text2pic"
)

func main() {
	http.Handle("/", http.HandlerFunc(generateSvg))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func generateSvg(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	usernames, present := query["username"]
	var username string
	if !present || len(usernames) == 0 {
		username = "flyingpot"
	}
	info, error := fetch(username)
	if error != nil {
		http.Error(w, "get stats error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	fontBytes, err := ioutil.ReadFile("font/Courier Prime.ttf")
	if err != nil {
		log.Println(err)
		return
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	pic := text2pic.NewTextPicture(text2pic.Configure{Width: 1500, BgColor: text2pic.ColorWhite})
	for _, line := range strings.Split(render(info), "\n") {
		if len(line) == 0 {
			continue
		}
		pic.AddTextLine(line, 8, f, text2pic.ColorBlack, text2pic.Padding{LineSpace: -15})
	}
	pic.Draw(w, text2pic.TypePng)
}
