// Code generated by "stringer -type=SagaOperationType"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MetaOperation-0]
	_ = x[PutOperation-1]
	_ = x[RemoveOperation-2]
	_ = x[TransferOperation-3]
}

const _SagaOperationType_name = "MetaOperationPutOperationRemoveOperationTransferOperation"

var _SagaOperationType_index = [...]uint8{0, 13, 25, 40, 57}

func (i SagaOperationType) String() string {
	if i < 0 || i >= SagaOperationType(len(_SagaOperationType_index)-1) {
		return "SagaOperationType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SagaOperationType_name[_SagaOperationType_index[i]:_SagaOperationType_index[i+1]]
}