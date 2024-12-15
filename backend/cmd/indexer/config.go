package indexer

import "time"

type EmailData struct {
	From            string    `json:"from"`
	To              string    `json:"to"`
	Subject         string    `json:"subject"`
	Content         string    `json:"content"`
	MessageID       string    `json:"message_id"`
	Date            time.Time `json:"date"`
	ContentType     string    `json:"content_type"`
	MimeVersion     string    `json:"mime_version"`
	ContentEncoding string    `json:"content_transfer_encoding"`
	XFrom           string    `json:"x_from"`
	XTo             string    `json:"x_to"`
	Xcc             string    `json:"x_cc"`
	Xbcc            string    `json:"x_bcc"`
	XFolder         string    `json:"x_folder"`
	XOrigin         string    `json:"x_origin"`
	XFilename       string    `json:"x_filename"`
}

type BulkIngestion struct {
	Index   string      `json:"index"`
	Records []EmailData `json:"records"`
}

type PropertyIndex struct {
	Type          string `json:"type"`
	Index         bool   `json:"index"`
	Store         bool   `json:"store"`
	Sortable      bool   `json:"sortable"`
	Aggregatable  bool   `json:"aggregatable"`
	Highlightable bool   `json:"highlightable"`
}

type Mapping struct {
	Properties map[string]PropertyIndex `json:"properties"`
}

type IndexerData struct {
	Name         string  `json:"name"`
	StorageType  string  `json:"storage_type"`
	ShardNum     int     `json:"shard_num"`
	MappingField Mapping `json:"mappings"`
}
