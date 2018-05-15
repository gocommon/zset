package zset

import (
	"github.com/gocommon/skiplist"
)

var _ skiplist.Interface = &Obj{}

// Obj Obj
type Obj struct {
	score  float64
	member string
}

func NewObj(score float64, member string) *Obj {
	return &Obj{score, member}
}

func (p *Obj) Score() float64 {
	return p.score
}

func (p *Obj) Member() string {
	return p.member
}

// Less Less
func (p *Obj) Less(other interface{}) bool {
	if p.score < other.(*Obj).score {
		return true
	}
	if p.score == other.(*Obj).score && p.member < other.(*Obj).member {
		return true
	}
	return false
}
