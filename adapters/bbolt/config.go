package bbolt

import "github.com/hermeznetwork/price-updater-service/core/ports"

type ProjectConfigRepository struct {
	db *Connection
}

type projectConfigModels struct {
	Origins string `json:"origins"`
}

func NewProjectConfigRepository(db *Connection) ports.ProjectConfigRepository {
	return &ProjectConfigRepository{
		db: db,
	}

}

func (pc *ProjectConfigRepository) SaveAllowedOrigin(origins string) error {
	pc.db.Start()
	defer pc.db.End()

	tx, err := pc.db.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bkt, err := tx.CreateBucketIfNotExists([]byte("projectConfig"))
	if err != nil {
		return err
	}

	err = bkt.Put([]byte("origins"), []byte(origins))
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (pc *ProjectConfigRepository) LoadAllowedOrigin() (string, error) {
	pc.db.Start()
	defer pc.db.End()
	tx, err := pc.db.db.Begin(true)
	if err != nil {
		return "", nil
	}
	defer tx.Rollback()
	bkt, err := tx.CreateBucketIfNotExists([]byte("projectConfig"))
	if err != nil {
		return "", err
	}

	originsRaw := bkt.Get([]byte("origins"))
	if originsRaw == nil {
		return "*", tx.Commit()
	}
	return string(originsRaw), tx.Commit()
}

func (pc *ProjectConfigRepository) SaveAPIKey(apiKey string) error {
	pc.db.Start()
	defer pc.db.End()

	tx, err := pc.db.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bkt, err := tx.CreateBucketIfNotExists([]byte("projectConfig"))
	if err != nil {
		return err
	}

	err = bkt.Put([]byte("apiKey"), []byte(apiKey))
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (pc *ProjectConfigRepository) LoadAPIKey() (string, error) {
	pc.db.Start()
	defer pc.db.End()
	tx, err := pc.db.db.Begin(true)
	if err != nil {
		return "", nil
	}
	defer tx.Rollback()
	bkt, err := tx.CreateBucketIfNotExists([]byte("projectConfig"))
	if err != nil {
		return "", err
	}

	apiKey := bkt.Get([]byte("apiKey"))
	if apiKey == nil {
		return "pr1c3upd4t3r", tx.Commit()
	}
	return string(apiKey), tx.Commit()
}
