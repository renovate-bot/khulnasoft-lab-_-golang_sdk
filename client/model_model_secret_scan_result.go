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
	"bytes"
	"fmt"
)

// checks if the ModelSecretScanResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSecretScanResult{}

// ModelSecretScanResult struct for ModelSecretScanResult
type ModelSecretScanResult struct {
	CloudAccountId string `json:"cloud_account_id"`
	CreatedAt int64 `json:"created_at"`
	DockerContainerName string `json:"docker_container_name"`
	DockerImageName string `json:"docker_image_name"`
	HostName string `json:"host_name"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	NodeType string `json:"node_type"`
	ScanId string `json:"scan_id"`
	Secrets []ModelSecret `json:"secrets"`
	SeverityCounts map[string]int32 `json:"severity_counts"`
	UpdatedAt int64 `json:"updated_at"`
}

type _ModelSecretScanResult ModelSecretScanResult

// NewModelSecretScanResult instantiates a new ModelSecretScanResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSecretScanResult(cloudAccountId string, createdAt int64, dockerContainerName string, dockerImageName string, hostName string, kubernetesClusterName string, nodeId string, nodeName string, nodeType string, scanId string, secrets []ModelSecret, severityCounts map[string]int32, updatedAt int64) *ModelSecretScanResult {
	this := ModelSecretScanResult{}
	this.CloudAccountId = cloudAccountId
	this.CreatedAt = createdAt
	this.DockerContainerName = dockerContainerName
	this.DockerImageName = dockerImageName
	this.HostName = hostName
	this.KubernetesClusterName = kubernetesClusterName
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.NodeType = nodeType
	this.ScanId = scanId
	this.Secrets = secrets
	this.SeverityCounts = severityCounts
	this.UpdatedAt = updatedAt
	return &this
}

// NewModelSecretScanResultWithDefaults instantiates a new ModelSecretScanResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSecretScanResultWithDefaults() *ModelSecretScanResult {
	this := ModelSecretScanResult{}
	return &this
}

// GetCloudAccountId returns the CloudAccountId field value
func (o *ModelSecretScanResult) GetCloudAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudAccountId
}

// GetCloudAccountIdOk returns a tuple with the CloudAccountId field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetCloudAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudAccountId, true
}

// SetCloudAccountId sets field value
func (o *ModelSecretScanResult) SetCloudAccountId(v string) {
	o.CloudAccountId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ModelSecretScanResult) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ModelSecretScanResult) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetDockerContainerName returns the DockerContainerName field value
func (o *ModelSecretScanResult) GetDockerContainerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerContainerName
}

// GetDockerContainerNameOk returns a tuple with the DockerContainerName field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetDockerContainerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerContainerName, true
}

// SetDockerContainerName sets field value
func (o *ModelSecretScanResult) SetDockerContainerName(v string) {
	o.DockerContainerName = v
}

// GetDockerImageName returns the DockerImageName field value
func (o *ModelSecretScanResult) GetDockerImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageName
}

// GetDockerImageNameOk returns a tuple with the DockerImageName field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetDockerImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerImageName, true
}

// SetDockerImageName sets field value
func (o *ModelSecretScanResult) SetDockerImageName(v string) {
	o.DockerImageName = v
}

// GetHostName returns the HostName field value
func (o *ModelSecretScanResult) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *ModelSecretScanResult) SetHostName(v string) {
	o.HostName = v
}

// GetKubernetesClusterName returns the KubernetesClusterName field value
func (o *ModelSecretScanResult) GetKubernetesClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesClusterName
}

// GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetKubernetesClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesClusterName, true
}

// SetKubernetesClusterName sets field value
func (o *ModelSecretScanResult) SetKubernetesClusterName(v string) {
	o.KubernetesClusterName = v
}

// GetNodeId returns the NodeId field value
func (o *ModelSecretScanResult) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelSecretScanResult) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelSecretScanResult) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelSecretScanResult) SetNodeName(v string) {
	o.NodeName = v
}

// GetNodeType returns the NodeType field value
func (o *ModelSecretScanResult) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *ModelSecretScanResult) SetNodeType(v string) {
	o.NodeType = v
}

// GetScanId returns the ScanId field value
func (o *ModelSecretScanResult) GetScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanId, true
}

// SetScanId sets field value
func (o *ModelSecretScanResult) SetScanId(v string) {
	o.ScanId = v
}

// GetSecrets returns the Secrets field value
// If the value is explicit nil, the zero value for []ModelSecret will be returned
func (o *ModelSecretScanResult) GetSecrets() []ModelSecret {
	if o == nil {
		var ret []ModelSecret
		return ret
	}

	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSecretScanResult) GetSecretsOk() ([]ModelSecret, bool) {
	if o == nil || IsNil(o.Secrets) {
		return nil, false
	}
	return o.Secrets, true
}

// SetSecrets sets field value
func (o *ModelSecretScanResult) SetSecrets(v []ModelSecret) {
	o.Secrets = v
}

// GetSeverityCounts returns the SeverityCounts field value
// If the value is explicit nil, the zero value for map[string]int32 will be returned
func (o *ModelSecretScanResult) GetSeverityCounts() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.SeverityCounts
}

// GetSeverityCountsOk returns a tuple with the SeverityCounts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSecretScanResult) GetSeverityCountsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.SeverityCounts) {
		return nil, false
	}
	return &o.SeverityCounts, true
}

// SetSeverityCounts sets field value
func (o *ModelSecretScanResult) SetSeverityCounts(v map[string]int32) {
	o.SeverityCounts = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelSecretScanResult) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanResult) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelSecretScanResult) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

func (o ModelSecretScanResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSecretScanResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud_account_id"] = o.CloudAccountId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["docker_container_name"] = o.DockerContainerName
	toSerialize["docker_image_name"] = o.DockerImageName
	toSerialize["host_name"] = o.HostName
	toSerialize["kubernetes_cluster_name"] = o.KubernetesClusterName
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	toSerialize["node_type"] = o.NodeType
	toSerialize["scan_id"] = o.ScanId
	if o.Secrets != nil {
		toSerialize["secrets"] = o.Secrets
	}
	if o.SeverityCounts != nil {
		toSerialize["severity_counts"] = o.SeverityCounts
	}
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ModelSecretScanResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud_account_id",
		"created_at",
		"docker_container_name",
		"docker_image_name",
		"host_name",
		"kubernetes_cluster_name",
		"node_id",
		"node_name",
		"node_type",
		"scan_id",
		"secrets",
		"severity_counts",
		"updated_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varModelSecretScanResult := _ModelSecretScanResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelSecretScanResult)

	if err != nil {
		return err
	}

	*o = ModelSecretScanResult(varModelSecretScanResult)

	return err
}

type NullableModelSecretScanResult struct {
	value *ModelSecretScanResult
	isSet bool
}

func (v NullableModelSecretScanResult) Get() *ModelSecretScanResult {
	return v.value
}

func (v *NullableModelSecretScanResult) Set(val *ModelSecretScanResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSecretScanResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSecretScanResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSecretScanResult(val *ModelSecretScanResult) *NullableModelSecretScanResult {
	return &NullableModelSecretScanResult{value: val, isSet: true}
}

func (v NullableModelSecretScanResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSecretScanResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


