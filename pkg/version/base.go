package version

var (
	gitTag       = ""
	gitCommit    = "$Format:%H$"    // sha1 from git . output of $(git rev-parse HEAD)
	gitTreeState = "not a git tree" // state of git tree, either "clean" or  "dirty"
	buildDate    = "1970-01-01T00:00:00Z"
)