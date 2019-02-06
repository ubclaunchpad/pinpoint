package database

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
)

const (
	clubAndUsersTable  = "ClubsAndUsers"
	clubTablePrefix    = "ClubData-"
	clubPrefix         = "Club-"
	userPrefix         = "User-"
	verificationPrefix = "Verification-"
	peidPrefix         = "PEID-"
	periodPrefix       = "Period-"
	applicantPrefix    = "Applicant-"
	tagPrefix          = "Tag-"
)

// returns a pointer to the main clubs and users table
func getClubsAndUsersTable() *string {
	return aws.String(clubAndUsersTable)
}

// returns a pointer the club's own table
func getClubTable(clubID string) *string {
	return aws.String(clubTablePrefix + clubID)
}

// prefixes the club ID for the clubs and users table
func prefixClubID(id string) string {
	return clubPrefix + id
}

// prefixes the user email for the clubs and users table
func prefixUserEmail(email string) string {
	return userPrefix + email
}

// prefixes the email verification hash for the clubs and users table
func prefixVerificationHash(hash string) string {
	return verificationPrefix + hash
}

// prefixes the period and event ID together for the club's table
func prefixPeriodEventID(period string, eventID string) string {
	return peidPrefix + period + "-" + eventID
}

// prefixes the period for the club's table
func prefixPeriodID(period string) string {
	return periodPrefix + period
}

// prefixes the applicant's email for the club's table
func prefixApplicantID(email string) string {
	return applicantPrefix + email
}

// prefixes the tag name for the for the club's table
func prefixTag(tag string) string {
	return tagPrefix + tag
}

// removes the prefix from the any table's key
func removePrefix(str string) string {
	s := strings.SplitN(str, "-", 2)
	if len(s) > 1 {
		return s[1]
	}
	return str
}

// extracts the preiod and event ID from the table's PEID
func getPeriodAndEventID(peid string) (period string, eventID string) {
	str := removePrefix(peid)
	s := strings.SplitN(str, "-", 2)
	if len(s) > 1 {
		return s[0], s[1]
	}
	return peid, ""
}
