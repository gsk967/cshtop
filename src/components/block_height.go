package components

import "github.com/gizak/termui/v3/widgets"

func BlockHeightComponent(refresh *bool) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Title = "Block Height"
	p.Text = "0"
	return p
}
