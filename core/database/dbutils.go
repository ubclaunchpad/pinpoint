package database

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
)

func prefixClubID(id string) *string {
	return aws.String("Club-" + id)
}

func prefixUserEmail(email string) *string {
	return aws.String("User-" + email)
}

func getClubTable(clubID string) *string {
	return aws.String(clubTablePrefix + clubID)
}

func prefixPeriodEventID(period string, eventID string) string {
	return peidPrefix + period + "-" + eventID
}

func prefixPeriodID(period string) string {
	return periodPrefix + period
}

func prefixApplicantID(email string) string {
	return applicantPrefix + email
}

func prefixTag(tag string) string {
	return tagPrefix + tag
}

func getPeriodAndEventID(peid string) (string, string) {
	str := removePrefix(peid)
	s := strings.SplitN(str, "-", 2)
	return s[0], s[1]
}

func removePrefix(str string) string {
	s := strings.SplitN(str, "-", 2)
	if len(s) > 1 {
		return s[1]
	}
	return str
}
