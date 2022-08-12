package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/z3orc/ender-cli/global"
)

// type Config struct {
//     Origin string
// }

func Get(config_path string, key string) string{
	content, err := os.ReadFile(config_path)
	if err != nil {
        log.Fatal("Error when opening file: ", err)
    }

	var payload map[string]string
	err = json.Unmarshal(content, &payload)
	if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }

	return payload[key]
}

func Set(config_path string, key string, value string) error{
	content, err := os.ReadFile(config_path)
	if err != nil {
        log.Fatal("Error when opening file: ", err)
    }

	var payload map[string]string
	err = json.Unmarshal(content, &payload)
	if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }

	payload[key] = value

	json_output, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = os.WriteFile(global.CONFIG_ENDER_PATH, json_output, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Create(path string, config map[string]string) error{
	json_output, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, json_output, 0644)
	if err != nil {
		return err
	}
	
	return nil
}