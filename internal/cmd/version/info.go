package version

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

// Info creates a Version struct for output
type Info struct {
	Version     string `json:"Version,omitempty" yaml:"Version,omitempty"`
	Commit      string `json:"Commit,omitempty" yaml:"Commit,omitempty"`
	BuildDate   string `json:"BuildDate,omitempty" yaml:"BuildDate,omitempty"`
	GoVersion   string `json:"GoVersion,omitempty" yaml:"GoVersion,omitempty"`
	OS          string `json:"OS/Arch,omitempty" yaml:"OS/Arch,omitempty"`
	ReleaseDate string `json:"Date,omitempty" yaml:"Date,omitempty"`
}

// ToJSON converts the Info into a JSON String
func (v *Info) toJSON() string {
	bytes, _ := json.Marshal(v)
	return string(bytes) + "\n"
}

// ToYAML converts the Info into a JSON String
func (v *Info) toYAML() string {
	bytes, _ := yaml.Marshal(v)
	return string(bytes)
}
