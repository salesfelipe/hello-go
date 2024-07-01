package gross

// Units stores the Gross Store unit measurements.

// quarter_of_a_dozen	3
// half_of_a_dozen	6
// dozen	12
// small_gross	120
// gross	144
// great_gross	1728
func Units() map[string]int {
	unitMap := make(map[string]int)

	unitMap["quarter_of_a_dozen"] = 3
	unitMap["half_of_a_dozen"] = 6
	unitMap["dozen"] = 12
	unitMap["small_gross"] = 120
	unitMap["gross"] = 144
	unitMap["great_gross"] = 1728

	return unitMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, existsUnit := units[unit]
	billItem, existsItemOnBill := bill[item]

	if !existsUnit {
		return false
	}

	if existsItemOnBill {
		bill[item] = billItem + unitValue
	} else {
		bill[item] = unitValue
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitValue, existsUnit := units[unit]
	billItem, existsItemOnBill := bill[item]

	if !existsItemOnBill || !existsUnit {
		return false
	}

	newQuantity := billItem - unitValue

	if newQuantity < 0 {
		return false
	}

	if newQuantity == 0 {
		delete(bill, item)
	} else {

		bill[item] = newQuantity
	}

	return true
}

// GetItem returns the quantitnewy of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billItem, existsItemOnBill := bill[item]

	if !existsItemOnBill {
		return 0, false
	}

	return billItem, true
}
