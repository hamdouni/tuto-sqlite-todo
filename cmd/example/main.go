package main

import (
	"fmt"
	"log"

	"github.com/hamdouni/tuto-sqlite-todo/task"
)

func main() {

	// First, let's have a new repository
	// and a way to teardown this repository
	store, teardown, err := getStore()
	if err != nil {
		log.Fatalf("could not get a store: %s", err)
	}
	defer teardown()

	// We inject this repository in our business layer
	// so our business know how to store and retrieve tasks
	task.Configure(&store)

	// Then we can use our business layer
	// without worrying about how the data is stored
	for i := 1; i < 10; i++ {
		t := fmt.Sprintf("faire %d café", i)
		if i > 1 {
			t = t + "s"
		}
		_, err := task.Create(t)
		if err != nil {
			log.Fatalf("could not create task: %s", err)
		}
	}

	// We can use our store to get business items
	items, err := store.GetAll()
	if err != nil {
		log.Fatalf("getting all tasks: %s", err)
	}
	for _, item := range items {
		fmt.Printf("%s\n", item)
	}
}
