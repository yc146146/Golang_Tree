package point

import "fmt"

type Point3D struct {
	X float64
	Y float64
	Z float64

}

//返回维度的数据
func (p *Point3D) Dimensions() int{
	return 3
}

func (p *Point3D) Dimension(i int)float64{
	switch i {
	case 0:
		return p.X
	case 1:
		return p.Y
	default:
		return p.Z

	}
}

//字符串返回
func (p *Point3D)String()string{
	return fmt.Sprintf("{%2.f %2.f %2.f}", p.X, p.Y, p.Z)
}