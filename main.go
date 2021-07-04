package main

import (
	"log"
	"net/http"

	"github.com/fujisawaryohei/goblueprints/chat"
)

func main() {
	http.Handle("/", &chat.TemplateHandler{Filename: "chat.html"})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
