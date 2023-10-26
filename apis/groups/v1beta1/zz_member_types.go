// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MemberInitParameters struct {
}

type MemberObservation struct {

	// The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
	// The object ID of the group you want to add the member to
	GroupObjectID *string `json:"groupObjectId,omitempty" tf:"group_object_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals
	MemberObjectID *string `json:"memberObjectId,omitempty" tf:"member_object_id,omitempty"`
}

type MemberParameters struct {

	// The object ID of the group you want to add the member to. Changing this forces a new resource to be created.
	// The object ID of the group you want to add the member to
	// +crossplane:generate:reference:type=Group
	// +kubebuilder:validation:Optional
	GroupObjectID *string `json:"groupObjectId,omitempty" tf:"group_object_id,omitempty"`

	// Reference to a Group to populate groupObjectId.
	// +kubebuilder:validation:Optional
	GroupObjectIDRef *v1.Reference `json:"groupObjectIdRef,omitempty" tf:"-"`

	// Selector for a Group to populate groupObjectId.
	// +kubebuilder:validation:Optional
	GroupObjectIDSelector *v1.Selector `json:"groupObjectIdSelector,omitempty" tf:"-"`

	// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals. Changing this forces a new resource to be created.
	// The object ID of the principal you want to add as a member to the group. Supported object types are Users, Groups or Service Principals
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/users/v1beta1.User
	// +kubebuilder:validation:Optional
	MemberObjectID *string `json:"memberObjectId,omitempty" tf:"member_object_id,omitempty"`

	// Reference to a User in users to populate memberObjectId.
	// +kubebuilder:validation:Optional
	MemberObjectIDRef *v1.Reference `json:"memberObjectIdRef,omitempty" tf:"-"`

	// Selector for a User in users to populate memberObjectId.
	// +kubebuilder:validation:Optional
	MemberObjectIDSelector *v1.Selector `json:"memberObjectIdSelector,omitempty" tf:"-"`
}

// MemberSpec defines the desired state of Member
type MemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MemberParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MemberInitParameters `json:"initProvider,omitempty"`
}

// MemberStatus defines the observed state of Member.
type MemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Member is the Schema for the Members API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type Member struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MemberSpec   `json:"spec"`
	Status            MemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemberList contains a list of Members
type MemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Member `json:"items"`
}

// Repository type metadata.
var (
	Member_Kind             = "Member"
	Member_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Member_Kind}.String()
	Member_KindAPIVersion   = Member_Kind + "." + CRDGroupVersion.String()
	Member_GroupVersionKind = CRDGroupVersion.WithKind(Member_Kind)
)

func init() {
	SchemeBuilder.Register(&Member{}, &MemberList{})
}
