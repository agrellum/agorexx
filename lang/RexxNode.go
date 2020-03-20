package lang

type RexxNode struct {
	Leaf     *Rexx
	InitLeaf *Rexx
}

func RxNode(argleaf *Rexx) (rcvr *RexxNode) {
	rcvr = &RexxNode{}
	rcvr.Leaf = argleaf
	rcvr.InitLeaf = argleaf
	return
}
