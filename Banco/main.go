package main

import (
	"banco/contas"
	"fmt"
	//"banco/clientes"
)


func main(){
	contaDoPec := contas.ContaCorrente{}
	contaDoPec.Depositar(300)
	PagarBoleto(&contaDoPec, 30)
	fmt.Println("Boleto pago com sucesso.")
	contaDoPec.Sacar(100)
	contaDoPec.Depositar(340)
	fmt.Println("Deposito realizado com sucesso")
	fmt.Println("Saldo restante:", contaDoPec.ObterSaldo())

	contaDoVeggeti := contas.ContaPoupanca{}
	contaDoVeggeti.Depositar(500)
	fmt.Println("Boleto pago com sucesso.")
	PagarBoleto(&contaDoVeggeti, 50)
	fmt.Println("Saldo restante:",contaDoVeggeti.ObterSaldo())


}

func PagarBoleto(contas verificarConta, valorDoBoleto float64) {
	contas.Sacar(valorDoBoleto)
}  

type verificarConta interface {
	Sacar(valor float64) (string, float64)
}