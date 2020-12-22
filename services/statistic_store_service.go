package services

import (
	"orderfood/models"
	"strconv"
)

func GetStatisticStore(sid int, startTime, endTime int64) (models.StatisticStore, error) {
	sIns, err := GetStoreInfo(sid)
	if err != nil {
		return models.StatisticStore{}, err
	}
	listLoyalCustomersName, err := GetListNameLoyalCustomers(sid, 3, startTime, endTime)
	if err != nil {
		return models.StatisticStore{}, err
	}
	listBestSellingFoods, err := GetBestSellingFood(sid, 3, startTime, endTime)
	if err != nil {
		return models.StatisticStore{}, err
	}
	r, err := GetRevenue(sid, startTime, endTime)
	if err != nil {
		return models.StatisticStore{}, err
	}
	sStore := models.StatisticStore{
		ID:                 sIns.ID,
		Name:               sIns.Name,
		RateAvg:            sIns.RateAvg,
		LoyalCustomersName: listLoyalCustomersName,
		BestSellingFoods:   listBestSellingFoods,
		Revenue:            r.TotalRevenue,
		NumberOfOrders:     r.TotalOrder,
	}

	return sStore, nil
}

func GetStoreInfo(sid int) (models.Store, error) {
	sIns := models.Store{
		ID: strconv.Itoa(sid),
	}
	return sIns.GetStoreInfo()
}

func GetListNameLoyalCustomers(sid, num int, startTime, endTime int64) ([]models.LoyalCustomers, error) {
	lc := models.LoyalCustomers{}
	listLoyalCustomersName, err := lc.GetLoyalCustomers(sid, startTime, endTime)
	if err != nil {
		return []models.LoyalCustomers{}, err
	}
	if len(listLoyalCustomersName) > num {
		listLoyalCustomersName = listLoyalCustomersName[:num]
	}
	for index, item := range listLoyalCustomersName {
		cusIns := models.Customers{
			ID: item.CustomersID,
		}
		cus, err := cusIns.GetCustomerName()
		if err != nil {
			return []models.LoyalCustomers{}, err
		}
		listLoyalCustomersName[index].CustomersName = cus.FullName
	}
	return listLoyalCustomersName, nil
}

func GetBestSellingFood(sid, num int, startTime, endTime int64) ([]models.BestSellingFoods, error) {
	lc := models.BestSellingFoods{}
	listBestSellingFoods, err := lc.GetBestSellingFoods(sid, startTime, endTime)
	if err != nil {
		return []models.BestSellingFoods{}, err
	}
	if len(listBestSellingFoods) > num {
		listBestSellingFoods = listBestSellingFoods[:num]
	}
	for index, item := range listBestSellingFoods {
		fIns := models.Foods{
			ID: item.FoodID,
		}
		f, err := fIns.GetFoodName()
		if err != nil {
			return []models.BestSellingFoods{}, err
		}
		listBestSellingFoods[index].FoodName = f.FullName
	}
	return listBestSellingFoods, nil
}

func GetRevenue(sid int, startTime, endTime int64) (models.Revenue, error) {
	rIns := models.Revenue{}
	r, err := rIns.GetRevenue(sid, startTime, endTime)
	if err != nil {
		return models.Revenue{}, err
	}
	return r, nil
}
