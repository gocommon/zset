package zset

import (
	"sync"

	"github.com/gocommon/skiplist"
)

// Zset Zset 倒序存
type Zset struct {
	dict *sync.Map
	list *skiplist.SkipList
}

// NewZset NewZset
func NewZset() *Zset {
	return &Zset{
		dict: new(sync.Map),
		list: skiplist.New(),
	}
}

// Len Len
func (p *Zset) Len() int {
	return p.list.Len()
}

// Add zadd
func (p *Zset) Add(score float64, member string) {

	obj := NewObj(score, member)

	oldObj, has := p.dict.Load(member)
	if has {
		// fmt.Println("数据已存在，删除原数据", oldObj)
		p.list.Delete(oldObj.(*Obj))
	}

	// fmt.Println("更新数据", obj)
	p.dict.Store(member, obj)
	p.list.Insert(obj)
}

// Range Range
// GetElementByRank finds an element by ites rank. The rank argument needs bo be 1-based.
// Note that is the first element e that GetRank(e.Value) == rank, and returns e or nil.
func (p *Zset) Range(start, stop int) []*Obj {

	start, stop = p.rangeZone(start, stop, false)
	// fmt.Println("Range", start, stop)

	start++
	stop++

	e := p.list.GetElementByRank(start)

	list := make([]*Obj, 0, stop-start)

	for i := start; e != nil && e.Value != nil && i < stop; i++ {
		// fmt.Println("Range append", p.list.Len(), i, start, stop)
		list = append(list, e.Value.(*Obj))
		e = e.Next()
	}

	return list
}

// RevRange RevRange
func (p *Zset) RevRange(start, stop int) []*Obj {
	start, stop = p.rangeZone(start, stop, true)
	// fmt.Println("RevRange", start, stop)

	list := make([]*Obj, 0, start-stop)

	e := p.list.GetElementByRank(start)
	// fmt.Println("RevRange e", e.Value)

	for i := start; e != nil && e.Value != nil && i > stop; i-- {
		// fmt.Println("start", i, e.Value)
		list = append(list, e.Value.(*Obj))
		e = e.Prev()
	}

	return list

}

func (p *Zset) rangeZone(start, stop int, rev bool) (int, int) {
	if start < 0 {
		start = p.list.Len() + start + 1
	}

	if stop < 0 {
		stop = p.list.Len() + stop + 1
	}

	if start > stop {
		return 0, 0
	}

	if rev {
		start = p.list.Len() - start
		stop = p.list.Len() - stop
	}

	if start < 0 {
		start = 0
	}
	if stop > p.list.Len() {
		stop = p.list.Len()
	}

	return start, stop
}

// Clean Clean
func (p *Zset) Clean() {
	p.dict = new(sync.Map)
	p.list = skiplist.New()
}
