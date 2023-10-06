//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane-contrib/provider-aws/apis/s3/common"
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPoint) DeepCopyInto(out *AccessPoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPoint.
func (in *AccessPoint) DeepCopy() *AccessPoint {
	if in == nil {
		return nil
	}
	out := new(AccessPoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessPoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointList) DeepCopyInto(out *AccessPointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessPoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointList.
func (in *AccessPointList) DeepCopy() *AccessPointList {
	if in == nil {
		return nil
	}
	out := new(AccessPointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessPointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointObservation) DeepCopyInto(out *AccessPointObservation) {
	*out = *in
	if in.AccessPointARN != nil {
		in, out := &in.AccessPointARN, &out.AccessPointARN
		*out = new(string)
		**out = **in
	}
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointObservation.
func (in *AccessPointObservation) DeepCopy() *AccessPointObservation {
	if in == nil {
		return nil
	}
	out := new(AccessPointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointParameters) DeepCopyInto(out *AccessPointParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.BucketAccountID != nil {
		in, out := &in.BucketAccountID, &out.BucketAccountID
		*out = new(string)
		**out = **in
	}
	if in.PublicAccessBlockConfiguration != nil {
		in, out := &in.PublicAccessBlockConfiguration, &out.PublicAccessBlockConfiguration
		*out = new(PublicAccessBlockConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCConfiguration != nil {
		in, out := &in.VPCConfiguration, &out.VPCConfiguration
		*out = new(VPCConfiguration)
		(*in).DeepCopyInto(*out)
	}
	in.CustomAccessPointParameters.DeepCopyInto(&out.CustomAccessPointParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointParameters.
func (in *AccessPointParameters) DeepCopy() *AccessPointParameters {
	if in == nil {
		return nil
	}
	out := new(AccessPointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointSpec) DeepCopyInto(out *AccessPointSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointSpec.
func (in *AccessPointSpec) DeepCopy() *AccessPointSpec {
	if in == nil {
		return nil
	}
	out := new(AccessPointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointStatus) DeepCopyInto(out *AccessPointStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointStatus.
func (in *AccessPointStatus) DeepCopy() *AccessPointStatus {
	if in == nil {
		return nil
	}
	out := new(AccessPointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPoint_SDK) DeepCopyInto(out *AccessPoint_SDK) {
	*out = *in
	if in.AccessPointARN != nil {
		in, out := &in.AccessPointARN, &out.AccessPointARN
		*out = new(string)
		**out = **in
	}
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		*out = new(string)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.BucketAccountID != nil {
		in, out := &in.BucketAccountID, &out.BucketAccountID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkOrigin != nil {
		in, out := &in.NetworkOrigin, &out.NetworkOrigin
		*out = new(string)
		**out = **in
	}
	if in.VPCConfiguration != nil {
		in, out := &in.VPCConfiguration, &out.VPCConfiguration
		*out = new(VPCConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPoint_SDK.
func (in *AccessPoint_SDK) DeepCopy() *AccessPoint_SDK {
	if in == nil {
		return nil
	}
	out := new(AccessPoint_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreateMultiRegionAccessPointInput) DeepCopyInto(out *CreateMultiRegionAccessPointInput) {
	*out = *in
	if in.PublicAccessBlock != nil {
		in, out := &in.PublicAccessBlock, &out.PublicAccessBlock
		*out = new(PublicAccessBlockConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreateMultiRegionAccessPointInput.
func (in *CreateMultiRegionAccessPointInput) DeepCopy() *CreateMultiRegionAccessPointInput {
	if in == nil {
		return nil
	}
	out := new(CreateMultiRegionAccessPointInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomAccessPointParameters) DeepCopyInto(out *CustomAccessPointParameters) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.BucketNameRef != nil {
		in, out := &in.BucketNameRef, &out.BucketNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.BucketNameSelector != nil {
		in, out := &in.BucketNameSelector, &out.BucketNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(common.BucketPolicyBody)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomAccessPointParameters.
func (in *CustomAccessPointParameters) DeepCopy() *CustomAccessPointParameters {
	if in == nil {
		return nil
	}
	out := new(CustomAccessPointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destination) DeepCopyInto(out *Destination) {
	*out = *in
	if in.Account != nil {
		in, out := &in.Account, &out.Account
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobManifestLocation) DeepCopyInto(out *JobManifestLocation) {
	*out = *in
	if in.ETag != nil {
		in, out := &in.ETag, &out.ETag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobManifestLocation.
func (in *JobManifestLocation) DeepCopy() *JobManifestLocation {
	if in == nil {
		return nil
	}
	out := new(JobManifestLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRegionAccessPointReport) DeepCopyInto(out *MultiRegionAccessPointReport) {
	*out = *in
	if in.PublicAccessBlock != nil {
		in, out := &in.PublicAccessBlock, &out.PublicAccessBlock
		*out = new(PublicAccessBlockConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRegionAccessPointReport.
func (in *MultiRegionAccessPointReport) DeepCopy() *MultiRegionAccessPointReport {
	if in == nil {
		return nil
	}
	out := new(MultiRegionAccessPointReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiRegionAccessPointRoute) DeepCopyInto(out *MultiRegionAccessPointRoute) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiRegionAccessPointRoute.
func (in *MultiRegionAccessPointRoute) DeepCopy() *MultiRegionAccessPointRoute {
	if in == nil {
		return nil
	}
	out := new(MultiRegionAccessPointRoute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicAccessBlockConfiguration) DeepCopyInto(out *PublicAccessBlockConfiguration) {
	*out = *in
	if in.BlockPublicACLs != nil {
		in, out := &in.BlockPublicACLs, &out.BlockPublicACLs
		*out = new(bool)
		**out = **in
	}
	if in.BlockPublicPolicy != nil {
		in, out := &in.BlockPublicPolicy, &out.BlockPublicPolicy
		*out = new(bool)
		**out = **in
	}
	if in.IgnorePublicACLs != nil {
		in, out := &in.IgnorePublicACLs, &out.IgnorePublicACLs
		*out = new(bool)
		**out = **in
	}
	if in.RestrictPublicBuckets != nil {
		in, out := &in.RestrictPublicBuckets, &out.RestrictPublicBuckets
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicAccessBlockConfiguration.
func (in *PublicAccessBlockConfiguration) DeepCopy() *PublicAccessBlockConfiguration {
	if in == nil {
		return nil
	}
	out := new(PublicAccessBlockConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Region) DeepCopyInto(out *Region) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.BucketAccountID != nil {
		in, out := &in.BucketAccountID, &out.BucketAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Region.
func (in *Region) DeepCopy() *Region {
	if in == nil {
		return nil
	}
	out := new(Region)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionReport) DeepCopyInto(out *RegionReport) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.BucketAccountID != nil {
		in, out := &in.BucketAccountID, &out.BucketAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionReport.
func (in *RegionReport) DeepCopy() *RegionReport {
	if in == nil {
		return nil
	}
	out := new(RegionReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionalBucket) DeepCopyInto(out *RegionalBucket) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionalBucket.
func (in *RegionalBucket) DeepCopy() *RegionalBucket {
	if in == nil {
		return nil
	}
	out := new(RegionalBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketDestination) DeepCopyInto(out *S3BucketDestination) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketDestination.
func (in *S3BucketDestination) DeepCopy() *S3BucketDestination {
	if in == nil {
		return nil
	}
	out := new(S3BucketDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3CopyObjectOperation) DeepCopyInto(out *S3CopyObjectOperation) {
	*out = *in
	if in.TargetKeyPrefix != nil {
		in, out := &in.TargetKeyPrefix, &out.TargetKeyPrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3CopyObjectOperation.
func (in *S3CopyObjectOperation) DeepCopy() *S3CopyObjectOperation {
	if in == nil {
		return nil
	}
	out := new(S3CopyObjectOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Grantee) DeepCopyInto(out *S3Grantee) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Identifier != nil {
		in, out := &in.Identifier, &out.Identifier
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Grantee.
func (in *S3Grantee) DeepCopy() *S3Grantee {
	if in == nil {
		return nil
	}
	out := new(S3Grantee)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3JobManifestGenerator) DeepCopyInto(out *S3JobManifestGenerator) {
	*out = *in
	if in.ExpectedBucketOwner != nil {
		in, out := &in.ExpectedBucketOwner, &out.ExpectedBucketOwner
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3JobManifestGenerator.
func (in *S3JobManifestGenerator) DeepCopy() *S3JobManifestGenerator {
	if in == nil {
		return nil
	}
	out := new(S3JobManifestGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ManifestOutputLocation) DeepCopyInto(out *S3ManifestOutputLocation) {
	*out = *in
	if in.ExpectedManifestBucketOwner != nil {
		in, out := &in.ExpectedManifestBucketOwner, &out.ExpectedManifestBucketOwner
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ManifestOutputLocation.
func (in *S3ManifestOutputLocation) DeepCopy() *S3ManifestOutputLocation {
	if in == nil {
		return nil
	}
	out := new(S3ManifestOutputLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ObjectMetadata) DeepCopyInto(out *S3ObjectMetadata) {
	*out = *in
	if in.CacheControl != nil {
		in, out := &in.CacheControl, &out.CacheControl
		*out = new(string)
		**out = **in
	}
	if in.ContentDisposition != nil {
		in, out := &in.ContentDisposition, &out.ContentDisposition
		*out = new(string)
		**out = **in
	}
	if in.ContentEncoding != nil {
		in, out := &in.ContentEncoding, &out.ContentEncoding
		*out = new(string)
		**out = **in
	}
	if in.ContentLanguage != nil {
		in, out := &in.ContentLanguage, &out.ContentLanguage
		*out = new(string)
		**out = **in
	}
	if in.ContentMD5 != nil {
		in, out := &in.ContentMD5, &out.ContentMD5
		*out = new(string)
		**out = **in
	}
	if in.ContentType != nil {
		in, out := &in.ContentType, &out.ContentType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ObjectMetadata.
func (in *S3ObjectMetadata) DeepCopy() *S3ObjectMetadata {
	if in == nil {
		return nil
	}
	out := new(S3ObjectMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3ObjectOwner) DeepCopyInto(out *S3ObjectOwner) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3ObjectOwner.
func (in *S3ObjectOwner) DeepCopy() *S3ObjectOwner {
	if in == nil {
		return nil
	}
	out := new(S3ObjectOwner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCConfiguration) DeepCopyInto(out *VPCConfiguration) {
	*out = *in
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCConfiguration.
func (in *VPCConfiguration) DeepCopy() *VPCConfiguration {
	if in == nil {
		return nil
	}
	out := new(VPCConfiguration)
	in.DeepCopyInto(out)
	return out
}
