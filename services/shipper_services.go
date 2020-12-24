package services

import (
	"go-orderfood/models"
)

func AcceptOrder(idship, idorder string) (bool, error) {
	oIns := models.Order{}
	status, err := oIns.GetStatus(idorder)
	if err != nil {
		return false, err
	}
	if status == "3" {
		err = oIns.UpStatus(idorder, idship)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

func EndOrder(idship, idorder string) (bool, error) {
	oIns := models.Order{}
	status, err := oIns.GetStatus(idorder)
	if err != nil {
		return false, err
	}
	if status == "4" {
		err = oIns.EndStatus(idorder, idship)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}
