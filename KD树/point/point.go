package point

import "fmt"

//点的集合
type Point struct {
	Coordinates []float64
	Data interface{}
}

func NewPoint(coordinates []float64, data interface{}) *Point {
	return &Point{
		Coordinates: coordinates,
		Data:data,
	}
}

//计算维度
func (p *Point) Dimensions() int{
	return len(p.Coordinates)
}

//根据维度返回x,y,z
func (p *Point) Dimension(i int) float64{
	return p.Coordinates[i]
}

func (p *Point)String()string{
	return fmt.Sprintf("{%v %v}",p.Coordinates, p.Data)
}
