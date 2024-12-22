package registry

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

const (
	ServerPort = ":3000"
)

// registry 用于存储服务信息集合的结构
type registry struct {
	registrations []Registration
	mutex         *sync.Mutex
}

// add 注册服务
func (r *registry) add(registration Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, registration)
	r.mutex.Unlock()
	return nil
}

var reg registry

type RegistryService struct{}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		rg := Registration{}
		dc := json.NewDecoder(r.Body)
		if err := dc.Decode(&rg); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			log.Printf("Decode registration failed, %s", err)
		}
		log.Printf("Adding service: %v with URL: %s\n", rg.ServiceName, rg.ServiceUrl)
		if err := reg.add(rg); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			log.Printf("Add registration failed, %s", err)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
