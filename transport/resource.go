package transport

import (
	"strconv"
	"strings"
)

type Resource struct {
	ID      string
	Level   int64
	Display string
}

func (r Resource) ToHeader() string {
	ret := r.ID
	if r.Level != 0 {
		ret += "." + strconv.FormatInt(r.Level, 10)
	}
	if r.Display != "" {
		ret += ":" + r.Display
	}
	return ret
}

func GetResources(from string) []Resource {
	var resources []Resource
	rawRes := strings.Split(from, ",")

	for _, res := range rawRes {
		// RawRes format = ID.Level:Display
		resource := Resource{}
		dis := strings.Split(res, ":")
		if len(dis) == 2 {
			resource.Display = dis[1]
			res = dis[0]
		}
		lvl := strings.Split(res, ".")
		if len(lvl) == 2 {
			resource.Level, _ = strconv.ParseInt(lvl[1], 10, 64)
		}
		resource.ID = lvl[0]
	}

	return resources
}
