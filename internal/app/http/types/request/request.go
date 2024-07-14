package request

type CreateDockerClientRequest struct {
	// 客户端名称
	Name string `json:"name"`
	// 客户端类型
	Type string `json:"type"`
	// 客户端地址
	IP string `json:"ip"`
	// 客户端端口
	Port int `json:"port"`
}

type TryConnectRequest struct {
	Mode string
}
