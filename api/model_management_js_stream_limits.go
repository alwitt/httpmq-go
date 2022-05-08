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

// ManagementJSStreamLimits struct for ManagementJSStreamLimits
type ManagementJSStreamLimits struct {
	// MaxAge is the max duration (ns) the stream will store a message  Messages breaching the limit will be removed.
	MaxAge *int64 `json:"max_age,omitempty"`
	// MaxBytes is the max number of message bytes the stream will store.  Oldest messages are removed once limit breached.
	MaxBytes *int64 `json:"max_bytes,omitempty"`
	// MaxConsumers is the max number of consumers allowed on the stream
	MaxConsumers *int64 `json:"max_consumers,omitempty"`
	// MaxMsgSize is the max size of a message allowed in this stream
	MaxMsgSize *int64 `json:"max_msg_size,omitempty"`
	// MaxMsgs is the max number of messages the stream will store.  Oldest messages are removed once limit breached.
	MaxMsgs *int64 `json:"max_msgs,omitempty"`
	// MaxMsgsPerSubject is the maximum number of subjects allowed on this stream
	MaxMsgsPerSubject *int64 `json:"max_msgs_per_subject,omitempty"`
}

// NewManagementJSStreamLimits instantiates a new ManagementJSStreamLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementJSStreamLimits() *ManagementJSStreamLimits {
	this := ManagementJSStreamLimits{}
	return &this
}

// NewManagementJSStreamLimitsWithDefaults instantiates a new ManagementJSStreamLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementJSStreamLimitsWithDefaults() *ManagementJSStreamLimits {
	this := ManagementJSStreamLimits{}
	return &this
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *ManagementJSStreamLimits) GetMaxAge() int64 {
	if o == nil || o.MaxAge == nil {
		var ret int64
		return ret
	}
	return *o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementJSStreamLimits) GetMaxAgeOk() (*int64, bool) {
	if o == nil || o.MaxAge == nil {
		return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *ManagementJSStreamLimits) HasMaxAge() bool {
	if o != nil && o.MaxAge != nil {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given int64 and assigns it to the MaxAge field.
func (o *ManagementJSStreamLimits) SetMaxAge(v int64) {
	o.MaxAge = &v
}

// GetMaxBytes returns the MaxBytes field value if set, zero value otherwise.
func (o *ManagementJSStreamLimits) GetMaxBytes() int64 {
	if o == nil || o.MaxBytes == nil {
		var ret int64
		return ret
	}
	return *o.MaxBytes
}

// GetMaxBytesOk returns a tuple with the MaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementJSStreamLimits) GetMaxBytesOk() (*int64, bool) {
	if o == nil || o.MaxBytes == nil {
		return nil, false
	}
	return o.MaxBytes, true
}

// HasMaxBytes returns a boolean if a field has been set.
func (o *ManagementJSStreamLimits) HasMaxBytes() bool {
	if o != nil && o.MaxBytes != nil {
		return true
	}

	return false
}

// SetMaxBytes gets a reference to the given int64 and assigns it to the MaxBytes field.
func (o *ManagementJSStreamLimits) SetMaxBytes(v int64) {
	o.MaxBytes = &v
}

// GetMaxConsumers returns the MaxConsumers field value if set, zero value otherwise.
func (o *ManagementJSStreamLimits) GetMaxConsumers() int64 {
	if o == nil || o.MaxConsumers == nil {
		var ret int64
		return ret
	}
	return *o.MaxConsumers
}

// GetMaxConsumersOk returns a tuple with the MaxConsumers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementJSStreamLimits) GetMaxConsumersOk() (*int64, bool) {
	if o == nil || o.MaxConsumers == nil {
		return nil, false
	}
	return o.MaxConsumers, true
}

// HasMaxConsumers returns a boolean if a field has been set.
func (o *ManagementJSStreamLimits) HasMaxConsumers() bool {
	if o != nil && o.MaxConsumers != nil {
		return true
	}

	return false
}

// SetMaxConsumers gets a reference to the given int64 and assigns it to the MaxConsumers field.
func (o *ManagementJSStreamLimits) SetMaxConsumers(v int64) {
	o.MaxConsumers = &v
}

// GetMaxMsgSize returns the MaxMsgSize field value if set, zero value otherwise.
func (o *ManagementJSStreamLimits) GetMaxMsgSize() int64 {
	if o == nil || o.MaxMsgSize == nil {
		var ret int64
		return ret
	}
	return *o.MaxMsgSize
}

// GetMaxMsgSizeOk returns a tuple with the MaxMsgSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementJSStreamLimits) GetMaxMsgSizeOk() (*int64, bool) {
	if o == nil || o.MaxMsgSize == nil {
		return nil, false
	}
	return o.MaxMsgSize, true
}

