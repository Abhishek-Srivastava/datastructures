package queues

import "fmt"

// Q se
type Q struct {
	Items []int
}

// Insert se
func (q *Q) Insert(item int) {
	q.Items = append(q.Items, item)
}

// Delete se
func (q *Q) Delete() (int, error) {
	if len(q.Items) == 0 {
		return -1, fmt.Errorf("Queue empty")
	}
	item := q.Items[0]
	q.Items = q.Items[1:len(q.Items)]
	return item, nil
}

// PrintQ se
func (q *Q) PrintQ() {
	fmt.Println(q.Items)
}
