package services

import (
	"go-orderfood/models"
)

func StoreAcceptOrder(ids, idorder string) (bool, error) {
	oIns := models.Order{}
	status, sid, err := oIns.GetStoreOrStatus(idorder)
	if err != nil {
		return false, err
	}
	if sid != ids {
		return false, err
	}
	if status == "2" {
		err = oIns.UpStatusStore(idorder)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}
