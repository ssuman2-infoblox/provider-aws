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
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// VPCLinkParameters defines the desired state of VPCLink
type VPCLinkParameters struct {
	// Region is which region the VPCLink will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The description of the VPC link.
	Description *string `json:"description,omitempty"`
	// The name used to label and identify the VPC link.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/].
	// The tag key can be up to 128 characters and must not start with aws:. The
	// tag value can be up to 256 characters.
	Tags map[string]*string `json:"tags,omitempty"`
	// The ARN of the network load balancer of the VPC targeted by the VPC link.
	// The network load balancer must be owned by the same Amazon Web Services account
	// of the API owner.
	// +kubebuilder:validation:Required
	TargetARNs              []*string `json:"targetARNs"`
	CustomVPCLinkParameters `json:",inline"`
}

// VPCLinkSpec defines the desired state of VPCLink
type VPCLinkSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VPCLinkParameters `json:"forProvider"`
}

// VPCLinkObservation defines the observed state of VPCLink
type VPCLinkObservation struct {
	// The identifier of the VpcLink. It is used in an Integration to reference
	// this VpcLink.
	ID *string `json:"id,omitempty"`
	// The status of the VPC link. The valid values are AVAILABLE, PENDING, DELETING,
	// or FAILED. Deploying an API will wait if the status is PENDING and will fail
	// if the status is DELETING.
	Status *string `json:"status,omitempty"`
	// A description about the VPC link status.
	StatusMessage *string `json:"statusMessage,omitempty"`
}

// VPCLinkStatus defines the observed state of VPCLink.
type VPCLinkStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VPCLinkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCLink is the Schema for the VPCLinks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCLinkSpec   `json:"spec"`
	Status            VPCLinkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCLinkList contains a list of VPCLinks
type VPCLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCLink `json:"items"`
}

// Repository type metadata.
var (
	VPCLinkKind             = "VPCLink"
	VPCLinkGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCLinkKind}.String()
	VPCLinkKindAPIVersion   = VPCLinkKind + "." + GroupVersion.String()
	VPCLinkGroupVersionKind = GroupVersion.WithKind(VPCLinkKind)
)

func init() {
	SchemeBuilder.Register(&VPCLink{}, &VPCLinkList{})
}
