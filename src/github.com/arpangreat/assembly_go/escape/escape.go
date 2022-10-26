package escape

type Addifier interface{ Add(a, b int32) int32 }

type Adder struct { name string }

func ( adder Adder ) Add(a, b int32) int32 { return a + b }

func main() {
  adder := Adder{name: "myAdder"}
  adder.Add(10, 32)
  Addifier(adder).Add(10, 32)

  panic("Hello")
}
