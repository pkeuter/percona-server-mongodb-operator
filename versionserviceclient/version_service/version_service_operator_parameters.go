// Code generated by go-swagger; DO NOT EDIT.

package version_service

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

// NewVersionServiceOperatorParams creates a new VersionServiceOperatorParams object
// with the default values initialized.
func NewVersionServiceOperatorParams() *VersionServiceOperatorParams {
	var ()
	return &VersionServiceOperatorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVersionServiceOperatorParamsWithTimeout creates a new VersionServiceOperatorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVersionServiceOperatorParamsWithTimeout(timeout time.Duration) *VersionServiceOperatorParams {
	var ()
	return &VersionServiceOperatorParams{

		timeout: timeout,
	}
}

// NewVersionServiceOperatorParamsWithContext creates a new VersionServiceOperatorParams object
// with the default values initialized, and the ability to set a context for a request
func NewVersionServiceOperatorParamsWithContext(ctx context.Context) *VersionServiceOperatorParams {
	var ()
	return &VersionServiceOperatorParams{

		Context: ctx,
	}
}

// NewVersionServiceOperatorParamsWithHTTPClient creates a new VersionServiceOperatorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVersionServiceOperatorParamsWithHTTPClient(client *http.Client) *VersionServiceOperatorParams {
	var ()
	return &VersionServiceOperatorParams{
		HTTPClient: client,
	}
}

/*VersionServiceOperatorParams contains all the parameters to send to the API endpoint
for the version service operator operation typically these are written to a http.Request
*/
type VersionServiceOperatorParams struct {

	/*BackupVersion*/
	BackupVersion *string
	/*ClusterWideEnabled*/
	ClusterWideEnabled *bool
	/*CustomResourceUID*/
	CustomResourceUID *string
	/*DatabaseVersion*/
	DatabaseVersion *string
	/*HaproxyVersion*/
	HaproxyVersion *string
	/*HashicorpVaultEnabled*/
	HashicorpVaultEnabled *bool
	/*KubeVersion*/
	KubeVersion *string
	/*LogCollectorVersion*/
	LogCollectorVersion *string
	/*NamespaceUID*/
	NamespaceUID *string
	/*OperatorVersion*/
	OperatorVersion string
	/*Platform*/
	Platform *string
	/*PmmVersion*/
	PmmVersion *string
	/*Product*/
	Product string
	/*ProxysqlVersion*/
	ProxysqlVersion *string
	/*ShardingEnabled*/
	ShardingEnabled *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the version service operator params
func (o *VersionServiceOperatorParams) WithTimeout(timeout time.Duration) *VersionServiceOperatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the version service operator params
func (o *VersionServiceOperatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the version service operator params
func (o *VersionServiceOperatorParams) WithContext(ctx context.Context) *VersionServiceOperatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the version service operator params
func (o *VersionServiceOperatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the version service operator params
func (o *VersionServiceOperatorParams) WithHTTPClient(client *http.Client) *VersionServiceOperatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the version service operator params
func (o *VersionServiceOperatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupVersion adds the backupVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithBackupVersion(backupVersion *string) *VersionServiceOperatorParams {
	o.SetBackupVersion(backupVersion)
	return o
}

// SetBackupVersion adds the backupVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetBackupVersion(backupVersion *string) {
	o.BackupVersion = backupVersion
}

// WithClusterWideEnabled adds the clusterWideEnabled to the version service operator params
func (o *VersionServiceOperatorParams) WithClusterWideEnabled(clusterWideEnabled *bool) *VersionServiceOperatorParams {
	o.SetClusterWideEnabled(clusterWideEnabled)
	return o
}

// SetClusterWideEnabled adds the clusterWideEnabled to the version service operator params
func (o *VersionServiceOperatorParams) SetClusterWideEnabled(clusterWideEnabled *bool) {
	o.ClusterWideEnabled = clusterWideEnabled
}

// WithCustomResourceUID adds the customResourceUID to the version service operator params
func (o *VersionServiceOperatorParams) WithCustomResourceUID(customResourceUID *string) *VersionServiceOperatorParams {
	o.SetCustomResourceUID(customResourceUID)
	return o
}

// SetCustomResourceUID adds the customResourceUid to the version service operator params
func (o *VersionServiceOperatorParams) SetCustomResourceUID(customResourceUID *string) {
	o.CustomResourceUID = customResourceUID
}

// WithDatabaseVersion adds the databaseVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithDatabaseVersion(databaseVersion *string) *VersionServiceOperatorParams {
	o.SetDatabaseVersion(databaseVersion)
	return o
}

// SetDatabaseVersion adds the databaseVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetDatabaseVersion(databaseVersion *string) {
	o.DatabaseVersion = databaseVersion
}

// WithHaproxyVersion adds the haproxyVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithHaproxyVersion(haproxyVersion *string) *VersionServiceOperatorParams {
	o.SetHaproxyVersion(haproxyVersion)
	return o
}

