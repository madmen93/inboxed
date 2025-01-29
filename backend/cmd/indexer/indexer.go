package indexer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

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

func deleteIndex(index string, zincURL string) error {
	zincURLPath := fmt.Sprintf("%s/api/index/%s", zincURL, index)

	err := getResponse("DELETE", zincURLPath, nil, "error en la eliminación del index", "Index eliminado exitosamente", false)
	if err != nil {
		return err
	}

	return nil
}

func createIndexZincSearch(indexer IndexerData, url string) error {
	zincURLPath := fmt.Sprintf("%s/api/index", url)

	data, err := json.Marshal(indexer)
	if err != nil {
		return fmt.Errorf("error al codear indexer: %v", err)
	}

	err = getResponse("POST", zincURLPath, bytes.NewReader(data), "error en la creación del index", "Index creado exitosamente", false)
	if err != nil {
		return err
	}

	return nil
}

func sendToZincSearch(records []EmailData, zincURL string) error {
	zincURLPath := fmt.Sprintf("%s/api/_bulkv2", zincURL)

	bulkData := BulkIngestion{
		Index:   "enron",
		Records: records,
	}

	jsonData, err := json.Marshal(bulkData)
	if err != nil {
		return fmt.Errorf("error al serializar datos: %v", err)
	}

	err = getResponse("POST", zincURLPath, bytes.NewReader(jsonData), "error al envíar datos a ZincSearch", "Lote enviado correctamente a ZincSearch", true)
	if err != nil {
		return err
	}

	return nil
}
