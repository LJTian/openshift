package echarts

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"net/http"
)

var Times []float64

// GenerateLineItems 图表元素生成
func GenerateLineItems() []opts.LineData {

	items := make([]opts.LineData, 0)
	// 生成一些示例数据
	for _, v := range Times {
		items = append(items, opts.LineData{Value: v})
	}
	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "时间曲线图",
			Subtitle: "时间曲线图",
		}))

	// Put data into instance
	x := make([]int, 0)
	for i := 1; i <= len(Times)/5; i++ {
		x = append(x, i*5)
	}
	line.SetXAxis(x).
		AddSeries("元素1 ", GenerateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func ShowWeb(times []float64) {

	Times = times
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)

}
