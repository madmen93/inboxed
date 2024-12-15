package main

import (
	"sync"

	"github.com/madmen93/inboxed/cmd/indexer"
	"github.com/madmen93/inboxed/cmd/server"
)

func main() {
	//Profiling
	// go func() {
	// 	log.Println("Starting pprof server on :6060")
	// 	log.Println(http.ListenAndServe(":6060", nil))
	// }()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Start()
	}()

	indexer.IndixingMainProcess()

	wg.Wait()
}
