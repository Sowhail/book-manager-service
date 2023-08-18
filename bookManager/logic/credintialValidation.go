package logic

import (
	"strings"

	"github.com/badoux/checkmail"
)

func PasswordValidation(password string) bool {
	if len(password) < 8 {
		return false
	}

	if strings.Contains(password, " ") {
		return false
	}

	return true

}

func UserNameValidation(userName string) bool {
	return len(userName) >= 5
}

func EmailValidation(email string) bool {

	err := checkmail.ValidateFormat(email)
	return err == nil
}

func PhoneNumberValidation(phoneNumber string) bool {
	numStr := "0987654321"
	for i := 0; i < len(phoneNumber); i++ {
		if !strings.Contains(numStr, string(phoneNumber[i])) {
			return false
		}
	}
	return len(phoneNumber) == 11
}

func NameValidation(firstName, lastName string) bool {
	if len(firstName) == 0 || len(lastName) == 0 {
		return false
	}
	return true
}
