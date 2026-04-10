package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40
// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actual int) int {
    RemainingOvenTime := OvenTime - actual
    return RemainingOvenTime
	panic("RemainingOvenTime not implemented")
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    PreparationTime := numberOfLayers * 2
    return PreparationTime
	panic("PreparationTime not implemented")
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    ElapsedTime := (numberOfLayers*2) + actualMinutesInOven
    return ElapsedTime
	panic("ElapsedTime not implemented")
}