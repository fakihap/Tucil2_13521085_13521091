package main

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/icza/gox/osx"
)

type Visualizer struct{
	points []Point
	solutionPoints [2]Point
}

func NewVisualizer(points []Point, solution [2]Point) *Visualizer {
	v := new(Visualizer)
	v.points = points
	v.solutionPoints = solution

	return v
}

func (v Visualizer) visualize() {
	scatter3d := charts.NewScatter3D()
	scatter3d.Chart3D.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			PageTitle: "Closest Pair Visualization",
			Width: "90vw",
			Height: "90vh",
			BackgroundColor: "#F2F4F7",
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Closest Pair Problem : 3D Visualization",
			Subtitle: "Author : Addin Munawwar Yusuf and Fakih Anugerah Pratama",
			Top:      "5%",
			Bottom:   "5%",
		}),
	)

	scatter3d.AddSeries("scatter3d", v.genScatter3dData(v.points, v.solutionPoints))
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		scatter3d.Render(w)
	})
	osx.OpenDefault("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func (v Visualizer) genScatter3dData(points []Point, solutionPoints [2]Point) []opts.Chart3DData {
	var data []opts.Chart3DData
	for i, point := range points {
		if point.Equals(solutionPoints[0]) || point.Equals(solutionPoints[1]) {}
		data = append(data, opts.Chart3DData{
			Name:  "Point " + string(i),
			Value: []interface{}{point.GetAxisValue(0), point.GetAxisValue(1), point.GetAxisValue(2)},
			ItemStyle: &opts.ItemStyle{
				Color:   "#1ECBE1",
				Opacity: 1,
			},
		})
	}

	for i, point := range solutionPoints {
		data = append(data, opts.Chart3DData{
			Name:  "Solution Point " + string(i),
			Value: []interface{}{point.GetAxisValue(0), point.GetAxisValue(1), point.GetAxisValue(2)},
			ItemStyle: &opts.ItemStyle{
				Color:   "#E1341E",
				Opacity: 1,
			},
		})
	}

	return data
}
