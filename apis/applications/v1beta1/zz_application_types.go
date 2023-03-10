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

type APIObservation struct {
}

type APIParameters struct {

	// A set of application IDs (client IDs), used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.
	// Used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app
	// +kubebuilder:validation:Optional
	KnownClientApplications []*string `json:"knownClientApplications,omitempty" tf:"known_client_applications,omitempty"`

	// Allows an application to use claims mapping without specifying a custom signing key. Defaults to false.
	// Allows an application to use claims mapping without specifying a custom signing key
	// +kubebuilder:validation:Optional
	MappedClaimsEnabled *bool `json:"mappedClaimsEnabled,omitempty" tf:"mapped_claims_enabled,omitempty"`

	// One or more oauth2_permission_scope blocks as documented below, to describe delegated permissions exposed by the web API represented by this application.
	// One or more `oauth2_permission_scope` blocks to describe delegated permissions exposed by the web API represented by this application
	// +kubebuilder:validation:Optional
	Oauth2PermissionScope []Oauth2PermissionScopeParameters `json:"oauth2PermissionScope,omitempty" tf:"oauth2_permission_scope,omitempty"`

	// The access token version expected by this resource. Must be one of 1 or 2, and must be 2 when sign_in_audience is either AzureADandPersonalMicrosoftAccount or PersonalMicrosoftAccount Defaults to 1.
	// The access token version expected by this resource
	// +kubebuilder:validation:Optional
	RequestedAccessTokenVersion *float64 `json:"requestedAccessTokenVersion,omitempty" tf:"requested_access_token_version,omitempty"`
}

type AccessTokenObservation struct {
}

type AccessTokenParameters struct {

	// List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
	// List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim
	// +kubebuilder:validation:Optional
	AdditionalProperties []*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
	// Whether the claim specified by the client is necessary to ensure a smooth authorization experience
	// +kubebuilder:validation:Optional
	Essential *bool `json:"essential,omitempty" tf:"essential,omitempty"`

	// The name of the optional claim.
	// The name of the optional claim
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The source of the claim. If source is absent, the claim is a predefined optional claim. If source is user, the value of name is the extension property from the user object.
	// The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type AppRoleObservation struct {
}

type AppRoleParameters struct {

	// Specifies whether this app role definition can be assigned to users and groups by setting to User, or to other applications (that are accessing this application in a standalone scenario) by setting to Application, or to both.
	// Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both
	// +kubebuilder:validation:Required
	AllowedMemberTypes []*string `json:"allowedMemberTypes" tf:"allowed_member_types,omitempty"`

	// Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
	// Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences
	// +kubebuilder:validation:Required
	Description *string `json:"description" tf:"description,omitempty"`

	// Display name for the app role that appears during app role assignment and in consent experiences.
	// Display name for the app role that appears during app role assignment and in consent experiences
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// Determines if the app role is enabled. Defaults to true.
	// Determines if the app role is enabled
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The unique identifier of the app role. Must be a valid UUID.
	// The unique identifier of the app role
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// The value that is used for the roles claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
	// The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ApplicationObservation struct {

	// A mapping of app role values to app role IDs, intended to be useful when referencing app roles in other resources in your configuration.
	// Mapping of app role names to UUIDs
	AppRoleIds map[string]*string `json:"appRoleIds,omitempty" tf:"app_role_ids,omitempty"`

	// The Application ID (also called Client ID).
	// The Application ID (also called Client ID)
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Whether Microsoft has disabled the registered application. If the application is disabled, this will be a string indicating the status/reason, e.g. DisabledDueToViolationOfServicesAgreement
	// Whether Microsoft has disabled the registered application
	DisabledByMicrosoft *string `json:"disabledByMicrosoft,omitempty" tf:"disabled_by_microsoft,omitempty"`

	// The unique identifier for an app role or OAuth2 permission scope published by the resource application.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// CDN URL to the application's logo, as uploaded with the logo_image property.
	// CDN URL to the application's logo
	LogoURL *string `json:"logoUrl,omitempty" tf:"logo_url,omitempty"`

	// A mapping of OAuth2.0 permission scope values to scope IDs, intended to be useful when referencing permission scopes in other resources in your configuration.
	// Mapping of OAuth2.0 permission scope names to UUIDs
	Oauth2PermissionScopeIds map[string]*string `json:"oauth2PermissionScopeIds,omitempty" tf:"oauth2_permission_scope_ids,omitempty"`

	// The application's object ID.
	// The application's object ID
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The verified publisher domain for the application.
	// The verified publisher domain for the application
	PublisherDomain *string `json:"publisherDomain,omitempty" tf:"publisher_domain,omitempty"`
}

