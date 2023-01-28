package client

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gizak/termui/v3/widgets"
	"github.com/gsk967/cshtop/src/types"
	"github.com/tendermint/tendermint/libs/log"
	tmclient "github.com/tendermint/tendermint/rpc/client/http"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	BlockSubscriber = "block_subscriber"
	TxSubscriber    = "tx_subscriber"
)

func init() {

}

func TMClient(logger log.Logger, nodeUris []string) *tmclient.HTTP {
	var client *tmclient.HTTP
	for _, nodeUri := range nodeUris {
		var err error
		client, err = tmclient.New(nodeUri, "/websocket")
		if err != nil {
			logger.Error("failed to connect websocket", "uri", nodeUri, "err", err.Error())
			continue
		}
		err = client.Start()
		if err != nil {
			logger.Error("failed to start websocket client", "err", err.Error())
			continue
		}
		logger.Info("👍 connect to websocket", "uri", nodeUri)
		return client
	}

	defer func() {
		unsubscribeList := []string{BlockSubscriber, TxSubscriber}
		for _, subscriber := range unsubscribeList {
			if err := client.UnsubscribeAll(context.Background(), subscriber); err != nil {
				logger.Error("failed to Unsubscribe websocket client", "err", err.Error())
				panic(err)
			}
		}

		if err := client.Stop(); err != nil {
			logger.Error("failed to stop websocket client", "err", err.Error())
			os.Exit(1)
		}
	}()

	return client
}

func BlocksAndTxProcess(logger log.Logger, client *tmclient.HTTP, refresh *bool, nbList *widgets.List,
	txList *widgets.List, bh *widgets.Paragraph, vals types.ValidatorMap) {
	nbEvents, err := client.Subscribe(context.Background(), BlockSubscriber, tmtypes.EventQueryNewBlock.String())
	if err != nil {
		logger.Error("failed to subscribe websocket", "query", tmtypes.EventQueryNewBlock.String(), "err", err.Error())
		os.Exit(1)
	}

	// logger.Info("👍 subscribed to websocket", "query", tmtypes.EventQueryNewBlock.String())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	nBSents := make([]string, 0, 30)
	txs := make([]string, 0, 30)
	for {
		select {
		case result := <-nbEvents:
			*refresh = true
			block := result.Data.(tmtypes.EventDataNewBlock).Block
			blist := blocksList(block, nBSents, vals)
			// nbList.TextStyle = ui.NewStyle(ui.ColorYellow)
			nbList.Rows = blist
			nBSents = blist
			bheight := fmt.Sprintf("%d at %s", block.Height, block.Time.Local().String())
			bh.Text = bheight

			// txs list
			txSents := txListProcess(logger, client, block.Txs)
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

func txListProcess(logger log.Logger, client *tmclient.HTTP, txs tmtypes.Txs) []string {
	txSents := make([]string, 0, len(txs))
	for _, txHash := range txs {
		resTx, err := client.Tx(context.Background(), txHash.Hash(), false)
		if err != nil {
			logger.Error("❌ err while fetching tx ", "txHash", txHash.Hash(), "err", err.Error())
			continue
		}
		var sent string
		if resTx.TxResult.Code == 0 {
			sent = fmt.Sprintf("✅ Height %d TxHash %s", resTx.Height, resTx.Hash.String())
		} else {
			sent = fmt.Sprintf("❌ Height %d TxHash %s", resTx.Height, resTx.Hash.String())
		}

		txSents = append(txSents, sent)
	}

	return txSents
}

func blocksList(block *tmtypes.Block, nBSents []string, vals types.ValidatorMap) []string {
	moniker, ok := vals[block.ProposerAddress.String()]
	if !ok {
		fmt.Println("could not find validator info..")
	}
	nbSentences := fmt.Sprintf("👉 %d is proposed by %s and no of txs=%d", block.Height, moniker, len(block.Txs))
	nBSents = append([]string{nbSentences}, nBSents...)
	return nBSents
}
