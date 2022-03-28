package swatch

import "fyne.io/fyne/v2/driver/desktop"

func (s *Swatch) MouseDown(ev *desktop.MouseEvent) {
	s.clickHandler(s)
	s.Selected = true
	s.Refresh()
}

func (s *Swatch) MouseUp(ev *desktop.MouseEvent) {

}
