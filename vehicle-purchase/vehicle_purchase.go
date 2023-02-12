package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch kind {
	case "car":
		return true
	case "truck":
		return true
	default:
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	endingMsg := " is clearly the better choice."
	// If option1 is first in alphabet order, then it is going to be the option.
	if option1 < option2 {
		return option1 + endingMsg
	}
	// Otherwise option2 is the option.
	return option2 + endingMsg
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		return originalPrice * 0.7
	} else {
		return originalPrice * 0.5
	}
}