// SetHaproxyVersion adds the haproxyVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetHaproxyVersion(haproxyVersion *string) {
	o.HaproxyVersion = haproxyVersion
}

// WithHashicorpVaultEnabled adds the hashicorpVaultEnabled to the version service operator params
func (o *VersionServiceOperatorParams) WithHashicorpVaultEnabled(hashicorpVaultEnabled *bool) *VersionServiceOperatorParams {
	o.SetHashicorpVaultEnabled(hashicorpVaultEnabled)
	return o
}

// SetHashicorpVaultEnabled adds the hashicorpVaultEnabled to the version service operator params
func (o *VersionServiceOperatorParams) SetHashicorpVaultEnabled(hashicorpVaultEnabled *bool) {
	o.HashicorpVaultEnabled = hashicorpVaultEnabled
}

// WithKubeVersion adds the kubeVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithKubeVersion(kubeVersion *string) *VersionServiceOperatorParams {
	o.SetKubeVersion(kubeVersion)
	return o
}

// SetKubeVersion adds the kubeVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetKubeVersion(kubeVersion *string) {
	o.KubeVersion = kubeVersion
}

// WithLogCollectorVersion adds the logCollectorVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithLogCollectorVersion(logCollectorVersion *string) *VersionServiceOperatorParams {
	o.SetLogCollectorVersion(logCollectorVersion)
	return o
}

// SetLogCollectorVersion adds the logCollectorVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetLogCollectorVersion(logCollectorVersion *string) {
	o.LogCollectorVersion = logCollectorVersion
}

// WithNamespaceUID adds the namespaceUID to the version service operator params
func (o *VersionServiceOperatorParams) WithNamespaceUID(namespaceUID *string) *VersionServiceOperatorParams {
	o.SetNamespaceUID(namespaceUID)
	return o
}

// SetNamespaceUID adds the namespaceUid to the version service operator params
func (o *VersionServiceOperatorParams) SetNamespaceUID(namespaceUID *string) {
	o.NamespaceUID = namespaceUID
}

// WithOperatorVersion adds the operatorVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithOperatorVersion(operatorVersion string) *VersionServiceOperatorParams {
	o.SetOperatorVersion(operatorVersion)
	return o
}

// SetOperatorVersion adds the operatorVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetOperatorVersion(operatorVersion string) {
	o.OperatorVersion = operatorVersion
}

// WithPlatform adds the platform to the version service operator params
func (o *VersionServiceOperatorParams) WithPlatform(platform *string) *VersionServiceOperatorParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the version service operator params
func (o *VersionServiceOperatorParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithPmmVersion adds the pmmVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithPmmVersion(pmmVersion *string) *VersionServiceOperatorParams {
	o.SetPmmVersion(pmmVersion)
	return o
}

// SetPmmVersion adds the pmmVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetPmmVersion(pmmVersion *string) {
	o.PmmVersion = pmmVersion
}

// WithProduct adds the product to the version service operator params
func (o *VersionServiceOperatorParams) WithProduct(product string) *VersionServiceOperatorParams {
	o.SetProduct(product)
	return o
}

// SetProduct adds the product to the version service operator params
func (o *VersionServiceOperatorParams) SetProduct(product string) {
	o.Product = product
}

// WithProxysqlVersion adds the proxysqlVersion to the version service operator params
func (o *VersionServiceOperatorParams) WithProxysqlVersion(proxysqlVersion *string) *VersionServiceOperatorParams {
	o.SetProxysqlVersion(proxysqlVersion)
	return o
}

// SetProxysqlVersion adds the proxysqlVersion to the version service operator params
func (o *VersionServiceOperatorParams) SetProxysqlVersion(proxysqlVersion *string) {
	o.ProxysqlVersion = proxysqlVersion
}

// WithShardingEnabled adds the shardingEnabled to the version service operator params
func (o *VersionServiceOperatorParams) WithShardingEnabled(shardingEnabled *bool) *VersionServiceOperatorParams {
	o.SetShardingEnabled(shardingEnabled)
	return o
}

// SetShardingEnabled adds the shardingEnabled to the version service operator params
func (o *VersionServiceOperatorParams) SetShardingEnabled(shardingEnabled *bool) {
	o.ShardingEnabled = shardingEnabled
}

