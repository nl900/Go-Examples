package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)


// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	if (t.Left != nil) {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	
	if (t.Right != nil) {
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	if t1 == nil || t2 == nil {
		return false
	}
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := 0; i<10; i++ {
		x:= <- c1
		y:= <- c2
		if x!= y {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
		for i := 0; i < 10 ; i++ {
			fmt.Println(<-ch)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
