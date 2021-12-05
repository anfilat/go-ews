package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/anfilat/go-ews"
	"github.com/anfilat/go-ews/enumerations/availabilityData"
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
)

func main() {
	login := secret("login")
	domain := secret("domain")
	password := secret("password")

	service := ews.NewExchangeService(exchangeVersion.Exchange2013)
	service.SetCredentials(ews.NewWebCredentials(login+domain, password))
	service.SetUrl("https://outlook.office365.com/Ews/Exchange.asmx")

	attendee := []ews.AttendeeInfo{
		ews.NewAttendeeInfo("adelev" + domain),
		ews.NewAttendeeInfo("akexw" + domain),
	}
	timeWindow := ews.NewTimeWindow(time.Now(), time.Now().Add(2*24*time.Hour))

	_, err := service.GetUserAvailability(attendee, timeWindow, availabilityData.FreeBusyAndSuggestions,
		ews.NewAvailabilityOptions())
	if err != nil {
		log.Fatal(err)
	}
}

func secret(fileName string) string {
	data, _ := os.ReadFile("./examples/secret/" + fileName)
	return strings.TrimSpace(string(data))
}
