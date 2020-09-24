package point

//处理距离 二维数组
type  Range [][2]float64

//处理距离
func New(limits ...float64) Range {
	if limits == nil || len(limits)%2 != 0{
		return nil
	}
	r :=make([][2]float64, len(limits)/2)
	for i:=range r{
		r[i]=[2]float64{limits[2*i], limits[2*i+1]}
	}
	return r
}
