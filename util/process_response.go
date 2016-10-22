package util

import (
	"gitlab.pramati.com/seshachalamm/fido/uaf/msg"
	"gitlab.pramati.com/seshachalamm/fido/uaf/ops"
	"gitlab.pramati.com/seshachalamm/fido/uaf/storage"
)

const SERVER_DATA_EXPIRY_IN_MS = 5 * 60 * 1000

func ProcessRegResponse(resp msg.RegistrationResponse) []storage.RegistrationRecord {
	notary := NotaryImpl{}
	result, err := ops.ProcessResponse(resp, SERVER_DATA_EXPIRY_IN_MS, notary)

	if err != nil {
		result = make([]storage.RegistrationRecord, 1)
		result[0].Status = err.Error()
	}

	return result
}
