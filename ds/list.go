package ds

import "fmt"

type Node struct {
	number int
	next *Node
}

func List() string {
	
	var list *Node
	n :=  &Node{
		number: 1,
		next: nil,
	}

	list = n

	n = &Node{
		number: 5,
		next: nil,
	}
	list.next = n

	for temp := list; temp != nil; temp = temp.next {
		fmt.Println(temp.number)
	}

	return ""


}