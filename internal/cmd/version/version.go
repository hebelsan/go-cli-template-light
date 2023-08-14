package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// GoVersion returns the version of the go runtime used to compile the binary
var goVersion = runtime.Version()

// OsArch returns the os and arch used to build the binary
var osArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)

// GenVersionString for printing in version command
func GenVersionString(printShort bool, outputType string) string {
	info := Info{
		Version:   getVersion(),
		Commit:    getCommit(),
		GoVersion: goVersion,
		OS:        osArch,
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

func getVersion() (version string) {
	version = "v0.0.0"
	info, ok := debug.ReadBuildInfo()
	if ok && info.Main.Version != "" {
		version = info.Main.Version
	}
	return
}

func getCommit() (commit string) {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}
	return
}
