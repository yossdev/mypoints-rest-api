package repositories

import (
	"github.com/spf13/viper"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"log"
	"strconv"
	"time"
)

type BodyReq struct {
	Name  string
	Email string
	Value float64
	Title string
	Desc  string // description from agent
}

func CreateInvoice(body BodyReq) (*xendit.Invoice, *xendit.Error) {
	secret := viper.GetString("MYPOINTSPG")
	xendit.Opt.SecretKey = secret

	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	cust := xendit.InvoiceCustomer{
		GivenNames: body.Name,
		Email:      body.Email,
	}

	items := []xendit.InvoiceItem{
		{
			Name:     body.Title,
			Price:    body.Value,
			Quantity: 1,
		},
	}

	//fees := []xendit.InvoiceFee{
	//	{
	//		Type:  "admin",
	//		Value: 5000, // problem with library
	//	},
	//}

	const adminFee float64 = 5000

	data := invoice.CreateParams{
		ExternalID:  "INV-" + timestamp,
		Amount:      body.Value + adminFee,
		PayerEmail:  body.Email,
		Description: body.Desc,
		Customer:    cust,
		Items:       items,
		//Fees:        fees,
		Currency: "IDR",
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		log.Fatal(err)

	}

	return resp, err
}
