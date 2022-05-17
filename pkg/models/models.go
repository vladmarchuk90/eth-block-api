package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jellydator/ttlcache/v3"
	"github.com/vladmarchuk90/eth-block-api/pkg/config"
)

// BlockInfoTotal holds number of transactions and amount in asked block
type BlockInfoTotal struct {
	Transactions int     `json:"transactions"`
	Amount       float64 `json:"amount"`
}

// EthResponse describes one of the root element in ethereum json response
type EthResponse struct {
	Result EthResult `json:"result"`
}

// EthResult describes main characteristics of ethereum block, included all transactions
type EthResult struct {
	Transactions []EthTransaction `json:"transactions"`
}

// EthTransaction describes particular transaction in ethereum block
type EthTransaction struct {
	Value string `json:"value"`
}

// wei is 1✕10-18 Ether
const wei = 0.000000000000000001

var app *config.AppConfig

func NewModels(appConfig *config.AppConfig) {
	app = appConfig
}

func GetBlockInfo(blockNumber string) (string, error) {

	// checking value in cache if caching is used
	if app.UseCache {
		item := app.Cache.Get(blockNumber)
		if item != nil {
			return item.Value(), nil
		}
	}

	var blockInfoJson string

	// converting string to int
	blockNumberInt, err := strconv.Atoi(blockNumber)
	if err != nil {
		return blockInfoJson, errors.New("could not convert passed block number to int")
	}

	// getting data by ethereum api
	blockInfo, err := getEthData(blockNumberInt)
	if err != nil {
		log.Fatal(err)
	}

	// converting to json
	byteValue, err := json.Marshal(blockInfo)
	if err != nil {
		log.Fatal(err)
	}

	blockInfoJson = string(byteValue)

	// caching results
	if app.UseCache {
		app.Cache.Set(blockNumber, blockInfoJson, ttlcache.NoTTL)
	}

	return blockInfoJson, nil
}

func getEthData(blockNumber int) (*BlockInfoTotal, error) {

	// it takes default api template that came from config file
	sourceUrl, err := url.Parse(app.ApiTemplateUrl)
	if err != nil {
		return nil, errors.New("could not parse api_template_url, check your config file")
	}

	// changing get query parameters
	values := sourceUrl.Query()
	values.Set("tag", fmt.Sprintf("0x%x", blockNumber))
	values.Set("apikey", app.ApiKey)

	sourceUrl.RawQuery = values.Encode()

	// calling ethereum api
	resp, err := http.Get(sourceUrl.String())
	if err != nil {
		return nil, errors.New("could not get success response from external ethereum api")
	}

	defer resp.Body.Close()

	byteValue, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("could not convert body response to bytes")
	}

	// getting EthResponse struct with needed information for analisys
	var ethResponse EthResponse
	err = json.Unmarshal(byteValue, &ethResponse)
	if err != nil {
		return nil, errors.New("error unmarshalling ethereum response")
	}

	// BlockInfoTotal is return type that holds number of transactions and them value
	var blockInfo BlockInfoTotal = BlockInfoTotal{
		Transactions: len(ethResponse.Result.Transactions),
	}

	// counting transactions' amount
	var total uint64
	for _, transaction := range ethResponse.Result.Transactions {
		value, _ := strconv.ParseUint(transaction.Value[2:], 16, 64)
		total += value
	}
	blockInfo.Amount = float64(total) * wei

	return &blockInfo, nil
}
