package main

// condicionais
func main() {
	//if
	a := 10
	b := 2
	c := 3

	if a > b && c > b {
		println(a)
	}

	//não existe else if
	if a > b {
		println(true)
	} else {
		println(false)
	}

	//switch
	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("d")
	}

}
