package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

var root string

func handler(w http.ResponseWriter, r *http.Request) {
	basepath := root + r.URL.Path
	filename := basepath

	if fi, err := os.Stat(filename); fi != nil && fi.IsDir() {
		filename = basepath + "/index.html"
	} else if os.IsNotExist(err) {
		filename = basepath + ".html"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		filename = root + "/404.html"
		page, _ := ioutil.ReadFile(filename)
		w.WriteHeader(http.StatusNotFound)
		w.Write(page)
		fmt.Printf("GET %s 404 NOT FOUND\n", r.URL.Path)
		return
	}

	filename = filepath.Clean(filename)
	renderFile(w, r, filename)
	fmt.Printf("GET %s (%s) 200 OK\n", r.URL.Path, filename)
}

func main() {
	var port int
	flag.StringVar(&root, "root", ".",
		"The root directory to serve files from "+
			"(your GitHub Pages repo)")
	flag.IntVar(&port, "port", 8080, "The port to serve over")
	flag.Parse()

	root, err := filepath.Abs(root)
	if err != nil {
		fmt.Printf("Unable to serve directory: %s\n", root)
		os.Exit(1)
	}

	fmt.Printf("Serving from %s on http://localhost:%d ...\n", root, port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func renderFile(w http.ResponseWriter, r *http.Request, filename string) {
	ext := path.Ext(filename)
	switch ext {
	case ".md":
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			http.ServeFile(w, r, filename)
		}

		htmlContent, err := convertMarkdown(content)
		if err != nil {
			http.ServeFile(w, r, filename)
		}

		w.Write(htmlContent)
	default:
		http.ServeFile(w, r, filename)
	}
}

func convertMarkdown(mdText []byte) ([]byte, error) {
	body := bytes.NewReader(mdText)
	req, err := http.NewRequest("POST", "https://api.github.com/markdown/raw", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/plain")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
