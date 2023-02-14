package elon

import "fmt"

// Drive increases the distance by the speed and decreases the battery by the batteryDrain.
func (c *Car) Drive() {
	if c.battery-c.batteryDrain < 0 {
		return
	}

	c.battery -= c.batteryDrain
	c.distance += c.speed
}

// DisplayDistance prints the meters the car has driven.
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery prints the battery percentage the car has left.
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish specifies whether the car can drive the provided speed.
func (c Car) CanFinish(trackDistance int) bool {
	// Calculating the maximum distance the car can make.
	possibleDistance := c.battery / c.batteryDrain * c.speed

	// If the track distance is less then the whole distance the car can make, then true.
	if trackDistance <= possibleDistance {
		return true
	}

	return false
}
