package sorting

import (
	"fmt"
	"strconv"
)

func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		num, ok := strconv.Atoi(fnb.Value())

		if ok != nil {
			return 0
		}

		return num
	}

	return 0
}

func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	fn := ExtractFancyNumber(fnb)

	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(fn))
}

func DescribeAnything(i interface{}) string {
	switch val := i.(type) {
	case int:
		return DescribeNumber(float64(val))
	case float64:
		return DescribeNumber(val)
	case NumberBox:
		return DescribeNumberBox(val)
	case FancyNumberBox:
		return DescribeFancyNumberBox(val)
	default:
		return "Return to sender"
	}
}
