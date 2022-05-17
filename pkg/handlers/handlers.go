package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vladmarchuk90/eth-block-api/pkg/config"
	"github.com/vladmarchuk90/eth-block-api/pkg/models"
)

// Repo the repository used by a handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a  new repository
func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{
		App: appConfig,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(repo *Repository) {
	Repo = repo
}

// GetBlockInfo is the handler for the eth block total info: number of transactions and value's sum
func (m *Repository) GetBlockInfo(w http.ResponseWriter, r *http.Request) {
	blockNumber := chi.URLParam(r, "blockNumber")
	blockInfo, err := models.GetBlockInfo(blockNumber)
	if err != nil {
		w.WriteHeader(404)
		errJSON, _ := json.Marshal(err)
		w.Write([]byte(errJSON))
	}

	w.Write([]byte(blockInfo))
}
