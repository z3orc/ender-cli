package endpoints

import (
	"encoding/json"
	"log"

	"github.com/z3orc/ender-cli/util"
)

type versions struct {
	Latest   map[string]string
	Versions []version
}

type version struct {
	Id string
}

func GetVersionsVanilla() versions{
	resp, err := util.GetJson("https://piston-meta.mojang.com/mc/game/version_manifest_v2.json")
	if err != nil {
		log.Fatal(err)
	}

	versions := versions{}

	err = json.Unmarshal(resp, &versions)
	if err != nil {
		log.Fatal(err)
	}
	
	return versions
}