type ApplicationParameters struct {

	// An api block as documented below, which configures API related settings for this application.
	// +kubebuilder:validation:Optional
	API []APIParameters `json:"api,omitempty" tf:"api,omitempty"`

	// A collection of app_role blocks as documented below. For more information see official documentation on Application Roles.
	// +kubebuilder:validation:Optional
	AppRole []AppRoleParameters `json:"appRole,omitempty" tf:"app_role,omitempty"`

	// A description of the application, as shown to end users.
	// Description of the application as shown to end users
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether this application supports device authentication without a user. Defaults to false.
	// Specifies whether this application supports device authentication without a user.
	// +kubebuilder:validation:Optional
	DeviceOnlyAuthEnabled *bool `json:"deviceOnlyAuthEnabled,omitempty" tf:"device_only_auth_enabled,omitempty"`

	// The display name for the application.
	// The display name for the application
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI. Defaults to false.
	// Specifies whether the application is a public client. Appropriate for apps using token grant flows that don't use a redirect URI
	// +kubebuilder:validation:Optional
	FallbackPublicClientEnabled *bool `json:"fallbackPublicClientEnabled,omitempty" tf:"fallback_public_client_enabled,omitempty"`

	// A feature_tags block as described below. Cannot be used together with the tags property.
	// Block of features to configure for this application using tags
	// +kubebuilder:validation:Optional
	FeatureTags []FeatureTagsParameters `json:"featureTags,omitempty" tf:"feature_tags,omitempty"`

	// Configures the groups claim issued in a user or OAuth 2.0 access token that the app expects. Possible values are None, SecurityGroup, DirectoryRole, ApplicationGroup or All.
	// Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects
	// +kubebuilder:validation:Optional
	GroupMembershipClaims []*string `json:"groupMembershipClaims,omitempty" tf:"group_membership_claims,omitempty"`

	// A set of user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
	// The user-defined URI(s) that uniquely identify an application within its Azure AD tenant, or within a verified custom domain if the application is multi-tenant
	// +kubebuilder:validation:Optional
	IdentifierUris []*string `json:"identifierUris,omitempty" tf:"identifier_uris,omitempty"`

	// A logo image to upload for the application, as a raw base64-encoded string. The image should be in gif, jpeg or png format. Note that once an image has been uploaded, it is not possible to remove it without replacing it with another image.
	// Base64 encoded logo image in gif, png or jpeg format
	// +kubebuilder:validation:Optional
	LogoImage *string `json:"logoImage,omitempty" tf:"logo_image,omitempty"`

	// URL of the application's marketing page.
	// URL of the application's marketing page
	// +kubebuilder:validation:Optional
	MarketingURL *string `json:"marketingUrl,omitempty" tf:"marketing_url,omitempty"`

	// User-specified notes relevant for the management of the application.
	// User-specified notes relevant for the management of the application
	// +kubebuilder:validation:Optional
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests. Defaults to false, which specifies that only GET requests are allowed.
	// Specifies whether, as part of OAuth 2.0 token requests, Azure AD allows POST requests, as opposed to GET requests.
	// +kubebuilder:validation:Optional
	Oauth2PostResponseRequired *bool `json:"oauth2PostResponseRequired,omitempty" tf:"oauth2_post_response_required,omitempty"`

	// An optional_claims block as documented below.
	// +kubebuilder:validation:Optional
	OptionalClaims []OptionalClaimsParameters `json:"optionalClaims,omitempty" tf:"optional_claims,omitempty"`

	// A set of object IDs of principals that will be granted ownership of the application. Supported object types are users or service principals. By default, no owners are assigned.
	// A list of object IDs of principals that will be granted ownership of the application
	// +kubebuilder:validation:Optional
	Owners []*string `json:"owners,omitempty" tf:"owners,omitempty"`

	// If true, will return an error if an existing application is found with the same name. Defaults to false.
	// If `true`, will return an error if an existing application is found with the same name
	// +kubebuilder:validation:Optional
	PreventDuplicateNames *bool `json:"preventDuplicateNames,omitempty" tf:"prevent_duplicate_names,omitempty"`

	// URL of the application's privacy statement.
	// URL of the application's privacy statement
	// +kubebuilder:validation:Optional
	PrivacyStatementURL *string `json:"privacyStatementUrl,omitempty" tf:"privacy_statement_url,omitempty"`

	// A public_client block as documented below, which configures non-web app or non-web API application settings, for example mobile or other public clients such as an installed application running on a desktop device.
	// +kubebuilder:validation:Optional
	PublicClient []PublicClientParameters `json:"publicClient,omitempty" tf:"public_client,omitempty"`

	// A collection of required_resource_access blocks as documented below.
	// +kubebuilder:validation:Optional
	RequiredResourceAccess []RequiredResourceAccessParameters `json:"requiredResourceAccess,omitempty" tf:"required_resource_access,omitempty"`

	// The Microsoft account types that are supported for the current application. Must be one of AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount or PersonalMicrosoftAccount. Defaults to AzureADMyOrg.
	// The Microsoft account types that are supported for the current application
	// +kubebuilder:validation:Optional
	SignInAudience *string `json:"signInAudience,omitempty" tf:"sign_in_audience,omitempty"`

	// A single_page_application block as documented below, which configures single-page application (SPA) related settings for this application.
	// +kubebuilder:validation:Optional
	SinglePageApplication []SinglePageApplicationParameters `json:"singlePageApplication,omitempty" tf:"single_page_application,omitempty"`

	// URL of the application's support page.
	// URL of the application's support page
	// +kubebuilder:validation:Optional
	SupportURL *string `json:"supportUrl,omitempty" tf:"support_url,omitempty"`

	// A set of tags to apply to the application for configuring specific behaviours of the application and linked service principals. Note that these are not provided for use by practitioners. Cannot be used together with the feature_tags block.
	// A set of tags to apply to the application
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Unique ID for a templated application in the Azure AD App Gallery, from which to create the application. Changing this forces a new resource to be created.
	// Unique ID of the application template from which this application is created
	// +kubebuilder:validation:Optional
	TemplateID *string `json:"templateId,omitempty" tf:"template_id,omitempty"`

	// URL of the application's terms of service statement.
	// URL of the application's terms of service statement
	// +kubebuilder:validation:Optional
	TermsOfServiceURL *string `json:"termsOfServiceUrl,omitempty" tf:"terms_of_service_url,omitempty"`

	// A web block as documented below, which configures web related settings for this application.
	// +kubebuilder:validation:Optional
	Web []WebParameters `json:"web,omitempty" tf:"web,omitempty"`
}

