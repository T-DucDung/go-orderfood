package services

import (
	"go-orderfood/models"
	"log"
	"strconv"
)

func GetLsOrderUser(id string) ([]models.Order, error) {
	lsor, err := GetOrUs(id)
	if err != nil {
		return []models.Order{}, err
	}
	for index, o := range lsor {
		s, err := GetInfoShipper(o.ShipperID)
		if err != nil {
			log.Println("err 1", err)
			return []models.Order{}, err
		}
		lsor[index].NameShipper = s.Name
		lsor[index].PhoneShipper = s.Phone
		sid, err := strconv.Atoi(o.ShopID)
		if err != nil {
			log.Println("err 2", err)

			return []models.Order{}, err
		}
		store, err := GetStoreInfo(sid)
		if err != nil {
			log.Println("err 3", err)

			return []models.Order{}, err
		}
		lsor[index].NameShop = store.Name
		lf, err := GetListFood(o.ID)
		if err != nil {
			log.Println("err 4", err)

			return []models.Order{}, err
		}
		lsor[index].ListFood = lf
	}

	return lsor, nil
}

func GetOrUs(id string) ([]models.Order, error) {
	o := models.Order{}
	ls, err := o.GetOrUs(id)
	if err != nil {
		return []models.Order{}, err
	}
	return ls, nil
}

func GetInfoShipper(id string) (models.Shipper, error) {
	sIns := models.Shipper{}
	s, err := sIns.GetInfoShipper(id)
	if err != nil {
		return models.Shipper{}, err
	}
	return s, nil
}

func GetListFood(id string) ([]models.FoodOrder, error) {
	fIns := models.FoodOrder{}
	lf, err := fIns.GetListFoodOrder(id)
	if err != nil {
		return []models.FoodOrder{}, err
	}
	return lf, nil
}

func GetLsOrderShipper() ([]models.Order, error) {
	lsor, err := GetOrSh()
	if err != nil {
		log.Println("err ", err)

		return []models.Order{}, err
	}
	for index, o := range lsor {

		info, err := GetInfoCus(o.UserID)
		if err != nil {
			log.Println("err 1", err)

			return []models.Order{}, err
		}
		lsor[index].FullName = info.FullName
		lsor[index].Mobile = info.Mobile
		sid, err := strconv.Atoi(o.ShopID)
		if err != nil {
			log.Println("err 2", err)

			return []models.Order{}, err
		}
		store, err := GetStoreInfo(sid)
		if err != nil {
			log.Println("err 3", err)

			return []models.Order{}, err
		}
		lsor[index].NameShop = store.Name
		lf, err := GetListFood(o.ID)
		if err != nil {
			log.Println("err 4", err)

			return []models.Order{}, err
		}
		lsor[index].ListFood = lf
	}
	return lsor, nil
}

func GetOrSh() ([]models.Order, error) {
	s := models.Order{}
	ls, err := s.GetOrSh()
	if err != nil {
		return []models.Order{}, err
	}
	return ls, nil
}

func GetInfoCus(id string) (models.Customers, error) {
	cus := models.Customers{}
	info, err := cus.GetInFo(id)
	if err != nil {
		return models.Customers{}, err
	}
	return info, nil
}

func GetLsOrderShop(id string) ([]models.Order, error) {
	lsor, err := GetOr(id)
	if err != nil {
		return []models.Order{}, err
	}
	for index, o := range lsor {
		s, err := GetInfoShipper(o.ShipperID)
		if err != nil {
			log.Println("err 1", err)
			return []models.Order{}, err
		}
		lsor[index].NameShipper = s.Name
		lsor[index].PhoneShipper = s.Phone
		info, err := GetInfoCus(o.UserID)
		if err != nil {
			log.Println("err 1", err)

			return []models.Order{}, err
		}
		lsor[index].FullName = info.FullName
		lsor[index].Mobile = info.Mobile
		lf, err := GetListFood(id)
		if err != nil {
			log.Println("err 4", err)

			return []models.Order{}, err
		}
		lsor[index].ListFood = lf
	}

	return lsor, nil
}

func GetOr(id string) ([]models.Order, error) {
	o := models.Order{}
	lo, err := o.GetOrS(id)
	if err != nil {
		return []models.Order{}, err
	}
	return lo, nil
}

func GetLsOrderShipper1(id string) ([]models.Order, error) {
	lsor, err := GetOrSh1(id)
	if err != nil {
		log.Println("err ", err)

		return []models.Order{}, err
	}
	for index, o := range lsor {

		info, err := GetInfoCus(o.UserID)
		if err != nil {
			log.Println("err 1", err)

			return []models.Order{}, err
		}
		lsor[index].FullName = info.FullName
		lsor[index].Mobile = info.Mobile
		sid, err := strconv.Atoi(o.ShopID)
		if err != nil {
			log.Println("err 2", err)

			return []models.Order{}, err
		}
		store, err := GetStoreInfo(sid)
		if err != nil {
			log.Println("err 3", err)

			return []models.Order{}, err
		}
		lsor[index].NameShop = store.Name
		lf, err := GetListFood(o.ShopID)
		if err != nil {
			log.Println("err 4", err)

			return []models.Order{}, err
		}
		lsor[index].ListFood = lf
	}
	return lsor, nil
}

func GetOrSh1(id string) ([]models.Order, error) {
	o := models.Order{}
	lo, err := o.GetOrSh1(id)
	if err != nil {
		return []models.Order{}, err
	}
	return lo, nil
}
