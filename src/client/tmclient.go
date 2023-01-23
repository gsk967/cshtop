package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	// nTxsEvents, err := client.Subscribe(context.Background(), TxSubscriber, tmtypes.EventQueryTx.String())
	// if err != nil {
	// 	os.Exit(1)
	// }

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	nBSents := make([]string, 0, 30)
	txs := make([]string, 0, 30)
	for {
		select {
		case result := <-nbEvents:
			*refresh = true
			block := result.Data.(tmtypes.EventDataNewBlock).Block
			blist := blocksList(block, nBSents)
			// nbList.TextStyle = ui.NewStyle(ui.ColorYellow)
			nbList.Rows = blist
			nBSents = blist
			bheight := fmt.Sprintf("%d at %s", block.Height, block.Time.Local().String())
			bh.Text = bheight

			// txs list
			txSents := txListProcess(client, block.Txs)
			txs = append(txSents, txs...)
			txList.Rows = txs

		// case <-nTxsEvents:
		// 	// *refresh = true
		// 	// // txList.TextStyle = ui.NewStyle(ui.ColorYellow)
		// 	// txsList := newTxsList(txResult.Data.(tmtypes.EventDataTx), txs)
		// 	// txList.Rows = txsList
		// 	// txs = txsList

		case <-quit:
			os.Exit(0)
		}
	}
}

func txListProcess(client *tmclient.HTTP, txs tmtypes.Txs) []string {
	txSents := make([]string, 0, len(txs))
	for _, txHash := range txs {
		resTx, err := client.Tx(context.Background(), txHash.Hash(), false)
		if err != nil {
			println("err while fetching tx ", err.Error())
			continue
		}
		var sent string
		if resTx.TxResult.Code == 0 {
			sent = fmt.Sprintf("✅ bHeight %d TxHash %s", resTx.Height, resTx.Hash.String())
		} else {
			sent = fmt.Sprintf("❌ bHeight %d TxHash %s", resTx.Height, resTx.Hash.String())
		}

		txSents = append(txSents, sent)
	}

	return txSents
}

func blocksList(block *tmtypes.Block, nBSents []string) []string {
	nbSentences := fmt.Sprintf("👉 %d is proposed by %s and no of txs=%d", block.Height, block.ProposerAddress, len(block.Txs))
	nBSents = append([]string{nbSentences}, nBSents...)
	return nBSents
}
