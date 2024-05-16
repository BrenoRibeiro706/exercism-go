package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {

	map_unity := make(map[string]int)

	map_unity["quarter_of_a_dozen"] = 3
	map_unity["half_of_a_dozen"] = 6
	map_unity["dozen"] = 12
	map_unity["small_gross"] = 120
	map_unity["gross"] = 144
	map_unity["great_gross"] = 1728

	return map_unity
}

// NewBill creates a new bill.
func NewBill() map[string]int {

	return map[string]int{}

}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if exists {
		bill[item] += value
		return true
	}
	return false
}

// RemoveItem remove item from customer bill
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if !exists {
		return false
	}
	if bill[item] >= value {
		bill[item] -= value
		if bill[item] == 0 {
			delete(bill, item)
		}
		return true
	}
	return false
}

// GetItem return the quantity of item that the customer has in his/her bill
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exist := bill[item]
	return qty, exist
}
