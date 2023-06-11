package TemplateMethod

import "fmt"

type Email struct {
}

func (email *Email) getRandomOtp(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL : generating otp %s \n", randomOTP)
	return randomOTP
}

func (email *Email) saveOtpCache(s string) {
	fmt.Printf("EMAIL : saving otp to cache %s \n", s)
}

func (email *Email) getMessage(s string) string {
	return "EMAIL OTP for login is " + s + "\n"
}

func (email *Email) sendNotification(s string) error {
	fmt.Printf("EMAIL : Sending EMAIL %s \n", s)
	return nil
}
