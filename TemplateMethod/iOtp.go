package TemplateMethod

type IOtp interface {
	getRandomOtp(int) string
	saveOtpCache(string)
	getMessage(string) string
	sendNotification(string) error
}
