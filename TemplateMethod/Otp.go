package TemplateMethod

type Otp struct {
	ObjOTP IOtp
}

func (o *Otp) GenAndSendOTP(l int) error {
	otp := o.ObjOTP.getRandomOtp(l)
	o.ObjOTP.saveOtpCache(otp)
	message := o.ObjOTP.getMessage(otp)
	err := o.ObjOTP.sendNotification(message)
	if err != nil {
		return err
	}

	return nil
}