type FeatureTagsObservation struct {
}

type FeatureTagsParameters struct {

	// Whether this application represents a custom SAML application for linked service principals. Enabling this will assign the WindowsAzureActiveDirectoryCustomSingleSignOnApplication tag. Defaults to false.
	// Whether this application represents a custom SAML application for linked service principals
	// +kubebuilder:validation:Optional
	CustomSingleSignOn *bool `json:"customSingleSignOn,omitempty" tf:"custom_single_sign_on,omitempty"`

	// Whether this application represents an Enterprise Application for linked service principals. Enabling this will assign the WindowsAzureActiveDirectoryIntegratedApp tag. Defaults to false.
	// Whether this application represents an Enterprise Application for linked service principals
	// +kubebuilder:validation:Optional
	Enterprise *bool `json:"enterprise,omitempty" tf:"enterprise,omitempty"`

	// Whether this application represents a gallery application for linked service principals. Enabling this will assign the WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1 tag. Defaults to false.
	// Whether this application represents a gallery application for linked service principals
	// +kubebuilder:validation:Optional
	Gallery *bool `json:"gallery,omitempty" tf:"gallery,omitempty"`

	// Whether this app is invisible to users in My Apps and Office 365 Launcher. Enabling this will assign the HideApp tag. Defaults to false.
	// Whether this application is invisible to users in My Apps and Office 365 Launcher
	// +kubebuilder:validation:Optional
	Hide *bool `json:"hide,omitempty" tf:"hide,omitempty"`
}

type IDTokenObservation struct {
}

type IDTokenParameters struct {

	// List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
	// List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim
	// +kubebuilder:validation:Optional
	AdditionalProperties []*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
	// Whether the claim specified by the client is necessary to ensure a smooth authorization experience
	// +kubebuilder:validation:Optional
	Essential *bool `json:"essential,omitempty" tf:"essential,omitempty"`

	// The name of the optional claim.
	// The name of the optional claim
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The source of the claim. If source is absent, the claim is a predefined optional claim. If source is user, the value of name is the extension property from the user object.
	// The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type ImplicitGrantObservation struct {
}

type ImplicitGrantParameters struct {

	// Whether this web application can request an access token using OAuth 2.0 implicit flow.
	// Whether this web application can request an access token using OAuth 2.0 implicit flow
	// +kubebuilder:validation:Optional
	AccessTokenIssuanceEnabled *bool `json:"accessTokenIssuanceEnabled,omitempty" tf:"access_token_issuance_enabled,omitempty"`

	// Whether this web application can request an ID token using OAuth 2.0 implicit flow.
	// Whether this web application can request an ID token using OAuth 2.0 implicit flow
	// +kubebuilder:validation:Optional
	IDTokenIssuanceEnabled *bool `json:"idTokenIssuanceEnabled,omitempty" tf:"id_token_issuance_enabled,omitempty"`
}

type Oauth2PermissionScopeObservation struct {
}

