// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RuleSpec defines the desired state of Rule.
type RuleSpec struct {

	// The retention rule description.
	Description *string `json:"description,omitempty"`
	// Information about the retention rule lock configuration.
	LockConfiguration *LockConfiguration `json:"lockConfiguration,omitempty"`
	// Specifies the resource tags to use to identify resources that are to be retained
	// by a tag-level retention rule. For tag-level retention rules, only deleted
	// resources, of the specified resource type, that have one or more of the specified
	// tag key and value pairs are retained. If a resource is deleted, but it does
	// not have any of the specified tag key and value pairs, it is immediately
	// deleted without being retained by the retention rule.
	//
	// You can add the same tag key and value pair to a maximum or five retention
	// rules.
	//
	// To create a Region-level retention rule, omit this parameter. A Region-level
	// retention rule does not have any resource tags specified. It retains all
	// deleted resources of the specified resource type in the Region in which the
	// rule is created, even if the resources are not tagged.
	ResourceTags []*ResourceTag `json:"resourceTags,omitempty"`
	// The resource type to be retained by the retention rule. Currently, only Amazon
	// EBS snapshots and EBS-backed AMIs are supported. To retain snapshots, specify
	// EBS_SNAPSHOT. To retain EBS-backed AMIs, specify EC2_IMAGE.
	// +kubebuilder:validation:Required
	ResourceType *string `json:"resourceType"`
	// Information about the retention period for which the retention rule is to
	// retain resources.
	// +kubebuilder:validation:Required
	RetentionPeriod *RetentionPeriod `json:"retentionPeriod"`
	// Information about the tags to assign to the retention rule.
	Tags []*Tag `json:"tags,omitempty"`
}

// RuleStatus defines the observed state of Rule
type RuleStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The unique ID of the retention rule.
	// +kubebuilder:validation:Optional
	Identifier *string `json:"identifier,omitempty"`
	// The lock state for the retention rule.
	//
	//    * locked - The retention rule is locked and can't be modified or deleted.
	//
	//    * pending_unlock - The retention rule has been unlocked but it is still
	//    within the unlock delay period. The retention rule can be modified or
	//    deleted only after the unlock delay period has expired.
	//
	//    * unlocked - The retention rule is unlocked and it can be modified or
	//    deleted by any user with the required permissions.
	//
	//    * null - The retention rule has never been locked. Once a retention rule
	//    has been locked, it can transition between the locked and unlocked states
	//    only; it can never transition back to null.
	// +kubebuilder:validation:Optional
	LockState *string `json:"lockState,omitempty"`
	// The state of the retention rule. Only retention rules that are in the available
	// state retain resources.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty"`
}

// Rule is the Schema for the Rules API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleSpec   `json:"spec,omitempty"`
	Status            RuleStatus `json:"status,omitempty"`
}

// RuleList contains a list of Rule
// +kubebuilder:object:root=true
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}
