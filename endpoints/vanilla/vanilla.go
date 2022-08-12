package vanilla

import (
	"encoding/json"
	"log"

	"github.com/z3orc/ender-cli/util"
)

type Versions struct {
	Latest   map[string]string
	Versions []Version
}

type Version struct {
	Id string
	Type string
}

func GetVersions() Versions{
	resp, err := util.GetJson("https://piston-meta.mojang.com/mc/game/version_manifest_v2.json")
	if err != nil {
		log.Fatal(err)
	}

	versions := Versions{}

	err = json.Unmarshal(resp, &versions)
	if err != nil {
		log.Fatal(err)
	}
	
	return versions
}