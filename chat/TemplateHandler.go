package chat

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

type TemplateHandler struct {
	once     sync.Once
	Filename string
	templ    *template.Template
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("template", t.Filename)))
	})
	fmt.Println(t.templ)
	t.templ.Execute(w, nil)
}
