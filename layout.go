package helpers

import (
  "html/template"
  "net/http"
  "path/filepath"
  "log"
)

var LayoutDir string = "./assets/view/layouts"
var MasterTemplate string = "bootstrap"

func OutLayout(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(append(loadLayoutFiles(), filename)...)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.ExecuteTemplate(w, MasterTemplate, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}


func loadLayoutFiles() []string {
  files, err := filepath.Glob(LayoutDir + "/*.gohtml")
  if err != nil {
    panic(err)
  }
  log.Println(files)
  return files
}