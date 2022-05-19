# Ethereum block info api
## General
This is small http web-service that allows get total information on particular block on Ethereum blockchain network. Internally for getting information used [Etherescan API](https://docs.etherscan.io/api-endpoints/geth-parity-proxy).

## External libraries used in this sevise
All actual external libraries used in application you can find in go.mod file. There are next:
1. [chi](https://github.com/go-chi/chi) is used as router. It is a lightweight, idiomatic and composable router for building Go HTTP services.
2. [ttlcache](github.com/jellydator/ttlcache) library is used as in-memory cache.

## API endpoint
One api endpoint is realized it's **/api/block/{blockNumber}/total**, where {blockNumber} is parameter and pointing to Ethereum network block you want getting information. Exploring blocks you are able on [etherscan](https://etherscan.io/blocks).
### Response
Http api response is in json format and has the next structure {"transactions":**NumberOfTransactionsInBlock**,"amount":**TotalAmountOfAllTransactions**}. For example for block 14797288 you get the next response {"transactions":160,"amount":10.49336032072198}.

## Installation and running
Clone the repository locally https://github.com/vladmarchuk90/eth-block-api.git after that you can run service by command _go run cmd/web/*.go_ or previously build it by command _go build cmd/web/*.go_ and after run _./main_ (if you built with default name). As well, you can build and run docker image by command **docker-compose up -d** (be sure you are in the project folder before running this comand).

## Application settings
You can provide some settings (see config.json). File config json should be in project root folder if you use run command or together in one folder with executable if you previously built it.
Default settings:
 1. "use_cache": true - turning on/off caching using, it's recommend to use cache.
 2. "api_key": "YourApiKeyToken" - you can get your personal api key at [https://etherscan.io/apis].
 3. "server_port": "80" - server port your service is running on.
 4. "api_template_url": "https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=0xafa01b&boolean=true&apikey=YourApiKeyToken" - ethereum api template which used as template to compose block info request: tag is blockNumber in hex format and apikey is your personal api key.
