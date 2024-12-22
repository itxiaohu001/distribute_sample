package log

import (
	"io"
	stdlog "log"
	"net/http"
	"os"
)

// 提供一个对象，用来将数据写入文件中

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_APPEND|os.O_RDONLY, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.Write(data)
}

var log *stdlog.Logger

// Run 初始化log
func Run(destination string) {
	log = stdlog.New(fileLog(destination), "go: ", stdlog.LstdFlags)
}

// RegisterHandler 注册服务
func RegisterHandler() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			data, err := io.ReadAll(r.Body)
			if err != nil || len(data) == 0 {
				_, _ = w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(data))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}
