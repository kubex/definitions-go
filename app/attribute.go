package app

import "github.com/kubex/definitions-go/translation"

type AttributeType string

const (
	// Base types

	AttributeTypeString AttributeType = "string" // string
	AttributeTypeBool   AttributeType = "bool"   // bool
	AttributeTypeInt    AttributeType = "int"    // int
	AttributeTypeFloat  AttributeType = "float"  // float

	// Additional formatting & indexing types

	AttributeTypeID        AttributeType = "id"        // string
	AttributeTypeTimestamp AttributeType = "timestamp" // unix timestamp int
	AttributeTypeLink      AttributeType = "link"      // string
)

type Attribute struct {
	// Key unique key for this attribute
	Key string
	// Name Label for this attribute within the UI
	Name translation.Text
	// Description
	Description translation.Text
	// Help Text do display to assist the user with data input
	Help translation.Text
	// Type the type of value stored
	Type AttributeType
	// Hidden Value of this attribute will not be visible within the UI
	Hidden bool
	// Required Value of this attribute is required to be set
	Required bool
	// Nullable allow null to be set/returned for this attribute
	Nullable bool
	// Index should indicate if the attribute can be searched
	Index bool
	// AnonymousData Attribute value can be guaranteed as anonymous
	AnonymousData bool
	// PersonalData Attribute contains personal data, encrypted on write, destroyed on a data delete request
	PersonalData bool
	// SecureData will be encrypted like personal data, but will not be removed on a data delete request
	SecureData bool
	// Annotations provide additional context
	Annotations map[string]string
	// AvailableValues provides a list of the possible options (if set, value must match)
	AvailableValues map[string]string
	// RegexMatch attribute is validated against this regex
	RegexMatch string
	// Priority allows you to specify the order of the attributes when in a slice
	Priority int
}
