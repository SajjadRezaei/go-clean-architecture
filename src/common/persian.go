package common

import (
	"log"
	"regexp"
)

const iranianMobileNumberPattern string = `^09(1[0-9]|2[0-2]|3[0-9])[0-9]{7}$`

func ValidateMobileNumber(mobile string) bool {
	res, err := regexp.MatchString(iranianMobileNumberPattern, mobile)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}
