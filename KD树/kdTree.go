package main

import (
	"fmt"
	"math"
	"sort"
	"yinchen.com/文件3/8.KD树/point"
	"yinchen.com/文件3/8.KD树/pq"
)

type Point interface {
	//维度
	Dimensions() int
	//根据维度取出数据
	Dimension(i int) float64
}

type KDTree struct {
	root *node
}

//显示树
func (t *KDTree) String() string {
	return fmt.Sprintf("[%s]", PrintTreeNode(t.root))
}

func PrintTreeNode(n *node) string {
	if n != nil && (n.Left != nil || n.Right != nil) {
		return fmt.Sprintf("[%s,%s,%s]", PrintTreeNode(n.Left), n.String(), PrintTreeNode(n.Right))
	}
	return fmt.Sprintf("%s", n)
}

func (t *KDTree) Insert(p Point) {
	if t.root == nil {
		t.root = &node{p, nil, nil}
	} else {
		t.root.Insert(p, 0)
	}
}

func (t *KDTree) Remove(p Point) Point {
	if t.root == nil || p == nil {
		return nil

	}
	n, sub := t.root.Remove(p, 0)
	if n == t.root {
		t.root = sub

	}

	if n == nil {
		return nil

	}
	return n.Point
}

func (t *KDTree) Points() []Point {
	if t.root == nil {
		return []Point{}
	}
	return t.root.Points()
}

func (t *KDTree) Balance() {
	//平衡
	t.root = newKDTree(t.Points(), 0)
}

func (t *KDTree) RangeSearch(r point.Range) []Point {
	if t.root == nil || t == nil || len(r) != t.root.Dimensions() {
		return []Point{}
	}

	return t.root.RangeSearch(r, 0)
}

func NewKDTree(points []Point) *KDTree {
	return &KDTree{root: newKDTree(points, 0)}
}

func newKDTree(points []Point, axis int) *node {
	if len(points) == 0 {
		return nil
	}

	if len(points) == 1 {
		return &node{points[0], nil, nil}
	}

	//排序
	sort.Sort(&byDismension{axis, points})

	//取得中间值
	mid := len(points) / 2

	//设定根节点
	root := points[mid]

	//1,2,3,4  维度循环
	nextDim := (axis + 1) % root.Dimensions()

	return &node{
		root,
		newKDTree(points[:mid], nextDim),
		newKDTree(points[mid+1:], nextDim),
	}
}

type byDismension struct {
	dismension int
	points     []Point
}

func (b *byDismension) Len() int {
	return len(b.points)
}

//设定比较大小规则
func (b *byDismension) Less(i, j int) bool {
	return b.points[i].Dimension(b.dismension) < b.points[j].Dimension(b.dismension)
}

//交换数据
func (b *byDismension) Swap(i, j int) {
	b.points[i], b.points[j] = b.points[j], b.points[i]
}

//二叉树
type node struct {
	Point
	Left  *node
	Right *node
}

//显示节点
func (n *node) String() string {
	return fmt.Sprintf("%v", n.Point)
}

//返回所有节点
func (n *node) Points() []Point {
	var points []Point
	if n.Left != nil {
		points = n.Left.Points()
	}

	points = append(points, n.Point)
	if n.Right != nil {
		points = append(points, n.Right.Points()...)
	}
	return points
}

func (n *node) Insert(p Point, axis int) {
	if p.Dimension(axis) < n.Point.Dimension(axis) {
		if n.Left == nil {
			n.Left = &node{p, nil, nil}
		} else {
			n.Left.Insert(p, (axis+1)%n.Point.Dimensions())
		}
	} else {
		if n.Right == nil {
			n.Right = &node{p, nil, nil}
		} else {
			n.Right.Insert(p, (axis+1)%n.Point.Dimensions())
		}
	}
}

//按照维度查找最小
func (n *node) FindMin(axis int, smallest *node) *node {
	if smallest == nil || n.Dimension(axis) < smallest.Dimension(axis) {
		smallest = n
	}
	if n.Left != nil {
		smallest = n.Left.FindMin(axis, smallest)
	}

	if n.Right != nil {
		smallest = n.Right.FindMin(axis, smallest)
	}

	return smallest
}

func (n *node) FindMax(axis int, biggest *node) *node {
	if biggest == nil || n.Dimension(axis) > biggest.Dimension(axis) {
		biggest = n
	}
	if n.Left != nil {
		biggest = n.Left.FindMin(axis, biggest)
	}

	if n.Right != nil {
		biggest = n.Right.FindMin(axis, biggest)
	}

	return biggest
}

