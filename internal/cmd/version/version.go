package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// filled by goreleaser
var (
	releaseVersion = "dev"
	releaseCommit  = ""
	releaseDate    = ""
)

// GenVersionString for printing in version command
func GenVersionString(printShort bool, outputType string) string {
	info := Info{
		Version:     releaseVersion,
		Commit:      commit(),
		GoVersion:   goVersion(),
		OS:          osArch(),
		ReleaseDate: releaseDate,
	}
	if printShort {
		info = Info{Version: info.Version}
	}

	oType, _ := parseOutputType(outputType)
	switch oType {
	case YAML:
		return info.toYAML()
	case JSON:
		return info.toJSON()
	default: // JSON as the default
		return info.toJSON()
	}
}

// commit reads the vcs revision from build info
func commit() string {
	if releaseCommit != "" {
		return releaseCommit
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}
	return ""
}

// goVersion returns the version of the go runtime used to compile the binary
func goVersion() string {
	return runtime.Version()
}

// osArch returns the os and arch used to build the binary
func osArch() string {
	return fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
}
