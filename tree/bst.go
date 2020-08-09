package tree

import "fmt"

type BST struct {
	root *Node
	//比对函数，0:v==nodeV,正数:v>nodeV,负数:v<nodeV
	compareFunc func(v, nodeV interface{}) int
}

func NewBST(rootV interface{}, compareFunc func(v, nodeV interface{}) int) *BST {
	if nil == compareFunc {
		return nil
	}
	return &BST{root: NewNode(rootV), compareFunc: compareFunc}
}

func (bst *BST) Find(v interface{}) *Node {
	p := bst.root
	for p != nil {
		rs := bst.compareFunc(v, p.data)
		if rs == 0 {
			return p
		} else if rs > 0 {
			p = p.right
		} else {
			p = p.left
		}
	}
	return nil
}

func (bst *BST) Insert(v interface{}) bool {
	p := bst.root
	for p != nil {
		rs := bst.compareFunc(v, p.data)
		if rs == 0 {
			return false
		} else if rs > 0 {
			if p.right == nil {
				p.right = NewNode(v)
				break
			}
			p = p.right
		} else {
			if p.left == nil {
				p.left = NewNode(v)
				break
			}
			p = p.left
		}
	}
	return true
}

func (bst *BST) Delete(v interface{}) bool {
	p := bst.root
	var pp *Node = nil
	for p != nil {
		rs := bst.compareFunc(v, p.data)
		if rs == 0 {
			break
		}
		pp = p
		if rs < 0 {
			p = p.left
		} else {
			p = p.right
		}
	}

	if p == nil {
		return false
	}

	if p.left != nil && p.right != nil {
		//找到右子树中的最小值
		minpp := p
		minp := p.right
		for minp.left != nil {
			minpp = minp
			minp = minp.left
		}

		p.data = minp.data
		p = minp
		pp = minpp
	}

	var child *Node
	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	} else {
		child = nil
	}

	if pp == nil {
		bst.root = child
	} else if pp.left == p {
		pp.left = child
	} else {
		pp.right = child
	}

	return true
}

func (bst *BST) InOrderTraverse(node *Node) {
	p := node
	if p.left != nil {
		bst.InOrderTraverse(p.left)
	}
	fmt.Printf("%+v ", p.data)
	if p.right != nil {
		bst.InOrderTraverse(p.right)
	}
}

func (bst *BST) PreOrderTraverse(node *Node) {
	p := node
	fmt.Printf("%+v ", p.data)
	if p.left != nil {
		bst.InOrderTraverse(p.left)
	}
	if p.right != nil {
		bst.InOrderTraverse(p.right)
	}
}

func (bst *BST) AfterOrderTraverse(node *Node) {
	p := node
	if p.left != nil {
		bst.InOrderTraverse(p.left)
	}
	if p.right != nil {
		bst.InOrderTraverse(p.right)
	}
	fmt.Printf("%+v ", p.data)
}
