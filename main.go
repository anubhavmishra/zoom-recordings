package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mitchellh/colorstring"

	zoom "github.com/anubhavmishra/zoom-lib-golang"
)

func main() {
	var (
		apiKey    = os.Getenv("ZOOM_API_KEY")
		apiSecret = os.Getenv("ZOOM_API_SECRET")
	)

	// Initialize command line files for the date range
	// for the recordings
	var from string
	var to string
	var email string
	var debug bool

	flag.StringVar(&from, "from", "", "The date and time to start looking up recordings. Example: 2019-03-26T19:51:10.661Z."+
		"The date range has to be within six months.")
	flag.StringVar(&to, "to", "", "The date and time to end looking up recordings. Example: 2019-04-26T19:51:10.661Z."+
		"The date range has to be within six months.")
	flag.StringVar(&email, "account-email", os.Getenv("ZOOM_ACCOUNT_EMAIL"), "Zoom account email. It can also be supplied by"+
		" using the \"ZOOM_ACCOUNT_EMAIL\" environment variable.")
	meetingID := flag.Int("meeting-id", 0, "Zoom meeting id to filter.")
	flag.BoolVar(&debug, "debug", false, "Enable or disable debugging. Set to false by default.")

	flag.Parse()

	if email == "" {
		fmt.Println("no account email is set.")
		os.Exit(1)
	}

	if *meetingID != 0 {
		colorstring.Printf("[green][bold]Meeting ID: [white]%d\n", *meetingID)
	}

	zoom.APIKey = apiKey
	zoom.APISecret = apiSecret
	if debug {
		colorstring.Printf("[grey]Debug is set to [white]%s\n", debug)
		zoom.Debug = debug
	}

	user, err := zoom.GetUser(zoom.GetUserOpts{EmailOrID: email})
	if err != nil {
		fmt.Printf("got error listing users: %+v\n", err)
		os.Exit(1)
	}

	hundred := int(100)
	recordings, err := zoom.ListAllRecordings(zoom.ListAllRecordingsOptions{
		UserID:   user.ID,
		From:     from,
		To:       to,
		PageSize: &hundred,
	})
	if err != nil {
		fmt.Printf("got error listing meeting recordings: %+v\n", err)
		os.Exit(1)
	}

	colorstring.Printf("[white]All cloud recordings from %d to %d:\n", from, to)

	fmt.Println("||Name||Date and Time||Meeting Recording URL||")
	for _, meeting := range recordings.Meetings {
		for _, recording := range meeting.RecordingFiles {
			if recording.RecordingType == "shared_screen_with_speaker_view" {
				if *meetingID != 0 {
					if meeting.ID == *meetingID {
						fmt.Printf("|%s|%s|%s|\n", meeting.Topic, meeting.StartTime, recording.PlayURL)
					}
				}
			}
		}
	}

}
