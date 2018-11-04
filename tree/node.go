package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//定义接收者 给node接收的
func (node Node) Print() {
	fmt.Println(node.Value)
}

//func print(node treeNode) {
//	fmt.Println(node.value)
//}
//func (node treeNode) setValue(value int) {
//	node.value = value
//}
//使用指针作为方法接收者
//只有使用指针才可以改变结构内容
//nil指针也可以调用方法
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil " + "node. Ignored")
		return
	}
	node.Value = value
}

//遍历函数
//func (node *Node) Traverse() {
//	if node == nil {
//		return
//	}
//	node.Left.Traverse()
//	node.Print()
//	node.Right.Traverse()
//}
func CreateNode(value int) *Node {
	return &Node{Value: value} //局部变量也可以返回给别人用
}

//值接收者VS指针接收者
//要改变内容必须使用指针接收者
//结构过大也考虑使用指针接收者
//一致性，如有指针接收者，最好都是指针接收者
//值接收者 是go语言特有的
//值/指针接收者均可以接收值/指针
