package define

// TCurl 命令参数
type TCurl struct {
	Uri       string
	Times     int
	Intervals int
	TimeOut   int
	SaveDB    bool
}

// DBInfo 数据库细信息
type DBInfo struct {
	DbConnectUri string
}

// Show 子命令参数
type Show struct {
	ClientName string
}
