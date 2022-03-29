package ds

import "fmt"

type Node struct {
	number int
	next *int
}

func List() string {
	node := Node{
		number: 1,
		next: nil,
	}

	return fmt.Sprintf("%d", node.number)

}