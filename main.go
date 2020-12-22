package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kpango/glg"
)

func main() {
	searcher := Searcher{}
	err := searcher.Load("completeworks.txt")
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/search", handleSearch(searcher))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Listening on port %s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Searcher struct {
	CompleteWorks string
	Index         *suffixarray.Index
}

func handleSearch(searcher Searcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query, ok := r.URL.Query()["q"]
		if !ok || len(query[0]) < 1 {
			err := "missing search query in URL params"
			glg.Error("handleSearch", "URL.Query", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err))
			return
		}
		results := searcher.Search(query[0])
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(results)
		if err != nil {
			glg.Error("handleSearch", "enc.Encode", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		glg.Printf("[handleSearch] found %d matches", len(results))
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}
}

func (s *Searcher) Load(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Load: %w", err)
	}
	s.CompleteWorks = strings.ToLower(string(dat))
	s.Index = suffixarray.New(dat)
	return nil
}

func (s *Searcher) Search(query string) (results []string) {
	qSplitted := filterText(query)
	for _, word := range qSplitted {
		idxs := s.Index.Lookup([]byte(word), -1)
		for _, idx := range idxs {
			maxIdx := idx + 250
			cwSize := len(s.CompleteWorks)
			if maxIdx > cwSize {
				maxIdx = cwSize - 1
			}
			minIdx := idx - 250
			if minIdx < 0 {
				minIdx = 0
			}
			results = append(results, s.CompleteWorks[minIdx:maxIdx])
		}
	}
	return results
}