// WriteToRequest writes these params to a swagger request
func (o *VersionServiceOperatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BackupVersion != nil {

		// query param backupVersion
		var qrBackupVersion string
		if o.BackupVersion != nil {
			qrBackupVersion = *o.BackupVersion
		}
		qBackupVersion := qrBackupVersion
		if qBackupVersion != "" {
			if err := r.SetQueryParam("backupVersion", qBackupVersion); err != nil {
				return err
			}
		}

	}

	if o.ClusterWideEnabled != nil {

		// query param clusterWideEnabled
		var qrClusterWideEnabled bool
		if o.ClusterWideEnabled != nil {
			qrClusterWideEnabled = *o.ClusterWideEnabled
		}
		qClusterWideEnabled := swag.FormatBool(qrClusterWideEnabled)
		if qClusterWideEnabled != "" {
			if err := r.SetQueryParam("clusterWideEnabled", qClusterWideEnabled); err != nil {
				return err
			}
		}

	}

	if o.CustomResourceUID != nil {

		// query param customResourceUid
		var qrCustomResourceUID string
		if o.CustomResourceUID != nil {
			qrCustomResourceUID = *o.CustomResourceUID
		}
		qCustomResourceUID := qrCustomResourceUID
		if qCustomResourceUID != "" {
			if err := r.SetQueryParam("customResourceUid", qCustomResourceUID); err != nil {
				return err
			}
		}

	}

	if o.DatabaseVersion != nil {

		// query param databaseVersion
		var qrDatabaseVersion string
		if o.DatabaseVersion != nil {
			qrDatabaseVersion = *o.DatabaseVersion
		}
		qDatabaseVersion := qrDatabaseVersion
		if qDatabaseVersion != "" {
			if err := r.SetQueryParam("databaseVersion", qDatabaseVersion); err != nil {
				return err
			}
		}

	}

	if o.HaproxyVersion != nil {

		// query param haproxyVersion
		var qrHaproxyVersion string
		if o.HaproxyVersion != nil {
			qrHaproxyVersion = *o.HaproxyVersion
		}
		qHaproxyVersion := qrHaproxyVersion
		if qHaproxyVersion != "" {
			if err := r.SetQueryParam("haproxyVersion", qHaproxyVersion); err != nil {
				return err
			}
		}

	}

	if o.HashicorpVaultEnabled != nil {

		// query param hashicorpVaultEnabled
		var qrHashicorpVaultEnabled bool
		if o.HashicorpVaultEnabled != nil {
			qrHashicorpVaultEnabled = *o.HashicorpVaultEnabled
		}
		qHashicorpVaultEnabled := swag.FormatBool(qrHashicorpVaultEnabled)
		if qHashicorpVaultEnabled != "" {
			if err := r.SetQueryParam("hashicorpVaultEnabled", qHashicorpVaultEnabled); err != nil {
				return err
			}
		}

	}

	if o.KubeVersion != nil {

		// query param kubeVersion
		var qrKubeVersion string
		if o.KubeVersion != nil {
			qrKubeVersion = *o.KubeVersion
		}
		qKubeVersion := qrKubeVersion
		if qKubeVersion != "" {
			if err := r.SetQueryParam("kubeVersion", qKubeVersion); err != nil {
				return err
			}
		}

	}

	if o.LogCollectorVersion != nil {

		// query param logCollectorVersion
		var qrLogCollectorVersion string
		if o.LogCollectorVersion != nil {
			qrLogCollectorVersion = *o.LogCollectorVersion
		}
		qLogCollectorVersion := qrLogCollectorVersion
		if qLogCollectorVersion != "" {
			if err := r.SetQueryParam("logCollectorVersion", qLogCollectorVersion); err != nil {
				return err
			}
		}

	}

	if o.NamespaceUID != nil {

		// query param namespaceUid
		var qrNamespaceUID string
		if o.NamespaceUID != nil {
			qrNamespaceUID = *o.NamespaceUID
		}
		qNamespaceUID := qrNamespaceUID
		if qNamespaceUID != "" {
			if err := r.SetQueryParam("namespaceUid", qNamespaceUID); err != nil {
				return err
			}
		}

	}

	// path param operatorVersion
	if err := r.SetPathParam("operatorVersion", o.OperatorVersion); err != nil {
		return err
	}

	if o.Platform != nil {

		// query param platform
		var qrPlatform string
		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {
			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}

	}

	if o.PmmVersion != nil {

		// query param pmmVersion
		var qrPmmVersion string
		if o.PmmVersion != nil {
			qrPmmVersion = *o.PmmVersion
		}
		qPmmVersion := qrPmmVersion
		if qPmmVersion != "" {
			if err := r.SetQueryParam("pmmVersion", qPmmVersion); err != nil {
				return err
			}
		}

	}

	// path param product
	if err := r.SetPathParam("product", o.Product); err != nil {
		return err
	}

	if o.ProxysqlVersion != nil {

		// query param proxysqlVersion
		var qrProxysqlVersion string
		if o.ProxysqlVersion != nil {
			qrProxysqlVersion = *o.ProxysqlVersion
		}
		qProxysqlVersion := qrProxysqlVersion
		if qProxysqlVersion != "" {
			if err := r.SetQueryParam("proxysqlVersion", qProxysqlVersion); err != nil {
				return err
			}
		}

	}

	if o.ShardingEnabled != nil {

		// query param shardingEnabled
		var qrShardingEnabled bool
		if o.ShardingEnabled != nil {
			qrShardingEnabled = *o.ShardingEnabled
		}
		qShardingEnabled := swag.FormatBool(qrShardingEnabled)
		if qShardingEnabled != "" {
			if err := r.SetQueryParam("shardingEnabled", qShardingEnabled); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
