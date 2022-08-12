package paper

import (
	"encoding/json"
	"log"

	"github.com/z3orc/ender-cli/util"
)

type Versions struct {
	Versions []string
}


func GetVersions() Versions {
	resp, err := util.GetJson("https://papermc.io/api/v2/projects/paper")
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