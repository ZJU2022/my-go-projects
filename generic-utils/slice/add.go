package slice

import "my-go-projects/generic-utils/internal/slice"

func Add[Src any](src []Src, index int, element Src) ([]Src, error) {
	res, err := slice.Add[Src](src, element, index)
	return res, err
}
