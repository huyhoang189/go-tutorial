package TemplateMethod

import "fmt"

type SMS struct {
}

func (sms *SMS) getRandomOtp(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS : generating otp %s \n", randomOTP)
	return randomOTP
}

func (sms *SMS) saveOtpCache(s string) {
	fmt.Printf("SMS : saving otp to cache %s \n", s)
}

func (sms *SMS) getMessage(s string) string {
	return "SMS OTP for login is " + s + "\n"
}

func (sms *SMS) sendNotification(s string) error {
	fmt.Printf("SMS : Sending sms %s \n", s)
	return nil
}
