package d7

import (
	"fmt"
)

type Veiculo interface {
	VelocidadeMaxima() int
}

type Motor struct {
	Potencia int
}

type Carro struct {
	Motor
}

func (c *Carro) VelocidadeMaxima() int {
	velocidadeCarro := c.Potencia * 3
	fmt.Println(velocidadeCarro)
	return c.Potencia * 3
}
