package adr

// ChangeStatus is the application service for changing the status in an ADR File
// It validates the status if there is a list of allowed statuses configured by user
//
func ChangeStatus(
	adrId int,
	status string,
	repository Repository,
	statusManager StatusManager,
	writer Writer,
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
