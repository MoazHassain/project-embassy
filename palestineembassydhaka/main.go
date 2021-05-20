package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var workingDirectory string
var websiteTemplateAbsPath string

func init() {

	workingDirectory, _ = os.Getwd()
	//fmt.Println("workingDirectory:", workingDirectory)
	websiteTemplateAbsPath = filepath.Join(workingDirectory, "templates", "website", "*.gohtml")
}
func main() {

	assetPath := filepath.Join(workingDirectory, "assets")
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir(assetPath))))

	//fmt.Println("Hello palestine")
	http.HandleFunc("/", index)

	http.ListenAndServe(":8085", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("home panicking >>", rec)
		}
	}()

	var baseurl string = GetBaseURL(r)

	ptmp, err := template.ParseGlob(websiteTemplateAbsPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	data := struct {
		Title string
		Base  string
	}{
		Title: "Palesine Embassy in Bangladesh",
		Base:  baseurl,
	}
	ptmp.Execute(w, data)
	//fmt.Fprintf(w, `Welcome to plaestine embassy dhaka`)

}
