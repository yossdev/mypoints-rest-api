package repositories

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"log"
	"time"
)

func CreateInvoice() {
	secret := viper.GetString("MYPOINTSPG")
	xendit.Opt.SecretKey = secret

	data := invoice.CreateParams{
		ExternalID:  "invoice-" + time.Now().String(),
		Amount:      200000,
		PayerEmail:  "customer@customer.com",
		Description: "invoice  #1",
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("created invoice: %+v\n", resp)
}
