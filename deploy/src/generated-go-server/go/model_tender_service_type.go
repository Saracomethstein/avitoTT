// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Tender Management API
 *
 * API для управления тендерами и предложениями.   Основные функции API включают управление тендерами (создание, изменение, получение списка) и управление предложениями (создание, изменение, получение списка).
 *
 * API version: 1.0
 */

package openapi

import (
	"fmt"
)

// TenderServiceType : Вид услуги, к которой относиться тендер
type TenderServiceType string

// List of TenderServiceType
const (
	CONSTRUCTION TenderServiceType = "Construction"
	DELIVERY     TenderServiceType = "Delivery"
	MANUFACTURE  TenderServiceType = "Manufacture"
	FREE 		 TenderServiceType = ""
)

// AllowedTenderServiceTypeEnumValues is all the allowed values of TenderServiceType enum
var AllowedTenderServiceTypeEnumValues = []TenderServiceType{
	"Construction",
	"Delivery",
	"Manufacture",
}

// validTenderServiceTypeEnumValue provides a map of TenderServiceTypes for fast verification of use input
var validTenderServiceTypeEnumValues = map[TenderServiceType]struct{}{
	"Construction": {},
	"Delivery":     {},
	"Manufacture":  {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TenderServiceType) IsValid() bool {
	_, ok := validTenderServiceTypeEnumValues[v]
	return ok
}

// NewTenderServiceTypeFromValue returns a pointer to a valid TenderServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTenderServiceTypeFromValue(v string) (TenderServiceType, error) {
	ev := TenderServiceType(v)
	if ev.IsValid() {
		return ev, nil
	}

	return "", fmt.Errorf("invalid value '%v' for TenderServiceType: valid values are %v", v, AllowedTenderServiceTypeEnumValues)
}

// AssertTenderServiceTypeRequired checks if the required fields are not zero-ed
func AssertTenderServiceTypeRequired(obj TenderServiceType) error {
	return nil
}

// AssertTenderServiceTypeConstraints checks if the values respects the defined constraints
func AssertTenderServiceTypeConstraints(obj TenderServiceType) error {
	return nil
}
