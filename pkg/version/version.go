package version

import (
	"fmt"
	"runtime"
)

// Info contains versioning information
type Info struct {
	GitTag       string `json:"gitTag"`
	GitComment   string `json:"gitComment"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (info Info) String() string {
	return info.GitTag
}

func Get() Info {
	return Info{
		GitTag:       gitTag,
		GitComment:   gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
