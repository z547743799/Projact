package link

import "fmt"

var head *node
var low *node

type node struct {
	key     string
	private *node
	netx    *node
}

func Add(s string) {
	node := new(node)
	if head == nil {
		node.key = s
		node.netx = nil
		head = node
		low = node
	} else {
		node.key = s
		low.netx = node
		low = node
	}
}

func Bianli() {
	node := new(node)
	node = head
	if head != nil {
		for {
			if node != nil {
				s := node.key
				fmt.Println(s)
				node = node.netx
			} else {
				break
			}

		}
	}

}
