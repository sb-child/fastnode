package nodes

const (
	NODE_TYPE_INPUT  = true
	NODE_TYPE_OUTPUT = false
)


func NewNodeBody() *NodeBody{
	r := NodeBody{}
	r.Init()
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
}

type NodeBody struct {
	inputs  map[string]*NodePoint
	outputs map[string]*NodePoint
}

func (body *NodeBody) Init(){
	body.inputs = make(map[string]*NodePoint)
	body.outputs = make(map[string]*NodePoint)
}

func (body *NodeBody) InsertPoint(nt bool, name string, p *NodePoint) bool {
	_, ok := body.GetPoint(nt, name)
	if ok {
		return false
	}
	if nt {
		body.inputs[name] = p
	}
	body.outputs[name] = p
	return true
}

func (body *NodeBody) GetPoint(nt bool, name string) (r *NodePoint, ok bool) {
	if nt {
		r, ok = body.inputs[name]
		return
	}
	r, ok = body.outputs[name]
	return
}
