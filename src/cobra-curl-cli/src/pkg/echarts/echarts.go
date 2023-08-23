package echarts

import (
	"cobra-curl-cli/pkg/db"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"net/http"
)

var Times []float64
var clientName string

// GenerateLineItems 图表元素生成
func GenerateLineItems(iNum int) []opts.LineData {

	name := fmt.Sprintf("%s_%d", clientName, iNum)
	timeline, err := db.ShowTimeLineLogsByClientName(name)
	if err != nil {
		fmt.Println(err)
	}

	Times = timeline
	items := make([]opts.LineData, 0)
	// 生成一些示例数据
	for _, v := range timeline {
		items = append(items, opts.LineData{Value: v})
	}
	return items
}

// GenerateScatterItems 散点图元素生成
func GenerateScatterItems(iNum int) []opts.ScatterData {

	name := fmt.Sprintf("%s_%d", clientName, iNum)
	timeline, err := db.ShowTimeLineLogsByClientName(name)
	if err != nil {
		fmt.Println(err)
	}

	Times = timeline

	items := make([]opts.ScatterData, 0)
	// 生成一些示例数据
	for _, v := range timeline {
		items = append(items, opts.ScatterData{Value: v,
			Symbol:       "roundRect",
			SymbolSize:   5,
			SymbolRotate: 5,
		})
	}
	return items
}

func lineBase() *charts.Line {
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
	for i := 1; i <= len(Times); i++ {
		x = append(x, i)
	}

	coroutineNum, err := db.GetCoroutineNumByClientName(clientName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//fmt.Println(coroutineNum)
	//fmt.Println(len(coroutineNum))

	line = line.SetXAxis(x)
	for i := 0; i < len(coroutineNum); i++ {
		sName := fmt.Sprintf("协程 %d", i)
		line = line.AddSeries(sName, GenerateLineItems(i))
	}
	line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	return line
}

func scatterBase() *charts.Scatter {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(charts.WithTitleOpts(
		opts.Title{
			Title: "时间散点图",
		}),
	)

	// Put data into instance
	x := make([]int, 0)
	for i := 1; i <= len(Times); i++ {
		x = append(x, i)
	}

	coroutineNum, err := db.GetCoroutineNumByClientName(clientName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//fmt.Println(coroutineNum)
	//fmt.Println(len(coroutineNum))

	scatter = scatter.SetXAxis(x)
	for i := 0; i < len(coroutineNum); i++ {
		sName := fmt.Sprintf("协程 %d", i)
		scatter = scatter.AddSeries(sName, GenerateScatterItems(i))
	}

	return scatter
}

func scatterShowLabel() *charts.Scatter {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(charts.WithTitleOpts(
		opts.Title{
			Title: "时间散点图",
		}),
	)

	// Put data into instance
	x := make([]int, 0)
	for i := 1; i <= len(Times); i++ {
		x = append(x, i)
	}

	coroutineNum, err := db.GetCoroutineNumByClientName(clientName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//fmt.Println(coroutineNum)
	//fmt.Println(len(coroutineNum))

	scatter = scatter.SetXAxis(x)
	for i := 0; i < len(coroutineNum); i++ {
		sName := fmt.Sprintf("协程 %d", i)
		scatter = scatter.AddSeries(sName, GenerateScatterItems(i))
	}
	scatter.SetSeriesOptions(charts.WithLabelOpts(
		opts.Label{
			Show:     true,
			Position: "right",
		}),
	)

	return scatter
}

func show(w http.ResponseWriter, _ *http.Request) {

	page := components.NewPage()
	page.AddCharts(
		lineBase(),
		scatterBase(),
		scatterShowLabel(),
	)

	page.Render(w)
}

func ShowWeb(name string) {

	clientName = name
	http.HandleFunc("/", show)

	fmt.Println("请访问【http://127.0.0.1:8081/】查看结果~")

	http.ListenAndServe(":8081", nil)
}
