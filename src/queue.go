package main

type Queue struct {
	data []Car
}

func (queue *Queue) push(element Car) {
	queue.data = append(queue.data, element)
}

func (queue *Queue) pop() {
	copy(queue.data, queue.data[1:])
}

func (queue *Queue) front() Car {
	var result = queue.data[0]
	queue.pop()
	return result
}

func (queue Queue) empty() bool {
	return len(queue.data) == 0
}
