package main

import (
	"log"
	"net/http"

	"github.com/fujisawaryohei/goblueprints/chat"
)

func main() {
	r := newRoom()
	http.Handle("/", &chat.TemplateHandler{Filename: "chat.html"})
	http.Handle("/room", r)
	// チャットルームの開始
	go r.run()

	// Webサーバーの起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
