/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ReplaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceRequest{}

// ReplaceRequest struct for ReplaceRequest
type ReplaceRequest struct {
	Files    []string `json:"files"`
	NewValue string   `json:"newValue"`
	Pattern  string   `json:"pattern"`
}

type _ReplaceRequest ReplaceRequest

// NewReplaceRequest instantiates a new ReplaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceRequest(files []string, newValue string, pattern string) *ReplaceRequest {
	this := ReplaceRequest{}
	this.Files = files
	this.NewValue = newValue
	this.Pattern = pattern
	return &this
}

// NewReplaceRequestWithDefaults instantiates a new ReplaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceRequestWithDefaults() *ReplaceRequest {
	this := ReplaceRequest{}
	return &this
}

// GetFiles returns the Files field value
func (o *ReplaceRequest) GetFiles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *ReplaceRequest) GetFilesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *ReplaceRequest) SetFiles(v []string) {
	o.Files = v
}

// GetNewValue returns the NewValue field value
func (o *ReplaceRequest) GetNewValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value
// and a boolean to check if the value has been set.
func (o *ReplaceRequest) GetNewValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewValue, true
}

// SetNewValue sets field value
func (o *ReplaceRequest) SetNewValue(v string) {
	o.NewValue = v
}

// GetPattern returns the Pattern field value
func (o *ReplaceRequest) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *ReplaceRequest) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *ReplaceRequest) SetPattern(v string) {
	o.Pattern = v
}

func (o ReplaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["files"] = o.Files
	toSerialize["newValue"] = o.NewValue
	toSerialize["pattern"] = o.Pattern
	return toSerialize, nil
}

func (o *ReplaceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"files",
		"newValue",
		"pattern",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varReplaceRequest := _ReplaceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReplaceRequest)

	if err != nil {
		return err
	}

	*o = ReplaceRequest(varReplaceRequest)

	return err
}

type NullableReplaceRequest struct {
	value *ReplaceRequest
	isSet bool
}

func (v NullableReplaceRequest) Get() *ReplaceRequest {
	return v.value
}

func (v *NullableReplaceRequest) Set(val *ReplaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceRequest(val *ReplaceRequest) *NullableReplaceRequest {
	return &NullableReplaceRequest{value: val, isSet: true}
}

func (v NullableReplaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}