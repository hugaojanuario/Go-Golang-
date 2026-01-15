package d5

import (
	"fmt"
)

type Animal interface{
	Falar () string
}

type Cachorro struct{
	Nome string
}

func (c *Cachorro) Falar () string{
	return "Au Au"
}

type Gato struct{
	Nome string
}

func (g *Gato) Falar () string{
	return "Miau"
}

func EmitirSom (a Animal){
	fmt.Println(a.Falar())
}