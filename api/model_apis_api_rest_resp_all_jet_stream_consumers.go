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

// ApisAPIRestRespAllJetStreamConsumers struct for ApisAPIRestRespAllJetStreamConsumers
type ApisAPIRestRespAllJetStreamConsumers struct {
	// Consumers the set of consumer details mapped against consumer name
	Consumers *map[string]ApisAPIRestRespConsumerInfo `json:"consumers,omitempty"`
	Error     *GoutilsErrorDetail                     `json:"error,omitempty"`
	// RequestID gives the request ID to match against logs
	RequestId string `json:"request_id"`
	// Success indicates whether the request was successful
	Success bool `json:"success"`
}

// NewApisAPIRestRespAllJetStreamConsumers instantiates a new ApisAPIRestRespAllJetStreamConsumers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApisAPIRestRespAllJetStreamConsumers(requestId string, success bool) *ApisAPIRestRespAllJetStreamConsumers {
	this := ApisAPIRestRespAllJetStreamConsumers{}
	this.RequestId = requestId
	this.Success = success
	return &this
}

// NewApisAPIRestRespAllJetStreamConsumersWithDefaults instantiates a new ApisAPIRestRespAllJetStreamConsumers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApisAPIRestRespAllJetStreamConsumersWithDefaults() *ApisAPIRestRespAllJetStreamConsumers {
	this := ApisAPIRestRespAllJetStreamConsumers{}
	return &this
}

// GetConsumers returns the Consumers field value if set, zero value otherwise.
func (o *ApisAPIRestRespAllJetStreamConsumers) GetConsumers() map[string]ApisAPIRestRespConsumerInfo {
	if o == nil || o.Consumers == nil {
		var ret map[string]ApisAPIRestRespConsumerInfo
		return ret
	}
	return *o.Consumers
}

// GetConsumersOk returns a tuple with the Consumers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespAllJetStreamConsumers) GetConsumersOk() (*map[string]ApisAPIRestRespConsumerInfo, bool) {
	if o == nil || o.Consumers == nil {
		return nil, false
	}
	return o.Consumers, true
}

// HasConsumers returns a boolean if a field has been set.
func (o *ApisAPIRestRespAllJetStreamConsumers) HasConsumers() bool {
	if o != nil && o.Consumers != nil {
		return true
	}

	return false
}

// SetConsumers gets a reference to the given map[string]ApisAPIRestRespConsumerInfo and assigns it to the Consumers field.
func (o *ApisAPIRestRespAllJetStreamConsumers) SetConsumers(v map[string]ApisAPIRestRespConsumerInfo) {
	o.Consumers = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ApisAPIRestRespAllJetStreamConsumers) GetError() GoutilsErrorDetail {
	if o == nil || o.Error == nil {
		var ret GoutilsErrorDetail
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespAllJetStreamConsumers) GetErrorOk() (*GoutilsErrorDetail, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ApisAPIRestRespAllJetStreamConsumers) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given GoutilsErrorDetail and assigns it to the Error field.
func (o *ApisAPIRestRespAllJetStreamConsumers) SetError(v GoutilsErrorDetail) {
	o.Error = &v
}

// GetRequestId returns the RequestId field value
func (o *ApisAPIRestRespAllJetStreamConsumers) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespAllJetStreamConsumers) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ApisAPIRestRespAllJetStreamConsumers) SetRequestId(v string) {
	o.RequestId = v
}

// GetSuccess returns the Success field value
func (o *ApisAPIRestRespAllJetStreamConsumers) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespAllJetStreamConsumers) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ApisAPIRestRespAllJetStreamConsumers) SetSuccess(v bool) {
	o.Success = v
}

func (o ApisAPIRestRespAllJetStreamConsumers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Consumers != nil {
		toSerialize["consumers"] = o.Consumers
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableApisAPIRestRespAllJetStreamConsumers struct {
	value *ApisAPIRestRespAllJetStreamConsumers
	isSet bool
}

func (v NullableApisAPIRestRespAllJetStreamConsumers) Get() *ApisAPIRestRespAllJetStreamConsumers {
	return v.value
}

func (v *NullableApisAPIRestRespAllJetStreamConsumers) Set(val *ApisAPIRestRespAllJetStreamConsumers) {
	v.value = val
	v.isSet = true
}

func (v NullableApisAPIRestRespAllJetStreamConsumers) IsSet() bool {
	return v.isSet
}

func (v *NullableApisAPIRestRespAllJetStreamConsumers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApisAPIRestRespAllJetStreamConsumers(val *ApisAPIRestRespAllJetStreamConsumers) *NullableApisAPIRestRespAllJetStreamConsumers {
	return &NullableApisAPIRestRespAllJetStreamConsumers{value: val, isSet: true}
}

func (v NullableApisAPIRestRespAllJetStreamConsumers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApisAPIRestRespAllJetStreamConsumers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
