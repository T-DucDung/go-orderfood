package queries

import "strconv"

func QueryGetLCN(sid int, startTime, endTime int64) string {
	return "SELECT COUNT(id) as total_count, user_id as customers_id from `order` where store_id = " + strconv.Itoa(sid) + " and order_status = 5 and updated_date BETWEEN " + strconv.FormatInt(startTime, 10) + " and " + strconv.FormatInt(endTime, 10) + " group by customers_id ORDER by total_count DESC "
}

func QueryGetIS(sid string) string {
	return "select id, created_date, updated_date, name, rate_avg from store where id = " + sid
}

func QueryGetUN(id string) string {
	return "select id, full_name from `user` where id  = " + id
}

func QueryGetBSF(sid int, startTime, endTime int64) string {
	return "SELECT COUNT(id) as total_count, food_id from orderfood where order_id IN (select id FROM `order` where order_status = 5 and store_id = " + strconv.Itoa(sid) + " and updated_date BETWEEN " + strconv.FormatInt(startTime, 10) + " and " + strconv.FormatInt(endTime, 10) + ") GROUP BY  food_id  ORDER  by total_count DESC "
}

func QueryGetFN(id string) string {
	return "select id, name as full_name from food where id = " + id
}

func QueryGetRS(sid int, startTime, endTime int64) string {
	return "SELECT COUNT(id) as total_order, SUM(total_amount) as total_revenue from `order` where store_id = " + strconv.Itoa(sid) + " and order_status = 5  and created_date BETWEEN " + strconv.FormatInt(startTime, 10) + " and " + strconv.FormatInt(endTime, 10)
}

func QueryGetIdS(startTime, endTime int64) string {
	return "SELECT id from store where created_date BETWEEN " + strconv.FormatInt(startTime, 10) + " and " + strconv.FormatInt(endTime, 10)
}

func QueryGetCC(startTime, endTime int64) string {
	return "select count(id) as soluong from account where type = 3 and created_date BETWEEN " + strconv.FormatInt(startTime, 10) + " and " + strconv.FormatInt(endTime, 10)
}

func QueryGetCS(startTime, endTime int64) string {
	return "select count(id) as soluong from account where type = 2 and created_date BETWEEN " + strconv.FormatInt(startTime, 10) + " and " + strconv.FormatInt(endTime, 10)
}

func QueryGetR(startTime, endTime int64) string {
	return "SELECT COUNT(id) as total_order, SUM(total_amount) as total_revenue from `order` where order_status = 5  and created_date BETWEEN " + strconv.FormatInt(startTime, 10) + " and " + strconv.FormatInt(endTime, 10)
}
