package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type VEB struct {
	//u数据 min 最小 max最大
	u, min, max int
	//总结
	summary *VEB
	//N个子节点
	cluster []*VEB
}

func (V VEB) Max() int {
	return V.max
}

func (V VEB) Min() int {
	return V.min
}

//给我一个x 计算存储的深度的族号
func (V VEB) High(x int) int {
	return int(math.Floor(float64(x) / float64(V.u)))
}

func LowerSqrt(u int) int {
	return int(math.Pow(2.0, math.Floor(math.Log2(float64(u))/2)))
}

func HigherSqrt(u int) int {
	return int(math.Pow(2.0, math.Ceil(math.Log2(float64(u))/2)))
}

func (V VEB) Low(x int) int {
	return x % LowerSqrt(V.u)
}

func (V VEB) Index(x, y int) int {
	return x*LowerSqrt(V.u) + y

}

//创造树
func CreateTree(size int) *VEB {
	if size < 0 {
		return nil
	}

	x := math.Ceil(math.Log2(float64(size)))
	u := int(math.Pow(2, x))

	//新建一个节点
	V := new(VEB)
	V.min, V.max = -1, -1
	V.u = u

	if u == 2 {
		return V
	}

	//计算cluster的数量
	clutercount := HigherSqrt(u)
	clustersize := LowerSqrt(u)

	for i := 0; i < clutercount; i++ {
		V.cluster = append(V.cluster, CreateTree(clustersize))
	}

	summarysize := HigherSqrt(u)
	V.summary = CreateTree(summarysize)

	return V
}

//判断节点是否存在
func (V VEB) IsMember(x int) bool {
	if x == V.min || x == V.max {
		return true
	} else if V.u == 2 {
		return false
	} else {
		return V.cluster[V.High(x)].IsMember(V.Low(x))
	}
}

//插入节点
func (V *VEB) Insert(x int) {
	if V.min == -1 {
		V.min, V.max = x, x

	} else {
		if x < V.min {
			V.min, x = x, V.min
		}

		if V.u > 2 {
			if V.cluster[V.High(x)].Min() == -1 {
				V.summary.Insert(V.High(x))
				V.cluster[V.High(x)].min, V.cluster[V.High(x)].max = V.Low(x), V.Low(x)
			} else {
				V.cluster[V.High(x)].Insert(V.Low(x))
			}
		}

		if x > V.max {
			V.max = x

		}
	}
}

//删除节点
func (V *VEB) Delete(x int) {
	if V.summary == nil || V.summary.Min() == -1 {
		//	无非空族
		if x == V.min && x == V.max {
			V.min, V.max = -1, -1
		} else if x == V.min {
			V.min = V.max
		} else {
			V.max = V.min
		}
	} else {
		//	存在非空族
		if x == V.min {
			//	取得最小在cluster
			y := V.Index(V.summary.min, V.cluster[V.summary.min].min)
			//取得最接近的, 赋值v.min
			V.min = y
			V.cluster[V.High(x)].Delete(V.Low(y))
			if V.cluster[V.High(y)].min == -1 {
				V.summary.Delete(V.High(y))
			}

		} else if x == V.max {
			y := V.Index(V.summary.max, V.cluster[V.summary.max].max)
			V.cluster[V.High(y)].Delete(V.Low(y))
			if V.cluster[V.High(y)].min == -1 {
				V.summary.Delete(V.High(y))
			}
			if V.summary == nil || V.summary.min == -1 {
				if V.min == y {
					V.min, V.max = -1, -1
				} else {
					//重合删除
					V.max = V.min
				}
			} else {
				V.max = V.Index(V.summary.max, V.cluster[V.summary.max].max)
			}
		} else {
			//删除节点
			V.cluster[V.High(x)].Delete(V.Low(x))
			if V.cluster[V.High(x)].min == -1 {
				V.summary.Delete(V.High(x))
			}
		}
	}
}

//找到X的位置
func (V VEB) Successor(x int) int {
	if V.u == 2 {
		if x == 0 && V.max == 1 {
			return 1
		} else {
			return -1
		}
	} else if V.min != -1 && x < V.min {
		//V.min就是x的后继
		return V.min
	} else {
		//最大值
		maxlow := V.cluster[V.High(x)].Max()
		if maxlow != -1 && V.Low(x) < maxlow {
			offset := V.cluster[V.High(x)].Successor(V.Low(x))
			return V.Index(V.High(x), offset)
		} else {
			succCluster := V.summary.Successor(V.High(x))
			if succCluster == -1 {
				return -1
			} else {
				offset := V.cluster[succCluster].Min()
				return V.Index(succCluster, offset)
			}

		}
	}
}

