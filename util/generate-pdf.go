package util

import (
	"eve/database"
	"eve/models"
	"fmt"
	"time"

	generator "github.com/angelodlfrtr/go-invoice-generator"
)

func GeneratePdf(event models.Event) error {

	var organizer models.User

	err := database.Database.Db.First(&organizer, event.OrganizerID)

	if err != nil {
		return err.Error
	}

	doc, _ := generator.New(generator.Invoice, &generator.Options{
		TextTypeInvoice: "FACTURE",
		AutoPrint:       true,
	})

	doc.SetHeader(&generator.HeaderFooter{
		Text:       "<center>Receipt from Eve",
		Pagination: true,
	})

	doc.SetFooter(&generator.HeaderFooter{
		Text:       "<center>Receipt from Eve",
		Pagination: true,
	})

	doc.SetRef("random")
	doc.SetVersion("1.0")

	doc.SetDescription(fmt.Sprintf("An invoice for the purchase of tickets for %s", event.Name))
	doc.SetDate(fmt.Sprintf("%d/%d/%d", time.Now().Day(), time.Now().Month(), time.Now().Year()))

	doc.SetCompany(&generator.Contact{
		Name: organizer.Username,
		Address: &generator.Address{
			Address:    "123 test str",
			Address2:   "Apartment 2",
			PostalCode: "12345",
			City:       "Test",
			Country:    "Test",
		},
	})

	doc.SetCustomer(&generator.Contact{
		Name: organizer.Username,
		Address: &generator.Address{
			Address:    "123 test str",
			Address2:   "Apartment 2",
			PostalCode: "12345",
			City:       "Test",
			Country:    "Test",
		},
	})

	doc.AppendItem(&generator.Item{
		Name:     event.Name,
		UnitCost: fmt.Sprintf("%d", event.Price),
		Quantity: fmt.Sprintf("%d", event.Tickets),
		Tax: &generator.Tax{
			Amount: "0",
		},
		Discount: &generator.Discount{
			Percent: "0",
		},
	})

	// pdf, err = doc.Build()
	// if err != nil {
	// 	return err.Error
	// }

	// err = pdf.OutputFileAndClose("out.pdf")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return nil
}
