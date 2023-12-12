package common

import "strconv"


// ConvertStringToInt convert string to int
func (c *CommonUsecase) ConvertStringToInt(numberStr string) (int64, error){
	number, err := strconv.ParseInt(numberStr, 10, 64)
	if err != nil{
		return 0, err
	}
	
	return number ,nil 
} 
