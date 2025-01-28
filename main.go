package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/dlinh31/go-text-search/utils"
)

func main(){
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "Hello world", "search query")
	flag.Parse()
	start := time.Now()
	docs, err := utils.LoadDocument(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Load %d documents in %v", len(docs), time.Since(start))

	
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	matchedIds := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIds), time.Since(start))

	for _, id := range matchedIds {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}