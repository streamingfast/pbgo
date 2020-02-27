package deos

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RAMOp_LegacyOperation(t *testing.T) {
	tests := []struct {
		in       RAMOp_Operation
		expected string
	}{
		{RAMOp_OPERATION_CREATE_TABLE, "create_table"},
		{RAMOp_OPERATION_DEFERRED_TRX_ADD, "deferred_trx_add"},
		{RAMOp_OPERATION_DEFERRED_TRX_CANCEL, "deferred_trx_cancel"},
		{RAMOp_OPERATION_DEFERRED_TRX_PUSHED, "deferred_trx_pushed"},
		{RAMOp_OPERATION_DEFERRED_TRX_RAM_CORRECTION, "deferred_trx_ram_correction"},
		{RAMOp_OPERATION_DEFERRED_TRX_REMOVED, "deferred_trx_removed"},
		{RAMOp_OPERATION_DELETEAUTH, "deleteauth"},
		{RAMOp_OPERATION_LINKAUTH, "linkauth"},
		{RAMOp_OPERATION_NEWACCOUNT, "newaccount"},
		{RAMOp_OPERATION_PRIMARY_INDEX_ADD, "primary_index_add"},
		{RAMOp_OPERATION_PRIMARY_INDEX_REMOVE, "primary_index_remove"},
		{RAMOp_OPERATION_PRIMARY_INDEX_UPDATE, "primary_index_update"},
		{RAMOp_OPERATION_PRIMARY_INDEX_UPDATE_ADD_NEW_PAYER, "primary_index_update_add_new_payer"},
		{RAMOp_OPERATION_PRIMARY_INDEX_UPDATE_REMOVE_OLD_PAYER, "primary_index_update_remove_old_payer"},
		{RAMOp_OPERATION_REMOVE_TABLE, "remove_table"},
		{RAMOp_OPERATION_SECONDARY_INDEX_ADD, "secondary_index_add"},
		{RAMOp_OPERATION_SECONDARY_INDEX_REMOVE, "secondary_index_remove"},
		{RAMOp_OPERATION_SECONDARY_INDEX_UPDATE_ADD_NEW_PAYER, "secondary_index_update_add_new_payer"},
		{RAMOp_OPERATION_SECONDARY_INDEX_UPDATE_REMOVE_OLD_PAYER, "secondary_index_update_remove_old_payer"},
		{RAMOp_OPERATION_SETABI, "setabi"},
		{RAMOp_OPERATION_SETCODE, "setcode"},
		{RAMOp_OPERATION_UNLINKAUTH, "unlinkauth"},
		{RAMOp_OPERATION_UPDATEAUTH_CREATE, "updateauth_create"},
		{RAMOp_OPERATION_UPDATEAUTH_UPDATE, "updateauth_update"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			op := &RAMOp{
				Operation: test.in,
			}

			assert.Equal(t, test.expected, op.LegacyOperation())
		})
	}
}
