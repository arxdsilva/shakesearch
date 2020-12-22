package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"index/suffixarray"
	"io/ioutil"
	"net/http"
	"os"
	"sort"

	"github.com/kpango/glg"
)

func main() {
	searcher := Searcher{}
	err := searcher.Load("completeworks.txt")
	if err != nil {
		glg.Fatal(err)
	}
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/search", handleSearch(searcher))
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	glg.Infof("Listening on port %s...", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		glg.Fatal(err)
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
			glg.Error("[handleSearch]", "URL.Query", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err))
			return
		}
		results := searcher.Search(query[0])
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err := enc.Encode(results)
		if err != nil {
			glg.Error("[handleSearch]", "enc.Encode", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("encoding failure"))
			return
		}
		glg.Infof("[handleSearch] found %d matches for %s", len(results), query[0])
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf.Bytes())
	}
}

func (s *Searcher) Load(filename string) error {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Load: %w", err)
	}
	s.CompleteWorks = string(dat)
	s.Index = suffixarray.New(bytes.ToLower(dat))
	return nil
}

func (s *Searcher) Search(query string) (results []string) {
	qSplitted := filterText(query)
	for _, word := range qSplitted {
		idxs := s.Index.Lookup([]byte(word), -1)
		sort.SliceStable(idxs, func(i, j int) bool {
			return idxs[i] < idxs[j]
		})
		idxs = removeDuplicates(idxs)
		for _, idx := range idxs {
			maxIdx := idx + 250
			cwLen := len(s.CompleteWorks)
			if maxIdx > cwLen {
				maxIdx = cwLen - 1
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
