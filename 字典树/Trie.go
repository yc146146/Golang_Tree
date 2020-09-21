package main

import "sort"

type Pair struct{
	K int64
	V int64
}
//map
type KeyWordKV map[int64]string
//字典树结构
type CharBeginKV map[string][]*KeyWordTreeNode

type PairList []Pair

func (p PairList)Len()int{
	return len(p)
}
func (p PairList)Less(i,j int)bool{
	return p[i].V>p[j].V
}

func (p PairList)Swap(i,j int){
	//实现数据交换
	p[i].V,p[j].V = p[j].V,p[i].V
}


//字符串压入树中
func (stree *KeyWordTree) Put(id int64, keyword string) {
	//后缀, keyword 反转

	keyword = rev(keyword)

	stree.rw.Lock()
	defer stree.rw.Unlock()
	//保存一份
	stree.kv[id] = keyword
	//备份root节点
	tmproot := stree.root


	for _, v := range keyword {
		//处理每个字符转化为字符串
		ch := string(v)
		if tmproot.SubKeyWordTreeNodes[ch] == nil {
			//开辟节点插入
			node := NewKeyWordTreeNodeWithParams(ch, tmproot)
			tmproot.SubKeyWordTreeNodes[ch] = node
			//加入每一个节点
			stree.char_begin_kv[ch] = append(stree.char_begin_kv[ch], node)
		} else {
			keywordtreenode := tmproot.SubKeyWordTreeNodes[ch]
			keywordtreenode.KeyWordIDs[id] = true
			//节点向前推进
			tmproot = tmproot.SubKeyWordTreeNodes[ch]
		}
	}

}

//搜索
func (stree *KeyWordTree) Search(keyword string, limit int) []string {
	stree.rw.Lock()
	defer stree.rw.Unlock()

	ids := make(map[int64]int64, 0)
	for pos, v := range keyword {
		ch := string(v)
		//取出映射字符的所有节点
		begins := stree.char_begin_kv[ch]

		for _, begin := range begins {
			//备份地址
			key_word_tmp_pt := begin
			//标记下一个位置
			next_pos := pos + 1
			for len(key_word_tmp_pt.SubKeyWordTreeNodes) > 0 && next_pos < len(keyword) {
				//	最大匹配
				//下一个字符
				next_ch := string(keyword[next_pos])
				if key_word_tmp_pt.SubKeyWordTreeNodes[next_ch] == nil {
					//跳出循环
					break
				}
				//abc
				//地推前进
				key_word_tmp_pt = key_word_tmp_pt.SubKeyWordTreeNodes[next_ch]
				next_pos++
			}
			//	保存结果
			for id, _ := range key_word_tmp_pt.KeyWordIDs {
				ids[id] = ids[id] + 1
			}
		}
	}
	//列表
	list := PairList{}
	for id, count := range ids{
		//加载数据
		list = append(list, Pair{id, count})
	}

	if !sort.IsSorted(list){
		//排序
		sort.Sort(list)
	}

	//limit 限制出现的数量
	if len(list) > limit{
		//数据截取
		list = list[:limit]
	}

	ret := make([]string, 0)
	for _,item := range list{
		//返回数组叠加
		//ret = append(ret, stree.kv[item.K])
		ret = append(ret, rev(stree.kv[item.K]))
	}


	return ret
}

//搜索提示 返回字符串 limit限制深度
func (stree *KeyWordTree) Sugg(keyword string, limit int) []string {
	stree.rw.Lock()
	defer stree.rw.Unlock()

	//根节点
	key_word_tmp_pt := stree.root
	//是否结束
	is_end := true

	for _, v := range keyword{
		ch := string(v)
		if key_word_tmp_pt.SubKeyWordTreeNodes[ch] == nil{
			is_end = false
			break
		}
		//孩子集合
		//循环的条件
		key_word_tmp_pt = key_word_tmp_pt.SubKeyWordTreeNodes[ch]
	}

	if is_end{

		ret := make([]string,0)
		ids := key_word_tmp_pt.KeyWordIDs
		for id,_ := range ids{
			//实现结果的追加
			//ret = append(ret, stree.kv[id])
			ret = append(ret, rev(stree.kv[id]))
			limit--
			if limit==0{
				break
			}
		}
		return ret


	}

	return make([]string, 0)
}



