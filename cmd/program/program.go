package program

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/smekuria1/GoSlow/binarysearchtree"
	"github.com/smekuria1/GoSlow/queue"
)

// Algorithm is the struct that holds the algorithm data to be run
type Algorithm struct {
	Name    string
	Exit    bool
	Package string
	Func    string
}

// DataStructure is the struct that holds the data structure data to be run
// //
// type DataStructure struct {
// 	Name    string
// 	Exit    bool
// 	Package string
// 	Func    string
// }

// ExitCli if true, exits the cli
func (a *Algorithm) ExitCli(tprogram *tea.Program) {
	if a.Exit {
		if err := tprogram.ReleaseTerminal(); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}

// Runs the chosen algorithm from the GoSlow package
func (a *Algorithm) Run() {
	if a.Name == "Binary Search" {
		bst := binarysearchtree.NewBST[int]()
		for i := 0; i < 10; i++ {
			randint := rand.Intn(100)
			bst.Add(randint)
		}
		// custom print function that pritns the tree in a nice format
		preorder := bst.LevelOrderTraversal()
		fmt.Printf("PreOrder Traversal: %v\n", preorder)
		PrettyPrintTree(preorder)
	} else if a.Name == "Queue" {
		q := queue.NewQueue[int]()
		for i := 0; i < 10; i++ {
			randint := rand.Intn(100)
			q.Enqueue(randint)
		}
		PrettyPrinQueue(q)
	}
}
