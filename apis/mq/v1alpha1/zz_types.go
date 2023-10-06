/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type ActionRequired struct {
	ActionRequiredCode *string `json:"actionRequiredCode,omitempty"`

	ActionRequiredInfo *string `json:"actionRequiredInfo,omitempty"`
}

// +kubebuilder:skipversion
type AvailabilityZone struct {
	Name *string `json:"name,omitempty"`
}

// +kubebuilder:skipversion
type BrokerEngineType struct {
	// The type of broker engine. Amazon MQ supports ActiveMQ and RabbitMQ.
	EngineType *string `json:"engineType,omitempty"`
}

// +kubebuilder:skipversion
type BrokerInstance struct {
	ConsoleURL *string `json:"consoleURL,omitempty"`

	Endpoints []*string `json:"endpoints,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty"`
}

// +kubebuilder:skipversion
type BrokerInstanceOption struct {
	// The type of broker engine. Amazon MQ supports ActiveMQ and RabbitMQ.
	EngineType *string `json:"engineType,omitempty"`

	HostInstanceType *string `json:"hostInstanceType,omitempty"`
	// The broker's storage type.
	//
	// EFS is not supported for RabbitMQ engine type.
	StorageType *string `json:"storageType,omitempty"`

	SupportedEngineVersions []*string `json:"supportedEngineVersions,omitempty"`
}

// +kubebuilder:skipversion
type BrokerSummary struct {
	BrokerARN *string `json:"brokerARN,omitempty"`

	BrokerID *string `json:"brokerID,omitempty"`

	BrokerName *string `json:"brokerName,omitempty"`
	// The broker's status.
	BrokerState *string `json:"brokerState,omitempty"`

	Created *metav1.Time `json:"created,omitempty"`
	// The broker's deployment mode.
	DeploymentMode *string `json:"deploymentMode,omitempty"`
	// The type of broker engine. Amazon MQ supports ActiveMQ and RabbitMQ.
	EngineType *string `json:"engineType,omitempty"`

	HostInstanceType *string `json:"hostInstanceType,omitempty"`
}

