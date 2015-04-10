package newmath

import (
	"fmt"
	"image"
)

func ImageWiter() {
	m := image.NewNRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
