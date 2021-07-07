package nodes

import (
	"errors"
	"math"
)

const (
	NODE_TYPE_INPUT  = true
	NODE_TYPE_OUTPUT = false
)

func NewNodeBody() *NodeBody {
	r := NodeBody{}
	r.Init()
	return &r
}

func NewNodeLine() *NodeLine {
	r := NodeLine{}
	return &r
}

func NewNodePoint(nb *NodeBody, nt bool, pid int64) *NodePoint {
	r := NodePoint{Parent: nb, NodeType: nt, Pid: pid}
	return &r
}

type NodeLine struct {
	left  *NodePoint
	right *NodePoint
}

func (line *NodeLine) Connect(l *NodePoint, r *NodePoint) {
	line.left = l
	line.right = r
}

type NodePoint struct {
	Parent   *NodeBody
	NodeType bool
	Pid      int64
}

func (p *NodePoint) GetName() (string, bool) {
	var points *map[string]int64
	if p.NodeType {
		points = &p.Parent.inputs_str
	} else {
		points = &p.Parent.outputs_str
	}
	for s, v := range *points {
		if v == p.Pid {
			return s, true
		}
	}
	return "", false
}

type NodeBody struct {
	points_counter int64
	points         map[int64]*NodePoint
	inputs_str     map[string]int64
	outputs_str    map[string]int64
}

func (body *NodeBody) Init() {
	body.points_counter = math.MinInt64
	body.points = make(map[int64]*NodePoint)
	body.inputs_str = make(map[string]int64)
	body.outputs_str = make(map[string]int64)
}

func (body *NodeBody) reset() (r int64, err error) {
	if len(body.points) == math.MaxInt64 {
		err = errors.New("no free pid left")
		return
	}
	dic := make(map[int64]*NodePoint)
	str1 := make(map[string]int64)
	str2 := make(map[string]int64)
	pc := (int64)(math.MinInt64)
	for _, j := range body.points {
		s, ok := j.GetName()
		if !ok {
			err = errors.New("failed to get name")
			return
		}
		j.Pid = pc
		dic[pc] = j
		if j.NodeType {
			str1[s] = pc
		} else {
			str2[s] = pc
		}
		pc++
	}
	r = pc
	return
}

func (body *NodeBody) getPid() (r int64) {
	r = body.points_counter
	if r == math.MaxInt64 {
		// r = math.MinInt64
		temp, err := body.reset()
		if err != nil {
			panic(err)
		}
		r = temp
	}
	body.points_counter++
	return
}

func (body *NodeBody) insertPoint(name string, p *NodePoint) {
	nt, pid := p.NodeType, p.Pid
	body.points[pid] = p
	if nt {
		body.inputs_str[name] = pid
	}
	body.outputs_str[name] = pid
}

func (body *NodeBody) InsertPoint(nt bool, name string) bool {
	pid := body.getPid()
	p := NewNodePoint(body, nt, pid)
	_, ok := body.GetPoint(nt, name)
	if ok {
		return false
	}
	body.insertPoint(name, p)
	return true
}

func (body *NodeBody) GetPoint(nt bool, name string) (r *NodePoint, ok bool) {
	pid := (int64)(0)
	temp_ok := false
	if nt {
		pid, temp_ok = body.inputs_str[name]
	} else {
		pid, temp_ok = body.outputs_str[name]
	}
	if !temp_ok {
		ok = false
		return
	}
	r, ok = body.points[pid]
	return
}
