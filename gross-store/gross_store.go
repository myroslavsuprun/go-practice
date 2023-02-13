package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	// Point of sale.
	pos := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return pos
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	if customerUnit, exists := units[unit]; exists {
		bill[item] += customerUnit

		return true
	}

	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	_, customerItemExists := bill[item]
	customerUnit, customerUnitExists := units[unit]

	// If any of the provided values do not exist, return false.
	if !customerItemExists || !customerUnitExists {
		return false
	}

	// Calculate the possible bill item quantity.
	newQuantity := bill[item] - customerUnit

	// If the possible bill item is a negative value, return false.
	if newQuantity < 0 {
		return false
		// If the possible bill item is zero, remove it.
	} else if newQuantity == 0 {
		delete(bill, item)
		// If the possible bill item is more than zero, update it.
	} else {
		bill[item] -= customerUnit
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billItem, exists := bill[item]
	return billItem, exists
}
