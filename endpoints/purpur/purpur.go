package purpur

import (
	"encoding/json"
	"log"

	"github.com/z3orc/ender-cli/util"
)

type Versions struct {
	Versions []string
}


func GetVersions() Versions {
	resp, err := util.GetJson("https://api.purpurmc.org/v2/purpur")
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