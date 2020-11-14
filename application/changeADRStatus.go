package application

import "github.com/asiermarques/adrgen/domain"


// ChangeADRStatus is the application service for changing the status in an ADR File
// It validates the status if there is a list of allowed statuses configured by user
//
func ChangeADRStatus(
	adrId int,
	status string,
	repository domain.ADRRepository,
	statusManager domain.ADRStatusManager,
	writer domain.ADRWriter,
) (string, error) {
	ADR, err := repository.FindById(adrId)
	if err != nil {
		return "", err
	}

	ADR, err = statusManager.ChangeStatus(ADR, status)
	if err != nil {
		return "", err
	}

	err = writer.Persist(ADR)
	return ADR.Filename().Value(), err
}
