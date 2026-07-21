package lib

import (
	"fmt"
	"log"
	"reflect"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/go-playground/validator/v10"
	eventid "github.com/oneweave/event-id"
)

var defaultValidator = validator.New(validator.WithRequiredStructEnabled())

func init() {
	_ = defaultValidator.RegisterValidation("eventid", ValidateEventID)
	_ = defaultValidator.RegisterValidation("notnil", ValidateNotNil)
}

// ValidateEventID validates that a string is a valid event-id.
func ValidateEventID(fl validator.FieldLevel) bool {
	_, err := eventid.Decode(fl.Field().String(), "")
	return err == nil
}

// ValidateNotNil validates that a field is not nil (allows empty maps, slices, etc.).
func ValidateNotNil(fl validator.FieldLevel) bool {
	field := fl.Field()
	switch field.Kind() {
	case reflect.Map, reflect.Slice, reflect.Chan, reflect.Func, reflect.Interface, reflect.Pointer:
		return !field.IsNil()
	default:
		return !field.IsZero()
	}
}

// ValidateStruct validates any struct using the default validator instance.
func ValidateStruct(s any) error {
	return defaultValidator.Struct(s)
}

// ParseAndValidate decodes a cloudevents.Event into the target event type and runs validation.
func ParseAndValidate[T any](event cloudevents.Event) (T, error) {
	var incomingData T
	err := event.DataAs(&incomingData)
	if err != nil {
		log.Printf("failed to decode event source=%s id=%s: %v", event.Source(), event.ID(), err)
		return incomingData, fmt.Errorf("decode %s payload: %w", event.Type(), err)
	}
	if err := ValidateStruct(&incomingData); err != nil {
		return incomingData, err
	}
	return incomingData, nil
}
