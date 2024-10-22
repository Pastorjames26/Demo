package main

import "fmt"

// Define the structure for a node in the BST
type BST struct {
	Value int
	Left  *BST
	Right *BST
}

// Insert a new value into the BST and return the root of the updated tree
func InsertIntoBST(root *BST, value int) *BST {
	// If the tree is empty, create a new node and return it
	if root == nil {
		return &BST{Value: value}
	}

	// Traverse the tree to find the correct position for the new value
	if value < root.Value {
		root.Left = InsertIntoBST(root.Left, value)
	} else if value > root.Value {
		root.Right = InsertIntoBST(root.Right, value)
	}

	//Return the root of the updated tree
	return root
}

// Helper function to print the BST in order (sorted order)
func InOrderTraversal(root *BST) {
	if root != nil {
		InOrderTraversal(root.Left)
		fmt.Print(root.Value, " ")
		InOrderTraversal(root.Right)
	}
}

func main() {
	// Create the BST based on the given example
	/*
			4
		   / \
		   2  5
		  / \
		 1   3
	*/
	root := &BST{Value: 4}
	root.Left = &BST{Value: 2}
	root.Right = &BST{Value: 5}
	root.Left.Left = &BST{Value: 1}
	root.Left.Right = &BST{Value: 3}

	// Insert the value 6 into the BST
	root = InsertIntoBST(root, 6)

	// Output the tree (in-ordertraversal should show values in sorted order)
	fmt.Print("In-order Traversal of the updated BST: ")
	InOrderTraversal(root)
}
