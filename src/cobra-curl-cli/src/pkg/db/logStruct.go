package db

type Logs struct {
	ClientName     string  `json:"clientName" gorm:"index;column:client_name"`
	NodeName       string  `json:"node name" gorm:"column:node_name"`
	Message        string  `json:"message" gorm:"column:message"`
	ProcessingTime string  `json:"processingTime" gorm:"column:processing_time"`
	ServerTime     string  `json:"serverTime" gorm:"column:server_time"`
	Code           int     `json:"code" gorm:"column:code"`
	TimeSinceLast  float64 `json:"time since last" gorm:"column:time_since_last"`
}
