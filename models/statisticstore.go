package models

type StatisticStore struct {
	ID                 string             `json:"id"`
	Name               string             `json:"name"`
	RateAvg            string             `json:"rate_avg"`
	LoyalCustomersName []LoyalCustomers   `json:"loyal_customers_name"`
	BestSellingFoods   []BestSellingFoods `json:"best_selling_foods"`
	Revenue            string             `json:"revenue"`
	NumberOfOrders     string             `json:"number_of_orders"`
}
