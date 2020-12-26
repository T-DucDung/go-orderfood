package queries

func QueryGetInfoCus(id string) string {
	return "select full_name, mobile from user where id = " + id
}

func QueryGetShip(id string) string {
	return "select id, name, phone from shipper where id = " + id
}

func QueryGetListFood(id string) string {
	return "select o.food_id as id_food, f.name as name , o.quantity as quantity, o.price as price from orderfood o, food f where o.food_id  = f.id and o.order_id  = " + id
}

func QueryGetIDOrUs(id string) string {
	return "select id, total_amount, shipper_id, store_id as shop_id, order_status, address from `order` where user_id = " + id + " order by order_status asc"
}

func QueryGetOrSh() string {
	return "select id, total_amount, user_id, store_id as shop_id, order_status, address from `order` where order_status = 3 order by id asc"
}

func QueryGetOrS(id string) string {
	return "select id, total_amount, user_id, shipper_id, order_status, address from `order` where store_id = " + id + " and order_status > 1 order by order_status asc"
}

func QueryGetOrSh1(id string) string {
	return "select id, total_amount, user_id, store_id as shop_id, order_status, address from `order` where shipper_id = " + id + " and order_status > 3 order by order_status asc"
}
