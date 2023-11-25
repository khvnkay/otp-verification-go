package api

import (
	"errors"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envACCOUNTSID(),
	Password: envAUTHTOKEN(),
})

func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")
	resp, err := client.VerifyV2.CreateVerification(envSERVICESID(), params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Status != nil {
			fmt.Println(*resp.Status)
		} else {
			fmt.Println(resp.Status)
		}
	}
	return *resp.Sid, nil

}

func (app *Config) twilioVerifyOTP(phoneNumber string, code string) error {

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(envSERVICESID(), params)

	if err != nil {
		fmt.Println(err.Error(), "=====")
	} else {
		if *resp.Status != "approved" {
			return errors.New("Not a valid code")

		}
		if resp.Status != nil {
			fmt.Println(*resp.Status)
		} else {
			fmt.Println(resp.Status)
		}
	}
	return nil
}
