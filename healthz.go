package main
import (
	"net/http"
	"encoding/json"
	"runtime"
	"os"
)

type HealthZ struct {
	Version   string `json:"version"`
	BuildDate string `json:"build_date"`
	BuildHost string `json:"build_host"`
	GoVersion string `json:"go_version"`
}

func (s *HealthZ) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		// what should we dooooOooo
		// w.Write([]byte("{}"))
		return
	}
	healthz := HealthZ{
		version,
		buildDate,
		hostname,
		runtime.Version(),
	}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(healthz)
    w.Write(response)
}

