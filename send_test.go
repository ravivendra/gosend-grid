package gosendgrid

import "testing"

var (
	key = "__fill_with_your_key__"

	testFrom     = "john@abc.com"
	testFromName = "John Smith"

	testTo     = "jane@abc.com"
	testToName = "Jane Doe"

	testSubject = "Test"
	testBody    = "Test"
)

func TestDoSend(t *testing.T) {
	if _, _, err := DoSend("", "text/xml", testFromName, testFrom, testToName, testTo, testSubject, testBody); err != nil {
		t.Errorf("Test false format using Sendgrid : %s", err.Error())
	}

	if _, _, err := DoSend("", "text/html", testFromName, testFrom, testToName, testTo, testSubject, testBody); err != nil {
		t.Errorf("Test empty key using Sendgrid : %s", err.Error())
	}

	if _, _, err := DoSend("abcdefghijklmnopqrstuvwxyz", "text/html", testFromName, testFrom, testToName, testTo, testSubject, testBody); err != nil {
		t.Errorf("Test invalid key using Sendgrid : %s", err.Error())
	}

	if _, _, err := DoSend(key, "text/html", testFromName, testFrom, testToName, testTo, testSubject, testBody); err != nil {
		t.Errorf("Test valid key using Sendgrid : %s", err.Error())
	}
}
