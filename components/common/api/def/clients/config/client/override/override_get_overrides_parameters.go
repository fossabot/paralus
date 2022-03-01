// Code generated by go-swagger; DO NOT EDIT.

package override

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewOverrideGetOverridesParams creates a new OverrideGetOverridesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOverrideGetOverridesParams() *OverrideGetOverridesParams {
	return &OverrideGetOverridesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOverrideGetOverridesParamsWithTimeout creates a new OverrideGetOverridesParams object
// with the ability to set a timeout on a request.
func NewOverrideGetOverridesParamsWithTimeout(timeout time.Duration) *OverrideGetOverridesParams {
	return &OverrideGetOverridesParams{
		timeout: timeout,
	}
}

// NewOverrideGetOverridesParamsWithContext creates a new OverrideGetOverridesParams object
// with the ability to set a context for a request.
func NewOverrideGetOverridesParamsWithContext(ctx context.Context) *OverrideGetOverridesParams {
	return &OverrideGetOverridesParams{
		Context: ctx,
	}
}

// NewOverrideGetOverridesParamsWithHTTPClient creates a new OverrideGetOverridesParams object
// with the ability to set a custom HTTPClient for a request.
func NewOverrideGetOverridesParamsWithHTTPClient(client *http.Client) *OverrideGetOverridesParams {
	return &OverrideGetOverridesParams{
		HTTPClient: client,
	}
}

/* OverrideGetOverridesParams contains all the parameters to send to the API endpoint
   for the override get overrides operation.

   Typically these are written to a http.Request.
*/
type OverrideGetOverridesParams struct {

	// ID.
	ID *string

	// BlueprintRef.
	BlueprintRef *string

	// ClusterID.
	ClusterID *string

	// Count.
	//
	// Format: int64
	Count *string

	// Deleted.
	Deleted *bool

	/* DisplayName.

	   displayName only used for update queries to set displayName (READONLY).
	*/
	DisplayName *string

	// Extended.
	Extended *bool

	/* GlobalScope.

	   globalScope sets partnerID,organizationID,projectID = 0.
	*/
	GlobalScope *bool

	// Groups.
	Groups []string

	/* IgnoreScopeDefault.

	     ignoreScopeDefault ignores default values for partnerID, organizationID and
	projectID.
	*/
	IgnoreScopeDefault *bool

	// IsSSOUser.
	IsSSOUser *bool

	// Limit.
	//
	// Format: int64
	Limit *string

	/* Name.

	     name is unique ID of a resource along with (partnerID, organizationID,
	projectID).
	*/
	Name *string

	// Offset.
	//
	// Format: int64
	Offset *string

	// Order.
	Order *string

	// OrderBy.
	OrderBy *string

	// OrganizationID.
	OrganizationID *string

	// PartnerID.
	PartnerID *string

	// ProjectID.
	ProjectID *string

	// PublishedVersion.
	PublishedVersion *string

	/* Selector.

	   selector is used to filter the labels of a resource.
	*/
	Selector *string

	/* URLScope.

	   urlScope is supposed to be passed in the URL as kind/HashID(value)
	*/
	URLScope string

	// Username.
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the override get overrides params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OverrideGetOverridesParams) WithDefaults() *OverrideGetOverridesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the override get overrides params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OverrideGetOverridesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the override get overrides params
func (o *OverrideGetOverridesParams) WithTimeout(timeout time.Duration) *OverrideGetOverridesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the override get overrides params
func (o *OverrideGetOverridesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the override get overrides params
func (o *OverrideGetOverridesParams) WithContext(ctx context.Context) *OverrideGetOverridesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the override get overrides params
func (o *OverrideGetOverridesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the override get overrides params
func (o *OverrideGetOverridesParams) WithHTTPClient(client *http.Client) *OverrideGetOverridesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the override get overrides params
func (o *OverrideGetOverridesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the override get overrides params
func (o *OverrideGetOverridesParams) WithID(id *string) *OverrideGetOverridesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the override get overrides params
func (o *OverrideGetOverridesParams) SetID(id *string) {
	o.ID = id
}

// WithBlueprintRef adds the blueprintRef to the override get overrides params
func (o *OverrideGetOverridesParams) WithBlueprintRef(blueprintRef *string) *OverrideGetOverridesParams {
	o.SetBlueprintRef(blueprintRef)
	return o
}

// SetBlueprintRef adds the blueprintRef to the override get overrides params
func (o *OverrideGetOverridesParams) SetBlueprintRef(blueprintRef *string) {
	o.BlueprintRef = blueprintRef
}

// WithClusterID adds the clusterID to the override get overrides params
func (o *OverrideGetOverridesParams) WithClusterID(clusterID *string) *OverrideGetOverridesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the override get overrides params
func (o *OverrideGetOverridesParams) SetClusterID(clusterID *string) {
	o.ClusterID = clusterID
}

// WithCount adds the count to the override get overrides params
func (o *OverrideGetOverridesParams) WithCount(count *string) *OverrideGetOverridesParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the override get overrides params
func (o *OverrideGetOverridesParams) SetCount(count *string) {
	o.Count = count
}

// WithDeleted adds the deleted to the override get overrides params
func (o *OverrideGetOverridesParams) WithDeleted(deleted *bool) *OverrideGetOverridesParams {
	o.SetDeleted(deleted)
	return o
}

// SetDeleted adds the deleted to the override get overrides params
func (o *OverrideGetOverridesParams) SetDeleted(deleted *bool) {
	o.Deleted = deleted
}

// WithDisplayName adds the displayName to the override get overrides params
func (o *OverrideGetOverridesParams) WithDisplayName(displayName *string) *OverrideGetOverridesParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the override get overrides params
func (o *OverrideGetOverridesParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithExtended adds the extended to the override get overrides params
func (o *OverrideGetOverridesParams) WithExtended(extended *bool) *OverrideGetOverridesParams {
	o.SetExtended(extended)
	return o
}

// SetExtended adds the extended to the override get overrides params
func (o *OverrideGetOverridesParams) SetExtended(extended *bool) {
	o.Extended = extended
}

// WithGlobalScope adds the globalScope to the override get overrides params
func (o *OverrideGetOverridesParams) WithGlobalScope(globalScope *bool) *OverrideGetOverridesParams {
	o.SetGlobalScope(globalScope)
	return o
}

// SetGlobalScope adds the globalScope to the override get overrides params
func (o *OverrideGetOverridesParams) SetGlobalScope(globalScope *bool) {
	o.GlobalScope = globalScope
}

// WithGroups adds the groups to the override get overrides params
func (o *OverrideGetOverridesParams) WithGroups(groups []string) *OverrideGetOverridesParams {
	o.SetGroups(groups)
	return o
}

// SetGroups adds the groups to the override get overrides params
func (o *OverrideGetOverridesParams) SetGroups(groups []string) {
	o.Groups = groups
}

// WithIgnoreScopeDefault adds the ignoreScopeDefault to the override get overrides params
func (o *OverrideGetOverridesParams) WithIgnoreScopeDefault(ignoreScopeDefault *bool) *OverrideGetOverridesParams {
	o.SetIgnoreScopeDefault(ignoreScopeDefault)
	return o
}

// SetIgnoreScopeDefault adds the ignoreScopeDefault to the override get overrides params
func (o *OverrideGetOverridesParams) SetIgnoreScopeDefault(ignoreScopeDefault *bool) {
	o.IgnoreScopeDefault = ignoreScopeDefault
}

// WithIsSSOUser adds the isSSOUser to the override get overrides params
func (o *OverrideGetOverridesParams) WithIsSSOUser(isSSOUser *bool) *OverrideGetOverridesParams {
	o.SetIsSSOUser(isSSOUser)
	return o
}

// SetIsSSOUser adds the isSSOUser to the override get overrides params
func (o *OverrideGetOverridesParams) SetIsSSOUser(isSSOUser *bool) {
	o.IsSSOUser = isSSOUser
}

// WithLimit adds the limit to the override get overrides params
func (o *OverrideGetOverridesParams) WithLimit(limit *string) *OverrideGetOverridesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the override get overrides params
func (o *OverrideGetOverridesParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithName adds the name to the override get overrides params
func (o *OverrideGetOverridesParams) WithName(name *string) *OverrideGetOverridesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the override get overrides params
func (o *OverrideGetOverridesParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the override get overrides params
func (o *OverrideGetOverridesParams) WithOffset(offset *string) *OverrideGetOverridesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the override get overrides params
func (o *OverrideGetOverridesParams) SetOffset(offset *string) {
	o.Offset = offset
}

// WithOrder adds the order to the override get overrides params
func (o *OverrideGetOverridesParams) WithOrder(order *string) *OverrideGetOverridesParams {
	o.SetOrder(order)
	return o
}

// SetOrder adds the order to the override get overrides params
func (o *OverrideGetOverridesParams) SetOrder(order *string) {
	o.Order = order
}

// WithOrderBy adds the orderBy to the override get overrides params
func (o *OverrideGetOverridesParams) WithOrderBy(orderBy *string) *OverrideGetOverridesParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the override get overrides params
func (o *OverrideGetOverridesParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrganizationID adds the organizationID to the override get overrides params
func (o *OverrideGetOverridesParams) WithOrganizationID(organizationID *string) *OverrideGetOverridesParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the override get overrides params
func (o *OverrideGetOverridesParams) SetOrganizationID(organizationID *string) {
	o.OrganizationID = organizationID
}

// WithPartnerID adds the partnerID to the override get overrides params
func (o *OverrideGetOverridesParams) WithPartnerID(partnerID *string) *OverrideGetOverridesParams {
	o.SetPartnerID(partnerID)
	return o
}

// SetPartnerID adds the partnerId to the override get overrides params
func (o *OverrideGetOverridesParams) SetPartnerID(partnerID *string) {
	o.PartnerID = partnerID
}

// WithProjectID adds the projectID to the override get overrides params
func (o *OverrideGetOverridesParams) WithProjectID(projectID *string) *OverrideGetOverridesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the override get overrides params
func (o *OverrideGetOverridesParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithPublishedVersion adds the publishedVersion to the override get overrides params
func (o *OverrideGetOverridesParams) WithPublishedVersion(publishedVersion *string) *OverrideGetOverridesParams {
	o.SetPublishedVersion(publishedVersion)
	return o
}

// SetPublishedVersion adds the publishedVersion to the override get overrides params
func (o *OverrideGetOverridesParams) SetPublishedVersion(publishedVersion *string) {
	o.PublishedVersion = publishedVersion
}

// WithSelector adds the selector to the override get overrides params
func (o *OverrideGetOverridesParams) WithSelector(selector *string) *OverrideGetOverridesParams {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the override get overrides params
func (o *OverrideGetOverridesParams) SetSelector(selector *string) {
	o.Selector = selector
}

// WithURLScope adds the uRLScope to the override get overrides params
func (o *OverrideGetOverridesParams) WithURLScope(uRLScope string) *OverrideGetOverridesParams {
	o.SetURLScope(uRLScope)
	return o
}

// SetURLScope adds the urlScope to the override get overrides params
func (o *OverrideGetOverridesParams) SetURLScope(uRLScope string) {
	o.URLScope = uRLScope
}

// WithUsername adds the username to the override get overrides params
func (o *OverrideGetOverridesParams) WithUsername(username *string) *OverrideGetOverridesParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the override get overrides params
func (o *OverrideGetOverridesParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *OverrideGetOverridesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param ID
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("ID", qID); err != nil {
				return err
			}
		}
	}

	if o.BlueprintRef != nil {

		// query param blueprintRef
		var qrBlueprintRef string

		if o.BlueprintRef != nil {
			qrBlueprintRef = *o.BlueprintRef
		}
		qBlueprintRef := qrBlueprintRef
		if qBlueprintRef != "" {

			if err := r.SetQueryParam("blueprintRef", qBlueprintRef); err != nil {
				return err
			}
		}
	}

	if o.ClusterID != nil {

		// query param clusterID
		var qrClusterID string

		if o.ClusterID != nil {
			qrClusterID = *o.ClusterID
		}
		qClusterID := qrClusterID
		if qClusterID != "" {

			if err := r.SetQueryParam("clusterID", qClusterID); err != nil {
				return err
			}
		}
	}

	if o.Count != nil {

		// query param count
		var qrCount string

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := qrCount
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

	if o.Deleted != nil {

		// query param deleted
		var qrDeleted bool

		if o.Deleted != nil {
			qrDeleted = *o.Deleted
		}
		qDeleted := swag.FormatBool(qrDeleted)
		if qDeleted != "" {

			if err := r.SetQueryParam("deleted", qDeleted); err != nil {
				return err
			}
		}
	}

	if o.DisplayName != nil {

		// query param displayName
		var qrDisplayName string

		if o.DisplayName != nil {
			qrDisplayName = *o.DisplayName
		}
		qDisplayName := qrDisplayName
		if qDisplayName != "" {

			if err := r.SetQueryParam("displayName", qDisplayName); err != nil {
				return err
			}
		}
	}

	if o.Extended != nil {

		// query param extended
		var qrExtended bool

		if o.Extended != nil {
			qrExtended = *o.Extended
		}
		qExtended := swag.FormatBool(qrExtended)
		if qExtended != "" {

			if err := r.SetQueryParam("extended", qExtended); err != nil {
				return err
			}
		}
	}

	if o.GlobalScope != nil {

		// query param globalScope
		var qrGlobalScope bool

		if o.GlobalScope != nil {
			qrGlobalScope = *o.GlobalScope
		}
		qGlobalScope := swag.FormatBool(qrGlobalScope)
		if qGlobalScope != "" {

			if err := r.SetQueryParam("globalScope", qGlobalScope); err != nil {
				return err
			}
		}
	}

	if o.Groups != nil {

		// binding items for groups
		joinedGroups := o.bindParamGroups(reg)

		// query array param groups
		if err := r.SetQueryParam("groups", joinedGroups...); err != nil {
			return err
		}
	}

	if o.IgnoreScopeDefault != nil {

		// query param ignoreScopeDefault
		var qrIgnoreScopeDefault bool

		if o.IgnoreScopeDefault != nil {
			qrIgnoreScopeDefault = *o.IgnoreScopeDefault
		}
		qIgnoreScopeDefault := swag.FormatBool(qrIgnoreScopeDefault)
		if qIgnoreScopeDefault != "" {

			if err := r.SetQueryParam("ignoreScopeDefault", qIgnoreScopeDefault); err != nil {
				return err
			}
		}
	}

	if o.IsSSOUser != nil {

		// query param isSSOUser
		var qrIsSSOUser bool

		if o.IsSSOUser != nil {
			qrIsSSOUser = *o.IsSSOUser
		}
		qIsSSOUser := swag.FormatBool(qrIsSSOUser)
		if qIsSSOUser != "" {

			if err := r.SetQueryParam("isSSOUser", qIsSSOUser); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset string

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := qrOffset
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Order != nil {

		// query param order
		var qrOrder string

		if o.Order != nil {
			qrOrder = *o.Order
		}
		qOrder := qrOrder
		if qOrder != "" {

			if err := r.SetQueryParam("order", qOrder); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// query param orderBy
		var qrOrderBy string

		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {

			if err := r.SetQueryParam("orderBy", qOrderBy); err != nil {
				return err
			}
		}
	}

	if o.OrganizationID != nil {

		// query param organizationID
		var qrOrganizationID string

		if o.OrganizationID != nil {
			qrOrganizationID = *o.OrganizationID
		}
		qOrganizationID := qrOrganizationID
		if qOrganizationID != "" {

			if err := r.SetQueryParam("organizationID", qOrganizationID); err != nil {
				return err
			}
		}
	}

	if o.PartnerID != nil {

		// query param partnerID
		var qrPartnerID string

		if o.PartnerID != nil {
			qrPartnerID = *o.PartnerID
		}
		qPartnerID := qrPartnerID
		if qPartnerID != "" {

			if err := r.SetQueryParam("partnerID", qPartnerID); err != nil {
				return err
			}
		}
	}

	if o.ProjectID != nil {

		// query param projectID
		var qrProjectID string

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {

			if err := r.SetQueryParam("projectID", qProjectID); err != nil {
				return err
			}
		}
	}

	if o.PublishedVersion != nil {

		// query param publishedVersion
		var qrPublishedVersion string

		if o.PublishedVersion != nil {
			qrPublishedVersion = *o.PublishedVersion
		}
		qPublishedVersion := qrPublishedVersion
		if qPublishedVersion != "" {

			if err := r.SetQueryParam("publishedVersion", qPublishedVersion); err != nil {
				return err
			}
		}
	}

	if o.Selector != nil {

		// query param selector
		var qrSelector string

		if o.Selector != nil {
			qrSelector = *o.Selector
		}
		qSelector := qrSelector
		if qSelector != "" {

			if err := r.SetQueryParam("selector", qSelector); err != nil {
				return err
			}
		}
	}

	// path param urlScope
	if err := r.SetPathParam("urlScope", o.URLScope); err != nil {
		return err
	}

	if o.Username != nil {

		// query param username
		var qrUsername string

		if o.Username != nil {
			qrUsername = *o.Username
		}
		qUsername := qrUsername
		if qUsername != "" {

			if err := r.SetQueryParam("username", qUsername); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamOverrideGetOverrides binds the parameter groups
func (o *OverrideGetOverridesParams) bindParamGroups(formats strfmt.Registry) []string {
	groupsIR := o.Groups

	var groupsIC []string
	for _, groupsIIR := range groupsIR { // explode []string

		groupsIIV := groupsIIR // string as string
		groupsIC = append(groupsIC, groupsIIV)
	}

	// items.CollectionFormat: "multi"
	groupsIS := swag.JoinByFormat(groupsIC, "multi")

	return groupsIS
}