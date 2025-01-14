package version

import (
	"strings"

	"github.com/blang/semver/v4"
)

const DefaultVersion = "0.0.0-unknown-version"

var (
	binaryVersion   string
	binarySHA       string
	binaryBuildDate string
)

func VersionString() string {
	// Remove the "v" prefix from the binary in case it is present
	binaryVersion = strings.TrimPrefix(binaryVersion, "v")
	versionString, err := semver.Make(binaryVersion)
	if err != nil {
		versionString = semver.MustParse(DefaultVersion)
	}

	metaData := []string{}
	if binarySHA != "" {
		metaData = append(metaData, binarySHA)
	}

	if binaryBuildDate != "" {
		metaData = append(metaData, binaryBuildDate)
	}

	versionString.Build = metaData

	return versionString.String()
}
