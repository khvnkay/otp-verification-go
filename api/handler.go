package api

import (
	"context"
	"fmt"
	"khvnkay/otp-verify/data"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 100

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData
		defer cancel()
		app.validateBody(ctx, &payload)

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		_, err := app.twilioSendOTP(newData.PhoneNumber)
		fmt.Println("twilioSendOTP", newData)
		if err != nil {
			app.errorJSON(ctx, err)
		}
		app.writeJSON(ctx, http.StatusAccepted, "OTP sent Successfully")

	}

}
func (app *Config) verifySMS() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.VerifyData
		defer cancel()

		app.validateBody(ctx, &payload)

		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}
		fmt.Println("newData", newData.User.PhoneNumber)

		err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		if err != nil {
			app.errorJSON(ctx, err)
			return
		}
		app.writeJSON(ctx, http.StatusAccepted, "OTP verify Successfully")

	}

}
