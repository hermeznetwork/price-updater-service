package command

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/ports"
	"github.com/hermeznetwork/price-updater-service/core/services"
)

type UpdateProviderConfigCommand struct {
	Provider           string
	LocationFileConfig string
	srv                *services.ConfigUpdaterService
}

type localProviderConfig struct {
	Provider string          `json:"provider"`
	Config   map[uint]string `json:"config"`
}

func NewUpdateProviderConfigCommand(provider, location string, srv *services.ConfigUpdaterService) ports.Command {
	return &UpdateProviderConfigCommand{
		Provider: provider, LocationFileConfig: location,
		srv: srv,
	}
}

func (cmd *UpdateProviderConfigCommand) Execute() error {
	lpc, err := cmd.loadJsonFile()
	if err != nil {
		return err
	}
	return cmd.srv.UpdateConfig(dbToDomain(lpc))
}

func (cmd *UpdateProviderConfigCommand) loadJsonFile() (*localProviderConfig, error) {
	var localConfig localProviderConfig
	configFile, err := os.Open(cmd.LocationFileConfig)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()
	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteValue, &localConfig)
	return &localConfig, err
}

func dbToDomain(lpc *localProviderConfig) *domain.PriceProvider {
	return &domain.PriceProvider{
		Provider:              lpc.Provider,
		MappingBetweenNetwork: lpc.Config,
	}
}
