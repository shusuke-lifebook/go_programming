package main

func do(i interface{}) {
	i.(type)
}

func main() {
	do("Mike")
}
