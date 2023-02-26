// package main

// import (
// 	"os"

// 	"github.com/go-echarts/go-echarts/v2/charts"
// 	"github.com/go-echarts/go-echarts/v2/opts"
// )


// func genScatter3dData(points Point[]) []opts.Chart3DData {
// 	var data []opts.Chart3DData
// 	for _, point := range points {
// 		data = append(data, opts.Chart3DData{Value: []interface{}{point.x, point.y, 0}})
// 	}
// 	return data
// }

// func scatter3DBase() *charts.Scatter3D {
// 	scatter3d := charts.NewScatter3D()
// 	scatter3d.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{Title: "basic Scatter3D example"}),
// 	)

// 	scatter3d.AddSeries("scatter3d", genScatter3dData())
// 	return scatter3d
// }

// func scatter3DDataItem() *charts.Scatter3D {
// 	scatter3d := charts.NewScatter3D()
// 	scatter3d.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{Title: "user-defined item style"}),
// 		charts.WithXAxis3DOpts(opts.XAxis3D{Name: "MY-X-AXIS", Show: true}),
// 		charts.WithYAxis3DOpts(opts.YAxis3D{Name: "MY-Y-AXIS"}),
// 		charts.WithZAxis3DOpts(opts.ZAxis3D{Name: "MY-Z-AXIS"}),
// 	)

// 	scatter3d.AddSeries("scatter3d", []opts.Chart3DData{
// 		{Name: "point1", Value: []interface{}{10, 10, 10}, ItemStyle: &opts.ItemStyle{Color: "green"}},
// 		{Name: "point2", Value: []interface{}{15, 15, 15}, ItemStyle: &opts.ItemStyle{Color: "blue"}},
// 		{Name: "point3", Value: []interface{}{20, 20, 20}, ItemStyle: &opts.ItemStyle{Color: "red"}},
// 	})

// 	return scatter3d
// }

// type Scatter3dExamples struct{}

// func (Scatter3dExamples) Examples() {
// 	page := components.NewPage()
// 	page.AddCharts(
// 		scatter3DBase(),
// 		scatter3DDataItem(),
// 	)

// 	f, err := os.Create("scatter3d.html")
// 	if err != nil {
// 		panic(err)
// 	}
// 	page.Render(io.MultiWriter(f))
// }

// func main() {
// 	s := Scatter3dExamples{}
// 	s.Examples()
// }