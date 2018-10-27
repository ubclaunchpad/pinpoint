package verifier

import "testing"

func TestVerifier_New(t *testing.T) {
	_, err := Init("test@pinpoint.com")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestVerifier_generateHash(t *testing.T) {
	hash, err := generateHash("test@pinpoint.com")
	if err != nil {
		t.Error(err)
		return
	}
	if hash != "NmSdjumzjHOF7IAnafAK74LAPug=" {
		t.Error("Unexpected hash")
		return
	}
}
