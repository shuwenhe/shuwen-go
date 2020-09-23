package main

type Integer int

func main() {
	var a Integer = 1
	a.Add(2)
}

func (a Integer) Add(b Integer) {
	a += b
}
