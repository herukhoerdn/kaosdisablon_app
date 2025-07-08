package usecase

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/kaosdisablon/entity"
)

func (u *usecase) CreateInvoice(ctx context.Context, checkoutId, userId int64, totalHarga float64, fileInvoice string) (int64, error) {
	invoice := entity.Invoice{
		CheckoutId:  checkoutId,
		UserId:      userId,
		TotalHarga:  totalHarga,
		TanggalBuat: time.Now(),
		FileInvoice: "",
	}
	id, err := u.repo.CreateInvoice(ctx, invoice)
	if err != nil {
		return 0, err
	}
	fileName := fmt.Sprintf("invoice_%d.pdf", id)
	filePath := "assets/" + fileName
	invoice.Id = id
	if err := u.GenerateInvoicePDF(filePath, invoice); err != nil {
		return id, err
	}
	if err := u.repo.UpdateInvoiceFilePath(ctx, id, fileName); err != nil {
		return id, err
	}

	return id, nil
}

func (u *usecase) GetInvoicesByUser(ctx context.Context, userId int64) ([]entity.Invoice, error) {
	return u.repo.GetInvoicesByUser(ctx, userId)
}
func (u *usecase) GetCheckoutByDesainID(ctx context.Context, desainId int64) (*entity.Checkout, error) {
	return u.repo.GetCheckoutByDesainID(ctx, desainId)
}
func (u *usecase) GenerateInvoicePDF(filePath string, inv entity.Invoice) error {

	if err := os.MkdirAll("assets", os.ModePerm); err != nil {
		return err
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "INVOICE PEMBAYARAN")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, fmt.Sprintf("Invoice ID: %d", inv.Id))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("User ID: %d", inv.UserId))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Checkout ID: %d", inv.CheckoutId))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Total: Rp %.2f", inv.TotalHarga))
	pdf.Ln(8)
	pdf.Cell(0, 10, fmt.Sprintf("Tanggal: %s", inv.TanggalBuat.Format("02 Jan 2006")))

	return pdf.OutputFileAndClose(filePath)
}
