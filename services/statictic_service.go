package services

import (
	"go-orderfood/models"
	"strconv"
)

func GetStatistic(startTime, endTime int64) (models.Statistic, error) {
	lsID, err := GetIdS(startTime, endTime)
	if err != nil {
		return models.Statistic{}, err
	}
	countC, err := GetCountCus(startTime, endTime)
	if err != nil {
		return models.Statistic{}, err
	}

	statistic := models.Statistic{
		TotalShop:     strconv.Itoa(len(lsID)),
		TotalCustomer: countC,
	}
	return statistic, nil
}

func GetIdS(startTime, endTime int64) ([]string, error) {
	s := models.Store{}
	lsID, err := s.GetID(startTime, endTime)
	if err != nil {
		return []string{}, err
	}
	return lsID, err
}

func GetCountCus(startTime, endTime int64) (string, error) {
	c := models.Customers{}
	count, err := c.GetTotalCus(startTime, endTime)
	if err != nil {
		return "", err
	}
	return count, nil
}
