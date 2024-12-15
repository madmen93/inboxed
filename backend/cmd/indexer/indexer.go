package indexer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func sendToZincSearch(records []EmailData, zincURL string) error {
	bulkData := BulkIngestion{
		Index:   "enron",
		Records: records,
	}

	jsonData, err := json.Marshal(bulkData)
	if err != nil {
		return fmt.Errorf("error al serializar datos: %v", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}

	bulkURL := fmt.Sprintf("%s/api/_bulkv2", zincURL)
	req, err := http.NewRequest("POST", bulkURL, bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("error al crear solicitud POST: %v", err)
	}

	req.SetBasicAuth(os.Getenv("ZINC_USER"), os.Getenv("ZINC_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error al enviar datos a ZincSearch: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error del servidor ZincSearch. StatusCode: %d, Body: %s", resp.StatusCode, string(body))
	}

	log.Println("Lote enviado correctamente a ZincSearch")
	return nil
}

func createIndexerFromJson(filepath string) (IndexerData, error) {
	var indexer IndexerData

	file, err := os.Open(filepath)
	if err != nil {
		return indexer, fmt.Errorf("error al abrir el archivo index.json: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&indexer)
	if err != nil {
		return indexer, fmt.Errorf("error al decodear archivo index.json: %v", err)
	}

	return indexer, nil
}

func createIndexZincSearch(indexer IndexerData, url string) error {
	data, err := json.Marshal(indexer)
	if err != nil {
		return fmt.Errorf("error al codear indexer: %v", err)
	}

	indexURL := fmt.Sprintf("%s/api/index", url)

	req, err := http.NewRequest("POST", indexURL, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("error en el método post en createIndexZincSearch: %v", err)
	}
	req.SetBasicAuth(os.Getenv("ZINC_USER"), os.Getenv("ZINC_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error en el retorno de la respuesta http: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error al crear el indexer. StatusCode: %v", resp.StatusCode)
	}

	return nil
}

func deleteIndex(index string) error {
	indexURL := fmt.Sprintf("%s/api/index/%s", os.Getenv("ZINC_URL"), index)

	req, err := http.NewRequest("DELETE", indexURL, nil)
	if err != nil {
		return fmt.Errorf("error en el request DELETE: %v", err)
	}
	req.SetBasicAuth(os.Getenv("ZINC_USER"), os.Getenv("ZINC_PASSWORD"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error en el retorno de la respuesta http: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error en la eliminación del index. StatusCode: %d", resp.StatusCode)
	}

	log.Println("Index deleted successfully")

	return nil
}

func parseDate(dateStr string) (time.Time, error) {
	const layout = "Mon, 02 Jan 2006 15:04:05 -0700 (MST)"

	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		const layoutTwo = "Mon, 2 Jan 2006 15:04:05 -0700 (MST)"
		parsedTime, err = time.Parse(layoutTwo, dateStr)
		if err != nil {
			return time.Time{}, fmt.Errorf("error al parsear la fecha: %v", err)
		}
	}

	return parsedTime, nil
}
