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

type CertificateInitParameters struct {

	// Specifies the encoding used for the supplied certificate data. Must be one of pem, base64 or hex. Defaults to pem.
	// Specifies the encoding used for the supplied certificate data
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). Changing this field forces a new resource to be created.
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`)
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// A relative duration for which the certificate is valid until, for example 240h (10 days) or 2400h30m. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"
	EndDateRelative *string `json:"endDateRelative,omitempty" tf:"end_date_relative,omitempty"`

	// A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
	// A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`

	// The type of key/certificate. Must be one of AsymmetricX509Cert or Symmetric. Changing this fields forces a new resource to be created.
	// The type of key/certificate
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CertificateObservation struct {

	// Specifies the encoding used for the supplied certificate data. Must be one of pem, base64 or hex. Defaults to pem.
	// Specifies the encoding used for the supplied certificate data
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). Changing this field forces a new resource to be created.
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`)
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// A relative duration for which the certificate is valid until, for example 240h (10 days) or 2400h30m. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"
	EndDateRelative *string `json:"endDateRelative,omitempty" tf:"end_date_relative,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
	// A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
	// The object ID of the service principal for which this certificate should be created
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`

	// The type of key/certificate. Must be one of AsymmetricX509Cert or Symmetric. Changing this fields forces a new resource to be created.
	// The type of key/certificate
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CertificateParameters struct {

	// Specifies the encoding used for the supplied certificate data. Must be one of pem, base64 or hex. Defaults to pem.
	// Specifies the encoding used for the supplied certificate data
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). Changing this field forces a new resource to be created.
	// The end date until which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`)
	// +kubebuilder:validation:Optional
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// A relative duration for which the certificate is valid until, for example 240h (10 days) or 2400h30m. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". Changing this field forces a new resource to be created.
	// A relative duration for which the certificate is valid until, for example `240h` (10 days) or `2400h30m`. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"
	// +kubebuilder:validation:Optional
	EndDateRelative *string `json:"endDateRelative,omitempty" tf:"end_date_relative,omitempty"`

	// A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated. Changing this field forces a new resource to be created.
	// A UUID used to uniquely identify this certificate. If not specified a UUID will be automatically generated
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The object ID of the service principal for which this certificate should be created. Changing this field forces a new resource to be created.
	// The object ID of the service principal for which this certificate should be created
	// +crossplane:generate:reference:type=Principal
	// +kubebuilder:validation:Optional
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// Reference to a Principal to populate servicePrincipalId.
	// +kubebuilder:validation:Optional
	ServicePrincipalIDRef *v1.Reference `json:"servicePrincipalIdRef,omitempty" tf:"-"`

	// Selector for a Principal to populate servicePrincipalId.
	// +kubebuilder:validation:Optional
	ServicePrincipalIDSelector *v1.Selector `json:"servicePrincipalIdSelector,omitempty" tf:"-"`

	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. 2018-01-01T01:02:03Z). If this isn't specified, the value is determined by Azure Active Directory and is usually the start date of the certificate for asymmetric keys, or the current timestamp for symmetric keys. Changing this field forces a new resource to be created.
	// The start date from which the certificate is valid, formatted as an RFC3339 date string (e.g. `2018-01-01T01:02:03Z`). If this isn't specified, the current date is used
	// +kubebuilder:validation:Optional
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`

	// The type of key/certificate. Must be one of AsymmetricX509Cert or Symmetric. Changing this fields forces a new resource to be created.
	// The type of key/certificate
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER. See also the encoding argument.
	// The certificate data, which can be PEM encoded, base64 encoded DER or hexadecimal encoded DER
	// +kubebuilder:validation:Optional
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
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
	InitProvider CertificateInitParameters `json:"initProvider,omitempty"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Certificate is the Schema for the Certificates API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.valueSecretRef)",message="spec.forProvider.valueSecretRef is a required parameter"
	Spec   CertificateSpec   `json:"spec"`
	Status CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
