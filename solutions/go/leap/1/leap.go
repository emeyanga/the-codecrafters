package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
    if year % 4 == 0 {
        if year % 100 != 0 {
            return true
        }
    }
    if year % 400 == 0 {
            return true
    } else {
        return false
    }
	panic("Please implement the IsLeapYear function")
}
