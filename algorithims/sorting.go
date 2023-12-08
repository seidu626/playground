package algorithms

func SelectionSort(elements *[]int) {
	for i := 0; i <= len(*elements) - 1; i ++ {
		x :=  i
		y := -1

		for j := i; j <= len(*elements) -2; j ++ {
			// get the smallest element
			if (*elements)[x] > (*elements)[j + 1] {
				y = i
				x = j + 1
			}
		}

		if y != -1 {
			Swap(elements, x, y)
		}

	}
}

func BubbleSort(elements *[]int) {
	skipCount := 0
	for i := 0; i <= len(*elements) - 1; i ++ {
		for j := i; j <= len(*elements) -2; j ++ {
			if (*elements)[i] > (*elements)[j + 1] {
				Swap(elements, i, j + 1)
				skipCount = 0
			} else {
				skipCount ++
			}
		}

		if skipCount + 1 == len(*elements) {
			break
		}
	}
}

// Swap / https://medium.com/swlh/golang-tips-why-pointers-to-slices-are-useful-and-how-ignoring-them-can-lead-to-tricky-bugs-cac90f72e77b
func Swap(elements *[]int, a int, b int) {
	(*elements)[a], (*elements)[b] = (*elements)[b], (*elements)[a]
}
