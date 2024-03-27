package annalyn

func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake {
		return false
	} else {
		return true
	}
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if (knightIsAwake) || (archerIsAwake) || (prisonerIsAwake) {
		return true
	} else {
		return false
	}
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if (!archerIsAwake) && (prisonerIsAwake) {
		return true
	} else {
		return false
	}
}
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if (!knightIsAwake) && (!archerIsAwake) && (prisonerIsAwake) {
		return true
	} else if (petDogIsPresent) && (!archerIsAwake) {
		return true
	} else {
		return false
	}
}
