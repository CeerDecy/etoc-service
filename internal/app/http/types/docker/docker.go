package docker_type

import (
	"encoding/json"
	"time"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	ApiVersionKey    = "ApiVersion"
	MinAPIVersionKey = "MinAPIVersion"
	GoVersionKey     = "GoVersion"
	ArchKey          = "Arch"
	KernelVersionKey = "KernelVersion"
	VersionKey       = "Version"
	GitCommitKey     = "GitCommit"
	OsKey            = "Os"
	BuildTimeKey     = "BuildTime"
	PlatformKey      = "Platform"
	ComponentsKey    = "Components"
)

// VersionInfo 表示 Docker 引擎的详细信息。
type VersionInfo struct {
	ApiVersion    string      `json:"ApiVersion"`
	MinAPIVersion string      `json:"MinAPIVersion"`
	GoVersion     string      `json:"GoVersion"`
	Arch          string      `json:"Arch"`
	KernelVersion string      `json:"KernelVersion"`
	BuildTime     time.Time   `json:"BuildTime"`
	Platform      Platform    `json:"Platform"`
	Components    []Component `json:"Components"`
	Version       string      `json:"Version"`
	GitCommit     string      `json:"GitCommit"`
	Os            string      `json:"Os"`
}

func ConvertEnvToVersionInfo(env *docker.Env) (*VersionInfo, error) {
	var vi VersionInfo
	vi.ApiVersion = env.Get(ApiVersionKey)
	vi.MinAPIVersion = env.Get(MinAPIVersionKey)
	vi.GoVersion = env.Get(GoVersionKey)
	vi.Arch = env.Get(ArchKey)
	vi.KernelVersion = env.Get(KernelVersionKey)
	vi.Version = env.Get(VersionKey)
	vi.GitCommit = env.Get(GitCommitKey)
	vi.Os = env.Get(OsKey)

	bt, err := time.Parse(time.RFC3339Nano, env.Get(BuildTimeKey))
	if err != nil {
		logrus.Error(err)
	}
	vi.BuildTime = bt
	err = json.Unmarshal([]byte(env.Get(PlatformKey)), &vi.Platform)
	if err != nil {
		return nil, errors.Errorf("failed to unmarshal Platform: %v", err)
	}
	err = json.Unmarshal([]byte(env.Get(ComponentsKey)), &vi.Components)
	if err != nil {
		return nil, errors.Errorf("failed to unmarshal Components: %v", err)
	}
	return &vi, nil
}

// Platform 表示 Docker 引擎的平台信息。
type Platform struct {
	Name string `json:"Name"`
}

// Component 表示 Docker 引擎的组件信息。
type Component struct {
	Name    string           `json:"Name"`
	Version string           `json:"Version"`
	Details ComponentDetails `json:"Details"`
}

// ComponentDetails 表示 Docker 引擎组件的详细配置信息。
type ComponentDetails struct {
	ApiVersion    string    `json:"ApiVersion"`
	Arch          string    `json:"Arch"`
	BuildTime     time.Time `json:"BuildTime"`
	Experimental  string    `json:"Experimental"`
	GitCommit     string    `json:"GitCommit"`
	GoVersion     string    `json:"GoVersion"`
	KernelVersion string    `json:"KernelVersion"`
	MinAPIVersion string    `json:"MinAPIVersion"`
	Os            string    `json:"Os"`
}
