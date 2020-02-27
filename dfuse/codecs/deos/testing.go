package deos

func NewTestAddEvent(idx int) *TransactionEvent_Addition {
	return &TransactionEvent_Addition{Addition: &TransactionEvent_Added{
		Receipt:    &TransactionReceipt{Index: uint64(idx)},
		PublicKeys: &PublicKeys{},
	}}
}

func NewTestIntAddEvent(prefix int) *TransactionEvent_InternalAddition {
	return &TransactionEvent_InternalAddition{InternalAddition: &TransactionEvent_AddedInternally{
		Transaction: &SignedTransaction{
			Transaction: &Transaction{
				Header: &TransactionHeader{
					RefBlockPrefix: uint32(prefix),
				},
			},
		},
	}}
}

func NewTestExecEvent(idx int) *TransactionEvent_Execution {
	return &TransactionEvent_Execution{Execution: &TransactionEvent_Executed{
		Trace: &TransactionTrace{
			Index: uint64(idx),
		},
	}}
}

func NewTestDtrxCreateEvent(src string) *TransactionEvent_DtrxScheduling {
	return &TransactionEvent_DtrxScheduling{DtrxScheduling: &TransactionEvent_DtrxScheduled{
		CreatedBy: &ExtDTrxOp{
			SourceTransactionId: src,
		},
	}}
}

func NewTestDtrxCancelEvent(src string) *TransactionEvent_DtrxCancellation {
	return &TransactionEvent_DtrxCancellation{DtrxCancellation: &TransactionEvent_DtrxCanceled{
		CanceledBy: &ExtDTrxOp{
			SourceTransactionId: src,
		},
	}}
}
