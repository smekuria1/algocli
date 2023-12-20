// Package: steps provides utility functions for
// algocli application
package steps

import "github.com/smekuria1/algocli/cmd/userinput"

type StepSchema struct {
	StepName string
	Options  []Item
	Headers  string
	Field    *string
}

type Steps struct {
	Steps []StepSchema
}

// Item contains the name of the item
// in StepSchema.Options
type Item struct {
	Title, Desc string
}

type Options struct {
	AlgorithmType string
	AlgorithmName string
	DataType      string
	DataName      *userinput.Output
	DataSize      *userinput.Output
}

// InitSteps initializes the steps and returns a pointer to the steps
func InitSteps(option *Options) *Steps {
	steps := &Steps{
		[]StepSchema{
			{
				StepName: "Algorithm/Datastructure Selection",
				Options: []Item{
					{
						Title: "Binary Search Tree",
						Desc:  "Binary Search Tree is a data structure that allows for fast lookup, insertion, and deletion of items.",
					},
					{
						Title: "Queue",
						Desc:  "Queue is a data structure that allows for FIFO (First In First Out) operations.",
					},

					{
						Title: "Hashtable",
						Desc:  "Hashtable is a data structure that implements an associative array, also called a dictionary, which is an abstract data type that maps keys to values. ",
					},
				},
				Headers: "Select an option",
				Field:   &option.AlgorithmType,
			},
			// {
			// 	StepName: "Data Structure Selection",
			// 	Options: []Item{
			// 		{
			// 			Title: "Array",
			// 			Desc:  "Select an array",
			// 		},
			// 		{
			// 			Title: "Linked List",
			// 			Desc:  "Select a linked list",
			// 		},
			// 		{
			// 			Title: "Stack",
			// 			Desc:  "Select a stack",
			// 		},
			// 		{
			// 			Title: "Queue",
			// 			Desc:  "Select a queue",
			// 		},
			// 		{
			// 			Title: "Tree",
			// 			Desc:  "Select a tree",
			// 		},
			// 		{
			// 			Title: "Hash Table",
			// 			Desc:  "Select a hash table",
			// 		},
			// 	},
			// 	Headers: "Select a data structure",
			// },
		},
	}
	return steps
}

func InitSecondSteps(option *Options) *Steps {
	secondSteps := &Steps{
		[]StepSchema{
			{
				StepName: "Input Data Type Selection",
				Options: []Item{
					{
						Title: "Input Integers",
						Desc:  "Input Integers is a data type that allows the user to input a list of integers.",
					},
					{
						Title: "Random Integers",
						Desc:  "Random Integers is a data type that generates a list of random integers.",
					},
				},
				Headers: "Select an option",
				Field:   &option.DataType,
			},
		},
	}
	return secondSteps

}
