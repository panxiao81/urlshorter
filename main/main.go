package main

import (
	"flag"
	"fmt"
	"github.com/panxiao81/urlshorter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	configFile := flag.String("c", "config/config.yaml", "Path of the config file")
	flag.Parse()

	absPath, _ := filepath.Abs(*configFile)

	mux := defaultMux()

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	file, err := os.Open(absPath)
	if err != nil {
		log.Fatalf("Failed to Open file %s", absPath)
	}
	defer file.Close()
	yaml, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	yamlHandler, err := urlshorter.YAMLHandler(yaml, mux)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, "Yet another naive URL shorter")
}
