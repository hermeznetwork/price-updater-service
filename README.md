# Price Updater Service

Price Updater is a web service to consult token and currency prices used by Hermez Node.

## How It Works

1. The system was built with Golang and designed to run through [commands](##usage).
2. Given a specific time to update (the time can be set in through configuration, or by using a default value of 30s), the system updates the prices of currency and tokens.
3. The prices of tokens and currencies are returned through an endpoint.

## Instalation

To install the system just clone and build it!

```bash=
$ git clone git@github.com:hermeznetwork/price-updater-service.git
$ cd price-updater-service/
$ go build -o priceupdater # or other name that you want
```

## Usage 

The system has been designed to run through the commands listed above


### Configuration

You can set up the configuration of the system by env file

- HTTP_HOST:
    Host to server
- HTTP_PORT:
    Port to server
- POSTGRES_USER (required):
    Username of hermez node database 
- POSTGRES_PASSWORD (required):
    Password of hermez node database
- POSTGRES_HOST (required):
    Host address of hermez node database
- POSTGRES_PORT (default: 5432):
    Port of hermez node database
- POSTGRES_DATABASE (required):
     Database name of hermez node database
- POSTGRES_SSL_ENABLED:
- POSTGRES_MAX_ID_CONNS:
- POSTGRES_MAX_OPEN_CONNS:
- ETH_NETWORK (required by Uniswap provider):
    URL address to Ethereum network (mainnet, rinkeby, goerli)
- ETH_HEZ_ROLLUP (required by Uniswap provider):
    Address of Smart Contract (mainnet, rinkeby, goerli)
- ETH_USDT_ADDRESS (required by Uniswap provider):
    Address of USDT token on network (mainnet, rinkeby, goerli)
- BBOLT_LOCATION (default: /tmp/price_updater.db):
    Local database used to system. You must choice a better location :+1: 
- MAIN_TIME_TO_UPDATE_PRICES (default: 30s):
    Time between updates executed by system.

### Provider Configuration

There are two commands related to providers:

1. `change-provider --provider <provider-name>`
    
    This command must be executed before the other commands. It tells the system what provider must be executed.
    
    It is used to change the provider in runtime.
    <details><summary><b>Show example</b></summary>
    
    ```bash
    $ priceupdater change-provider --provider bitfinex
    ```
2. `update-config --provider <provider-name> --configFile <path-to-config.json>`
    
    After running `change-provider` the first time, you should load the provider's configuration
    <details><summary><b>Show example</b></summary>
    
    
    ```bash
    $ priceupdater update-config --provider bitfinex --configFile assets/bitfinex.json
    ```
###### Note: This action must be done for each provider we want to configure.

### Priority configuration (Requirement)
The next command allows configure the provider priority. The first provider will be chosen by default but if it fails during retrieving the price for a specific token, the service will try to get the price for this token using the next provider in the list.
Run: 
```
change-priority --priority <provider_1>,<provider_2>,<provider_3>
```
###### Note: <provider_1> has the highest priority and <provider_2> the lowest

### Project Settings

There is just one command for project settings:

1. `setup-origin --origins "items,comma,separated"`
    This command will set up hostnames that the system will accept as incoming requests. The default value is `*`, which means it accepts all requests. 

### Update Static Tokens

There is just one command for update static tokens:

1. `update-static-token --tokenID 1 --price 1.531`
    This command will update the price of given tokenID

### Server

There is just one command for running the server:

1. `server`

    This command will start the Price Updater server. It's possible to pass some configurations to the system.
    
    <details><summary><b>Show example</b></summary>
    
    ```bash
    $  priceupdater server --help

    Usage:
      price-updater server [flags]

    Flags:
          --eth-network string   ethereum address
          --fiat-key string      fiat api key
      -h, --help                 help for server
          --pg-dbname string     postgresql database name
          --pg-host string       postgresql host
          --pg-pass string       postgresql password
          --pg-port int          postgresql port (default 5432)
          --pg-user string       postgresql username
   ```


## Providers

There a list of implemented providers by Price Updater

- Bitfinex
- Coingecko
- Uniswap

### Provider Config

We have created a pattern using a JSON file to put the necessary configuration for running the providers on the system.

```json
{
    "provider": "providerName",
    "config": {
        "0": "0xETHEREUM MAINNET ADDRESS",
        "1": "0xETHEREUM MAINNET ADDRESS",
        "2": "SYMBOL"
    }
}
```

The config object means: The key is the `tokenID` from your hermez-node database (possibly in Goerli, Rinkeby or Mainnet), and the value can be the Address of the token on Mainnet Ethereum network or the symbol.


## API

There is a list of endpoints (and examples) that can be called on the Price Update Service. You can see these examples in the Postman Collection file at the project's root folder.
First of all, the API has a requirement for all requests, the user must set the **Origin** header even if the **Origin** is not registered on the project.

- _All examples have been run with [Httpie](https://httpie.io/)_


### Health endpoint

- Endpoint: /v1/health
- HTTP Method: GET
- Details: This will expose if the server is running.

```cmd
☁  ~  http https://priceupdater.hermez.io/v1/health Origin:my-origin
HTTP/1.1 200 OK
CF-Cache-Status: DYNAMIC
CF-RAY: 678fdf2f28c825f9-GIG
Connection: keep-alive
Content-Length: 7
Content-Type: text/plain; charset=utf-8
Date: Tue, 03 Aug 2021 13:25:26 GMT
Expect-CT: max-age=604800, report-uri="https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct"
Last-Modified: Tue, 03 Aug 2021 13:25:26 GMT
NEL: {"report_to":"cf-nel","max_age":604800}
Report-To: {"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report\/v3?s=NgEZxHGZ2mqgZn%2Fcyk%2FSwwjCNEhKp23K2rI5fjAYT5IwufaNr37oS5ZeSVZxIGq1uVQGCtzCIj1j3NJqjZO6O025FRrzxSlDGsMr%2Fdvkia30teNuQY%2Fny0%2F5j1gDI%2Fd9q%2FBFPkEYG4xowtc5d5yPrwkQqnQ%3D"}],"group":"cf-nel","max_age":604800}
Server: cloudflare

Healthy
```


### Tokens endpoint

- Endpoint: /v1/tokens
- HTTP Method: GET
- Details: Returns a list of tokens prices.

_We omitted some tokens in the response to make it more readable_

```cmd
☁  ~  http https://priceupdater.hermez.io/v1/tokens Origin:my-origin
HTTP/1.1 200 OK
CF-Cache-Status: DYNAMIC
CF-RAY: 678fe041ddd6f84b-GIG
Connection: keep-alive
Content-Encoding: gzip
Content-Type: application/json
Date: Tue, 03 Aug 2021 13:26:10 GMT
Expect-CT: max-age=604800, report-uri="https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct"
Last-Modified: Tue, 03 Aug 2021 13:26:10 GMT
NEL: {"report_to":"cf-nel","max_age":604800}
Report-To: {"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report\/v3?s=hsJzqKJPtCm%2B0SvhPgiLHry1X1PNnGwmaXYiL%2Bm9JJ9urh7OEs%2BlMBO0%2Fb5VY8YNJQnmbJDxXAigF9goLX%2FP37gZ92L8NUYPeMztczd%2FQUUlnB5jZ4YXmzaCD%2F5pLzQBGtQOVDYLu7z5mWlIoOUJhtIjo9Y%3D"}],"group":"cf-nel","max_age":604800}
Server: cloudflare
Transfer-Encoding: chunked

{
    "tokens": [
        {
            "USD": 21.044,
            "decimals": 18,
            "ethereumAddress": "0xba100000625a3754423978a60c9317c58a424e3D",
            "ethereumBlockNum": 12197953,
            "id": 11,
            "itemId": 12,
            "name": "Balancer",
            "symbol": "BAL",
            "usdUpdate": "2021-08-03T13:25:22.212488Z"
        },
        {
            "USD": 2508.8,
            "decimals": 18,
            "ethereumAddress": "0x0000000000000000000000000000000000000000",
            "ethereumBlockNum": 0,
            "id": 0,
            "itemId": 1,
            "name": "Ether",
            "symbol": "ETH",
            "usdUpdate": "2021-08-03T13:25:22.243868Z"
        },
        {
            "USD": 31737,
            "decimals": 18,
            "ethereumAddress": "0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e",
            "ethereumBlockNum": 12198043,
            "id": 13,
            "itemId": 14,
            "name": "yearn.finance",
            "symbol": "YFI",
            "usdUpdate": "2021-08-03T13:25:22.360085Z"
        },
        {
            "USD": 21.362,
            "decimals": 18,
            "ethereumAddress": "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
            "ethereumBlockNum": 12197927,
            "id": 8,
            "itemId": 9,
            "name": "Uniswap",
            "symbol": "UNI",
            "usdUpdate": "2021-08-03T13:25:22.373092Z"
        }
    ]
}

```

### Get specific Token

- Endpoint: /v1/tokens/<TOKEN_ID>
- HTTP Method: GET
- Details: Returns the price for a specific token.


```cmd
☁  ~  http https://priceupdater.hermez.io/v1/tokens/17 Origin:my-origin
HTTP/1.1 200 OK
CF-Cache-Status: DYNAMIC
CF-RAY: 678fe6ec6910275e-GIG
Connection: keep-alive
Content-Encoding: gzip
Content-Type: application/json
Date: Tue, 03 Aug 2021 13:30:43 GMT
Expect-CT: max-age=604800, report-uri="https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct"
Last-Modified: Tue, 03 Aug 2021 13:30:43 GMT
NEL: {"report_to":"cf-nel","max_age":604800}
Report-To: {"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report\/v3?s=T0ihq%2BtQhUqsVnXh9dTtB93Tm8S7MMYFCKOfa8tdN6wq9Sis8okHYcZhh06DgcSC3bS99qeHRCPR%2BOLyTXbX%2FZEQ%2F0PLtjS%2FQKTNq%2FjzBXNwN14CSf6Z2cXHujsribuD3Jaoln5RL%2BlpgAnLrj6GtqiDUIM%3D"}],"group":"cf-nel","max_age":604800}
Server: cloudflare
Transfer-Encoding: chunked

{
    "USD": 7.44,
    "decimals": 18,
    "ethereumAddress": "0xDe30da39c46104798bB5aA3fe8B9e0e1F348163F",
    "ethereumBlockNum": 12605976,
    "id": 17,
    "itemId": 18,
    "name": "Gitcoin",
    "symbol": "GTC",
    "usdUpdate": "2021-08-03T13:30:37.330874Z"
}
```

If you try to get a token that doesn't exists on Price Updater:


```cmd
☁  ~  http https://priceupdater.hermez.io/v1/tokens/171 Origin:my-origin
HTTP/1.1 404 Not Found
CF-Cache-Status: DYNAMIC
CF-RAY: 678fe7d70f2a2603-GIG
Connection: keep-alive
Content-Length: 29
Content-Type: application/json
Date: Tue, 03 Aug 2021 13:31:20 GMT
Expect-CT: max-age=604800, report-uri="https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct"
NEL: {"report_to":"cf-nel","max_age":604800}
Report-To: {"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report\/v3?s=W2oI3RMtIsukqqB%2FPZvpJBNijXlGsRNhivwtCfL8HIFwb2MBDNjtE%2BqGIevp08dbyhPAzIYmvMk4Wet5ag5XTmG85oIepqLk0uwsUcfV2XMsz9fZnBM4h88zC5URtZcfH7VqTvPywJgnB9n2FMf3xmnjP%2FM%3D"}],"group":"cf-nel","max_age":604800}
Server: cloudflare

{
    "message": "token not found"
}
```

### Currencies

- Endpoint: /v1/currencies
- HTTP Method: GET
- Details: Returns a list of currencies prices.

```cmd
☁  ~  http https://priceupdater.hermez.io/v1/currencies Origin:my-origin
HTTP/1.1 200 OK
CF-Cache-Status: DYNAMIC
CF-RAY: 678fe9531fa32599-GIG
Connection: keep-alive
Content-Encoding: gzip
Content-Type: application/json
Date: Tue, 03 Aug 2021 13:32:21 GMT
Expect-CT: max-age=604800, report-uri="https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct"
Last-Modified: Tue, 03 Aug 2021 13:32:21 GMT
NEL: {"report_to":"cf-nel","max_age":604800}
Report-To: {"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report\/v3?s=Z%2FkQ4PDo%2F6j7tMUzzdOsrS9qNpJlpuzalOy8zy%2BYzJwYplD0%2ByuCU8hFxJo1IuZq2TLPP2e7Q7CEUp8Extjjx3joGBCxKgpaYU5skIZSGxGmGCcpVodgVBi9yGEGKfdQUgA1aI2rjiwQL2OXcl8HQLi%2Byng%3D"}],"group":"cf-nel","max_age":604800}
Server: cloudflare
Transfer-Encoding: chunked

{
    "currencies": [
        {
            "baseCurrency": "USD",
            "currency": "CNY",
            "lastUpdate": "2021-08-03T13:32:05.317295Z",
            "price": 6.464898
        },
        {
            "baseCurrency": "USD",
            "currency": "EUR",
            "lastUpdate": "2021-08-03T13:32:05.32069Z",
            "price": 0.842115
        },
        {
            "baseCurrency": "USD",
            "currency": "JPY",
            "lastUpdate": "2021-08-03T13:32:05.323926Z",
            "price": 108.989498
        },
        {
            "baseCurrency": "USD",
            "currency": "GBP",
            "lastUpdate": "2021-08-03T13:32:05.327223Z",
            "price": 0.718935
        }
    ]
}
```

### Get specific Currency

- Endpoint: /v1/currencies/<CURRENCY>
- HTTP Method: GET
- Details: Returns the price for a specific currency.

```cmd
☁  ~  http https://priceupdater.hermez.io/v1/currencies/EUR Origin:my-origin
HTTP/1.1 200 OK
CF-Cache-Status: DYNAMIC
CF-RAY: 678fea8b9912275e-GIG
Connection: keep-alive
Content-Encoding: gzip
Content-Type: application/json
Date: Tue, 03 Aug 2021 13:33:11 GMT
Expect-CT: max-age=604800, report-uri="https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct"
Last-Modified: Tue, 03 Aug 2021 13:33:11 GMT
NEL: {"report_to":"cf-nel","max_age":604800}
Report-To: {"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report\/v3?s=NDfgP002wUZsZ6ijFLgrqJcJxIsiLzw7iyiiJhFxnOqeSFejn9cvgDkcEscUShMNH2shqTv1dEn2aLK%2FIQ%2BDyN7UGUT%2BBndt9jT2XzVPeMGX9Cz3U0ZAtq5dbRyAFAUm3j0ujqdaiPpw8C3t4hUNNzUHiK4%3D"}],"group":"cf-nel","max_age":604800}
Server: cloudflare
Transfer-Encoding: chunked

{
    "baseCurrency": "USD",
    "currency": "EUR",
    "lastUpdate": "2021-08-03T13:33:07.03667Z",
    "price": 0.842115
}
```

If you try to get a currency that doesn't exists on Price Updater:

```cmd
☁  ~  http https://priceupdater.hermez.io/v1/currencies/AOE Origin:my-origin
HTTP/1.1 404 Not Found
CF-Cache-Status: DYNAMIC
CF-RAY: 678feb8eec6b2616-GIG
Connection: keep-alive
Content-Length: 36
Content-Type: application/json
Date: Tue, 03 Aug 2021 13:33:52 GMT
Expect-CT: max-age=604800, report-uri="https://report-uri.cloudflare.com/cdn-cgi/beacon/expect-ct"
NEL: {"report_to":"cf-nel","max_age":604800}
Report-To: {"endpoints":[{"url":"https:\/\/a.nel.cloudflare.com\/report\/v3?s=vxorZ11iIOlGNJVQVPSL%2Fs89fqiktVPHV6svVX4W96IxlrlqchVt%2BuOPcVp3BApiAd3lt4QGUMWKT%2FE75g0X8bpPyhnMuFBgLJ1Do4ENsWG9cMXxisTMUb%2Bp76mfp138bPA3O5t1rPEEOweZzQeY0lpD60Q%3D"}],"group":"cf-nel","max_age":604800}
Server: cloudflare

{
    "message": "currency AOE not found"
}
```
