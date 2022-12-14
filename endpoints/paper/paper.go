package paper

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/z3orc/ender-cli/util"
)

type Versions struct {
	Versions []string
}

type Version struct {
	Builds []int
}


func GetVersions() Versions {
	resp, err := util.GetJson("https://api.papermc.io/v2/projects/paper")
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

func GetVersion(version string) Version{
	resp, err := util.GetJson("https://api.papermc.io/v2/projects/paper/versions/" + version)
	if err != nil {
		log.Fatal(err)
	}

	builds := Version{}

	err = json.Unmarshal(resp, &builds)
	if err != nil {
		log.Fatal(err)
	}

	return builds
}

func GetLatestBuild(currentVersion string) string{
	version := GetVersion(currentVersion)
	builds := version.Builds

	latest := builds[len(builds) - 1]
	latestAsString := fmt.Sprintf("%v", latest)

	return latestAsString
}