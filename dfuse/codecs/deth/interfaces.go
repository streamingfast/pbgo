package deth

func (t *TransactionTrace) ToTransaction() *Transaction {
	return &Transaction{
		To:       t.To,
		Nonce:    t.Nonce,
		GasPrice: t.GasPrice,
		GasLimit: t.GasLimit,
		Value:    t.Value,
		Input:    t.Input,
		V:        t.V,
		R:        t.R,
		S:        t.S,
		Hash:     t.Hash,
		From:     t.From,
	}
}
