package esp32healthcheck

import (
	"io/ioutil"
	"net/http"
	"time"
	"ueckoken/plarail2021-soft-internal/pkg/station2espIp"

	"github.com/prometheus/client_golang/prometheus"
)

type PingHandler struct {
	station2espIp.Stations
	Esp32HealthCheck *prometheus.GaugeVec
}

func (p *PingHandler) Start() {
	ticker := time.NewTicker(1 * time.Second)
	client := http.Client{}
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for _, s := range p.Stations.Enumerate() {
				resp, err := client.Get(s.Station.Address + "/health")
				if err != nil {
					p.Esp32HealthCheck.With(prometheus.Labels{"esp32Addr": s.Station.Name}).Set(0)
					ioutil.ReadAll(resp.Body)
					resp.Body.Close()
					continue
				}
				defer ioutil.ReadAll(resp.Body)
				defer resp.Body.Close()
				if 200 <= resp.StatusCode && resp.StatusCode < 300 {
					p.Esp32HealthCheck.With(prometheus.Labels{"esp32Addr": s.Station.Name}).Set(1)
					continue
				}
				p.Esp32HealthCheck.With(prometheus.Labels{"esp32Addr": s.Station.Name}).Set(0)
			}
		}
	}
}
