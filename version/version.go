package version

import (
	"fmt"
	"runtime"
)

var (
	Commit        = ""
	Version       = ""
	VendorDirHash = ""
	BuildTags     = ""
)

type versionInfo struct {
	RelayerApp    string `json:"relayerApp"`
	GitCommit     string `json:"commit"`
	VendorDirHash string `json:"vendorHash"`
	BuildTags     string `json:"buildTags"`
	GoVersion     string `json:"go"`
}

func (v versionInfo) String() string {
	return fmt.Sprintf(`relayer: %s
git commit: %s
vendor hash: %s
build tags: %s
%s`, v.RelayerApp, v.GitCommit, v.VendorDirHash, v.BuildTags, v.GoVersion)
}

func newVersionInfo() versionInfo {
	return versionInfo{
		RelayerApp:    Version,
		GitCommit:     Commit,
		VendorDirHash: VendorDirHash,
		BuildTags:     BuildTags,
		GoVersion:     fmt.Sprintf("go version %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH),
	}
}
