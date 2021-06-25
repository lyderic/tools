package tools

import (
	"os"
	"testing"
)

func TestWirepusher(t *testing.T) {
	id := os.Getenv("WIREPUSHER_ID")
	title := "Tools Testing Report By Wirepusher"
	message := `This is a test.
Please delete!`
	system := "testing"
	err := Wirepusher(id, title, message, system)
	if err != nil {
		t.Errorf("Wirepusher failed with error: %v\n", err)
	}
}

func TestSendgrid(t *testing.T) {
	recipient := os.Getenv("RECIPIENT")
	sendgridApiKey := os.Getenv("SENDGRID_API_KEY")
	for _, s := range []string{recipient, sendgridApiKey} {
		if len(s) == 0 {
			t.Errorf("Missing environment variable (RECIPIENT or SENDGRID_API_KEY)\n")
			return
		}
	}
	hostname, err := os.Hostname()
	if err != nil {
		t.Errorf("%v", err)
	}
	email := Email{
		From:    "testing@" + hostname,
		To:      recipient,
		Subject: "This is a test of tools Lyderic's API",
		Body: `This is a test email.
Please delete!`,
	}
	t.Logf("email: %#v", email)
	t.Logf("key: %s", sendgridApiKey)
	err = Sendgrid(email, sendgridApiKey)
	if err != nil {
		t.Errorf("Sendgrid failed with error: %v\n", err)
	}
}
