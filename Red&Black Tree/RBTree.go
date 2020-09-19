package main

const (
	RED   = true
	BLACK = false
)

//红黑树的结构
type RBNode struct {
	//节点
	Left   *RBNode
	Right  *RBNode
	Parent *RBNode
	//颜色
	Color bool
	//数据接口
	//DataItem interface{}
	Item
}

//数据接口
type Item interface {
	Less(than Item) bool
}

type RBtree struct {
	NIL   *RBNode
	Root  *RBNode
	count uint
}

//比大小
func less(x, y Item) bool {
	return x.Less(y)
}

//初始化内存
func NewRBTree() *RBtree {
	return new(RBtree).Init()
}

//初始化红黑树
func (rbt *RBtree) Init() *RBtree {
	node := &RBNode{nil, nil, nil, BLACK, nil}
	return &RBtree{node, node, 0}
}

//获取红黑树的长度
func (rbt *RBtree) Len() uint {
	return rbt.count
}

//取得红黑树的极大值
func (rbt *RBtree) max(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}

	for x.Right != rbt.NIL {
		x = x.Right
	}
	return x
}

//取得红黑树的极小值
func (rbt *RBtree) min(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}

	for x.Left != rbt.NIL {
		x = x.Left
	}
	return x
}

//搜索红黑树
func (rbt *RBtree) search(x *RBNode) *RBNode {
	pnode := rbt.Root
	for pnode != rbt.NIL {
		if less(pnode.Item, x.Item) {
			pnode = pnode.Right
		} else if less(x.Item, pnode.Item) {
			pnode = pnode.Left
		} else {
			//找到
			break
		}
	}
	return pnode
}

func (rbt *RBtree) leftRotate(x *RBNode) {
	if x.Right == rbt.NIL {
		//左旋转 逆时针
		return
	}

	y := x.Right
	//实现旋转的左旋
	x.Right = y.Left
	if y.Left != rbt.NIL {
		//设定父亲节点
		y.Left.Parent = x
	}

	//传递父节点
	y.Parent = x.Parent
	if x.Parent == rbt.NIL {
		//	根节点
		rbt.Root = y

	} else if x == x.Parent.Left {
		//x在根节点的左边
		x.Parent.Left = y

	} else {
		//x在根节点的右边
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y

}

func (rbt *RBtree) rightRotate(x *RBNode) {
	if x.Left == nil {
		return
	}
	y := x.Left
	x.Left = y.Right

	if y.Right != rbt.NIL {
		//设置祖先
		y.Right.Parent = x

	}
	//y保存x的父节点
	y.Parent = x.Parent

	if x.Parent == rbt.NIL {
		rbt.Root = y

	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y

	}

	y.Right = x
	x.Parent = y

}

//插入一条数据
func (rbt *RBtree) Insert(item Item) *RBNode {
	if item == nil {
		return nil

	}

	return rbt.insert(&RBNode{rbt.NIL, rbt.NIL, rbt.NIL, RED, item})
}

//插入
func (rbt *RBtree) insert(z *RBNode) *RBNode {
	//	寻找插入位置
	x := rbt.Root
	y := rbt.NIL

	for x != rbt.NIL {
		//备份位置,数据插入x,y之间
		y = x
		if less(z.Item, x.Item) {
			x = x.Left
		} else if less(x.Item, z.Item) {
			x = x.Right
		} else {
			//数据已经存在无法插入
			return x
		}
	}
	z.Parent = y
	if y == rbt.NIL {
		rbt.Root = z

	} else if less(z.Item, y.Item) {
		//小于左边插入
		y.Left = z
	} else {
		//大于右边插入
		y.Right = z
	}
	rbt.count++
	rbt.insertFixup(z)
	return z
}

//插入之后,调整平衡
func (rbt *RBtree) insertFixup(z *RBNode) {
	//一直循环下去 直到根节点
	for z.Parent.Color == RED {
		//父亲节点在爷爷左边
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				//循环判断
				z = z.Parent.Parent
			} else {
				//z 比父亲小
				if z == z.Parent.Right {
					z = z.Parent
					//左旋
					rbt.leftRotate(z)
				} else {
					//	z比父亲大
					z.Parent.Color = BLACK
					z.Parent.Parent.Color = RED
					rbt.rightRotate(z.Parent.Parent)
				}
			}
		} else {
			y := z.Parent.Parent.Left
			if y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				//循环前进
				z = z.Parent.Parent

			} else {
				if z == z.Parent.Left {
					z = z.Parent
					rbt.rightRotate(z)
				} else {
					z.Parent.Color = BLACK
					z.Parent.Parent.Color = RED
					rbt.leftRotate(z.Parent.Parent)
				}

			}
		}
	}
	rbt.Root.Color = BLACK
}

