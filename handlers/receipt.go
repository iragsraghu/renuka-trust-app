package handlers

import (
	"fmt"
	"net/http"

	"github.com/iragsraghu/renuka-trust-app/config"
	"github.com/iragsraghu/renuka-trust-app/helper"
	"github.com/iragsraghu/renuka-trust-app/models"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func GenerateReceipt(c *gin.Context) {

	receiptNumber := c.Param("receipt_number")

	var donation models.Donation

	err := config.DB.
		Preload("Donor").
		Preload("Donor.Village").
		Where("receipt_number = ?", receiptNumber).
		First(&donation).Error

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Receipt not found",
		})

		return
	}

	// PDF creation starts here

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Background template
	pdf.Image(
		"assets/logo4.png",
		0,
		0,
		210,
		297,
		false,
		"",
		0,
		"",
	)

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "B", 12)

	// Receipt Number
	pdf.SetXY(60, 87)
	pdf.Cell(50, 8, donation.ReceiptNumber)

	// Date
	pdf.SetXY(150, 87)
	pdf.Cell(30, 8, donation.CreatedAt.Format("02-01-2006"))

	// Donor Name
	pdf.SetFont("Arial", "B", 12)
	pdf.SetXY(82, 115)
	pdf.Cell(80, 8, donation.Donor.Name)

	// Mobile
	pdf.SetXY(82, 127)
	pdf.Cell(80, 8, donation.Donor.Mobile)

	// Village
	pdf.SetXY(82, 139)
	pdf.Cell(80, 8, donation.Donor.Village.Name)

	// Address
	pdf.SetXY(82, 150)
	pdf.Cell(90, 8, donation.Donor.Address)

	// Amount
	pdf.SetFont("Arial", "B", 15)
	pdf.SetTextColor(180, 0, 0)

	pdf.SetXY(82, 177)
	pdf.Cell(
		60,
		8,
		fmt.Sprintf("Rs. %.2f", donation.Amount),
	)

	// Payment Mode
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "B", 12)

	pdf.SetXY(82, 189)
	pdf.Cell(80, 8, donation.PaymentMode)

	// Purpose
	pdf.SetXY(82, 200)
	pdf.Cell(80, 8, donation.Purpose)

	// Amount In Words
	pdf.SetFont("Times", "BI", 15)

	pdf.SetXY(40, 224)
	pdf.Cell(
		120,
		8,
		helper.RupeesInWords(donation.Amount),
	)

	filename := "receipt_" + donation.ReceiptNumber + ".pdf"

	c.Header("Content-Type", "application/pdf")
	c.Header(
		"Content-Disposition",
		"attachment; filename="+filename,
	)

	err = pdf.Output(c.Writer)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}
}
