package app

type AttributeType string

const (
	AttributeTypeText   AttributeType = "text"
	AttributeTypeBool   AttributeType = "bool"
	AttributeTypeNumber AttributeType = "number"
)

type Attribute struct {
	// Key unique key for this attribute
	Key string
	// Type the type of value stored
	Type AttributeType
	// Required Value of this attribute is required to be set
	Required bool
	// AnonymousData Attribute value can be guaranteed as anonymous
	AnonymousData bool
	// PersonalData Attribute contains personal data
	PersonalData bool
	// Annotations provide additional context
	Annotations map[string]string
	// AvailableValues provides a list of the possible options (if set, value must match)
	AvailableValues map[string]string
	// RegexMatch attribute is validated against this regex
	RegexMatch string
}
