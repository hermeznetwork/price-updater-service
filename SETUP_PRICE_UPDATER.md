# Setup Price Updater

## Instalation

To install the system just clone and build it!
```bash=
$ git clone git@github.com:hermeznetwork/price-updater-service.git
$ cd price-updater-service/
$ go build -o priceupdater # or other name that you want
```

## Running as Linux Service

Before execute this steps, you should [configure](https://github.com/hermeznetwork/price-updater-service#configuration) the service.

### Quick Guide

```cmd=
# First, define the priority execution of providers
$ priceupdater change-priority --priority bifinex,uniswap

# Second, load the configuration to providers
$ priceupdater update-config --provider bitfinex --configFile pathToConfig.json

# Third, create a Key for API
$ priceupdater setup-apikey --apiKey Th3P0w3r4p1K3y
```

1. Copy service file to systemd folder:

```cmd=
$ sudo cp etc/priceupdater.service /etc/systemd/system/
```

2. Reload the services file to include the Price Updater service

```cmd=
$ sudo systemctl daemon-reload
```

3. Start the Price Updater service

```cmd=
$ sudo systemctl start priceupdater.service
```

4. To enable the Price Updater service on every reboot

```cmd=
$ sudo systemctl enable priceupdater.service
```

5. To disable the Price Updater service on every reboot

```cmd=
$ sudo systemctl disable priceupdater.service
```
