/*
httpmq

HTTP/2 based message broker built around NATS JetStream

API version: v0.4.0-rc.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GoutilsErrorDetail struct for GoutilsErrorDetail
type GoutilsErrorDetail struct {
	// Code is the response code
	Code int32 `json:"code"`
	// Detail is an optional descriptive message providing additional details on the error
	Detail *string `json:"detail,omitempty"`
	// Msg is an optional descriptive message
	Message *string `json:"message,omitempty"`
}

// NewGoutilsErrorDetail instantiates a new GoutilsErrorDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoutilsErrorDetail(code int32) *GoutilsErrorDetail {
	this := GoutilsErrorDetail{}
	this.Code = code
	return &this
}

// NewGoutilsErrorDetailWithDefaults instantiates a new GoutilsErrorDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoutilsErrorDetailWithDefaults() *GoutilsErrorDetail {
	this := GoutilsErrorDetail{}
	return &this
}

// GetCode returns the Code field value
func (o *GoutilsErrorDetail) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GoutilsErrorDetail) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GoutilsErrorDetail) SetCode(v int32) {
	o.Code = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *GoutilsErrorDetail) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoutilsErrorDetail) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *GoutilsErrorDetail) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *GoutilsErrorDetail) SetDetail(v string) {
	o.Detail = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GoutilsErrorDetail) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoutilsErrorDetail) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GoutilsErrorDetail) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GoutilsErrorDetail) SetMessage(v string) {
	o.Message = &v
}

func (o GoutilsErrorDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableGoutilsErrorDetail struct {
	value *GoutilsErrorDetail
	isSet bool
}

func (v NullableGoutilsErrorDetail) Get() *GoutilsErrorDetail {
	return v.value
}

func (v *NullableGoutilsErrorDetail) Set(val *GoutilsErrorDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableGoutilsErrorDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableGoutilsErrorDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoutilsErrorDetail(val *GoutilsErrorDetail) *NullableGoutilsErrorDetail {
	return &NullableGoutilsErrorDetail{value: val, isSet: true}
}

func (v NullableGoutilsErrorDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoutilsErrorDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
