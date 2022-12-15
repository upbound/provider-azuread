/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azuread/apis/users/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this RoleAssignment.
func (mg *RoleAssignment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrincipalObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PrincipalObjectIDRef,
		Selector:     mg.Spec.ForProvider.PrincipalObjectIDSelector,
		To: reference.To{
			List:    &v1beta1.UserList{},
			Managed: &v1beta1.User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrincipalObjectID")
	}
	mg.Spec.ForProvider.PrincipalObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrincipalObjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleID),
		Extract:      resource.ExtractParamPath("template_id", true),
		Reference:    mg.Spec.ForProvider.RoleIDRef,
		Selector:     mg.Spec.ForProvider.RoleIDSelector,
		To: reference.To{
			List:    &RoleList{},
			Managed: &Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleID")
	}
	mg.Spec.ForProvider.RoleID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleIDRef = rsp.ResolvedReference

	return nil
}
