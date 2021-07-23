# Price Updater Service

Price Updater is an web service to consult token and currency prices used to Hermez Node.

## How It Works

1. The system was built with Golang and designed to running through [commands](##usage).
2. Given a time to update (time set in through configuration, or using default value of 30s), the system update the prices of currency and tokens.
3. The prices of tokens and currencies are returned through an endpoint.

## Instalation

To install system just clone and build it!

```bash=
$ git clone git@github.com:hermeznetwork/price-updater-service.git
$ cd price-updater-service/
$ cd go build -o priceupdater # or other name that you want
```

## Usage 

The system has designed to running through commands listed above


### Configuration

You can setup the configuration to system on environment variables or an .env file

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

### Providers configuration

There are two commands related to providers:

1. `change-provider --provider <provider-name>`
    
    This command must be executed before the other commands. It tell's to system what provider must be executed.
    
    Also, he is used to change the provider in runtime.
    <details><summary><b>Show example</b></summary>
    
    ```bash
    $ priceupdater change-provider --provider bitfinex
    ```
2. `update-config --provider <provider-name> --configFile <path-to-config.json>`
    
    After running `change-provider` first time, you should load the provider's configuration
    <details><summary><b>Show example</b></summary>
    
    
    ```bash
    $ priceupdater update-config --provider bitfinex --configFile assets/bitfinex.json
    ```


### Project Settings

There is just one command to project settings:

1. `setup-origin --origins "items,comma,separated"`
    This command will setup hostnames that the system will accept incoming requests. The default value is `*`, that's means accept everyone.

### Server

There is just one command to running the server:

1. `server`

    This command will start the server of Price Updater. It's possible pass some configuration to system.
    
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

We had create a pattern using a JSON file to put the necessary configuration to running the providers on system.

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

The config object means: The key is the `tokenID` from database of your hermez-node (maybe in goerli, rinkeby or mainnet), and the value can be the Address of token on Mainnet Ethereum network or the symbol.
