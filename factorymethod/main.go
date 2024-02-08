package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypay struct {
}

func (p Paypay) Pay() {
	fmt.Println("Pagado por Paypal")
}

type Cash struct{}

func (c Cash) Pay() {
	fmt.Println("Pagado por Efectivo")
}

type CreditCar struct{}

func (c CreditCar) Pay() {
	fmt.Println("Pagado por Tarjeta de Crédito")
}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypay{}
	case 2:
		return Cash{}
	case 3:
		return CreditCar{}
	default:
		return nil
	}
}

func main() {
	var method uint
	fmt.Println("Ingrese un numero de metodo de pago:")
	fmt.Println("\t 1-paypal 2-efectivo 3-tarjeta de credito ")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("debe digitar un metodo valido")
	}

	if method > 3 {
		panic("debe digitar un metodo valido")
	}
	payMethod := Factory(method)
	payMethod.Pay()
}
