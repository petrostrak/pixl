package ui

func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)

	app.PixlWindow.SetContent(swatchesContainer)
}
