package backup

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
)

type BackupRegistry struct {
	Latest  uuid.UUID
	Backups map[string]*Backup
}

func LoadRegistry(path string) (*BackupRegistry, error) {
	br := &BackupRegistry{}

	if _, err := os.Stat(BACKUP_DIR + "/backups.json"); err == nil {
		file, err := os.Open(BACKUP_DIR + "/backups.json")
		if err != nil {
			return nil, fmt.Errorf("could not open backups.json. %s", err)
		}
		bytes, err := io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("reading backups.json failed. %s", err)
		}

		json.Unmarshal(bytes, &br)
		return br, nil
	}

	return br, nil
}

func (br *BackupRegistry) Write() error {
	content, err := json.Marshal(br)
	if err != nil {
		return fmt.Errorf("marshalling json failed. %s", err)
	}

	err = os.WriteFile(BACKUP_DIR+"/backups.json", content, 0644)
	if err != nil {
		return fmt.Errorf("writing to file failed. %s", err)
	}

	return nil
}

func (br *BackupRegistry) Load() error {
	if _, err := os.Stat(BACKUP_DIR + "/backups.json"); err == nil {
		file, err := os.Open(BACKUP_DIR + "/backups.json")
		if err != nil {
			return fmt.Errorf("could not open backups.json. %s", err)
		}
		bytes, err := io.ReadAll(file)
		if err != nil {
			return fmt.Errorf("reading backups.json failed. %s", err)
		}

		json.Unmarshal(bytes, &br)
		return nil
	}

	return nil
}

func (br *BackupRegistry) Add(newBackup Backup) {

}

func (br *BackupRegistry) Get(uuid uuid.UUID) {

}

func (br *BackupRegistry) Remove(uuid uuid.UUID) {

}

func (br *BackupRegistry) PurgeAll() {

}

func (br *BackupRegistry) PurgeOld() {

}

func ReadOverview() (map[string]*Backup, error) {
	var backups map[string]*Backup

	if _, err := os.Stat(BACKUP_DIR + "/backups.json"); err == nil {
		file, err := os.Open(BACKUP_DIR + "/backups.json")
		if err != nil {
			return nil, fmt.Errorf("could not open backups.json. %s", err)
		}
		bytes, err := io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("reading backups.json failed. %s", err)
		}

		json.Unmarshal(bytes, &backups)
		return backups, nil
	}

	return make(map[string]*Backup), nil
}

func WriteOverview(backups map[string]*Backup) error {
	content, err := json.Marshal(backups)
	if err != nil {
		return fmt.Errorf("marshalling json failed. %s", err)
	}

	err = os.WriteFile(BACKUP_DIR+"/backups.json", content, 0644)
	if err != nil {
		return fmt.Errorf("writing to file failed. %s", err)
	}

	return nil
}