//找到X的位置
func (V VEB) Predecessor(x int) int {
	if V.u == 2 {
		if x == 1 && V.min == 0 {
			return 0
		} else {
			return -1
		}
	} else if V.min != -1 && x > V.max {
		//V.min就是x的后继
		return V.max
	} else {
		//最大值
		minlow := V.cluster[V.High(x)].Max()
		if minlow != -1 && V.Low(x) > minlow {
			offset := V.cluster[V.High(x)].Successor(V.Low(x))
			return V.Index(V.High(x), offset)
		} else {
			preCluster := V.summary.Predecessor(V.High(x))
			if preCluster == -1 {

				if V.min != -1 && x > V.min {
					return V.min

				} else {
					return -1
				}
			} else {
				offset := V.cluster[preCluster].Max()
				return V.Index(preCluster, offset)
			}

		}
	}
}

//统计数的节点
func (V VEB) Count() int {
	if V.u == 2 {
		return 1
	}

	sum := 1
	for i := 0; i < len(V.cluster); i++ {
		//统计次数
		sum += V.cluster[i].Count()
	}

	sum += V.summary.Count()
	return sum

}

func (V VEB) Print() {
	V.PrintFunc(0, 0, false)
}

//递归打印层级
func (V VEB) PrintFunc(level int, clusterNo int, summary bool) {
	space := " "
	for i := 0; i < level; i++ {
		space += "\t"
	}
	if level == 0 {
		fmt.Printf("%vR:{u:%v,min:%v;max:%v,cluster:%v}\n", space, V.u, V.min, V.max, len(V.cluster))

	} else {
		if summary {
			fmt.Printf("%vS:{u:%v,min:%v;max:%v,cluster:%v}\n", space, V.u, V.min, V.max, len(V.cluster))

		} else {
			fmt.Printf("%vC[%v]:{u:%v,min:%v;max:%v,cluster:%v}\n", space, clusterNo, V.u, V.min, V.max, len(V.cluster))

		}

	}

	if len(V.cluster) > 0 {
		V.summary.PrintFunc(level+1, 0, true)
		for i := 0; i < len(V.cluster); i++ {
			V.cluster[i].PrintFunc(level+1, i, false)
		}
	}
}

//清空
func (V *VEB) Clear() {
	V.min, V.max = -1, -1
	if V.u == 2 {
		return
	}
	for i := 0; i < len(V.cluster); i++ {
		//统计次数
		V.cluster[i].Clear()
	}
	V.summary.Clear()

}

//填充
func (V *VEB) Fill() {
	for i := 0; i < V.u; i++ {
		V.Insert(i)
	}
}

//返回元素
func (V VEB) Members() []int {
	members := []int{}
	for i := 0; i < V.u; i++ {
		if V.IsMember(i) {
			members = append(members, i)
		}
	}
	return members
}

func arrayContains(arr []int, value int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return true
		}
	}
	return false
}

//创建随机数数组 不重复
func makerandom(max int) []int {
	myrand := rand.New(rand.NewSource(int64(time.Now().Nanosecond() * 1)))
	keys := []int{}
	keyNo := myrand.Intn(max)
	for i := 0; i < keyNo; i++ {
		mykey := myrand.Intn(max - 1)
		if arrayContains(keys, mykey) == false {
			keys = append(keys, mykey)
		}
	}
	sort.Ints(keys)
	return keys

}

func main() {
	maxUpower := 10
	for i := 1; i < maxUpower; i++ {
		u := int(math.Pow(2, float64(i)))
		V := CreateTree(u)

		keys := makerandom(u)
		

		fmt.Println("keys", keys)
		for j := 0; j < len(keys); j++ {
			//插入数据
			V.Insert(keys[j])
		}

		for j:=0;j<u;j++{
			if j>0 && j<len(keys){
				fmt.Println(V.IsMember(keys[j]))
			}

		}


		V.Print()
		fmt.Println("-----------------------")
	}
}
