package app

import (
	"errors"
	"regexp"
	"strings"
)

var (
	// ErrInvalidGlobalAppID invalid Global App ID
	ErrInvalidGlobalAppID = errors.New("the Global App ID specified is invalid")
	// ErrInvalidID Error for an invalid ID
	ErrInvalidID = errors.New("the ID specified is invalid")
)

// GlobalAppID
type GlobalAppID struct {
	VendorID  string
	AppID     string
	remainder string
}

func (g *GlobalAppID) String() string {
	return g.VendorID + "/" + g.AppID
}

// New Create a Global App ID from your vendor and application IDs
func NewID(vendorID string, applicationID string) *GlobalAppID {
	resp := &GlobalAppID{VendorID: vendorID, AppID: applicationID}
	return resp
}

// FromString Take a string starting with a GlobalAppID, and extract the vendor, app and remainder
func IDFromString(input string) *GlobalAppID {
	glapid := &GlobalAppID{}
	parts := strings.SplitN(input, "/", 3)
	if len(parts) > 1 {
		glapid = NewID(parts[0], parts[1])
		if len(parts) > 2 {
			glapid.remainder = parts[2]
		}
	}
	return glapid
}

// Validate Validates a Global App ID
func (g *GlobalAppID) Validate(strict bool) error {

	if err := ValidateID(g.VendorID); err != nil {
		return errors.New("Invalid Vendor ID " + g.VendorID)
	}

	if err := ValidateID(g.AppID); err != nil {
		return errors.New("Invalid App ID " + g.AppID)
	}
	if strict && g.remainder != "" {
		return ErrInvalidGlobalAppID
	}
	return nil

}

// CreateID converts a string to a valid ID
func CreateID(input string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	output := reg.ReplaceAllString(input, "-")
	output = strings.ToLower(output)
	output = strings.Trim(output, "-")
	return output
}

// ValidateID verifies an ID is a valid string
func ValidateID(input string) error {
	if !regexp.MustCompile("^[a-z0-9][a-z0-9\\-]+[a-z0-9]$").MatchString(input) {
		return ErrInvalidID
	}
	return nil
}
