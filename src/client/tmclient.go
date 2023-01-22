package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	BlockSubscriber = "block_subscriber"
	TxSubscriber    = "tx_subscriber"
)

func init() {

}

func TMClient(uri string, refresh *bool, nbList *widgets.List, txList *widgets.List, bh *widgets.Paragraph) {
	client, err := tmclient.New(uri, "/websocket")
	if err != nil {
		log.Fatalf("failed to connect websocket : %v", err)
		os.Exit(1)
	}
	err = client.Start()
	if err != nil {
		log.Fatalf("failed to start websocket client : %v", err)
		os.Exit(1)
	}

	defer func() {
		if err := client.Stop(); err != nil {
			panic(err)
		}
		unsubscribeList := []string{BlockSubscriber, TxSubscriber}
		for _, subscriber := range unsubscribeList {
			if err := client.UnsubscribeAll(context.Background(), subscriber); err != nil {
				panic(err)
			}
		}
	}()

	nbEvents, err := client.Subscribe(context.Background(), BlockSubscriber, tmtypes.EventQueryNewBlock.String())
	if err != nil {
		os.Exit(1)
	}

	nTxsEvents, err := client.Subscribe(context.Background(), TxSubscriber, tmtypes.EventQueryTx.String())
	if err != nil {
		os.Exit(1)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	nBSents := make([]string, 0, 30)
	txs := make([]string, 0, 30)
	for {
		select {
		case result := <-nbEvents:
			*refresh = true
			// logger.Info("got new block header", "height", result.Data.(tmtypes.EventDataNewBlock).Block.Height)
			blist := blocksList(result.Data.(tmtypes.EventDataNewBlock).Block, nBSents)
			nbList.Rows = blist
			nbList.TextStyle = ui.NewStyle(ui.ColorYellow)
			nBSents = blist
			bheight := fmt.Sprintf("%d", result.Data.(tmtypes.EventDataNewBlock).Block.Height)
			bh.Text = bheight

		case txResult := <-nTxsEvents:
			*refresh = true
			txsList := newTxsList(txResult.Data.(tmtypes.EventDataTx), txs)
			txList.Rows = txsList
			txList.TextStyle = ui.NewStyle(ui.ColorYellow)
			txs = txsList

		case <-quit:
			os.Exit(0)
		}
	}
}

func newTxsList(tx tmtypes.EventDataTx, txs []string) []string {
	var sent string
	if tx.Result.Code == 0 {
		sent = fmt.Sprintf("✅ bHeight %d Tx Hash %s", tx.Height, "")
	} else {
		sent = fmt.Sprintf("😑 bHeight %d Tx Hash %s", tx.Height, "")
	}
	txs = append([]string{sent}, txs...)
	return txs
}

func blocksList(block *tmtypes.Block, nBSents []string) []string {
	nbSentences := fmt.Sprintf("🦧 %d proposed by %s and no of Txs=%d", block.Height, block.ProposerAddress, len(block.Txs))
	nBSents = append([]string{nbSentences}, nBSents...)
	return nBSents
}
