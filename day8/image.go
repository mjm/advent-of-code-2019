package day8

type Image struct {
	Width  int
	Height int
	Layers [][]int
}

func ImageFromString(s string, w, h int) (*Image, error) {
	img := &Image{
		Width:  w,
		Height: h,
	}

	layerLen := w * h
	for i := 0; i < len(s); i += layerLen {
		layerStr := s[i : i+layerLen]
		nums := make([]int, 0, len(layerStr))
		for _, c := range layerStr {
			nums = append(nums, int(c)-48)
		}
		img.Layers = append(img.Layers, nums)
	}

	return img, nil
}

func (img *Image) DigitCounts() []map[int]int {
	var counts []map[int]int
	for _, layer := range img.Layers {
		layerCounts := make(map[int]int)
		for _, digit := range layer {
			layerCounts[digit]++
		}
		counts = append(counts, layerCounts)
	}
	return counts
}

func (img *Image) Composite() []int {
	ret := make([]int, img.Width*img.Height)
	copy(ret, img.Layers[0])

	for _, layer := range img.Layers[1:] {
		for i, color := range layer {
			if ret[i] == 2 {
				ret[i] = color
			}
		}
	}

	return ret
}
