package response

type CreateDockerClientResp struct {
}

type GetLocalDockerConnectionResp struct {
	EngineVersion string `json:"engineVersion"`
	Arch          string `json:"arch"`
	Os            string `json:"os"`
	KernelVersion string `json:"kernelVersion"`
	Platform      string `json:"platform"`
}
