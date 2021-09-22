package pointerplay

func Double(input *int) {
	*input *= 2
}

type MyInt int

func (input *MyInt) Double() {
	*input *= 2
}
