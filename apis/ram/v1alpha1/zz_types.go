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
type AssociatedPermission struct {
	ARN *string `json:"arn,omitempty"`

	DefaultVersion *bool `json:"defaultVersion,omitempty"`

	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`

	PermissionVersion *string `json:"permissionVersion,omitempty"`

	ResourceShareARN *string `json:"resourceShareARN,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`

	Status *string `json:"status,omitempty"`
}

// +kubebuilder:skipversion
type Principal struct {
	CreationTime *metav1.Time `json:"creationTime,omitempty"`

	External *bool `json:"external,omitempty"`

	ID *string `json:"id,omitempty"`

	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`

	ResourceShareARN *string `json:"resourceShareARN,omitempty"`
}

// +kubebuilder:skipversion
type ReplacePermissionAssociationsWork struct {
	CreationTime *metav1.Time `json:"creationTime,omitempty"`

	FromPermissionARN *string `json:"fromPermissionARN,omitempty"`

	FromPermissionVersion *string `json:"fromPermissionVersion,omitempty"`

	ID *string `json:"id,omitempty"`

	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`

	StatusMessage *string `json:"statusMessage,omitempty"`

	ToPermissionARN *string `json:"toPermissionARN,omitempty"`

	ToPermissionVersion *string `json:"toPermissionVersion,omitempty"`
}

// +kubebuilder:skipversion
type Resource struct {
	ARN *string `json:"arn,omitempty"`

	CreationTime *metav1.Time `json:"creationTime,omitempty"`

	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`

	ResourceGroupARN *string `json:"resourceGroupARN,omitempty"`

	ResourceShareARN *string `json:"resourceShareARN,omitempty"`

	StatusMessage *string `json:"statusMessage,omitempty"`

	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type ResourceShareAssociation struct {
	AssociatedEntity *string `json:"associatedEntity,omitempty"`

	CreationTime *metav1.Time `json:"creationTime,omitempty"`

	External *bool `json:"external,omitempty"`

	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`

	ResourceShareARN *string `json:"resourceShareARN,omitempty"`

	ResourceShareName *string `json:"resourceShareName,omitempty"`

	StatusMessage *string `json:"statusMessage,omitempty"`
}

// +kubebuilder:skipversion
type ResourceShareInvitation struct {
	InvitationTimestamp *metav1.Time `json:"invitationTimestamp,omitempty"`

	ReceiverAccountID *string `json:"receiverAccountID,omitempty"`

	ReceiverARN *string `json:"receiverARN,omitempty"`

	ResourceShareARN *string `json:"resourceShareARN,omitempty"`

	ResourceShareInvitationARN *string `json:"resourceShareInvitationARN,omitempty"`

	ResourceShareName *string `json:"resourceShareName,omitempty"`

	SenderAccountID *string `json:"senderAccountID,omitempty"`
}

// +kubebuilder:skipversion
type ResourceSharePermissionDetail struct {
	ARN *string `json:"arn,omitempty"`

	CreationTime *metav1.Time `json:"creationTime,omitempty"`

	DefaultVersion *bool `json:"defaultVersion,omitempty"`

	IsResourceTypeDefault *bool `json:"isResourceTypeDefault,omitempty"`

	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`

	Name *string `json:"name,omitempty"`

	Permission *string `json:"permission,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	Version *string `json:"version,omitempty"`
}

// +kubebuilder:skipversion
type ResourceSharePermissionSummary struct {
	ARN *string `json:"arn,omitempty"`

	CreationTime *metav1.Time `json:"creationTime,omitempty"`

	DefaultVersion *bool `json:"defaultVersion,omitempty"`

	IsResourceTypeDefault *bool `json:"isResourceTypeDefault,omitempty"`

	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`

	Name *string `json:"name,omitempty"`

	ResourceType *string `json:"resourceType,omitempty"`

	Status *string `json:"status,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	Version *string `json:"version,omitempty"`
}

// +kubebuilder:skipversion
type ResourceShare_SDK struct {
	AllowExternalPrincipals *bool `json:"allowExternalPrincipals,omitempty"`

	CreationTime *metav1.Time `json:"creationTime,omitempty"`

	FeatureSet *string `json:"featureSet,omitempty"`

	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`

	Name *string `json:"name,omitempty"`

	OwningAccountID *string `json:"owningAccountID,omitempty"`

	ResourceShareARN *string `json:"resourceShareARN,omitempty"`

	Status *string `json:"status,omitempty"`

	StatusMessage *string `json:"statusMessage,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type ServiceNameAndResourceType struct {
	ResourceType *string `json:"resourceType,omitempty"`

	ServiceName *string `json:"serviceName,omitempty"`
}

// +kubebuilder:skipversion
type Tag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type TagFilter struct {
	TagKey *string `json:"tagKey,omitempty"`

	TagValues []*string `json:"tagValues,omitempty"`
}