type Oauth2PermissionScopeParameters struct {

	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
	// Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users
	// +kubebuilder:validation:Optional
	AdminConsentDescription *string `json:"adminConsentDescription,omitempty" tf:"admin_consent_description,omitempty"`

	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
	// Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users
	// +kubebuilder:validation:Optional
	AdminConsentDisplayName *string `json:"adminConsentDisplayName,omitempty" tf:"admin_consent_display_name,omitempty"`

	// Determines if the permission scope is enabled. Defaults to true.
	// Determines if the permission scope is enabled
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The unique identifier of the delegated permission. Must be a valid UUID.
	// The unique identifier of the delegated permission
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to User. Possible values are User or Admin.
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
	// Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf
	// +kubebuilder:validation:Optional
	UserConsentDescription *string `json:"userConsentDescription,omitempty" tf:"user_consent_description,omitempty"`

	// Display name for the delegated permission that appears in the end user consent experience.
	// Display name for the delegated permission that appears in the end user consent experience
	// +kubebuilder:validation:Optional
	UserConsentDisplayName *string `json:"userConsentDisplayName,omitempty" tf:"user_consent_display_name,omitempty"`

	// The value that is used for the scp claim in OAuth 2.0 access tokens.
	// The value that is used for the `scp` claim in OAuth 2.0 access tokens
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type OptionalClaimsObservation struct {
}

type OptionalClaimsParameters struct {

	// One or more access_token blocks as documented below.
	// +kubebuilder:validation:Optional
	AccessToken []AccessTokenParameters `json:"accessToken,omitempty" tf:"access_token,omitempty"`

	// One or more id_token blocks as documented below.
	// +kubebuilder:validation:Optional
	IDToken []IDTokenParameters `json:"idToken,omitempty" tf:"id_token,omitempty"`

	// One or more saml2_token blocks as documented below.
	// +kubebuilder:validation:Optional
	Saml2Token []Saml2TokenParameters `json:"saml2Token,omitempty" tf:"saml2_token,omitempty"`
}

type PublicClientObservation struct {
}

type PublicClientParameters struct {

	// A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid https or ms-appx-web URL.
	// The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent
	// +kubebuilder:validation:Optional
	RedirectUris []*string `json:"redirectUris,omitempty" tf:"redirect_uris,omitempty"`
}

type RequiredResourceAccessObservation struct {
}

type RequiredResourceAccessParameters struct {

	// A collection of resource_access blocks as documented below, describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
	// +kubebuilder:validation:Required
	ResourceAccess []ResourceAccessParameters `json:"resourceAccess" tf:"resource_access,omitempty"`

	// The unique identifier for the resource that the application requires access to. This should be the Application ID of the target application.
	// +kubebuilder:validation:Required
	ResourceAppID *string `json:"resourceAppId" tf:"resource_app_id,omitempty"`
}

type ResourceAccessObservation struct {
}

type ResourceAccessParameters struct {

	// The unique identifier for an app role or OAuth2 permission scope published by the resource application.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// Specifies whether the id property references an app role or an OAuth2 permission scope. Possible values are Role or Scope.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type Saml2TokenObservation struct {
}

type Saml2TokenParameters struct {

	// List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
	// List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim
	// +kubebuilder:validation:Optional
	AdditionalProperties []*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
	// Whether the claim specified by the client is necessary to ensure a smooth authorization experience
	// +kubebuilder:validation:Optional
	Essential *bool `json:"essential,omitempty" tf:"essential,omitempty"`

	// The name of the optional claim.
	// The name of the optional claim
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The source of the claim. If source is absent, the claim is a predefined optional claim. If source is user, the value of name is the extension property from the user object.
	// The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type SinglePageApplicationObservation struct {
}

type SinglePageApplicationParameters struct {

	// A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid https URL.
	// The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent
	// +kubebuilder:validation:Optional
	RedirectUris []*string `json:"redirectUris,omitempty" tf:"redirect_uris,omitempty"`
}

type WebObservation struct {
}

type WebParameters struct {

	// Home page or landing page of the application.
	// Home page or landing page of the application
	// +kubebuilder:validation:Optional
	HomepageURL *string `json:"homepageUrl,omitempty" tf:"homepage_url,omitempty"`

	// An implicit_grant block as documented above.
	// +kubebuilder:validation:Optional
	ImplicitGrant []ImplicitGrantParameters `json:"implicitGrant,omitempty" tf:"implicit_grant,omitempty"`

	// The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
	// The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols
	// +kubebuilder:validation:Optional
	LogoutURL *string `json:"logoutUrl,omitempty" tf:"logout_url,omitempty"`

	// A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid http URL or a URN.
	// The URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent
	// +kubebuilder:validation:Optional
	RedirectUris []*string `json:"redirectUris,omitempty" tf:"redirect_uris,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Application is the Schema for the Applications API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSpec   `json:"spec"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
