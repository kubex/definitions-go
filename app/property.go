package app

import (
	"github.com/kubex/definitions-go/translation"
	"time"
)

type Property struct {
	Key        string `json:"key"`
	Type       PropertyType
	Definition *PropertyDefinition
}

type PropertyType string

const (
	PropertyTypeText    PropertyType = "text"
	PropertyTypeNumber  PropertyType = "number"
	PropertyTypeBoolean PropertyType = "boolean"
	PropertyTypeFloat   PropertyType = "float"
	PropertyTypeTime    PropertyType = "time"
	PropertyTypeMap     PropertyType = "map"
	PropertyTypeSet     PropertyType = "set"

	// Extended Types
	// TODO: Type > Map {json def}

	PropertyTypeAmount PropertyType = "amount"
)

type PropertyValue struct {
	Text   string
	Number int64
	Bool   bool
	Float  float64
	Time   time.Time
	Set    []string
	Map    map[string][]byte
}

type PropertyDefinition struct {
	// Name Label for this attribute within the UI
	Name translation.Text
	// Description
	Description translation.Text
	// Help Text do display to assist the user with data input
	Help translation.Text
	//Caption text to display under the input
	Caption translation.Text
	// Placeholder text to display in the input placeholder / hints the input value
	Placeholder translation.Text

	// Required Value of this attribute is required to be set
	Required bool
	// Nullable allow null to be set/returned for this attribute
	Nullable bool
	// PersonalData property contains personal data, encrypted on write, destroyed on a data delete request
	PersonalData bool

	// Immutable Value of this attribute cannot be changed
	Immutable bool

	// WriteOnly Allows this value to be written, but not displayed within the UI
	WriteOnly bool

	// Annotations provide additional context
	Annotations map[string]string

	// AvailableValues provides a list of the possible options (if set, value must match)
	AvailableValues map[string]string
	// RegexMatch attribute is validated against this regex
	RegexMatch string

	DefaultValue PropertyValue

	//DisplayType is the type of input to display
	DisplayType PropertyDisplayType

	//DisplayFormat Format the output following the pattern xX0-()/\[]{}#%&*+~,.:;<>|@!?
	DisplayFormat string

	// MaxOptions is the maximum number of elements that can be selected in a multi select
	MaxOptions int

	//TODO: Display when other property conditions are met
}

type PropertyDisplayType string

const (
	PropertyDisplayTypeText      PropertyDisplayType = "text"       // text
	PropertyDisplayTypeTextBlock PropertyDisplayType = "text-block" // long text
	PropertyDisplayTypeToggle    PropertyDisplayType = "toggle"     // bool
	PropertyDisplayTypeInt       PropertyDisplayType = "int"        // int
	PropertyDisplayTypeFloat     PropertyDisplayType = "float"      // float
	PropertyDisplayTypeDate      PropertyDisplayType = "date"       // date
	PropertyDisplayTypeTime      PropertyDisplayType = "time"       // time
	PropertyDisplayTypeTimestamp PropertyDisplayType = "timestamp"  // unix timestamp int
	PropertyDisplayTypeLink      PropertyDisplayType = "link"       // string
	PropertyDisplayTypeCheckbox  PropertyDisplayType = "checkbox"   // checkbox
	PropertyDisplayTypeRadio     PropertyDisplayType = "radio"      // radio
	PropertyDisplayTypeEmail     PropertyDisplayType = "email"      // email
	PropertyDisplayTypePhone     PropertyDisplayType = "phone"      // phone
	PropertyDisplayTypeAmount    PropertyDisplayType = "amount"     // amount + currency
	PropertyDisplayTypeAddress   PropertyDisplayType = "address"    // string
)
