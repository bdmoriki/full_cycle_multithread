package main

import "fmt"

type Usuario struct {
	nome  string
	idade int
}

func main() {
	canal := make(chan Usuario)

	go func() {
		canal <- Usuario{nome: "Bruno", idade: 30}
	}()

	usuario := <-canal
	fmt.Printf("O nome do usuário é %s e a idade é %d\n", usuario.nome, usuario.idade)
}
