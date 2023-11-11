// Package: steps provides utility functions for
// algocli application
package steps

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
	DataName      string
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
