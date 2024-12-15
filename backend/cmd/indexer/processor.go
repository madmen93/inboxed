package indexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/mail"
	"os"
)

func processFile(filepath string) (EmailData, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return EmailData{}, fmt.Errorf("no se pudo abrir el archivo: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	message, err := mail.ReadMessage(reader)
	if err != nil {
		return EmailData{}, fmt.Errorf("formato inválido: %v", err)
	}

	var body bytes.Buffer
	_, err = io.Copy(&body, message.Body)
	if err != nil {
		log.Printf("Error al leer cuerpo del mensaje en archivo: %s. Error: %v\n", filepath, err)
	}

	rawDate := message.Header.Get("Date")
	parsedDate, err := parseDate(rawDate)
	if err != nil {
		log.Printf("Error al parsear la fecha '%s': %v. Se guardará sin formato", rawDate, err)
	}

	return EmailData{
		MessageID:       message.Header.Get("Message-ID"),
		Date:            parsedDate,
		From:            message.Header.Get("From"),
		To:              message.Header.Get("To"),
		Subject:         message.Header.Get("Subject"),
		MimeVersion:     message.Header.Get("Mime-Version"),
		ContentType:     message.Header.Get("Content-Type"),
		ContentEncoding: message.Header.Get("Content-Transfer-Encoding"),
		XFrom:           message.Header.Get("X-From"),
		XTo:             message.Header.Get("X-To"),
		Xcc:             message.Header.Get("X-cc"),
		Xbcc:            message.Header.Get("X-bcc"),
		XFolder:         message.Header.Get("X-Folder"),
		XOrigin:         message.Header.Get("X-Origin"),
		XFilename:       message.Header.Get("X-FileName"),
		Content:         body.String(),
	}, nil
}
