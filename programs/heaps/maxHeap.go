package heaps

import "fmt"

// MaxHeap se
type MaxHeap struct {
	Heap []int
}

func (mh *MaxHeap) heapify(length int) {
	fmt.Println("Length to heapify: ", length)
	for i := length - 1; i >= 0; i-- {
		child := mh.Heap[i]
		parent := mh.Heap[i/2]
		if child > parent {
			mh.Heap[i] = parent
			mh.Heap[i/2] = child
		}
	}
}

// CreateHeap se
func (mh *MaxHeap) CreateHeap(input []int) {
	for _, element := range input {
		mh.Heap = append(mh.Heap, element)
		mh.heapify(len(mh.Heap))
	}
}

// Insert se
func (mh *MaxHeap) Insert(input int) {
	mh.Heap = append(mh.Heap, input)
	mh.heapify(len(mh.Heap))
}

// Delete se
func (mh *MaxHeap) Delete() (int, error) {
	if len(mh.Heap) == 0 {
		return -1, fmt.Errorf("Heap empty")
	}
	deleteElement := mh.Heap[0]
	mh.Heap = mh.Heap[1:len(mh.Heap)]
	mh.heapify(len(mh.Heap))
	return deleteElement, nil
}

// HeapSort se
func (mh *MaxHeap) HeapSort() {
	heapLen := len(mh.Heap)
	for i := 0; i < heapLen; i++ {
		deletedItem := mh.Heap[0]
		mh.Heap = mh.Heap[1:heapLen]

		mh.heapify(heapLen - i - 1)
		fmt.Println("Deleted item: ", deletedItem)
		mh.Heap = append(mh.Heap, deletedItem)
		fmt.Println("New heap: ", mh.Heap)
	}
	for i := 0; i < len(mh.Heap)/2; i++ {
		mh.Heap[i], mh.Heap[len(mh.Heap)-1-i] = mh.Heap[len(mh.Heap)-1-i], mh.Heap[i]
	}
}

// Run se
func Run() {
	mh := &MaxHeap{}
	// input := []int{10, 5, 20, 1, 25, 15}
	input := []int{0, 25, 50, 6, 12, 1, 3, 100}
	mh.CreateHeap(input)
	fmt.Println(mh.Heap)
	//mh.Delete()
	//fmt.Println(mh.Heap)
	// mh.Insert(25)
	// fmt.Println(mh.Heap)
	mh.HeapSort()
	fmt.Println(mh.Heap)
}
