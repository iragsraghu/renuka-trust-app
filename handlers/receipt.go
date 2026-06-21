package handlers

import (
	"fmt"
	"net/http"

	"github.com/iragsraghu/renuka-trust-app/config"
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

	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(190, 10, "Shri Renuka Yellamma Devi Trust")

	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)

	pdf.Cell(190, 10, "Donation Receipt")

	pdf.Ln(15)

	pdf.Cell(190, 10, "Receipt No: "+donation.ReceiptNumber)

	pdf.Ln(10)

	pdf.Cell(190, 10, "Donor Name: "+donation.Donor.Name)

	pdf.Ln(10)

	pdf.Cell(190, 10, "Mobile: "+donation.Donor.Mobile)

	pdf.Ln(10)

	pdf.Cell(190, 10, "Village: "+donation.Donor.Village.Name)

	pdf.Ln(10)

	pdf.Cell(190, 10, "Amount: Rs. "+fmt.Sprintf("%.2f", donation.Amount))

	pdf.Ln(10)

	pdf.Cell(190, 10, "Purpose: "+donation.Purpose)

	pdf.Ln(10)

	pdf.Cell(190, 10, "Payment Mode: "+donation.PaymentMode)

	pdf.Ln(20)

	pdf.Cell(190, 10, "Thank You For Your Donation")

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
