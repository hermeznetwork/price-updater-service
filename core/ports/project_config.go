package ports

type ProjectConfigRepository interface {
	SaveAllowedOrigin(origins string) error
	LoadAllowedOrigin() (string, error)
	SaveAPIKey(origins string) error
	LoadAPIKey() (string, error)
}
