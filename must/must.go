package must

func Return[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func Do(err error) {
	if err != nil {
		panic(err)
	}
}
