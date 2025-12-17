package main

func Add(a, b int32) int32 {
	return a + b
}

type Adder struct{ id int32 }

func (adder *Adder) AddPtr(a, b int32) int32 {
	return a + b
}

func (adder *Adder) AddVal(a, b int32) int32 {
	return a + b
}

func main() {
	Add(10, 32)

	adder := Adder{id: 6784}

	adder.AddPtr(10, 32)
	adder.AddVal(10, 32)

	(&adder).AddVal(10, 32)
}
