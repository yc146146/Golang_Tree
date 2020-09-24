package point

import "fmt"

type Point2D struct {
	X float64
	Y float64
}

//2维
func (p *Point2D)Dimensions()int{
	return 2
}

//根据维度返回数据
func (p *Point2D) Dimension(i int) float64{
	if i==0 {
		return p.X
	}
	return p.Y
}

// 字符串返回
func (p *Point2D) String() string{
	return fmt.Sprintf("{%.2f %.2f}", p.X, p.Y)
}

