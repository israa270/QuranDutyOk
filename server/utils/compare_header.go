package utils

// // CompareStrChunk  compare string chunk
// func CompareStrChunk(a, b []string) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	if (b == nil) != (a == nil) {
// 		return false
// 	}
// 	for key, value := range a {
// 		if value != b[key] {
// 			return false
// 		}
// 	}
// 	return true
// }




// CompareStrSlice compare string used in excel
func CompareStrSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (b == nil) != (a == nil) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}