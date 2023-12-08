package main

import (
	"fmt"
	algorithms "github.com/seidu626/playground/algorithims"
	"github.com/seidu626/playground/rss"
)

type product struct {
	ID    int32 `json:"int,omitempty"`
	Name  string
	Price float32
}

func (product product) Cost(quantity int32) (float32, error) {
	if quantity <= 0 {
		return 0, nil
	}
	return product.Price * float32(quantity), nil
}

func main() {

}

func __main() {
	rss.Rss()
	return
	// _ = algorithms.AddStrings("1234", "35644")

	data := []int{7, 6, 1, 2, 3, 4, 5}
	algorithms.SelectionSort(&data)
	fmt.Printf("SelectionSort: %d\n", data)
	return
	str := []byte("hello world")
	//var result []byte
	algorithms.ReverseString(str)

	fmt.Printf("Reverse: %s\n", str)
	return
	products := []product{
		{ID: 1, Name: "Cranberry", Price: 34.3},
		{ID: 2, Name: "Pineapple", Price: 3.5},
		{ID: 3, Name: "Short Bread", Price: 20.3},
	}

	for i := 0; i <= len(products)-1; i++ {
		prod := products[i]
		cost, _ := prod.Cost(3)
		fmt.Printf("Name: %s - Cost: %f \n", prod.Name, cost)
	}

	fmt.Println("Done!")
}
