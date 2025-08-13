package forms

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/marcelocg/bambah-sdk"
	"github.com/marcelocg/bambah-sdk/types"
)

type EntryForm struct {
	service bambah.FinancialService
	
	// Form widgets
	typeRadio       *widget.RadioGroup
	accountSelect   *widget.Select
	categorySelect  *widget.Select
	valueEntry      *widget.Entry
	dateEntry       *widget.Entry
	descriptionEntry *widget.Entry
	
	// Buttons
	saveButton   *widget.Button
	clearButton  *widget.Button
	
	// Container
	container *container.VBox
}

func NewEntryForm(service bambah.FinancialService) *EntryForm {
	form := &EntryForm{
		service: service,
	}
	
	form.initializeWidgets()
	form.buildContainer()
	form.loadDropdownData()
	
	return form
}

func (f *EntryForm) initializeWidgets() {
	// Entry type radio buttons
	f.typeRadio = widget.NewRadioGroup([]string{"Receita", "Despesa"}, func(value string) {
		// Update button text based on selection
		if value == "Receita" {
			f.saveButton.SetText("Salvar Receita")
		} else {
			f.saveButton.SetText("Salvar Despesa")
		}
	})
	f.typeRadio.SetSelected("Despesa") // Default to expense
	
	// Account dropdown
	f.accountSelect = widget.NewSelect([]string{}, nil)
	f.accountSelect.PlaceHolder = "Selecione uma conta..."
	
	// Category dropdown  
	f.categorySelect = widget.NewSelect([]string{}, nil)
	f.categorySelect.PlaceHolder = "Selecione uma categoria..."
	
	// Value entry
	f.valueEntry = widget.NewEntry()
	f.valueEntry.SetPlaceHolder("0.00")
	f.valueEntry.Validator = func(text string) error {
		if text == "" {
			return nil
		}
		if _, err := strconv.ParseFloat(strings.ReplaceAll(text, ",", "."), 64); err != nil {
			return fmt.Errorf("valor deve ser um número válido")
		}
		return nil
	}
	
	// Date entry with current date as default
	f.dateEntry = widget.NewEntry()
	f.dateEntry.SetText(time.Now().Format("2006-01-02"))
	f.dateEntry.Validator = func(text string) error {
		if text == "" {
			return nil
		}
		if _, err := time.Parse("2006-01-02", text); err != nil {
			return fmt.Errorf("data deve estar no formato YYYY-MM-DD")
		}
		return nil
	}
	
	// Description entry
	f.descriptionEntry = widget.NewMultiLineEntry()
	f.descriptionEntry.SetPlaceHolder("Descrição opcional...")
	f.descriptionEntry.Resize(fyne.NewSize(0, 60))
	
	// Buttons
	f.saveButton = widget.NewButton("Salvar Despesa", f.handleSave)
	f.clearButton = widget.NewButton("Limpar", f.handleClear)
}

func (f *EntryForm) buildContainer() {
	f.container = container.NewVBox(
		widget.NewCard("", "Tipo de Lançamento", f.typeRadio),
		
		widget.NewForm(
			widget.NewFormItem("Conta", f.accountSelect),
			widget.NewFormItem("Categoria", f.categorySelect),
			widget.NewFormItem("Valor", f.valueEntry),
			widget.NewFormItem("Data", f.dateEntry),
			widget.NewFormItem("Descrição", f.descriptionEntry),
		),
		
		container.NewGridWithColumns(2,
			f.saveButton,
			f.clearButton,
		),
	)
}

func (f *EntryForm) loadDropdownData() {
	// Load accounts
	if accounts, err := f.service.ListAccounts(); err == nil {
		f.accountSelect.Options = accounts
		f.accountSelect.Refresh()
	}
	
	// Load categories
	if categories, err := f.service.ListCategories(); err == nil {
		f.categorySelect.Options = categories
		f.categorySelect.Refresh()
	}
}

func (f *EntryForm) handleSave() {
	// Validate required fields
	if f.accountSelect.Selected == "" {
		f.showError("Por favor, selecione uma conta")
		return
	}
	
	if f.categorySelect.Selected == "" {
		f.showError("Por favor, selecione uma categoria")
		return
	}
	
	if f.valueEntry.Text == "" {
		f.showError("Por favor, informe o valor")
		return
	}
	
	// Parse value
	valueText := strings.ReplaceAll(f.valueEntry.Text, ",", ".")
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		f.showError("Valor inválido")
		return
	}
	
	// Adjust value sign based on type
	isExpense := f.typeRadio.Selected == "Despesa"
	if isExpense && value > 0 {
		value = -value
	} else if !isExpense && value < 0 {
		value = -value
	}
	
	// Parse date
	dateStr := f.dateEntry.Text
	if dateStr == "" {
		dateStr = time.Now().Format("2006-01-02")
	}
	
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		f.showError("Data inválida. Use o formato YYYY-MM-DD")
		return
	}
	
	// Format date as YYYYMMDD for SDK
	formattedDate := parsedDate.Format("20060102")
	
	// Create entry
	entry := types.FinancialEntry{
		Account:     f.accountSelect.Selected,
		Category:    f.categorySelect.Selected,
		Value:       value,
		Date:        formattedDate,
		Year:        parsedDate.Year(),
		Month:       int(parsedDate.Month()),
		Day:         parsedDate.Day(),
		Description: f.descriptionEntry.Text,
	}
	
	// Save entry
	if err := f.service.CreateFinancialEntry(entry); err != nil {
		f.showError(fmt.Sprintf("Erro ao salvar: %v", err))
		return
	}
	
	// Show success and clear form
	f.showSuccess(fmt.Sprintf("%s salva com sucesso!", f.typeRadio.Selected))
	f.handleClear()
}

func (f *EntryForm) handleClear() {
	f.accountSelect.ClearSelected()
	f.categorySelect.ClearSelected()
	f.valueEntry.SetText("")
	f.dateEntry.SetText(time.Now().Format("2006-01-02"))
	f.descriptionEntry.SetText("")
	f.typeRadio.SetSelected("Despesa")
	f.saveButton.SetText("Salvar Despesa")
}

func (f *EntryForm) showError(message string) {
	dialog.ShowError(fmt.Errorf(message), fyne.CurrentApp().Driver().AllWindows()[0])
}

func (f *EntryForm) showSuccess(message string) {
	dialog.ShowInformation("Sucesso", message, fyne.CurrentApp().Driver().AllWindows()[0])
}

func (f *EntryForm) GetContainer() *container.VBox {
	return f.container
}