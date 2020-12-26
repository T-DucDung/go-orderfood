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
	countS, err := GetCountShip(startTime, endTime)
	if err != nil {
		return models.Statistic{}, err
	}
	r, err := GetRevenue(startTime, endTime)
	if err != nil {
		return models.Statistic{}, err
	}
	ls := []models.StatisticStore{}
	s := models.StatisticStore{}
	for _, id := range lsID {
		ids, err := strconv.Atoi(id.ID)
		if err != nil {
			return models.Statistic{}, err
		}
		s, err = GetStatisticStore(ids, startTime, endTime)
		if err != nil {
			return models.Statistic{}, err
		}
		ls = append(ls, s)
	}

	statistic := models.Statistic{
		TotalShop:         strconv.Itoa(len(lsID)),
		TotalCustomer:     countC,
		TotalShipper:      countS,
		TotalOrder:        r.TotalOrder,
		Revenue:           r.TotalRevenue,
		StatisticAllStore: ls,
	}
	return statistic, nil
}

func GetIdS(startTime, endTime int64) ([]models.LString, error) {
	s := models.Store{}
	lsID, err := s.GetID(startTime, endTime)
	if err != nil {
		return []models.LString{}, err
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

func GetCountShip(startTime, endTime int64) (string, error) {
	c := models.Shipper{}
	count, err := c.GetTotalShi(startTime, endTime)
	if err != nil {
		return "", err
	}
	return count, nil
}

func GetRevenue(startTime, endTime int64) (models.Revenue, error) {
	rIns := models.Revenue{}
	r, err := rIns.GetRevenue(startTime, endTime)
	if err != nil {
		return models.Revenue{}, err
	}
	return r, nil
}
