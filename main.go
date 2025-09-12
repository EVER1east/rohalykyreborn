package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func main() {
	// Створюємо новий застосунок
	myApp := app.New()

	// Створюємо нове вікно з заголовком
	myWindow := myApp.NewWindow("Fyne App")

	var megatitle string
	var label = widget.NewLabel("")
	var input = widget.NewEntry()
	input.SetPlaceHolder("Напиши свої спогади про обригалики...")

	megatitle = "ROHALYKY ARE BACK!!!"
	// Створюємо кнопку
	helloBtn := widget.NewButton("Натисни мене", func() {
		label.SetText("Hello World!" + input.Text)
		myWindow.SetTitle(megatitle)
	})
	var content = container.NewVBox(
		input,
		helloBtn,
		label,
	)
	// Додаємо кнопку у вікно
	myWindow.SetContent(content)

	// Встановлюємо розмір і показуємо вікно
	myWindow.Resize(fyne.NewSize(500, 300))
	myWindow.ShowAndRun()
}
