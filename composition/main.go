package main

import (
	"fmt"

	"github.com/alvarollopi/composition/pkg/customer"
	"github.com/alvarollopi/composition/pkg/invoice"
	"github.com/alvarollopi/composition/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Colombia",
		"Popayán",
		customer.New("Alejandro", "Cl 123 # 23-4", "12345"),
		invoiceitem.NewItems(
			invoiceitem.New(1, "Curso de Go", 12.34),
			invoiceitem.New(1, "Curso de POO con Go", 54.23),
			invoiceitem.New(1, "Curso de Testing con Go", 90.00),
		),
	)
	i.SetTotal()
	fmt.Printf("%+v", i)
}
