package backup

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadOverview() (map[string]*Backup, error) {
	var backups map[string]*Backup

	if _, err := os.Stat("./testing/backups/backups.json"); err == nil {
		file, err := os.Open("./testing/backups/backups.json")
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

	err = os.WriteFile("./testing/backups/backups.json", content, 0644)
	if err != nil {
		return fmt.Errorf("writing to file failed. %s", err)
	}

	return nil
}
