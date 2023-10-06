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

// JobParameters defines the desired state of Job
type JobParameters struct {
	// Region is which region the Job will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// This parameter is deprecated. Use MaxCapacity instead.
	//
	// The number of Glue data processing units (DPUs) to allocate to this Job.
	// You can allocate a minimum of 2 DPUs; the default is 10. A DPU is a relative
	// measure of processing power that consists of 4 vCPUs of compute capacity
	// and 16 GB of memory. For more information, see the Glue pricing page (https://aws.amazon.com/glue/pricing/).
	AllocatedCapacity *int64 `json:"allocatedCapacity,omitempty"`
	// The representation of a directed acyclic graph on which both the Glue Studio
	// visual component and Glue Studio code generation is based.
	CodeGenConfigurationNodes map[string]*CodeGenConfigurationNode `json:"codeGenConfigurationNodes,omitempty"`
	// The JobCommand that runs this job.
	// +kubebuilder:validation:Required
	Command *JobCommand `json:"command"`
	// The default arguments for every run of this job, specified as name-value
	// pairs.
	//
	// You can specify arguments here that your own job-execution script consumes,
	// as well as arguments that Glue itself consumes.
	//
	// Job arguments may be logged. Do not pass plaintext secrets as arguments.
	// Retrieve secrets from a Glue Connection, Secrets Manager or other secret
	// management mechanism if you intend to keep them within the Job.
	//
	// For information about how to specify and consume your own Job arguments,
	// see the Calling Glue APIs in Python (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html)
	// topic in the developer guide.
	//
	// For information about the arguments you can provide to this field when configuring
	// Spark jobs, see the Special Parameters Used by Glue (https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html)
	// topic in the developer guide.
	//
	// For information about the arguments you can provide to this field when configuring
	// Ray jobs, see Using job parameters in Ray jobs (https://docs.aws.amazon.com/glue/latest/dg/author-job-ray-job-parameters.html)
	// in the developer guide.
	DefaultArguments map[string]*string `json:"defaultArguments,omitempty"`
	// Description of the job being defined.
	Description *string `json:"description,omitempty"`
	// Indicates whether the job is run with a standard or flexible execution class.
	// The standard execution-class is ideal for time-sensitive workloads that require
	// fast job startup and dedicated resources.
	//
	// The flexible execution class is appropriate for time-insensitive jobs whose
	// start and completion times may vary.
	//
	// Only jobs with Glue version 3.0 and above and command type glueetl will be
	// allowed to set ExecutionClass to FLEX. The flexible execution class is available
	// for Spark jobs.
	ExecutionClass *string `json:"executionClass,omitempty"`
	// An ExecutionProperty specifying the maximum number of concurrent runs allowed
	// for this job.
	ExecutionProperty *ExecutionProperty `json:"executionProperty,omitempty"`
	// In Spark jobs, GlueVersion determines the versions of Apache Spark and Python
	// that Glue available in a job. The Python version indicates the version supported
	// for jobs of type Spark.
	//
	// Ray jobs should set GlueVersion to 4.0 or greater. However, the versions
	// of Ray, Python and additional libraries available in your Ray job are determined
	// by the Runtime parameter of the Job command.
	//
	// For more information about the available Glue versions and corresponding
	// Spark and Python versions, see Glue version (https://docs.aws.amazon.com/glue/latest/dg/add-job.html)
	// in the developer guide.
	//
	// Jobs that are created without specifying a Glue version default to Glue 0.9.
	GlueVersion *string `json:"glueVersion,omitempty"`
	// This field is reserved for future use.
	LogURI *string `json:"logURI,omitempty"`
	// For Glue version 1.0 or earlier jobs, using the standard worker type, the
	// number of Glue data processing units (DPUs) that can be allocated when this
	// job runs. A DPU is a relative measure of processing power that consists of
	// 4 vCPUs of compute capacity and 16 GB of memory. For more information, see
	// the Glue pricing page (https://aws.amazon.com/glue/pricing/).
	//
	// For Glue version 2.0+ jobs, you cannot specify a Maximum capacity. Instead,
	// you should specify a Worker type and the Number of workers.
	//
	// Do not set MaxCapacity if using WorkerType and NumberOfWorkers.
	//
	// The value that can be allocated for MaxCapacity depends on whether you are
	// running a Python shell job, an Apache Spark ETL job, or an Apache Spark streaming
	// ETL job:
	//
	//    * When you specify a Python shell job (JobCommand.Name="pythonshell"),
	//    you can allocate either 0.0625 or 1 DPU. The default is 0.0625 DPU.
	//
	//    * When you specify an Apache Spark ETL job (JobCommand.Name="glueetl")
	//    or Apache Spark streaming ETL job (JobCommand.Name="gluestreaming"), you
	//    can allocate from 2 to 100 DPUs. The default is 10 DPUs. This job type
	//    cannot have a fractional DPU allocation.
	MaxCapacity *float64 `json:"maxCapacity,omitempty"`
	// The maximum number of times to retry this job if it fails.
	MaxRetries *int64 `json:"maxRetries,omitempty"`
	// Arguments for this job that are not overridden when providing job arguments
	// in a job run, specified as name-value pairs.
	NonOverridableArguments map[string]*string `json:"nonOverridableArguments,omitempty"`
	// Specifies configuration properties of a job notification.
	NotificationProperty *NotificationProperty `json:"notificationProperty,omitempty"`
	// The number of workers of a defined workerType that are allocated when a job
	// runs.
	NumberOfWorkers *int64 `json:"numberOfWorkers,omitempty"`
	// The details for a source control configuration for a job, allowing synchronization
	// of job artifacts to or from a remote repository.
	SourceControlDetails *SourceControlDetails `json:"sourceControlDetails,omitempty"`
	// The tags to use with this job. You may use tags to limit access to the job.
	// For more information about tags in Glue, see Amazon Web Services Tags in
	// Glue (https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html) in the
	// developer guide.
	Tags map[string]*string `json:"tags,omitempty"`
	// The job timeout in minutes. This is the maximum time that a job run can consume
	// resources before it is terminated and enters TIMEOUT status. The default
	// is 2,880 minutes (48 hours).
	Timeout *int64 `json:"timeout,omitempty"`
	// The type of predefined worker that is allocated when a job runs. Accepts
	// a value of G.1X, G.2X, G.4X, G.8X or G.025X for Spark jobs. Accepts the value
	// Z.2X for Ray jobs.
	//
	//    * For the G.1X worker type, each worker maps to 1 DPU (4 vCPUs, 16 GB
	//    of memory) with 84GB disk (approximately 34GB free), and provides 1 executor
	//    per worker. We recommend this worker type for workloads such as data transforms,
	//    joins, and queries, to offers a scalable and cost effective way to run
	//    most jobs.
	//
	//    * For the G.2X worker type, each worker maps to 2 DPU (8 vCPUs, 32 GB
	//    of memory) with 128GB disk (approximately 77GB free), and provides 1 executor
	//    per worker. We recommend this worker type for workloads such as data transforms,
	//    joins, and queries, to offers a scalable and cost effective way to run
	//    most jobs.
	//
	//    * For the G.4X worker type, each worker maps to 4 DPU (16 vCPUs, 64 GB
	//    of memory) with 256GB disk (approximately 235GB free), and provides 1
	//    executor per worker. We recommend this worker type for jobs whose workloads
	//    contain your most demanding transforms, aggregations, joins, and queries.
	//    This worker type is available only for Glue version 3.0 or later Spark
	//    ETL jobs in the following Amazon Web Services Regions: US East (Ohio),
	//    US East (N. Virginia), US West (Oregon), Asia Pacific (Singapore), Asia
	//    Pacific (Sydney), Asia Pacific (Tokyo), Canada (Central), Europe (Frankfurt),
	//    Europe (Ireland), and Europe (Stockholm).
	//
	//    * For the G.8X worker type, each worker maps to 8 DPU (32 vCPUs, 128 GB
	//    of memory) with 512GB disk (approximately 487GB free), and provides 1
	//    executor per worker. We recommend this worker type for jobs whose workloads
	//    contain your most demanding transforms, aggregations, joins, and queries.
	//    This worker type is available only for Glue version 3.0 or later Spark
	//    ETL jobs, in the same Amazon Web Services Regions as supported for the
	//    G.4X worker type.
	//
	//    * For the G.025X worker type, each worker maps to 0.25 DPU (2 vCPUs, 4
	//    GB of memory) with 84GB disk (approximately 34GB free), and provides 1
	//    executor per worker. We recommend this worker type for low volume streaming
	//    jobs. This worker type is only available for Glue version 3.0 streaming
	//    jobs.
	//
	//    * For the Z.2X worker type, each worker maps to 2 M-DPU (8vCPUs, 64 GB
	//    of memory) with 128 GB disk (approximately 120GB free), and provides up
	//    to 8 Ray workers based on the autoscaler.
	WorkerType          *string `json:"workerType,omitempty"`
	CustomJobParameters `json:",inline"`
}

// JobSpec defines the desired state of Job
type JobSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       JobParameters `json:"forProvider"`
}

// JobObservation defines the observed state of Job
type JobObservation struct {
	// The time and date that this job definition was created.
	CreatedOn *metav1.Time `json:"createdOn,omitempty"`
	// The last point in time when this job definition was modified.
	LastModifiedOn *metav1.Time `json:"lastModifiedOn,omitempty"`
	// The unique name that was provided for this job definition.
	Name *string `json:"name,omitempty"`
}

// JobStatus defines the observed state of Job.
type JobStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          JobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Job is the Schema for the Jobs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Job struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobSpec   `json:"spec"`
	Status            JobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobList contains a list of Jobs
type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Job `json:"items"`
}

// Repository type metadata.
var (
	JobKind             = "Job"
	JobGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: JobKind}.String()
	JobKindAPIVersion   = JobKind + "." + GroupVersion.String()
	JobGroupVersionKind = GroupVersion.WithKind(JobKind)
)

func init() {
	SchemeBuilder.Register(&Job{}, &JobList{})
}
