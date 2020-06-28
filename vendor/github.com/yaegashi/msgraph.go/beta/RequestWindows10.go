// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// Windows10GeneralConfigurationRequestBuilder is request builder for Windows10GeneralConfiguration
type Windows10GeneralConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns Windows10GeneralConfigurationRequest
func (b *Windows10GeneralConfigurationRequestBuilder) Request() *Windows10GeneralConfigurationRequest {
	return &Windows10GeneralConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Windows10GeneralConfigurationRequest is request for Windows10GeneralConfiguration
type Windows10GeneralConfigurationRequest struct{ BaseRequest }

// Get performs GET request for Windows10GeneralConfiguration
func (r *Windows10GeneralConfigurationRequest) Get(ctx context.Context) (resObj *Windows10GeneralConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Windows10GeneralConfiguration
func (r *Windows10GeneralConfigurationRequest) Update(ctx context.Context, reqObj *Windows10GeneralConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Windows10GeneralConfiguration
func (r *Windows10GeneralConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Windows10ImportedPFXCertificateProfileRequestBuilder is request builder for Windows10ImportedPFXCertificateProfile
type Windows10ImportedPFXCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns Windows10ImportedPFXCertificateProfileRequest
func (b *Windows10ImportedPFXCertificateProfileRequestBuilder) Request() *Windows10ImportedPFXCertificateProfileRequest {
	return &Windows10ImportedPFXCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Windows10ImportedPFXCertificateProfileRequest is request for Windows10ImportedPFXCertificateProfile
type Windows10ImportedPFXCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for Windows10ImportedPFXCertificateProfile
func (r *Windows10ImportedPFXCertificateProfileRequest) Get(ctx context.Context) (resObj *Windows10ImportedPFXCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Windows10ImportedPFXCertificateProfile
func (r *Windows10ImportedPFXCertificateProfileRequest) Update(ctx context.Context, reqObj *Windows10ImportedPFXCertificateProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Windows10ImportedPFXCertificateProfile
func (r *Windows10ImportedPFXCertificateProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Windows10PkcsCertificateProfileRequestBuilder is request builder for Windows10PkcsCertificateProfile
type Windows10PkcsCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns Windows10PkcsCertificateProfileRequest
func (b *Windows10PkcsCertificateProfileRequestBuilder) Request() *Windows10PkcsCertificateProfileRequest {
	return &Windows10PkcsCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Windows10PkcsCertificateProfileRequest is request for Windows10PkcsCertificateProfile
type Windows10PkcsCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for Windows10PkcsCertificateProfile
func (r *Windows10PkcsCertificateProfileRequest) Get(ctx context.Context) (resObj *Windows10PkcsCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Windows10PkcsCertificateProfile
func (r *Windows10PkcsCertificateProfileRequest) Update(ctx context.Context, reqObj *Windows10PkcsCertificateProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Windows10PkcsCertificateProfile
func (r *Windows10PkcsCertificateProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Windows10VpnConfigurationRequestBuilder is request builder for Windows10VpnConfiguration
type Windows10VpnConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns Windows10VpnConfigurationRequest
func (b *Windows10VpnConfigurationRequestBuilder) Request() *Windows10VpnConfigurationRequest {
	return &Windows10VpnConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Windows10VpnConfigurationRequest is request for Windows10VpnConfiguration
type Windows10VpnConfigurationRequest struct{ BaseRequest }

// Get performs GET request for Windows10VpnConfiguration
func (r *Windows10VpnConfigurationRequest) Get(ctx context.Context) (resObj *Windows10VpnConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Windows10VpnConfiguration
func (r *Windows10VpnConfigurationRequest) Update(ctx context.Context, reqObj *Windows10VpnConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Windows10VpnConfiguration
func (r *Windows10VpnConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}