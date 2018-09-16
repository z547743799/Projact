package linked

import "fmt"

type node struct {
	prixnode  *node
    nextnode  *node
    key string
}

var haed *node
var curr *node

func Add(ss string) {
	   node:=new(node)
	if haed == nil{
        node.key=ss
        node.nextnode=nil
		haed=node
		curr=node
	}else {
		node.key=ss
		curr.nextnode=node
		curr=node
	}
}
func Ergodic()  {
	node:=haed
	for  {
		fmt.Println(node.key)
		if node.nextnode==nil {
			break
		}
		node=node.nextnode
	}

}


