package mathlib

import (
	"github.com/tiyee/gokit/internal/constraints"
	"math"
)

// R 地球半径，单位米
const R = 6367000

// Distance
// lonA, latA分别为A点的纬度和经度
// lonB, latB分别为B点的纬度和经度
// 返回的距离单位为米
func point(a int64) float64 {
	return float64(a) / 1000000
}
func Distance(lngA, latA, lngB, latB float64) float64 {
	c := math.Sin(latA)*math.Sin(latB)*math.Cos(lngA-lngB) + math.Cos(latA)*math.Cos(latB)
	return R * math.Acos(c) * math.Pi / 180
}
func DistanceI[T constraints.Integer](lngA, latA, lngB, latB int64) T {
	return T(Distance(point(lngA), point(latA), point(lngB), point(latB)))
}

// BDToGCJ 百度坐标系 (BD-09) 与 火星坐标系 (GCJ-02)的转换
func BDToGCJ(lat, lng float64) (float64, float64) {
	const xPi = math.Pi * 3000.0 / 180.0
	x := lng - 0.0065

	y := lat - 0.006
	z := math.Sqrt(x*x+y*y) - 0.00002*math.Sin(y*xPi)
	theta := math.Atan2(y, x) - 0.000003*math.Cos(x*xPi)

	g2lng := z * math.Cos(theta)

	g2lat := z * math.Sin(theta)
	return g2lat, g2lng
}