// +kubebuilder:skipversion
type Configuration struct {
	ARN *string `json:"arn,omitempty"`
	// Optional. The authentication strategy used to secure the broker. The default
	// is SIMPLE.
	AuthenticationStrategy *string `json:"authenticationStrategy,omitempty"`

	Created *metav1.Time `json:"created,omitempty"`

	Description *string `json:"description,omitempty"`
	// The type of broker engine. Amazon MQ supports ActiveMQ and RabbitMQ.
	EngineType *string `json:"engineType,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Tags map[string]*string `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type ConfigurationID struct {
	ID *string `json:"id,omitempty"`

	Revision *int64 `json:"revision,omitempty"`
}

// +kubebuilder:skipversion
type ConfigurationRevision struct {
	Created *metav1.Time `json:"created,omitempty"`

	Description *string `json:"description,omitempty"`

	Revision *int64 `json:"revision,omitempty"`
}

// +kubebuilder:skipversion
type Configurations struct {
	// A list of information about the configuration.
	Current *ConfigurationID `json:"current,omitempty"`

	History []*ConfigurationID `json:"history,omitempty"`
	// A list of information about the configuration.
	Pending *ConfigurationID `json:"pending,omitempty"`
}

// +kubebuilder:skipversion
type DataReplicationCounterpart struct {
	BrokerID *string `json:"brokerID,omitempty"`

	Region *string `json:"region,omitempty"`
}

// +kubebuilder:skipversion
type DataReplicationMetadataOutput struct {
	// Specifies a broker in a data replication pair.
	DataReplicationCounterpart *DataReplicationCounterpart `json:"dataReplicationCounterpart,omitempty"`

	DataReplicationRole *string `json:"dataReplicationRole,omitempty"`
}

// +kubebuilder:skipversion
type EncryptionOptions struct {
	KMSKeyID *string `json:"kmsKeyID,omitempty"`

	UseAWSOwnedKey *bool `json:"useAWSOwnedKey,omitempty"`
}

// +kubebuilder:skipversion
type EngineVersion struct {
	Name *string `json:"name,omitempty"`
}

// +kubebuilder:skipversion
type LDAPServerMetadataInput struct {
	Hosts []*string `json:"hosts,omitempty"`

	RoleBase *string `json:"roleBase,omitempty"`

	RoleName *string `json:"roleName,omitempty"`

	RoleSearchMatching *string `json:"roleSearchMatching,omitempty"`

	RoleSearchSubtree *bool `json:"roleSearchSubtree,omitempty"`

	ServiceAccountPassword *string `json:"serviceAccountPassword,omitempty"`

	ServiceAccountUsername *string `json:"serviceAccountUsername,omitempty"`

	UserBase *string `json:"userBase,omitempty"`

	UserRoleName *string `json:"userRoleName,omitempty"`

	UserSearchMatching *string `json:"userSearchMatching,omitempty"`

	UserSearchSubtree *bool `json:"userSearchSubtree,omitempty"`
}

// +kubebuilder:skipversion
type LDAPServerMetadataOutput struct {
	Hosts []*string `json:"hosts,omitempty"`

	RoleBase *string `json:"roleBase,omitempty"`

	RoleName *string `json:"roleName,omitempty"`

	RoleSearchMatching *string `json:"roleSearchMatching,omitempty"`

	RoleSearchSubtree *bool `json:"roleSearchSubtree,omitempty"`

	ServiceAccountUsername *string `json:"serviceAccountUsername,omitempty"`

	UserBase *string `json:"userBase,omitempty"`

	UserRoleName *string `json:"userRoleName,omitempty"`

	UserSearchMatching *string `json:"userSearchMatching,omitempty"`

	UserSearchSubtree *bool `json:"userSearchSubtree,omitempty"`
}

// +kubebuilder:skipversion
type Logs struct {
	Audit *bool `json:"audit,omitempty"`

	General *bool `json:"general,omitempty"`
}

// +kubebuilder:skipversion
type LogsSummary struct {
	Audit *bool `json:"audit,omitempty"`

	AuditLogGroup *string `json:"auditLogGroup,omitempty"`

	General *bool `json:"general,omitempty"`

	GeneralLogGroup *string `json:"generalLogGroup,omitempty"`
	// The list of information about logs to be enabled for the specified broker.
	Pending *PendingLogs `json:"pending,omitempty"`
}

// +kubebuilder:skipversion
type PendingLogs struct {
	Audit *bool `json:"audit,omitempty"`

	General *bool `json:"general,omitempty"`
}

// +kubebuilder:skipversion
type SanitizationWarning struct {
	AttributeName *string `json:"attributeName,omitempty"`

	ElementName *string `json:"elementName,omitempty"`
}

// +kubebuilder:skipversion
type UserPendingChanges struct {
	ConsoleAccess *bool `json:"consoleAccess,omitempty"`

	Groups []*string `json:"groups,omitempty"`
	// The type of change pending for the ActiveMQ user.
	PendingChange *string `json:"pendingChange,omitempty"`
}

// +kubebuilder:skipversion
type UserSummary struct {
	// The type of change pending for the ActiveMQ user.
	PendingChange *string `json:"pendingChange,omitempty"`

	Username *string `json:"username,omitempty"`
}

// +kubebuilder:skipversion
type User_SDK struct {
	ConsoleAccess *bool `json:"consoleAccess,omitempty"`

	Groups []*string `json:"groups,omitempty"`

	Password *string `json:"password,omitempty"`

	ReplicationUser *bool `json:"replicationUser,omitempty"`

	Username *string `json:"username,omitempty"`
}

// +kubebuilder:skipversion
type WeeklyStartTime struct {
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	TimeOfDay *string `json:"timeOfDay,omitempty"`

	TimeZone *string `json:"timeZone,omitempty"`
}
