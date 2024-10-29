package main

// quando eu uso o * eu estou falando que estou apontando a referencia da memória
// passa a não ser mais uma cópia
func soma(a, b *int) int {
	//fez uma alteração na memória
	*a = 50
	return *a + *b
}

func main() {
	minhavar1 := 10
	minhavar2 := 20
	println(soma(&minhavar1, &minhavar2))
	println(minhavar1)
}
