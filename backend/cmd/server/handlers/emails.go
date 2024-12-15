package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

// GetEmailsFiltered: Carga, filtra, empagina y ordena los correos.
func GetEmailsFiltered(w http.ResponseWriter, r *http.Request) {
	pageParam := r.URL.Query().Get("page")
	sizeParam := r.URL.Query().Get("size")
	sort := r.URL.Query().Get("sort")
	order := r.URL.Query().Get("order")
	field := r.URL.Query().Get("field")

	page := 1
	size := 10

	if pageParam != "" {
		if p, err := strconv.Atoi(pageParam); err == nil {
			page = p
		}
	}
	if sizeParam != "" {
		if s, err := strconv.Atoi(sizeParam); err == nil {
			size = s
		}
	}
	if sort == "" {
		sort = "date"
	}
	sortPrefix := ""
	if order == "desc" {
		sortPrefix = "-"
	} else {
		sortPrefix = ""
	}
	if field == "" {
		field = "content"
	}

	from := (page - 1) * size

	query := r.URL.Query().Get("q")
	query = strings.ToLower(query)
	var searchBody string

	if query == "" {
		searchBody = fmt.Sprintf(`{
			"search_type": "matchall",
			"sort_fields": ["%s%s"],
			"from": %d,
			"max_results": %d
		}`, sortPrefix, sort, from, size)
	} else {
		searchBody = fmt.Sprintf(`{
			"search_type": "match",
			"query": {
				"term": "%s",
				"field": "%s"
			},
			"sort_fields": ["%s%s"],
			"from": %d,
			"max_results": %d
		}`, query, field, sortPrefix, sort, from, size)
	}

	zincURL := os.Getenv("ZINC_URL") + "/api/enron/_search"

	req, err := http.NewRequest("POST", zincURL, bytes.NewBuffer([]byte(searchBody)))
	if err != nil {
		http.Error(w, "Fallo en la creación de la solicitud", http.StatusInternalServerError)
		return
	}
	req.SetBasicAuth(os.Getenv("ZINC_USER"), os.Getenv("ZINC_PASSWORD"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error al conectar con ZincSearch", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("ZincSearch devolvió status: %d", resp.StatusCode), resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error al leer la respuesta de ZincSearch", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

// GetEmabilByID: Trae el correo según el ID generado por ZincSearch
func GetEmailByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "El ID del email es requerido ", http.StatusBadRequest)
		return
	}

	zincURL := os.Getenv("ZINC_URL") + fmt.Sprintf("/api/enron/_doc/%s", id)

	req, err := http.NewRequest("GET", zincURL, nil)
	if err != nil {
		http.Error(w, "Fallo en la creación de la solicitud", http.StatusInternalServerError)
		return
	}

	req.SetBasicAuth(os.Getenv("ZINC_USER"), os.Getenv("ZINC_PASSWORD"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error al conectar con ZincSearch", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("ZincSearch devolvió status: %d", resp.StatusCode), resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error al leer la respuesta de ZincSearch", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
