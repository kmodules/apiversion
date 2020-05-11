package apiversion

import (
	"fmt"
	"regexp"
	"strconv"
)

type Version struct {
	X int
	Y string
	Z int
}

type InvalidVersion struct {
	v string
}

func (e InvalidVersion) Error() string {
	return fmt.Sprintf("invalid version %s", e.v)
}

var (
	re = regexp.MustCompile(`^v(\d+)(alpha|beta|rc)?(\d*)$`)
)

func NewVersion(s string) (*Version, error) {
	groups := re.FindStringSubmatch(s)
	if len(groups) == 0 {
		return nil, InvalidVersion{v: s}
	}

	var out Version

	x, err := strconv.Atoi(groups[1])
	if err != nil {
		return nil, err
	}
	out.X = x

	out.Y = groups[2]

	if groups[3] != "" {
		z, err := strconv.Atoi(groups[3])
		if err != nil {
			return nil, err
		}
		out.Z = z
	}

	return &out, nil
}

func (v Version) Compare(other Version) int {
	diffX := v.X - other.X
	switch {
	case diffX > 0:
		return 1
	case diffX < 0:
		return -1
	}

	if v.Y != other.Y {
		if v.Y == "" {
			return 1
		} else if other.Y == "" {
			return -1
		} else if v.Y > other.Y {
			return 1
		} else {
			return -1
		}
	}

	diffZ := v.Z - other.Z
	switch {
	case diffZ > 0:
		return 1
	case diffZ < 0:
		return -1
	}
	return 0
}
