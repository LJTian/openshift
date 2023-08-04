package echarts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
)

func generateLineItems() []opts.LineData {
	// 生成一些示例数据
	data := []opts.LineData{
		{Value: 20},
		{Value: 50},
		{Value: 80},
		{Value: 70},
		{Value: 60},
		{Value: 30},
		{Value: 10},
	}
	return data
}

//
//func ShowWeb(data []opts.LineData) error {
//	// 创建一个折线图实例
//	line := charts.NewLine()
//
//	// 设置图表的标题和数据
//	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
//		Title: "数据折线展示",
//	}))
//
//	line.AddSeries("数据折线", data)
//
//	// 创建一个简单的 HTTP 服务器并将图表输出到浏览器
//	http.Handle("/line", line)
//	http.Handle("/", http.FileServer(http.Dir(".")))
//	port := 8080
//	addr := fmt.Sprintf(":%d", port)
//	log.Printf("请在浏览器中访问 http://localhost:%d/line 查看数据折线图", port)
//	log.Fatal(http.ListenAndServe(addr, nil))
//}
