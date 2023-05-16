package sort

import (
	"image/gif"

	"github.com/invzhi/sorting-visualization/animation"
)

// SelectionSort will record a frame for every selection.
func SelectionSort(a []uint8, y int, g *gif.GIF) {
	frame := 1
	l := len(a)

	for i := 0; i < l-1; i++ {
		minIndex := i
		for j := i + 1; j < l; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[minIndex], a[i] = a[i], a[minIndex]

		animation.SetLine(g, y, frame, a)
		frame++
	}
}
