package invoice

import (
	"github.com/alvarollopi/composition/pkg/customer"
	"github.com/alvarollopi/composition/pkg/invoiceitem"
)

// Invoice es la estructura de una factura(invoice)
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   invoiceitem.Items
}

// Retorna una nueva Factura(Invoice)
func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// Settotal es un setter de Invoice.total
func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}
