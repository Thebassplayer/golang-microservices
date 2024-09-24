package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWostCase(t *testing.T) {

	//Initialization
	elements := []int{9, 8, 7, 6, 5}

	//Execution
	BubbleSort(elements)

	//Validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 8, elements[3])
	assert.EqualValues(t, 9, elements[4])
}
func TestBubbleSortBestCase(t *testing.T) {

	//Initialization
	elements := []int{5, 6, 7, 8, 9}

	//Execution
	BubbleSort(elements)

	//Validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 8, elements[3])
	assert.EqualValues(t, 9, elements[4])
}
func TestBubbleSortNilSlice(t *testing.T) {

	//Initialization

	//Execution
	BubbleSort(nil)

	//Validation
}

func getElements(n int) []int {
	result := make([]int, n)

	for j := 0; j < n; j++ {
		result[j] = n - j
	}

	return result
}

func TestGetElements(t *testing.T) {
	numOfElements := 5
	elements := getElements(numOfElements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, numOfElements, len(elements))
	assert.EqualValues(t, numOfElements, elements[0])
	assert.EqualValues(t, numOfElements-1, elements[1])
	assert.EqualValues(t, numOfElements-2, elements[2])
	assert.EqualValues(t, numOfElements-3, elements[3])
	assert.EqualValues(t, numOfElements-4, elements[4])
}

func BenchmarkBubbleSort10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}
func BenchmarkSort10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
func BenchmarkBubbleSort1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}
func BenchmarkSort1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}
func BenchmarkSort100000(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
