package purpur

import (
	"encoding/json"
	"log"

	"github.com/z3orc/ender-cli/util"
)

type Versions struct {
	Versions []string
}

type Version struct {
	Builds Builds
	Project string
	Version string
}

type Builds struct {
	All []string
	Latest string
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

func GetVersion(version string) Version{
	resp, err := util.GetJson("https://api.purpurmc.org/v2/purpur/" + version)
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

	return version.Builds.Latest
}