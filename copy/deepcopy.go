package copy

import "github.com/jinzhu/copier"

func DeepCopy[T any](src *T) (*T, error) {
	target := new(T)
	if err := copier.Copy(target, src); err != nil {
		return nil, err
	}
	return target, nil
}
