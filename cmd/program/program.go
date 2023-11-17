package program

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/smekuria1/GoSlow/binarysearchtree"
	"github.com/smekuria1/GoSlow/queue"
)

var sliceStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("30")).Bold(true)
var preorderStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("31")).Bold(true)
var inorderStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("32")).Bold(true)
var postorderStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("33")).Bold(true)

// Algorithm is the struct that holds the algorithm data to be run
type Algorithm struct {
	Name string
	Exit bool
	Data DataStructure
}

// DataStructure is the struct that holds the data structure data to be run
type DataStructure struct {
	Name  string
	Exit  bool
	Value string
	Size  int
}

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
	if a.Name == "Binary Search Tree" {
		if a.Data.Name == "Random Integers" {
			bst := binarysearchtree.NewBST[int]()
			randomSlice := make([]int, a.Data.Size)
			for i := 0; i < a.Data.Size; i++ {
				randint := rand.Intn(100)
				randomSlice[i] = randint
				bst.Add(randint)
			}
			fmt.Println(sliceStyle.Render(fmt.Sprintf("Random Slice: %v", randomSlice)))
			preorder := bst.PreOrderTraversal()
			inorder := bst.InOrderTraversal()
			postorder := bst.PostOrderTraversal()
			fmt.Println(preorderStyle.Render(fmt.Sprintf("PreOrder Traversal: %v", preorder)))
			fmt.Println(inorderStyle.Render(fmt.Sprintf("InOrder Traversal: %v", inorder)))
			fmt.Println(postorderStyle.Render(fmt.Sprintf("PostOrder Traversal: %v", postorder)))
			PrettyPrintTree(preorder)
		} else if a.Data.Name == "Input Integers" {
			bst := binarysearchtree.NewBST[int]()
			s := strings.Split(a.Data.Value, " ")
			for _, v := range s {
				i, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
				bst.Add(i)
			}
			preorder := bst.PreOrderTraversal()
			inorder := bst.InOrderTraversal()
			postorder := bst.PostOrderTraversal()
			fmt.Printf("PreOrder Traversal: %v\n", preorder)
			fmt.Printf("InOrder Traversal: %v\n", inorder)
			fmt.Printf("PostOrder Traversal: %v\n", postorder)
			PrettyPrintTree(preorder)
		}

	} else if a.Name == "Queue" {
		q := queue.NewQueue[int]()
		for i := 0; i < 10; i++ {
			randint := rand.Intn(100)
			q.Enqueue(randint)
		}
		PrettyPrinQueue(q)
	}
}