// HasMaxMsgSize returns a boolean if a field has been set.
func (o *ManagementJSStreamLimits) HasMaxMsgSize() bool {
	if o != nil && o.MaxMsgSize != nil {
		return true
	}

	return false
}

// SetMaxMsgSize gets a reference to the given int64 and assigns it to the MaxMsgSize field.
func (o *ManagementJSStreamLimits) SetMaxMsgSize(v int64) {
	o.MaxMsgSize = &v
}

// GetMaxMsgs returns the MaxMsgs field value if set, zero value otherwise.
func (o *ManagementJSStreamLimits) GetMaxMsgs() int64 {
	if o == nil || o.MaxMsgs == nil {
		var ret int64
		return ret
	}
	return *o.MaxMsgs
}

// GetMaxMsgsOk returns a tuple with the MaxMsgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementJSStreamLimits) GetMaxMsgsOk() (*int64, bool) {
	if o == nil || o.MaxMsgs == nil {
		return nil, false
	}
	return o.MaxMsgs, true
}

// HasMaxMsgs returns a boolean if a field has been set.
func (o *ManagementJSStreamLimits) HasMaxMsgs() bool {
	if o != nil && o.MaxMsgs != nil {
		return true
	}

	return false
}

// SetMaxMsgs gets a reference to the given int64 and assigns it to the MaxMsgs field.
func (o *ManagementJSStreamLimits) SetMaxMsgs(v int64) {
	o.MaxMsgs = &v
}

// GetMaxMsgsPerSubject returns the MaxMsgsPerSubject field value if set, zero value otherwise.
func (o *ManagementJSStreamLimits) GetMaxMsgsPerSubject() int64 {
	if o == nil || o.MaxMsgsPerSubject == nil {
		var ret int64
		return ret
	}
	return *o.MaxMsgsPerSubject
}

// GetMaxMsgsPerSubjectOk returns a tuple with the MaxMsgsPerSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementJSStreamLimits) GetMaxMsgsPerSubjectOk() (*int64, bool) {
	if o == nil || o.MaxMsgsPerSubject == nil {
		return nil, false
	}
	return o.MaxMsgsPerSubject, true
}

// HasMaxMsgsPerSubject returns a boolean if a field has been set.
func (o *ManagementJSStreamLimits) HasMaxMsgsPerSubject() bool {
	if o != nil && o.MaxMsgsPerSubject != nil {
		return true
	}

	return false
}

// SetMaxMsgsPerSubject gets a reference to the given int64 and assigns it to the MaxMsgsPerSubject field.
func (o *ManagementJSStreamLimits) SetMaxMsgsPerSubject(v int64) {
	o.MaxMsgsPerSubject = &v
}

func (o ManagementJSStreamLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxAge != nil {
		toSerialize["max_age"] = o.MaxAge
	}
	if o.MaxBytes != nil {
		toSerialize["max_bytes"] = o.MaxBytes
	}
	if o.MaxConsumers != nil {
		toSerialize["max_consumers"] = o.MaxConsumers
	}
	if o.MaxMsgSize != nil {
		toSerialize["max_msg_size"] = o.MaxMsgSize
	}
	if o.MaxMsgs != nil {
		toSerialize["max_msgs"] = o.MaxMsgs
	}
	if o.MaxMsgsPerSubject != nil {
		toSerialize["max_msgs_per_subject"] = o.MaxMsgsPerSubject
	}
	return json.Marshal(toSerialize)
}

type NullableManagementJSStreamLimits struct {
	value *ManagementJSStreamLimits
	isSet bool
}

func (v NullableManagementJSStreamLimits) Get() *ManagementJSStreamLimits {
	return v.value
}

func (v *NullableManagementJSStreamLimits) Set(val *ManagementJSStreamLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementJSStreamLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementJSStreamLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementJSStreamLimits(val *ManagementJSStreamLimits) *NullableManagementJSStreamLimits {
	return &NullableManagementJSStreamLimits{value: val, isSet: true}
}

func (v NullableManagementJSStreamLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementJSStreamLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
