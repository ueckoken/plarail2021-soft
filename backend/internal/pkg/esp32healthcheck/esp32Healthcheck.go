package esp32healthcheck

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
	"ueckoken/plarail2021-soft-internal/pkg/station2espIp"
)

type PingHandler struct{
	station2espIp.Stations
	Esp32HealthCheck *prometheus.GaugeVec
}

func (p *PingHandler)Start(){
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <- ticker.C:
			for _, s := range p.Stations.Enumerate() {
				resp, err := http.Get(s.Station.Address)
				if err != nil {
					p.Esp32HealthCheck.With(prometheus.Labels{"esp32Addr": s.Station.Name}).Set(0)
					continue
				}
				if 200 <= resp.StatusCode && resp.StatusCode < 300{
					p.Esp32HealthCheck.With(prometheus.Labels{"esp32Addr": s.Station.Name}).Set(1)
					continue
				}
				p.Esp32HealthCheck.With(prometheus.Labels{"esp32Addr": s.Station.Name}).Set(0)
			}
		}
	}
}