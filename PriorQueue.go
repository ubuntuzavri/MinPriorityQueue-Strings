package PriorQueue

import (
	"errors"
	"fmt"
)

type queue []string

var size int
var (
	errorEmptyQueue = errors.New("Queue Is Empty!")
)

func NewQueue() queue {

	size = 0
	q := make(queue, 1, 1)
	return q
}
func NewQueueWithSlice(elements []string) queue {
	if len(elements) > 0 {
		size = 0
		q := make(queue, 1, 1)
		for _, tmp := range elements {
			q.Enqueue(tmp)
		}
		
		return q
	} else {
		return NewQueue()
	}

}

func (q queue) Size() int {
	return size
}

func (q queue) IsEmpty() bool {
	return (size == 0)
}

func (q *queue) Enqueue(value string) {
	//fmt.Println(cap(*q))
	if size+1 == cap(*q) {
		newQueue := make(queue, cap(*q)*2, cap(*q)*2)
		for i := 0; i < size; i++ {
			newQueue[i] = (*q)[i]
		}
		*q = newQueue
	}

	(*q)[size] = value
	size++
	k := size

	for k > 1 {

		if value < (*q)[k/2-1] {
			(*q)[k-1], (*q)[k/2-1] = (*q)[k/2-1], (*q)[k-1]

		}
		k /= 2
	}

}

func (q *queue) DequeueMin() (string, error) {
	if size == 0 {
		return "", fmt.Errorf("Error %w", errorEmptyQueue)
	}
	s := (*q)[0]
	(*q)[0], (*q)[size-1] = (*q)[size-1], (*q)[0]
	size -= 1
	index := 1
	var leftChild, rightChild int = 0, 0
	//fmt.Println(q)
	for true {

		leftChild = 2*index - 1
		rightChild = 2 * index

		if rightChild >= size || ((*q)[leftChild] >= (*q)[index-1] && (*q)[rightChild] >= (*q)[index-1]) {

			break
		}

		if (*q)[index-1] > (*q)[leftChild] && (*q)[leftChild] < (*q)[rightChild] {
			(*q)[index-1], (*q)[leftChild] = (*q)[leftChild], (*q)[index-1]
			index = leftChild + 1
		} else {

			(*q)[index-1], (*q)[rightChild] = (*q)[rightChild], (*q)[index-1]
			index = rightChild + 1

		}
		if index >= size {
			break
		}

	}
	return s, nil
}
