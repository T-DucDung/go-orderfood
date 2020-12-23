package services

import (
	"go-orderfood/models"
)

func GetLsOrderUser(id string) ([]models.Order, error) {
	lsid, err := GetIdOrUs(id)
	if err != nil {
		return []models.Order{}, err
	}
	lsor := []models.Order{}
	for _, item := range lsid {
		or, err := GetOrderUser(item)
		if err != nil {
			return []models.Order{}, err
		}
		lsor = append(lsor, or)
	}
	return lsor, err
}

func GetOrderUser(id string) (models.Order, error) {
	return models.Order{}, nil
}

func GetIdOrUs(id string) ([]string, error) {
	o := models.Order{}
	ls, err := o.GetOrUs(id)
	if err != nil {
		return []string{}, err
	}
	return ls, nil
}
