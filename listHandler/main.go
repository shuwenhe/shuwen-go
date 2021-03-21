package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

const (
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
	ListDir      = 0x0001
)

var templates = make(map[string]*template.Template)

func init() {
	fileInfoAll, err := ioutil.ReadDir(TEMPLATE_DIR)
	checkError(err)
	var templateName, templatePath string
	for _, fileInfo := range fileInfoAll {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templatePath] = t
	}
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets", "./public", 0)
	http.HandleFunc("/", ListHandler)
	http.ListenAndServe(":8080", nil)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoAll, err := ioutil.ReadDir(UPLOAD_DIR)
	checkError(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoAll {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	renderHtml(w, "list", locals)
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	err := templates[tmpl].Execute(w, locals)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func staticDirHandler(mux *http.ServeMux, prefix, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if exists := isExists(file); exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
