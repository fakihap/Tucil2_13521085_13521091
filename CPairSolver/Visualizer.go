package CPairSolver

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/icza/gox/osx"
)

func visualizePoints(points []Point) {
	scatter3d := charts.NewScatter3D()
	scatter3d.Chart3D.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Closest Pair Problem : 3D Visualization",
			Subtitle: "Author : Addin Munawwar Yusuf and Fakih Anugerah Pratama",
			Top:      "5%",
			Bottom:   "5%",
		}),
	)

	scatter3d.AddSeries("scatter3d", genScatter3dData(points))
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		scatter3d.Render(w)
	})
	osx.OpenDefault("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func genScatter3dData(points []Point) []opts.Chart3DData {
	var data []opts.Chart3DData
	for i, point := range points {
		data = append(data, opts.Chart3DData{
			Name:  "Point " + string(i),
			Value: []interface{}{point.GetAxisValue(0), point.GetAxisValue(1), point.GetAxisValue(2)},
			ItemStyle: &opts.ItemStyle{
				Color:   "#1ecbe1",
				Opacity: 1,
			},
		})
	}
	return data
}

// func main() {
// 	points := []Point{ *NewPoint(1, 2, 3), *NewPoint(4, 5, 6), *NewPoint(7, 8, 9), *NewPoint(10, 11, 12)}
// 	visualizePoints(points)
// }
