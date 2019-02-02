package database

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
)

const (
	clubPrefix         = "Club-"
	userPrefix         = "User-"
	verificationPrefix = "Verification-"
	clubTablePrefix    = "ClubData-"
	peidPrefix         = "PEID-"
	periodPrefix       = "Period-"
	applicantPrefix    = "Applicant-"
	tagPrefix          = "Tag-"
)

func getClubsAndUsersTable() *string {
	return aws.String("ClubsAndUsers")
}

func getClubTable(clubID string) *string {
	return aws.String(clubTablePrefix + clubID)
}

func prefixClubID(id string) string {
	return clubPrefix + id
}

func prefixUserEmail(email string) string {
	return userPrefix + email
}

func prefixVerificationHash(hash string) string {
	return verificationPrefix + hash
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

func removePrefix(str string) string {
	s := strings.SplitN(str, "-", 2)
	if len(s) > 1 {
		return s[1]
	}
	return str
}

func getPeriodAndEventID(peid string) (period string, eventID string) {
	str := removePrefix(peid)
	s := strings.SplitN(str, "-", 2)
	if len(s) > 1 {
		return s[0], s[1]
	}
	return peid, ""
}
