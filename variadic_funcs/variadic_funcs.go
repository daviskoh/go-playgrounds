package variadic_funcs

func Add(n ...int) int {
	var result int

	for _, num := range n {
		result += num
	}

	return result
}
