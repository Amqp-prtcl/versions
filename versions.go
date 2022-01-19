package versions

import (
	"regexp"
	"strconv"
	"strings"
)

/*
MAJOR.MINOR.PATCH
MAJOR version when you make incompatible API changes,
MINOR version when you add functionality in a backwards compatible manner, and
PATCH version when you make backwards compatible bug fixes.
*/
type Version string

var versionRegex = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

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

func (v Version) GetMajor() int {
	return v.GetSubInt()[0]
}

func (v Version) GetMinor() int {
	return v.GetSubInt()[1]
}

func (v Version) GetPatch() int {
	return v.GetSubInt()[2]
}

func (v Version) GetSubInt() []int {
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

func (v Version) IsCompatible(vrs Version) bool {
	return v.GetMajor() == vrs.GetMajor()
}

func (v Version) IsNewerThan(vrs Version) bool {
	b := vrs.GetSubInt()
	a := v.GetSubInt()
	if a[0] > b[0] {
		return true
	} else if a[0] < b[0] {
		return false
	} else {
		if a[1] > b[1] {
			return true
		} else if a[1] < b[1] {
			return false
		} else {
			return a[2] > b[2]
		}
	}
}

type Versions []Version

func (vrs Versions) GetLastVersion() Version {
	var last Version = Version("0.0.0")
	for _, v := range vrs {
		if v.IsNewerThan(last) {
			last = v
		}
	}
	return last
}

func (vrs Versions) IsLast(v Version) bool {
	return vrs.GetLastVersion() == v
}
