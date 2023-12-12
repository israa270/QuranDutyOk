package utils

// import "errors"

// Contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// CheckStringValue Validate string is contains only numbers
func CheckStringValue(s string) bool {
	b := true
	for _, c := range s {
		if c < '0' || c > '9' {
			b = false
			break
		}
	}
	return b
}

func ContainsInt(s []uint, num uint) bool {
	for _, v := range s {
		if v == num {
			return true
		}
	}
	return false
}



// func paginateSlice(slice []int, page int, pageSize int) ([]int, error) {
//     startIndex := (page - 1) * pageSize
//     if startIndex >= len(slice) {
//         return nil, errors.New("invalid page number")
//     }
//     endIndex := startIndex + pageSize
//     if endIndex > len(slice) {
//         endIndex = len(slice)
//     }
//     return slice[startIndex:endIndex], nil
// }