package main

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
	"sync"
)

// Item interface{}
type Item generic.Type

// 节点结构体
type Node struct {
	key   int   //key,根据key进行节点插入
	value Item  //Item interface{}
	left  *Node //left
	right *Node //right
}

// 二叉树结构体
type ItemBinarySearchTree struct {
	root *Node        // 根节点
	lock sync.RWMutex //读写锁
}

/*

------------------------------------------------
                     ---[ 1
              ---[ 2
                     ---[ 3
       ---[ 4
                     ---[ 5
              ---[ 6
                     ---[ 7
---[ 8
              ---[ 9
       ---[ 10
              ---[ 11
------------------------------------------------

*/

// 格式化输出二叉搜索树
func (tree *ItemBinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	fmt.Println("------------------------------------------------")
	stringify(tree.root, 0)
	fmt.Println("------------------------------------------------")
}

// 通过递归进行格式化输出
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++

		stringify(n.left, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.right, level)

	}
}

// 构建ItemBinarySearchTree搜索二叉树插入节点的方法
// 有两个参数，一个是节点的key，另一个是value
func (tree *ItemBinarySearchTree) Insert(key int, value Item) {

	// 上锁
	tree.lock.Lock()
	// 解锁
	defer tree.lock.Unlock()

	// 创建一个节点
	node := &Node{key, value, nil, nil}

	if tree.root == nil {
		tree.root = node
	} else {
		// 以根节点为入口，插入新节点node
		insertNode(tree.root, node)
	}
}

func insertNode(node, newNode *Node) {
	// 判断节点插入到左边还是右边
	// 如果新节点的key小于根节点的key，插入到左边，否则插入到右边
	if newNode.key < node.key {

		if node.left == nil {
			node.left = newNode
		} else {
			// 通过递归插入新节点到左边
			insertNode(node.left, newNode)
		}

	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			// 通过递归插入新节点到右边
			insertNode(node.right, newNode)
		}

	}
}

// 搜索二叉树中最小的值
func (tree *ItemBinarySearchTree) Min() *Item {
	// 读写锁上锁
	tree.lock.RLock()
	// 读写锁解锁
	defer tree.lock.RUnlock()

	// 取出二叉树根节点
	node := tree.root

	if node == nil {
		return nil
	}
	// 最小的值一定在左边，所以只需要处理左边的节点即可
	// while
	for {
		if node.left == nil {
			return &node.value
		}
		node = node.left
	}
}

// 搜索二叉树中最小的值
func (tree *ItemBinarySearchTree) Max() *Item {
	// 读写锁上锁
	tree.lock.RLock()
	// 读写锁解锁
	defer tree.lock.RUnlock()

	// 取出二叉树根节点
	node := tree.root

	if node == nil {
		return nil
	}
	// 最小的值一定在左边，所以只需要处理左边的节点即可
	// while
	for {
		if node.right == nil {
			return &node.value
		}
		node = node.right
	}
}

// 传入某个key，查询这个key是否存在于二叉树中
func (tree *ItemBinarySearchTree) Search(key int) bool {
	// 上锁
	tree.lock.RLock()
	// 解锁
	defer tree.lock.RUnlock()
	// 搜索，返回搜索的结果
	return search(tree.root, key)
}

// 10

//13
/*
------------------------------------------------
                     ---[ 1
              ---[ 2
                     ---[ 3
       ---[ 4
                     ---[ 5
              ---[ 6
                     ---[ 7
---[ 8
              ---[ 9
       ---[ 10
              ---[ 11
------------------------------------------------
*/
func search(n *Node, key int) bool {
	if n == nil {
		return false
	}

	if key < n.key {
		// 如果传入的key小于节点的key，往左边搜索
		return search(n.left, key)
	}

	if key > n.key {
		// 如果传入的key大于节点的key，往右边边搜索
		return search(n.right, key)
	}
	// 两个key相等返回true
	return true
}

// 中序遍历
func (bst *ItemBinarySearchTree) InOrderTraverse(f func(Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	inOrderTraverse(bst.root, f)
}

func inOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		inOrderTraverse(n.left, f)  // L
		f(n.value)                  // D
		inOrderTraverse(n.right, f) // R
	}
}

// 后序遍历
func (bst *ItemBinarySearchTree) postOrderTraverse(f func(Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	postOrderTraverse(bst.root, f)
}

func postOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		postOrderTraverse(n.left, f)  // L
		postOrderTraverse(n.right, f) // R
		f(n.value)                    // D
	}
}

// 前序遍历
func (bst *ItemBinarySearchTree) preOrderTraverse(f func(Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()

	preOrderTraverse(bst.root, f)
}

func preOrderTraverse(n *Node, f func(Item)) {
	if n != nil {
		f(n.value)                   // D
		preOrderTraverse(n.left, f)  // L
		preOrderTraverse(n.right, f) // R
	}
}

func main() {
	// 创建二叉搜索树对象
	// 0xc4200a4020
	var tree ItemBinarySearchTree
	// 打印tree
	//fmt.Println(tree)
	//fmt.Printf("%p",&tree)
	// 插入第一个节点
	tree.Insert(8, "8")
	tree.Insert(4, "4")
	tree.Insert(10, "10")
	tree.Insert(2, "2")
	tree.Insert(6, "6")
	tree.Insert(1, "1")
	tree.Insert(3, "3")
	tree.Insert(5, "5")
	tree.Insert(7, "7")
	tree.Insert(9, "9")
	tree.Insert(11, "11")

	tree.String()

	fmt.Printf("二叉树key最小的节点的value值：%v\n", *tree.Min())
	fmt.Printf("二叉树key最大的节点的value值：%v\n", *tree.Max())

	fmt.Println(tree.Search(13))

	f := func(i Item) {
		fmt.Println(i)
	}
	fmt.Println("---------------中序遍历--------------------")
	tree.InOrderTraverse(f)
	fmt.Println("---------------前序遍历--------------------")
	tree.preOrderTraverse(f)
	fmt.Println("---------------后序遍历--------------------")
	tree.postOrderTraverse(f)

}