func (rbt *RBtree) GetDepth() int {
	var getDeepth func(node *RBNode) int
	getDeepth = func(node *RBNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return 1
		}
		var leftdeep int = getDeepth(node.Left)
		var rightdeep int = getDeepth(node.Right)

		if leftdeep > rightdeep {
			return leftdeep + 1
		} else {
			return rightdeep + 1
		}
	}
	return getDeepth(rbt.Root)
}

//近似查找
func (rbt *RBtree) searchle(x *RBNode) *RBNode {
	//根节点
	p := rbt.Root
	//备份根节点
	n := p
	for n != rbt.NIL {
		if less(n.Item, x.Item) {
			p = n
			//大鱼
			n = n.Right
		} else if less(x.Item, n.Item) {
			p = n
			//小鱼
			n = n.Left
		} else {
			return n
			break

		}
	}

	if less(p.Item, x.Item) {
		return p

	}

	//近似查找
	p = rbt.desuccessor(p)
	return p

}

func (rbt *RBtree) successor(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}

	if x.Right != rbt.NIL {
		//取得右边最小
		return rbt.min(x.Right)
	}

	y := x.Parent
	for y != rbt.NIL && x == y.Right {
		x = y
		y = y.Parent
	}

	return y

}
func (rbt *RBtree) desuccessor(x *RBNode) *RBNode {
	if x == rbt.NIL {
		return rbt.NIL
	}

	if x.Left != rbt.NIL {
		//取得左边最大
		return rbt.max(x.Left)
	}

	y := x.Parent
	for y != rbt.NIL && x == y.Left {
		x = y
		y = y.Parent
	}

	return y
}


//最小最大查找 修改 近似查找
func (rbt *RBtree)Delete(item Item)Item{
	if item == nil{
		return nil
	}
	return rbt.delete(&RBNode{rbt.NIL,rbt.NIL,rbt.NIL,RED,item}).Item
}

func (rbt *RBtree) delete(key *RBNode) *RBNode {
	z := rbt.search(key)
	if z == rbt.NIL {
		return rbt.NIL
	}

	//新建节点备份 x,y 备份 夹闭
	var x *RBNode
	var y *RBNode
	//节点备份
	ret := &RBNode{rbt.NIL, rbt.NIL, rbt.NIL, z.Color, z.Item}

	if z.Left == rbt.NIL || z.Right == rbt.NIL {
		//直接替换删除
		//电接点 y z 重合
		y = z
	} else {
		//找到最接近的 右边最小
		y = rbt.successor(z)
	}

	if y.Left != rbt.NIL {
		x = y.Left
	} else {
		x = y.Right
	}

	x.Parent = y.Parent

	if y.Parent == rbt.NIL {
		rbt.Root = x

	} else if y == y.Parent.Left {
		y.Parent.Left = x

	} else {
		y.Parent.Right = x

	}

	if y != z {
		z.Item = y.Item
	}

	if y.Color == BLACK {
		rbt.deleteFixup(x)
	}

	rbt.count--
	return ret

}

func (rbt *RBtree) deleteFixup(x *RBNode) {

	for x != rbt.Root && x.Color == BLACK {
		//x在左边
		if x == x.Parent.Left {
			w := x.Parent.Right
			//左旋
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color=RED
				rbt.leftRotate(x.Parent)
				w=x.Parent.Right

			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color=RED
				x = x.Parent

			} else {
				if w.Right.Color == BLACK{
					w.Left.Color=BLACK
					w.Color=RED
					//右旋
					rbt.rightRotate(w)
					w=x.Parent.Right
				}
				w.Color=x.Parent.Color
				x.Parent.Color=BLACK
				w.Right.Color=BLACK
				rbt.leftRotate(x.Parent)
				x=rbt.Root

			}
		} else {
			w := x.Parent.Left
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color=RED
				rbt.rightRotate(x.Parent)
				w=x.Parent.Right
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color=RED
				x = x.Parent
			} else {
				if w.Right.Color == BLACK{
					w.Left.Color=BLACK
					w.Color=RED
					//右旋
					rbt.leftRotate(w)
					w=x.Parent.Left
				}
				w.Color=x.Parent.Color
				x.Parent.Color=BLACK
				w.Right.Color=BLACK
				rbt.rightRotate(x.Parent)
				x=rbt.Root
			}
		}
	}

	//循环到最后就是根节点
	x.Color = BLACK

}
