package main

import (
	"fmt"

	g "github.com/OpenDiablo2/giu"
	"github.com/OpenDiablo2/giu/imgui"
)

var (
	editor     imgui.TextEditor
	errMarkers imgui.ErrorMarkers
)

func loop() {
	g.SingleWindow("Code Editor", g.Layout{
		g.Line(
			g.Button("Get Text", func() {
				if editor.HasSelection() {
					fmt.Println(editor.GetSelectedText())
				} else {
					fmt.Println(editor.GetText())
				}

				column, line := editor.GetCursorPos()
				fmt.Println("Cursor pos:", column, line)

				column, line = editor.GetSelectionStart()
				fmt.Println("Selection start:", column, line)

				fmt.Println("Current line is", editor.GetCurrentLineText())
			}),
			g.Button("Set Text", func() {
				editor.SetText("Set text")
			}),
			g.Button("Set Error Marker", func() {
				errMarkers.Clear()
				errMarkers.Insert(1, "Error message")
				fmt.Println("ErrMarkers Size:", errMarkers.Size())

				editor.SetErrorMarkers(errMarkers)
			}),
		),
		g.Custom(func() {
			editor.Render("Hello", imgui.Vec2{X: 0, Y: 0}, true)
		}),
	})
}

func main() {
	errMarkers = imgui.NewErrorMarkers()

	editor = imgui.NewTextEditor()
	editor.SetShowWhitespaces(false)
	editor.SetTabSize(2)
	editor.SetText("select * from greeting\nwhere date > current_timestamp\norder by date")
	editor.SetLanguageDefinitionSQL()

	wnd := g.NewMasterWindow("Code Editor", 800, 600, 0, nil)
	wnd.Main(loop)
}
