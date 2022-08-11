package withings

import (
	"fmt"
	"os"
	"time"

	"github.com/zono-dev/withings-go/withings"
)

const (
	tokenFile = "access_token.json"
	layout    = "2006-01-02"
	layout2   = "2006-01-02 15:04:05"
	isnotify  = false
)

var (
	jst        *time.Location
	t          time.Time
	adayago    time.Time
	lastupdate time.Time
	ed         string
	sd         string
	client     *(withings.Client)
	settings   map[string]string
)

func auth(settings map[string]string) {
	var err error
	client, err = withings.New(settings["CID"], settings["Secret"], settings["RedirectURL"])

	if err != nil {
		fmt.Println("Failed to create New client")
		fmt.Println(err)
		return
	}

	if _, err := os.Open(tokenFile); err != nil {
		var e error

		client.Token, e = withings.AuthorizeOffline(client.Conf)
		client.Client = withings.GetClient(client.Conf, client.Token)

		if e != nil {
			fmt.Println("Failed to authorize offline.")
		}
		fmt.Println("~~ authorized. Let's check the token file!")
	} else {
		_, err = client.ReadToken(tokenFile)

		if err != nil {
			fmt.Println("Failed to read token file.")
			fmt.Println(err)
			return
		}
	}
}

func tokenFuncs() {
	// Show token
	client.PrintToken()

	// Refresh Token if you need
	_, rf, err := client.RefreshToken()
	if err != nil {
		fmt.Println("Failed to RefreshToken")
		fmt.Println(err)
		return
	}
	if rf {
		fmt.Println("You got new token!")
		client.PrintToken()
	}

	// Save Token if you need
	err = client.SaveToken(tokenFile)
	if err != nil {
		fmt.Println("Failed to RefreshToken")
		fmt.Println(err)
		return
	}
}

func mainSetup() {
	jst = time.FixedZone("Asis/Tokyo", 9*60*60)
	t = time.Now()
	// to get sample data from 2 days ago to now
	adayago = t.Add(-48 * time.Hour)
	ed = t.Format(layout)
	sd = adayago.Format(layout)
	lastupdate = withings.OffsetBase
	//lastupdate = time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC)
}

func printMeas(v withings.MeasureData) {
	fmt.Printf("%.1f\n", v.Value)
}

func getWeight() {
	mym, err := client.GetMeas(withings.Real, adayago, t, lastupdate, 0, false, true, withings.Weight, withings.Height, withings.FatFreeMass, withings.BoneMass, withings.FatRatio, withings.FatMassWeight, withings.Temp, withings.HeartPulse, withings.Hydration)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range mym.SerializedData.Weights {
		fmt.Println(v.Value)
	}
}

func getFat() {
	mym, err := client.GetMeas(withings.Real, adayago, t, lastupdate, 0, false, true, withings.Weight, withings.Height, withings.FatFreeMass, withings.BoneMass, withings.FatRatio, withings.FatMassWeight, withings.Temp, withings.HeartPulse, withings.Hydration)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range mym.SerializedData.FatRatios {
		fmt.Println(v)
	}
}

func GetInfo() {
	settings = withings.ReadSettings(".test_settings.yaml")

	auth(settings)
	tokenFuncs()
	mainSetup()

	getWeight()
}
