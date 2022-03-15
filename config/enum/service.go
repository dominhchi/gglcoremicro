package enum

import (
	"database/sql/driver"
	"io"
)

type ServiceType string

func (p ServiceType) UnmarshalGQL(v interface{}) error {
	panic("implement me")
}

func (p ServiceType) MarshalGQL(w io.Writer) {
	panic("implement me")
}

const (
	Unknown                ServiceType = "UNKNOWN"
	AirConditionerCleaning ServiceType = "AC_CLEANING"
)

func (p ServiceType) String() string {
	switch p {
	case AirConditionerCleaning:
		return "AC_CLEANING"
	default:
		return "UNKNOWN"
	}
}

// Values provides list valid values for Enum.
func (ServiceType) Values() (types []string) {
	for _, s := range []ServiceType{Unknown, AirConditionerCleaning} {
		types = append(types, string(s))
	}
	return
}

// Value provides the DB a string from int.
func (p ServiceType) Value() (driver.Value, error) {
	return p.String(), nil
}

// Scan tells our code how to read the enum into our type.
func (p *ServiceType) Scan(val interface{}) error {
	var s string
	switch v := val.(type) {
	case nil:
		return nil
	case string:
		s = v
	case []uint8:
		s = string(v)
	}
	switch s {
	case "AC_CLEANING":
		*p = AirConditionerCleaning
	default:
		*p = Unknown
	}
	return nil
}
