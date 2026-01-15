package d2

import "fmt"

type ContaBancaria struct{
	saldo float64
}

func (c *ContaBancaria) Depositar (valorDeposito float64){
	if valorDeposito > 0 && valorDeposito <10000 {
		c.saldo += valorDeposito
	}else{
		fmt.Println("Deposito com erro X")
	}

}

func (c *ContaBancaria) ExibirSaldo (){
	fmt.Println("Saldo Atual:", c.saldo)
}