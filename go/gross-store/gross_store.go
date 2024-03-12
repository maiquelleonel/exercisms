package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	mesurements := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return mesurements
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, exists := units[unit]; !exists {
		return false
	}

	if value, exists := bill[item]; exists {
		bill[item] = value + units[unit]
	} else {
		bill[item] = units[unit]
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, exists := bill[item]; !exists {
		return false
	}

	if _, exists := units[unit]; !exists {
		return false
	}

	if bill[item]-units[unit] < 0 {
		return false
	}

	bill[item] -= units[unit]
	if bill[item] == 0 {
		delete(bill, item)
	}

	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if value, exists := bill[item]; exists {
		return value, true
	}
	return 0, false

}
