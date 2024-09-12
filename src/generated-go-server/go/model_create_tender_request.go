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
	"github.com/google/uuid"
	"unicode/utf8"
)

type CreateTenderRequest struct {

	// Полное название тендера
	Name string `json:"name"`

	// Описание тендера
	Description string `json:"description"`

	ServiceType TenderServiceType `json:"serviceType"`

	// Уникальный идентификатор организации, присвоенный сервером.
	OrganizationId string `json:"organizationId"`

	// Уникальный slug пользователя.
	CreatorUsername string `json:"creatorUsername"`
}

// AssertCreateTenderRequestRequired checks if the required fields are not zero-ed
func AssertCreateTenderRequestRequired(obj CreateTenderRequest) error {
	elements := map[string]interface{}{
		"name":            obj.Name,
		"description":     obj.Description,
		"serviceType":     obj.ServiceType,
		"organizationId":  obj.OrganizationId,
		"creatorUsername": obj.CreatorUsername,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertCreateTenderRequestConstraints checks if the values respects the defined constraints
func AssertCreateTenderRequestConstraints(obj CreateTenderRequest) error {
	if len(obj.Name) == 0 {
		return fmt.Errorf("name is required")
	}
	if utf8.RuneCountInString(obj.Name) > 100 {
		return fmt.Errorf("name exceeds the maximum length of 100 characters")
	}

	if utf8.RuneCountInString(obj.Description) == 0 {
		return fmt.Errorf("description is required")
	}
	if utf8.RuneCountInString(obj.Description) > 500 {
		return fmt.Errorf("description exceeds the maximum length of 500 characters")
	}

	if ok := obj.ServiceType.IsValid(); !ok {
		return fmt.Errorf("invalid serviceType: must be one of 'Construction', 'Delivery', 'Manufacture'")
	}

	if _, err := uuid.Parse(obj.OrganizationId); err != nil {
		return fmt.Errorf("organizationId is required and must be a valid UUID")
	}

	if utf8.RuneCountInString(obj.CreatorUsername) == 0 {
		return fmt.Errorf("creatorUsername is required")
	}

	return nil
}
