package versions

import (
	"regexp"
	"strconv"
	"strings"
)

type Version string

var versionRegex = regexp.MustCompile("^\\d+\\.\\d+\\.\\d+$")

func matchRegex(ver string) bool {
	return versionRegex.Match([]byte(ver))
}

func IsVersion(ver string) bool {
	return matchRegex(ver)
}

func IsVersionAndParse(ver string) (Version, bool) {
	if matchRegex(ver) {
		return Version(ver), true
	}
	return "", false
}

func (v Version) ToString() string {
	return string(v)
}

func (v Version) GetMainVersion() int {
	return v.GetVersions()[0]
}

func (v Version) GetSecondaryVersion() int {
	return v.GetVersions()[1]
}

func (v Version) GetTertiaryVersion() int {
	return v.GetVersions()[2]
}

func (v Version) GetVersions() []int {
	ints := make([]int, 3)
	if !matchRegex(string(v)) {
		return ints
	}
	strs := strings.Split(string(v), ".")
	ints[0], _ = strconv.Atoi(strs[0])
	ints[1], _ = strconv.Atoi(strs[1])
	ints[2], _ = strconv.Atoi(strs[2])
	return ints
}
