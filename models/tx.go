package models

import (
	"fmt"
	"github.com/bianjieai/irita-sync/confs/server"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNameTx = "sync_tx"
)

type (
	Tx struct {
		Time      int64       `bson:"time"`
		Height    int64       `bson:"height"`
		TxHash    string      `bson:"tx_hash"`
		Type      string      `bson:"type"` // parse from first msg
		Memo      string      `bson:"memo"`
		Status    uint32      `bson:"status"`
		Log       string      `bson:"log"`
		Fee       *Fee        `bson:"fee"`
		Types     []string    `bson:"types"`
		Events    []Event     `bson:"events"`
		Signers   []string    `bson:"signers"`
		DocTxMsgs []DocTxMsg  `bson:"msgs"`
		Addrs     []string    `bson:"addrs"`
		TxIndex   uint32      `bson:"tx_index"` // sequence tx of this block
		Ext       interface{} `bson:"ext"`
	}

	Event struct {
		Type       string   `bson:"type"`
		Attributes []KvPair `bson:"attributes"`
	}

	KvPair struct {
		Key   string `bson:"key"`
		Value string `bson:"value"`
	}

	DocTxMsg struct {
		Type string `bson:"type"`
		Msg  Msg    `bson:"msg"`
	}

	Fee struct {
		Amount []Coin `bson:"amount" json:"amount"`
		Gas    int64  `bson:"gas" json:"gas"`
	}

	Msg interface {
		GetType() string
		BuildMsg(msg interface{})
	}
)

func (d Tx) Name() string {
	if server.SvrConf.ChainId == "" {
		return CollectionNameTx
	}
	return fmt.Sprintf("sync_%v_tx", server.SvrConf.ChainId)
}

func (d Tx) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{"-tx_hash"},
		Unique:     true,
		Background: true,
	})

	indexes = append(indexes, mgo.Index{
		Key:        []string{"-height", "tx_index"},
		Unique:     true,
		Background: true,
	})

	ensureIndexes(d.Name(), indexes)
}

func (d Tx) PkKvPair() map[string]interface{} {
	return bson.M{"tx_hash": d.TxHash}
}
