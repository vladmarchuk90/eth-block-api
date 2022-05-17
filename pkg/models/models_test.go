package models

import (
	"path/filepath"
	"testing"

	"github.com/jellydator/ttlcache/v3"
	"github.com/vladmarchuk90/eth-block-api/pkg/config"
)

func TestGetBlockInfo(t *testing.T) {
	blockNumber := "11509797"
	expectedResult := `{"transactions":155,"amount":2.285404805647828}`

	configFilename, _ := filepath.Abs("../../config.json")
	app := config.NewConfig(configFilename)

	cache := ttlcache.New[string, string]()
	app.Cache = cache

	NewModels(app)

	result, err := GetBlockInfo(blockNumber)
	if err != nil {
		t.Error("There error was got when it shouldn't")
	}

	if result != expectedResult {
		t.Error("Actual result didn't equal expected result")
	}
}
