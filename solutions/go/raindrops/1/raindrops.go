package raindrops

import "strconv"

func Convert(number int) string {
    if number % 3 == 0 && number % 5 == 0 && number % 7 == 0 {
        return "PlingPlangPlong"
    } else if number % 3 == 0 && number % 5 == 0 {
        return "PlingPlang"
    } else if number % 3 == 0 && number % 7 == 0 {
        return "PlingPlong"
    } else if number % 5 == 0 && number % 3 == 0 {
        return "PlangPling"
    } else if number % 5 == 0 && number % 7 == 0 {
        return "PlangPlong"
    } else if number % 7 == 0 && number % 3 == 0 {
        return "PlongPling"
    } else if number % 7 == 0 && number % 5 == 0 {
        return "PlongPlang"
    } else if number % 3 == 0 {
        return "Pling"
    } else if number % 5 == 0 {
        return "Plang"
    } else if number % 7 == 0 {
        return "Plong"
    } else if number % 3 != 0 || number % 5 != 0 || number % 7 != 0 {
        return strconv.Itoa(number)
    }
	panic("Please implement the Convert function")
}
