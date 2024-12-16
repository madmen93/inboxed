package indexer

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func IndixingMainProcess() {
	log.Println("Creando indexer...")

	zincURL := os.Getenv("ZINC_URL")
	if zincURL == "" {
		zincURL = "http://localhost:4080"
	}

	log.Printf("Conectando a ZincSearch en %s...\n", zincURL)

	indexer, err := createIndexerFromJson("/app/cmd/indexer/index.json")
	if err != nil {
		log.Printf("Error al crear indexer desde el archivo JSON: %v", err)
		return
	}

	log.Println("Eliminando index en caso exista...")
	deleted := deleteIndex("enron")
	if deleted != nil {
		log.Println("Index no existe. Creando uno nuevo")
	}

	err = createIndexZincSearch(indexer, zincURL)
	if err != nil {
		log.Printf("Error al crear indexer en ZincSearch: %v", err)
		return
	}

	log.Println("Indexer creado exitosamente")
	log.Println("Empezando indexación...")

	start := time.Now()

	var records = make([]EmailData, 0, 1000)
	var m sync.Mutex
	var wg sync.WaitGroup

	slicePool := sync.Pool{
		New: func() interface{} {
			batch := make([]EmailData, 0, 1000)
			return &batch
		},
	}

	sem := make(chan struct{}, 4)

	err = filepath.Walk("/app/data/maildir/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			wg.Add(1)
			sem <- struct{}{}

			go func(p string) {
				defer wg.Done()
				defer func() { <-sem }()

				emailData, err := processFile(p)
				if err != nil {
					log.Printf("Error procesando archivo %s: %v\n", p, err)
					return
				}
				m.Lock()
				records = append(records, emailData)

				if len(records) >= 1000 {
					batchPtr := slicePool.Get().(*[]EmailData)
					*batchPtr = append((*batchPtr)[:0], records...)
					records = records[:0]
					m.Unlock()
					sendToZincSearch(*batchPtr, zincURL)
					slicePool.Put(batchPtr)
				} else {
					m.Unlock()
				}
			}(path)
		}
		return nil
	})
	if err != nil {
		log.Printf("Error al recorrer la estructura de directorios: %v", err)
		return
	}

	wg.Wait()

	if len(records) > 0 {
		sendToZincSearch(records, zincURL)
	}
	duration := time.Since(start)
	log.Println("La indexación ha concluido en ", duration)
}

// func IndixingMainProcess() {
// 	log.Println("Creando indexer...")

// 	zincURL := os.Getenv("ZINC_URL")
// 	if zincURL == "" {
// 		zincURL = "http://localhost:4080"
// 	}

// 	log.Printf("Conectando a ZincSearch en %s...\n", zincURL)

// 	indexer, err := createIndexerFromJson("/app/cmd/indexer/index.json")
// 	if err != nil {
// 		log.Printf("Error al crear indexer desde el archivo JSON: %v", err)
// 		return
// 	}

// 	log.Println("Eliminando index en caso exista...")
// 	deleted := deleteIndex("enron")
// 	if deleted != nil {
// 		log.Println("Index no existe. Creando uno nuevo")
// 	}

// 	err = createIndexZincSearch(indexer, zincURL)
// 	if err != nil {
// 		log.Printf("Error al crear indexer en ZincSearch: %v", err)
// 		return
// 	}

// 	log.Println("Indexer creado exitosamente")
// 	log.Println("Empezando indexación...")

// 	start := time.Now()

// 	var records = make([]EmailData, 0, 1000)
// 	var m sync.Mutex
// 	var wg sync.WaitGroup

// 	slicePool := sync.Pool{
// 		New: func() interface{} {
// 			batch := make([]EmailData, 0, 1000)
// 			return &batch
// 		},
// 	}

// 	err = filepath.Walk("/app/data/maildir/", func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		if !info.IsDir() {
// 			wg.Add(1)
// 			go func(p string) {
// 				defer wg.Done()
// 				emailData, err := processFile(p)
// 				if err != nil {
// 					log.Printf("Error procesando archivo %s: %v\n", p, err)
// 					return
// 				}
// 				m.Lock()
// 				records = append(records, emailData)

// 				if len(records) >= 1000 {
// 					batchPtr := slicePool.Get().(*[]EmailData)
// 					*batchPtr = append((*batchPtr)[:0], records...)
// 					records = records[:0]
// 					m.Unlock()
// 					sendToZincSearch(*batchPtr, zincURL)
// 					slicePool.Put(batchPtr)
// 				} else {
// 					m.Unlock()
// 				}
// 			}(path)
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		log.Printf("Error al recorrer la estructura de directorios: %v", err)
// 		return
// 	}

// 	wg.Wait()

// 	if len(records) > 0 {
// 		sendToZincSearch(records, zincURL)
// 	}
// 	duration := time.Since(start)
// 	log.Println("La indexación ha concluido en ", duration)
// }
