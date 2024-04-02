package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

var reg = regexp.MustCompile(`\D`)

func Number(phoneNumber string) (string, error) {
	var onlyNumbers string = reg.ReplaceAllString(phoneNumber, "")

	if len(onlyNumbers) <= 9 || len(onlyNumbers) > 11 {
		return "", errors.New("invalid phone number")
	}

	if len(onlyNumbers) == 11 {
		if onlyNumbers[0] != '1' {
			return "", errors.New("invalid country code")
		}

		if onlyNumbers[1] < '2' {
			return "", errors.New("invalid area code")
		}

		if onlyNumbers[4] < '2' {
			return "", errors.New("invalid exchange code")
		}

		return onlyNumbers[1:], nil
	}

	if onlyNumbers[0] < '2' {
		return "", errors.New("invalid area code")
	}

	if onlyNumbers[3] < '2' {
		return "", errors.New("invalid exchange code")
	}

	return onlyNumbers, nil
}

func AreaCode(phoneNumber string) (string, error) {
	if phone, err := Number(phoneNumber); err == nil {
		return phone[0:3], nil
	} else {
		return "", err
	}
}

func Format(phoneNumber string) (string, error) {
	if phone, err := Number(phoneNumber); err == nil {
		return fmt.Sprintf("(%s) %s-%s", phone[0:3], phone[3:6], phone[6:]), nil
	} else {
		return "", err
	}
}
