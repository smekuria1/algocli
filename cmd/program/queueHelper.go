package program

import (
	"fmt"
	"strings"

	"github.com/smekuria1/GoSlow/queue"
)

func PrettyPrinQueue(q *queue.Queue[int]) {
	fmt.Println("Queue:")
	fmt.Println(strings.Repeat("-", 10))
	fmt.Println("Front")
	for i := 0; i < q.Size(); i++ {
		fmt.Printf("%v\n", q.Dequeue())
	}
	fmt.Println("Back")
	fmt.Println(strings.Repeat("-", 10))
}
