package types

import "github.com/aymantaybi/ronin/common"

type BlobSidecars []*BlobSidecar

type BlobSidecar struct {
	BlobTxSidecar
	TxHash common.Hash `json:"txHash"`
}

func NewBlobSidecarFromTx(tx *Transaction) *BlobSidecar {
	if tx.BlobTxSidecar() == nil {
		return nil
	}
	return &BlobSidecar{
		BlobTxSidecar: *tx.BlobTxSidecar(),
		TxHash:        tx.Hash(),
	}
}
