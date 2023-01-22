package src

import (
	"log"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/gsk967/cshtop/src/client"
	"github.com/gsk967/cshtop/src/components"
)

// DrawMainMenu
func DrawMainMenu(appName, pName, cid string, node string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	refresh := false
	priceTicker := time.NewTicker(TI_PRICE * time.Second)

	// components
	priceComponent := *components.PriceComponent(appName, pName)
	blockHeight := *components.BlockHeightComponent(&refresh)

	lists := make([]*widgets.List, 2)
	for i := range lists {
		lists[i] = widgets.NewList()
	}
	lists[0].Title = "Blocks"
	lists[1].Title = "Txs"

	// Start the client
	go client.TMClient(node, &refresh, lists[0], lists[1], &blockHeight)

	grid.Set(
		ui.NewRow(0.1,
			ui.NewCol(1.0/3, &priceComponent),
			ui.NewCol(1.0/3, &blockHeight),
			ui.NewCol(1.0/3, components.ChainIdComponent(appName, cid)),
		),

		ui.NewRow(0.9,
			ui.NewCol(1.0/2, lists[0]),
			ui.NewCol(1.0/2, lists[1]),
		),
	)

	ui.Render(grid)
	tick := time.NewTicker(100 * time.Millisecond)
	uiEvents := ui.PollEvents()

	for {
		select {
		case <-tick.C:
			if !refresh {
				continue
			}
			refresh = false
			ui.Render(grid)

		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				ui.Clear()
				ui.Close()
				os.Exit(0)
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				ui.Render(grid)
			}

		case <-priceTicker.C:
			refresh = true
			priceComponent = *components.PriceComponent(appName, pName)
		}
	}
}
