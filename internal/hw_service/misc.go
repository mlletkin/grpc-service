package hwservice

func mapper[A, B any](items []A, converter func(A) B) []B {
	newItems := []B{}
	for _, item := range items {
		newItems = append(newItems, converter(item))
	}
	return newItems
}
