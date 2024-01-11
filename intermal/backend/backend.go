package backend

import (
	"encoding/json"
	"eoda/intermal/db"
	"eoda/intermal/models"
	"fmt"
	"io"
	"net/http"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

func (s *Server) SetContainers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "not allowed method")
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Ошибка чтения тела запроса")
		return
	}

	var payload map[string][]string

	err = json.Unmarshal(body, &payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Ошибка разбора JSON")
		return
	}
	var images []*db.Image
	for _, name := range payload["images"] {
		images = append(images, &db.Image{Name: name})

	}
	// fmt.Println(images)
	s.db.Clauses(clause.OnConflict{DoNothing: true}).Create(images)

	fmt.Println("Получено JSON-тело:")
	fmt.Println(string(body))

	w.WriteHeader(http.StatusOK)
}
func (s *Server) CheckContainers(w http.ResponseWriter, r *http.Request) {
	var images []models.Image

	err := json.NewDecoder(r.Body).Decode(&images)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(images)
	for i, image := range images {
		var count int64
		s.db.Model(&db.Image{}).Where("name = ?", image.Name).Count(&count)
		if count > 0 {

			images[i].Accept = true
		}

	}
	err = json.NewEncoder(w).Encode(&images)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
