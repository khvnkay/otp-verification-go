package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envACCOUNTSID() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func envAUTHTOKEN() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("TWILIO_AUTHTOKEN")
}

func envSERVICESID() string {
	println(godotenv.Unmarshal(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("TWILIO_SERVICES_ID")
}
