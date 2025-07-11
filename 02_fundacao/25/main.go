package main

// if e switch, há o select tbm
func main() {
	a := 8
	b := 2
	c := 3

	// no GO não existe else IF
	// ==, >, <, <= ou >=
	if a > b {
		println(a)
	} else {
		println(b)
	}

	// para usar o or usamos ||
	if a > b && c > a {
		println("a >b && c> a")
	}

	//switch
	switch a {
	case 1:
		println(a)
	case 2:
		println(b)
	case 3:
		println(c)
	default:
		println("d")
	}

}
