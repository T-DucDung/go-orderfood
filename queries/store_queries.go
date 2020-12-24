package queries

func QueryGetStoreOrStaus(id string) string {
	return "select order_status, store_id from `order` where id = " + id
}

func QueryUpStoreOrStatus(ido string) string {
	return "update `order` set order_status = 3 where id = " + ido
}
