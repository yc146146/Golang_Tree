package pq

import "sort"

type PriorityQueue struct {
	items *items
	sizeOption Option

}

type optionEnum int

const (
	none optionEnum = iota
	minPrioSize
	maxPrioSize
)

type Option struct {
	option optionEnum
	value int

}

func NewPriorityQueue(options ...Option)*PriorityQueue{
	pq := PriorityQueue{
		items:&items{},
	}
	for _, o := range options {
		switch o.option {
		case none, minPrioSize,maxPrioSize:
			pq.sizeOption = o
		}
	}
	return &pq

}

func WithMinPrioSize(size int) Option{
	return Option{
		option: minPrioSize,
		value: size,
	}
}

func WithMaxPrioSize(size int) Option{
	return Option{
		option: maxPrioSize,
		value: size,
	}
}

func (p *PriorityQueue)Len()int{
	return p.items.Len()
}


func (p *PriorityQueue)Insert(v interface{}, priority float64){
	*p.items = append(*p.items, &item{value:v, priority: priority})
	sort.Sort(p.items)
	switch p.sizeOption.option {
	case minPrioSize:
		if p.sizeOption.value < len(*p.items){
			*p.items = (*p.items)[:p.sizeOption.value]
		}
	case maxPrioSize:
		diff := len(*p.items) - p.sizeOption.value
		if diff > 0{
			*p.items = (*p.items)[diff:]
		}
	}
}

func (p *PriorityQueue)PopLowest()interface{}{
	if len(*p.items) == 0{
		return nil

	}

	x := (*p.items)[0]
	*p.items = (*p.items)[1:]
	return x.value
}

func (p *PriorityQueue)PopHighest()interface{}{
	l := len(*p.items) - 1
	if l<0{
		return  nil

	}

	x := (*p.items)[l]
	*p.items = (*p.items)[:l]
	return x.value
}


func (p *PriorityQueue)Get(i int)(interface{}, float64){
	x := (*p.items)[i]
	return x.value, x.priority
}



type items []*item

func (p items) Len()int {return len(p)}
func (p items) Less(i,j int)bool {return p[i].priority < p[j].priority}
func (p items) Swap(i,j int) {p[i],p[j]=p[j],p[i]}

type item struct {
	value interface{}
	priority float64
}