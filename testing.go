package timber

func Reset() []options {
	writers = []options{}

	return writers
}

func Writers() []options {
	return writers
}
