package messenger

import (
	"net/http"
	"sync"
)

var (
	httpMu     sync.RWMutex
	HttpClient *http.Client
)

func init() {
	HttpClient = &http.Client{}
}

func SetHttpClient(client *http.Client) {
	httpMu.Lock()
	HttpClient = client
	httpMu.Unlock()
}

func GetHttpClient() (client *http.Client) {
	httpMu.RLock()
	client = HttpClient
	httpMu.RUnlock()
	return
}
