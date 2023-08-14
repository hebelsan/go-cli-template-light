package version

import "strings"

type Output int

const (
	JSON Output = iota
	YAML
)

var (
	capabilitiesMap = map[string]Output{
		"json": JSON,
		"yaml": YAML,
		"yml":  YAML,
	}
)

func parseOutputType(str string) (Output, bool) {
	c, ok := capabilitiesMap[strings.ToLower(str)]
	return c, ok
}
