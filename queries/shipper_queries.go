package queries

func QueryGetStatus(id string) string {
	return "select order_status from `order` where id = " + id
}

func QueryUpStatus(ido, ids string) string {
	return "update `order` set order_status = 4, shipper_id = " + ids + " where id = " + ido
}

func QueryEndStatus(ido, ids string) string {
	return "update `order` set order_status = 5 where shipper_id = " + ids + " and id = " + ido
}
