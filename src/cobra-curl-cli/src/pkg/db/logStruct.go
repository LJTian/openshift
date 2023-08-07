package db

type Logs struct {
	ClientName     string  `json:"clientName" gorm:"client_name:index"`
	NodeName       string  `json:"node name" gorm:"node_name"`
	Message        string  `json:"message" gorm:"message"`
	ProcessingTime string  `json:"processingTime" gorm:"processing_time"`
	ServerTime     string  `json:"serverTime" gorm:"server_time"`
	Code           int     `json:"code" gorm:"code"`
	TimeSinceLast  float64 `json:"time since last" gorm:"time_since_last"`
}
