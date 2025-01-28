package indexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"os"
	"time"
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
		return EmailData{}, fmt.Errorf("error copiando contenido: %s, %v", filepath, err)
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

func getResponse(method string, zincURLPath string, body io.Reader, errorMessage string, successMessage string, bulkIngestionFunction bool) error {
	req, err := http.NewRequest(method, zincURLPath, body)
	if err != nil {
		return fmt.Errorf("error en el request %s: %v", method, err)
	}
	req.SetBasicAuth(os.Getenv("ZINC_USER"), os.Getenv("ZINC_PASSWORD"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error al obtener la respuesta de la solicitud %s: %v", method, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && !bulkIngestionFunction {
		return fmt.Errorf("%s. StatusCode: %v", errorMessage, err)
	} else if resp.StatusCode != http.StatusOK && bulkIngestionFunction {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error del servidor ZincSearch. StatusCode: %d, Body: %s", resp.StatusCode, string(body))
	}

	log.Println(successMessage)
	return nil
}
