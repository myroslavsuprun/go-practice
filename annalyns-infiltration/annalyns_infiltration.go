package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	return knightIsAwake == false
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	// If anyone in the group is awake, then true.
	groupIsAwake := knightIsAwake || archerIsAwake || prisonerIsAwake
	return groupIsAwake
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return !archerIsAwake && prisonerIsAwake
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	// In case archer is asleep and the dog is present.
	if !archerIsAwake && petDogIsPresent {
		return true
	}

	// In case any of the kidnappers are asleep and the prisoner is awake.
	kidnappersAreAwake := knightIsAwake || archerIsAwake
	if !kidnappersAreAwake && prisonerIsAwake {
		return true
	}

	// No match. return false value.
	return false
}
