package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// RegisterService 通过请求服务注册服务注册当前服务
func RegisterService(r Registration) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	resp, err := http.Post(ServerAddress, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. responed with code %v", resp.StatusCode)
	}

	return nil
}
