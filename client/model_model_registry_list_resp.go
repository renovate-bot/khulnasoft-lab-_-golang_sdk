/*
Khulnasoft Kengine

Khulnasoft Runtime API provides programmatic control over Khulnasoft microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.1
Contact: community@khulnasoft.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelRegistryListResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRegistryListResp{}

// ModelRegistryListResp struct for ModelRegistryListResp
type ModelRegistryListResp struct {
	CreatedAt *int32 `json:"created_at,omitempty"`
	Id *int32 `json:"id,omitempty"`
	IsSyncing *bool `json:"is_syncing,omitempty"`
	Name *string `json:"name,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	NonSecret interface{} `json:"non_secret,omitempty"`
	RegistryType *string `json:"registry_type,omitempty"`
	UpdatedAt *int32 `json:"updated_at,omitempty"`
}

// NewModelRegistryListResp instantiates a new ModelRegistryListResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRegistryListResp() *ModelRegistryListResp {
	this := ModelRegistryListResp{}
	return &this
}

// NewModelRegistryListRespWithDefaults instantiates a new ModelRegistryListResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRegistryListRespWithDefaults() *ModelRegistryListResp {
	this := ModelRegistryListResp{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelRegistryListResp) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRegistryListResp) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelRegistryListResp) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *ModelRegistryListResp) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelRegistryListResp) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRegistryListResp) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelRegistryListResp) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelRegistryListResp) SetId(v int32) {
	o.Id = &v
}

// GetIsSyncing returns the IsSyncing field value if set, zero value otherwise.
func (o *ModelRegistryListResp) GetIsSyncing() bool {
	if o == nil || IsNil(o.IsSyncing) {
		var ret bool
		return ret
	}
	return *o.IsSyncing
}

// GetIsSyncingOk returns a tuple with the IsSyncing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRegistryListResp) GetIsSyncingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSyncing) {
		return nil, false
	}
	return o.IsSyncing, true
}

// HasIsSyncing returns a boolean if a field has been set.
func (o *ModelRegistryListResp) HasIsSyncing() bool {
	if o != nil && !IsNil(o.IsSyncing) {
		return true
	}

	return false
}

// SetIsSyncing gets a reference to the given bool and assigns it to the IsSyncing field.
func (o *ModelRegistryListResp) SetIsSyncing(v bool) {
	o.IsSyncing = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelRegistryListResp) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRegistryListResp) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelRegistryListResp) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelRegistryListResp) SetName(v string) {
	o.Name = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ModelRegistryListResp) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRegistryListResp) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ModelRegistryListResp) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *ModelRegistryListResp) SetNodeId(v string) {
	o.NodeId = &v
}

// GetNonSecret returns the NonSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelRegistryListResp) GetNonSecret() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NonSecret
}

// GetNonSecretOk returns a tuple with the NonSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelRegistryListResp) GetNonSecretOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NonSecret) {
		return nil, false
	}
	return &o.NonSecret, true
}

// HasNonSecret returns a boolean if a field has been set.
func (o *ModelRegistryListResp) HasNonSecret() bool {
	if o != nil && !IsNil(o.NonSecret) {
		return true
	}

	return false
}

// SetNonSecret gets a reference to the given interface{} and assigns it to the NonSecret field.
func (o *ModelRegistryListResp) SetNonSecret(v interface{}) {
	o.NonSecret = v
}

// GetRegistryType returns the RegistryType field value if set, zero value otherwise.
func (o *ModelRegistryListResp) GetRegistryType() string {
	if o == nil || IsNil(o.RegistryType) {
		var ret string
		return ret
	}
	return *o.RegistryType
}

// GetRegistryTypeOk returns a tuple with the RegistryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRegistryListResp) GetRegistryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RegistryType) {
		return nil, false
	}
	return o.RegistryType, true
}

// HasRegistryType returns a boolean if a field has been set.
func (o *ModelRegistryListResp) HasRegistryType() bool {
	if o != nil && !IsNil(o.RegistryType) {
		return true
	}

	return false
}

// SetRegistryType gets a reference to the given string and assigns it to the RegistryType field.
func (o *ModelRegistryListResp) SetRegistryType(v string) {
	o.RegistryType = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelRegistryListResp) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRegistryListResp) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelRegistryListResp) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *ModelRegistryListResp) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

func (o ModelRegistryListResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRegistryListResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsSyncing) {
		toSerialize["is_syncing"] = o.IsSyncing
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if o.NonSecret != nil {
		toSerialize["non_secret"] = o.NonSecret
	}
	if !IsNil(o.RegistryType) {
		toSerialize["registry_type"] = o.RegistryType
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableModelRegistryListResp struct {
	value *ModelRegistryListResp
	isSet bool
}

func (v NullableModelRegistryListResp) Get() *ModelRegistryListResp {
	return v.value
}

func (v *NullableModelRegistryListResp) Set(val *ModelRegistryListResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRegistryListResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRegistryListResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRegistryListResp(val *ModelRegistryListResp) *NullableModelRegistryListResp {
	return &NullableModelRegistryListResp{value: val, isSet: true}
}

func (v NullableModelRegistryListResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRegistryListResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


