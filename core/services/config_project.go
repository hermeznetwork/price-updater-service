package services

import (
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type ProjectConfigService struct {
	pcr ports.ProjectConfigRepository
}

func NewProjectConfigServices(pcr ports.ProjectConfigRepository) *ProjectConfigService {
	return &ProjectConfigService{
		pcr: pcr,
	}
}

func (pc *ProjectConfigService) LoadOriginValue() (string, error) {
	return pc.pcr.LoadAllowedOrigin()
}

func (pc *ProjectConfigService) SaveOriginValue(origin string) error {
	return pc.pcr.SaveAllowedOrigin(origin)
}

func (pc *ProjectConfigService) LoadAPIKey() (string, error) {
	return pc.pcr.LoadAPIKey()
}

func (pc *ProjectConfigService) SaveAPIKey(apiKey string) error {
	return pc.pcr.SaveAPIKey(apiKey)
}