//返回节点 替换节点
func (n *node) Remove(p Point, axis int) (*node, *node) {
	for i := 0; i < n.Dimensions(); i++ {
		if n.Dimension(i) != p.Dimension(i) {
			if n.Left != nil {
				//左子树循环
				returnNode, subNode := n.Left.Remove(p, (axis+1)%n.Dimensions())
				if returnNode != nil {
					if returnNode == n.Left {
						n.Left = subNode
					}
					return returnNode, nil
				}
			}

			if n.Right != nil {
				returnNode, subNode := n.Right.Remove(p, (axis+1)%n.Dimensions())
				if returnNode != nil {
					if returnNode == n.Right {
						n.Right = subNode
					}
					return returnNode, nil
				}
			}
			//不等无需删除
			return nil, nil
		}
	}

	if n.Left != nil {
		biggest := n.Left.FindMax(axis, nil)
		removed, sub := n.Left.Remove(biggest, axis)
		removed.Left = n.Left
		removed.Right = n.Right
		if n.Left == removed {
			removed.Left = sub

		}
		return n, removed
	}

	if n.Right != nil {
		smallest := n.Right.FindMin(axis, nil)
		removed, sub := n.Right.Remove(smallest, axis)
		removed.Left = n.Left
		removed.Right = n.Right
		if n.Right == removed {
			removed.Right = sub

		}
		return n, removed
	}

	//left = right = nil
	return n, nil

}

//按照维度搜索 范围内的数据
func (n *node) RangeSearch(r point.Range, axis int) []Point {
	//节点集合
	points := []Point{}
	for dim, limit := range r {
		if limit[0] > n.Dimension(dim) || limit[1] < n.Dimension(dim) {
			//	节点在我的范围之内
			goto ChildCheck
		}
	}
	//节点叠加
	points = append(points, n.Point)
ChildCheck:
	if n.Left != nil && n.Dimension(axis) >= r[axis][0] {
		points = append(points, n.Left.RangeSearch(r, (axis+1)%n.Dimensions())...)
	}

	if n.Right != nil && n.Dimension(axis) <= r[axis][0] {
		points = append(points, n.Right.RangeSearch(r, (axis+1)%n.Dimensions())...)

	}

	return points
}

//计算距离
func distance(p1, p2 Point) float64 {
	sum := 0.0
	for i := 0; i < p1.Dimensions(); i++ {
		sum += math.Pow(p1.Dimension(i)-p2.Dimension(i), 2.0)
	}

	return math.Sqrt(sum)
}

//维度距离的绝对值
func planeDistance(p Point, plane float64, dim int) float64 {
	return math.Abs(plane - p.Dimension(dim))
}

//弹出数据
func popLast(arr []*node) ([]*node, *node) {
	length := len(arr) - 1
	if length < 0 {
		return arr, nil

	}

	return arr[:length], arr[length]
}

//队列中取出第n大的数据
func GetKthDistance(nearestPQ *pq.PriorityQueue, i int) float64 {
	if nearestPQ.Len() <= i {
		return math.MaxFloat64
	}

	_, prio := nearestPQ.Get(i)
	return prio

}

func (t *KDTree) KNN(p Point, k int) []Point {
	if t.root == nil || p == nil || k == 0 {
		return []Point{}
	}

	//队列 处理距离
	nearestPQ := pq.NewPriorityQueue(pq.WithMinPrioSize(k))

	//遍历整个kd树
	knn(p, k, t.root, 0, nearestPQ)

	points := make([]Point, 0, k)
	for i := 0; i < k && 0 < nearestPQ.Len(); i++ {
		o := nearestPQ.PopLowest().(*node).Point
		points = append(points, o)
	}

	return points
}

func knn(p Point, k int, start *node, curAxis int, nearestPQ *pq.PriorityQueue) {
	if p == nil || k == 0 || start == nil {
		return
	}

	//路径
	var path []*node
	//当前节点
	curnode := start

	//向下移动
	for curnode != nil {
		//记录路径
		path = append(path, curnode)
		if p.Dimension(curAxis) < curnode.Dimension(curAxis) {
			curnode = curnode.Left
		} else {
			curnode = curnode.Right
		}
		curAxis = (curAxis + 1) % p.Dimensions()
	}

	//向上移动 维度倒退
	curAxis = (curAxis - 1 + p.Dimensions()) % p.Dimensions()
	for path, curnode := popLast(path); curnode != nil; path, curnode = popLast(path) {
		//	计算当前距离
		curDistance := distance(p, curnode)
		checkedDistance := GetKthDistance(nearestPQ, k-1)

		//淘汰长距离
		if curDistance < checkedDistance {
			//插入当前节点 当前距离
			nearestPQ.Insert(curnode, curDistance)
			checkedDistance = GetKthDistance(nearestPQ, k-1)
		}

		if planeDistance(p, curnode.Dimension(curAxis), curAxis) < checkedDistance {
			var next *node
			if p.Dimension(curAxis) < curnode.Dimension(curAxis) {
				next = curnode.Right
			} else {
				next = curnode.Left
			}

			//knn算法
			knn(p, k, next, (curAxis+1)%p.Dimensions(), nearestPQ)
		}

		curAxis = (curAxis - 1 + p.Dimensions()) % p.Dimensions()
	}

}
