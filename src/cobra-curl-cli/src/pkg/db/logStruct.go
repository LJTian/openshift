package db

type Logs struct {
	NodeName       string `json:"node name" gorm:"node_name"`
	Message        string `json:"message" gorm:"message"`
	ProcessingTime string `json:"processingTime" gorm:"processing_time"`
	ServerTime     string `json:"serverTime" gorm:"server_time"`
	Code           string `json:"code" gorm:"code"`
}
