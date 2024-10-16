package main

import (
	"fmt"
	"sync"
)

type Product struct {
	Name string
	Len  int
	Wid  int
	Hei  int
}

type Box struct {
	Len int
	Wid int
	Hei int
}

type Result struct {
	box    Box
	volume int
}

// Helper function to check if a box can fit a single product
func canFit(box Box, product Product) bool {
	return box.Len >= product.Len && box.Wid >= product.Wid && box.Hei >= product.Hei
}

// Helper function to calculate the volume of a box
func volume(box Box) int {
	return box.Len * box.Wid * box.Hei
}

// Function to check if a box can fit all products in the list
func canFitAll(box Box, products []Product) bool {
	for _, product := range products {
		if !canFit(box, product) {
			return false
		}
	}
	return true
}

// Worker function that runs in a goroutine and checks if a box can fit all products
func checkBox(box Box, products []Product, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure the WaitGroup is decremented when the goroutine completes
	if canFitAll(box, products) {
		results <- Result{box: box, volume: volume(box)}
	}
}

// Concurrent version of getBestBox that works with multiple products
func getBestBox(availableBoxes []Box, products []Product) Box {
	results := make(chan Result, len(availableBoxes)) // Buffered channel to collect results
	var wg sync.WaitGroup                             // WaitGroup to wait for all goroutines

	// Launch a goroutine for each box
	for _, box := range availableBoxes {
		wg.Add(1)
		go checkBox(box, products, results, &wg)
	}

	// Close the results channel after all goroutines have finished
	go func() {
		wg.Wait()
		close(results)
	}()

	// Now collect results and find the best box
	var bestBox Box
	found := false

	for result := range results {
		if !found || result.volume < volume(bestBox) {
			bestBox = result.box
			found = true
		}
	}

	// If no box is found, return an empty Box (optional behavior)
	if !found {
		return Box{}
	}
	return bestBox
}

func main() {
	// Example of available boxes
	availableBoxes := []Box{
		{Len: 10, Wid: 10, Hei: 10},
		{Len: 15, Wid: 12, Hei: 8},
		{Len: 20, Wid: 20, Hei: 20},
		{Len: 25, Wid: 25, Hei: 25},
	}

	// Example of products
	products := []Product{
		{Name: "Toy1", Len: 20, Wid: 8, Hei: 8},
		{Name: "Toy2", Len: 9, Wid: 9, Hei: 9},
	}

	// Find the best box for all products
	bestBox := getBestBox(availableBoxes, products)

	// Print result
	if bestBox.Len == 0 && bestBox.Wid == 0 && bestBox.Hei == 0 {
		fmt.Println("No suitable box found.")
	} else {
		fmt.Printf("Best box for all products: Len=%d, Wid=%d, Hei=%d\n", bestBox.Len, bestBox.Wid, bestBox.Hei)
	}
}
