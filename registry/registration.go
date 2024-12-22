package registry

type ServiceName string

type Registration struct {
	ServiceName ServiceName
	ServiceUrl  string
}

// 已有服务列表

const (
	LogService = ServiceName("LogService")
)
