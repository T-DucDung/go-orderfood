package queries

func QueryGetOrder(id string) string {
	return "select id, user_id, shipper_id, shop_id, order_ "
}

func QueryGetShip(id string) string {
	return "select id, name, phone from shipper where id = " + id
}

func QueryGetIDOrUs(id string) string {
	return "select id from `order` where user_id = " + id + "order by order_status asc"
}
