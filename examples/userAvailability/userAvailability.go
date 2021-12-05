package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/anfilat/go-ews"
	"github.com/anfilat/go-ews/enumerations/availabilityData"
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
	"github.com/anfilat/go-ews/ewsCredentials"
	"github.com/anfilat/go-ews/ewsType"
)

func main() {
	login := secret("login")
	domain := secret("domain")
	password := secret("password")

	service := ews.New(exchangeVersion.Exchange2013)
	service.SetCredentials(ewsCredentials.NewWebCredentials(login+domain, password))
	service.SetUrl("https://outlook.office365.com/Ews/Exchange.asmx")

	attendee := []ewsType.AttendeeInfo{
		ewsType.NewAttendeeInfo("adelev" + domain),
		ewsType.NewAttendeeInfo("akexw" + domain),
	}
	timeWindow := ewsType.NewTimeWindow(time.Now(), time.Now().Add(2*24*time.Hour))

	_, err := service.GetUserAvailability(context.Background(),
		attendee, timeWindow, availabilityData.FreeBusyAndSuggestions, ewsType.NewAvailabilityOptions())
	if err != nil {
		log.Fatal(err)
	}
}

func secret(fileName string) string {
	data, _ := os.ReadFile("./examples/secret/" + fileName)
	return strings.TrimSpace(string(data))
}
