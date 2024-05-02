package inteiros

import "fmt"

func ExampleAdiciona() {
	soma := Adiciona(1, 5)
	fmt.Println(soma)
	//Output: 6
}

// Adiciona recebe dois inteiros e retorna a soma deles
func Adiciona(x, y int) int {
	return x + y
}
