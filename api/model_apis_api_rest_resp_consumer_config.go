/*
httpmq

HTTP/2 based message broker built around NATS JetStream

API version: v0.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ApisAPIRestRespConsumerConfig struct for ApisAPIRestRespConsumerConfig
type ApisAPIRestRespConsumerConfig struct {
	// AckWait duration (ns) to wait for an ACK for the delivery of a message
	AckWait int64 `json:"ack_wait"`
	// DeliverGroup is the delivery group if this consumer uses delivery group  A consumer using delivery group allows multiple clients to subscribe under the same consumer and group name tuple. For subjects this consumer listens to, the messages will be shared amongst the connected clients.
	DeliverGroup *string `json:"deliver_group,omitempty"`
	// DeliverSubject subject this consumer is listening on
	DeliverSubject *string `json:"deliver_subject,omitempty"`
	// FilterSubject sets the consumer to filter for subjects matching this NATs subject string  See https://docs.nats.io/running-a-nats-service/nats_admin/jetstream_admin/naming
	FilterSubject *string `json:"filter_subject,omitempty"`
	// MaxAckPending controls the max number of un-ACKed messages permitted in-flight
	MaxAckPending *int64 `json:"max_ack_pending,omitempty"`
	// MaxDeliver max number of times a message can be deliveried (including retry) to this consumer
	MaxDeliver *int64 `json:"max_deliver,omitempty"`
	// MaxWaiting NATS JetStream does not clearly document this
	MaxWaiting *int64 `json:"max_waiting,omitempty"`
	// Description an optional description of the consumer
	Notes *string `json:"notes,omitempty"`
}

// NewApisAPIRestRespConsumerConfig instantiates a new ApisAPIRestRespConsumerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApisAPIRestRespConsumerConfig(ackWait int64) *ApisAPIRestRespConsumerConfig {
	this := ApisAPIRestRespConsumerConfig{}
	this.AckWait = ackWait
	return &this
}

// NewApisAPIRestRespConsumerConfigWithDefaults instantiates a new ApisAPIRestRespConsumerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApisAPIRestRespConsumerConfigWithDefaults() *ApisAPIRestRespConsumerConfig {
	this := ApisAPIRestRespConsumerConfig{}
	return &this
}

// GetAckWait returns the AckWait field value
func (o *ApisAPIRestRespConsumerConfig) GetAckWait() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AckWait
}

// GetAckWaitOk returns a tuple with the AckWait field value
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerConfig) GetAckWaitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AckWait, true
}

// SetAckWait sets field value
func (o *ApisAPIRestRespConsumerConfig) SetAckWait(v int64) {
	o.AckWait = v
}

// GetDeliverGroup returns the DeliverGroup field value if set, zero value otherwise.
func (o *ApisAPIRestRespConsumerConfig) GetDeliverGroup() string {
	if o == nil || o.DeliverGroup == nil {
		var ret string
		return ret
	}
	return *o.DeliverGroup
}

// GetDeliverGroupOk returns a tuple with the DeliverGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerConfig) GetDeliverGroupOk() (*string, bool) {
	if o == nil || o.DeliverGroup == nil {
		return nil, false
	}
	return o.DeliverGroup, true
}

// HasDeliverGroup returns a boolean if a field has been set.
func (o *ApisAPIRestRespConsumerConfig) HasDeliverGroup() bool {
	if o != nil && o.DeliverGroup != nil {
		return true
	}

	return false
}

// SetDeliverGroup gets a reference to the given string and assigns it to the DeliverGroup field.
func (o *ApisAPIRestRespConsumerConfig) SetDeliverGroup(v string) {
	o.DeliverGroup = &v
}

// GetDeliverSubject returns the DeliverSubject field value if set, zero value otherwise.
func (o *ApisAPIRestRespConsumerConfig) GetDeliverSubject() string {
	if o == nil || o.DeliverSubject == nil {
		var ret string
		return ret
	}
	return *o.DeliverSubject
}

// GetDeliverSubjectOk returns a tuple with the DeliverSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerConfig) GetDeliverSubjectOk() (*string, bool) {
	if o == nil || o.DeliverSubject == nil {
		return nil, false
	}
	return o.DeliverSubject, true
}

// HasDeliverSubject returns a boolean if a field has been set.
func (o *ApisAPIRestRespConsumerConfig) HasDeliverSubject() bool {
	if o != nil && o.DeliverSubject != nil {
		return true
	}

	return false
}

// SetDeliverSubject gets a reference to the given string and assigns it to the DeliverSubject field.
func (o *ApisAPIRestRespConsumerConfig) SetDeliverSubject(v string) {
	o.DeliverSubject = &v
}

// GetFilterSubject returns the FilterSubject field value if set, zero value otherwise.
func (o *ApisAPIRestRespConsumerConfig) GetFilterSubject() string {
	if o == nil || o.FilterSubject == nil {
		var ret string
		return ret
	}
	return *o.FilterSubject
}

// GetFilterSubjectOk returns a tuple with the FilterSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerConfig) GetFilterSubjectOk() (*string, bool) {
	if o == nil || o.FilterSubject == nil {
		return nil, false
	}
	return o.FilterSubject, true
}

// HasFilterSubject returns a boolean if a field has been set.
func (o *ApisAPIRestRespConsumerConfig) HasFilterSubject() bool {
	if o != nil && o.FilterSubject != nil {
		return true
	}

	return false
}

// SetFilterSubject gets a reference to the given string and assigns it to the FilterSubject field.
func (o *ApisAPIRestRespConsumerConfig) SetFilterSubject(v string) {
	o.FilterSubject = &v
}

// GetMaxAckPending returns the MaxAckPending field value if set, zero value otherwise.
func (o *ApisAPIRestRespConsumerConfig) GetMaxAckPending() int64 {
	if o == nil || o.MaxAckPending == nil {
		var ret int64
		return ret
	}
	return *o.MaxAckPending
}

// GetMaxAckPendingOk returns a tuple with the MaxAckPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerConfig) GetMaxAckPendingOk() (*int64, bool) {
	if o == nil || o.MaxAckPending == nil {
		return nil, false
	}
	return o.MaxAckPending, true
}

// HasMaxAckPending returns a boolean if a field has been set.
func (o *ApisAPIRestRespConsumerConfig) HasMaxAckPending() bool {
	if o != nil && o.MaxAckPending != nil {
		return true
	}

	return false
}

// SetMaxAckPending gets a reference to the given int64 and assigns it to the MaxAckPending field.
func (o *ApisAPIRestRespConsumerConfig) SetMaxAckPending(v int64) {
	o.MaxAckPending = &v
}

// GetMaxDeliver returns the MaxDeliver field value if set, zero value otherwise.
func (o *ApisAPIRestRespConsumerConfig) GetMaxDeliver() int64 {
	if o == nil || o.MaxDeliver == nil {
		var ret int64
		return ret
	}
	return *o.MaxDeliver
}

// GetMaxDeliverOk returns a tuple with the MaxDeliver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerConfig) GetMaxDeliverOk() (*int64, bool) {
	if o == nil || o.MaxDeliver == nil {
		return nil, false
	}
	return o.MaxDeliver, true
}

// HasMaxDeliver returns a boolean if a field has been set.
func (o *ApisAPIRestRespConsumerConfig) HasMaxDeliver() bool {
	if o != nil && o.MaxDeliver != nil {
		return true
	}

	return false
}

// SetMaxDeliver gets a reference to the given int64 and assigns it to the MaxDeliver field.
func (o *ApisAPIRestRespConsumerConfig) SetMaxDeliver(v int64) {
	o.MaxDeliver = &v
}

// GetMaxWaiting returns the MaxWaiting field value if set, zero value otherwise.
func (o *ApisAPIRestRespConsumerConfig) GetMaxWaiting() int64 {
	if o == nil || o.MaxWaiting == nil {
		var ret int64
		return ret
	}
	return *o.MaxWaiting
}

// GetMaxWaitingOk returns a tuple with the MaxWaiting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerConfig) GetMaxWaitingOk() (*int64, bool) {
	if o == nil || o.MaxWaiting == nil {
		return nil, false
	}
	return o.MaxWaiting, true
}

// HasMaxWaiting returns a boolean if a field has been set.
func (o *ApisAPIRestRespConsumerConfig) HasMaxWaiting() bool {
	if o != nil && o.MaxWaiting != nil {
		return true
	}

	return false
}

// SetMaxWaiting gets a reference to the given int64 and assigns it to the MaxWaiting field.
func (o *ApisAPIRestRespConsumerConfig) SetMaxWaiting(v int64) {
	o.MaxWaiting = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ApisAPIRestRespConsumerConfig) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApisAPIRestRespConsumerConfig) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ApisAPIRestRespConsumerConfig) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ApisAPIRestRespConsumerConfig) SetNotes(v string) {
	o.Notes = &v
}

func (o ApisAPIRestRespConsumerConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ack_wait"] = o.AckWait
	}
	if o.DeliverGroup != nil {
		toSerialize["deliver_group"] = o.DeliverGroup
	}
	if o.DeliverSubject != nil {
		toSerialize["deliver_subject"] = o.DeliverSubject
	}
	if o.FilterSubject != nil {
		toSerialize["filter_subject"] = o.FilterSubject
	}
	if o.MaxAckPending != nil {
		toSerialize["max_ack_pending"] = o.MaxAckPending
	}
	if o.MaxDeliver != nil {
		toSerialize["max_deliver"] = o.MaxDeliver
	}
	if o.MaxWaiting != nil {
		toSerialize["max_waiting"] = o.MaxWaiting
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	return json.Marshal(toSerialize)
}

type NullableApisAPIRestRespConsumerConfig struct {
	value *ApisAPIRestRespConsumerConfig
	isSet bool
}

func (v NullableApisAPIRestRespConsumerConfig) Get() *ApisAPIRestRespConsumerConfig {
	return v.value
}

func (v *NullableApisAPIRestRespConsumerConfig) Set(val *ApisAPIRestRespConsumerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableApisAPIRestRespConsumerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableApisAPIRestRespConsumerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApisAPIRestRespConsumerConfig(val *ApisAPIRestRespConsumerConfig) *NullableApisAPIRestRespConsumerConfig {
	return &NullableApisAPIRestRespConsumerConfig{value: val, isSet: true}
}

func (v NullableApisAPIRestRespConsumerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApisAPIRestRespConsumerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}