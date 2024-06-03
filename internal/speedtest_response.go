package internal

import "time"

type SpeedtestResponse struct {
	ID         int64
	Timestamp  time.Time `json:"timestamp"`
	ISP        string    `json:"isp"`
	PackatLoss string    `json:"packatLoss"`
	Ping       struct {
		Latency float32 `json:"latency"`
		Jitter  float32 `json:"jitter"`
	}
	Interface struct {
		ExternalIp  string `json:"externalIp"`
		ContainerIp string `json:"internalIp"`
		IsVPN       bool   `json:"isVpn"`
	}
	Server struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Country  string `json:"country"`
	}
	Download struct {
		Bandwidth float64 `json:"bandwidth"`
		Bytes     float64 `json:"bytes"`
		Latency   struct {
			Ping   float32 `json:"iqm"`
			Jitter float32 `json:"jitter"`
		}
	}
	Upload struct {
		Bandwidth float64 `json:"bandwidth"`
		Bytes     float64 `json:"bytes"`
		Latency   struct {
			Ping   float32 `json:"iqm"`
			Jitter float32 `json:"jitter"`
		}
	}
	Result struct {
		SpeedtestResponseUrl string `json:"url"`
		ID                   string `json:"id"`
	}
}
