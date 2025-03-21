/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package model

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/wso2/apk/adapter/config"
	"github.com/wso2/apk/adapter/internal/interceptor"
	logger "github.com/wso2/apk/adapter/internal/loggers"
	"github.com/wso2/apk/adapter/internal/oasparser/constants"
)

// MgwSwagger represents the object structure holding the information related to the
// openAPI object. The values are populated from the extensions/properties mentioned at
// the root level of the openAPI definition. The pathItem level information is represented
// by the resources array which contains the MgwResource entries.
type MgwSwagger struct {
	id                       string
	UUID                     string
	apiType                  string
	description              string
	title                    string
	version                  string
	vendorExtensions         map[string]interface{}
	xWso2Endpoints           map[string]*EndpointCluster
	resources                []*Resource
	xWso2Basepath            string
	xWso2HTTP2BackendEnabled bool
	xWso2Cors                *CorsConfig
	securityScheme           []SecurityScheme
	security                 []map[string][]string
	xWso2ThrottlingTier      string
	xWso2AuthHeader          string
	disableSecurity          bool
	OrganizationID           string
	IsPrototyped             bool
	EndpointType             string
	LifecycleStatus          string
	xWso2RequestBodyPass     bool
	IsDefaultVersion         bool
	clientCertificates       []Certificate
	xWso2MutualSSL           string
	xWso2ApplicationSecurity bool
	EnvType                  string
	// GraphQLSchema              string
	// GraphQLComplexities        GraphQLComplexityYaml
	IsSystemAPI bool
}

// EndpointCluster represent an upstream cluster
type EndpointCluster struct {
	EndpointPrefix string
	Endpoints      []Endpoint
	// EndpointType enum {failover, loadbalance}. if any other value provided, consider as the default value; which is loadbalance
	EndpointType string
	Config       *EndpointConfig
	// Is http2 protocol enabled
	HTTP2BackendEnabled bool
}

// Endpoint represents the structure of an endpoint.
type Endpoint struct {
	// Host name
	Host string
	// BasePath (which would be added as prefix to the path mentioned in openapi definition)
	// In openAPI v2, it is determined from the basePath property
	// In openAPi v3, it is determined from the server object's suffix
	Basepath string
	// https, http, ws, wss
	// In openAPI v2, it is fetched from the schemes entry
	// In openAPI v3, it is extracted from the server property under servers object
	// only https and http are supported at the moment.
	URLType string
	// Port of the endpoint.
	// If the port is not specified, 80 is assigned if URLType is http
	// 443 is assigned if URLType is https
	Port uint32
	//ServiceDiscoveryQuery consul query for service discovery
	ServiceDiscoveryString string
	RawURL                 string
	// Trusted CA Cerificate for the endpoint
	Certificate []byte
	// Subject Alternative Names to verify in the public certificate
	AllowedSANs []string
}

// EndpointSecurity contains parameters of endpoint security at api.json
type EndpointSecurity struct {
	Password         string            `json:"password,omitempty" mapstructure:"password"`
	Type             string            `json:"type,omitempty" mapstructure:"type"`
	Enabled          bool              `json:"enabled,omitempty" mapstructure:"enabled"`
	Username         string            `json:"username,omitempty" mapstructure:"username"`
	CustomParameters map[string]string `json:"customparameters,omitempty" mapstructure:"customparameters"`
}

// EndpointConfig holds the configs such as timeout, retry, etc. for the EndpointCluster
type EndpointConfig struct {
	RetryConfig     *RetryConfig     `mapstructure:"retryConfig"`
	TimeoutInMillis uint32           `mapstructure:"timeoutInMillis"`
	CircuitBreakers *CircuitBreakers `mapstructure:"circuitBreakers"`
}

// RetryConfig holds the parameters for retries done by cc to the EndpointCluster
type RetryConfig struct {
	Count       int32    `mapstructure:"count"`
	StatusCodes []uint32 `mapstructure:"statusCodes"`
}

// CircuitBreakers holds the parameters for retries done by cc to the EndpointCluster
type CircuitBreakers struct {
	MaxConnections     int32 `mapstructure:"maxConnections"`
	MaxRequests        int32 `mapstructure:"maxRequests"`
	MaxPendingRequests int32 `mapstructure:"maxPendingRequests"`
	MaxRetries         int32 `mapstructure:"maxRetries"`
	MaxConnectionPools int32 `mapstructure:"maxConnectionPools"`
}

// SecurityScheme represents the structure of an security scheme.
type SecurityScheme struct {
	DefinitionName string // Arbitrary name used to define the security scheme. ex: default, myApikey
	Type           string // Type of the security scheme. Valid: apiKey, api_key, oauth2
	Name           string // Used for API key. Name of header or query. ex: x-api-key, apikey
	In             string // Where the api key found in. Valid: query, header
}

// CorsConfig represents the API level Cors Configuration
type CorsConfig struct {
	Enabled                       bool     `mapstructure:"corsConfigurationEnabled"`
	AccessControlAllowCredentials bool     `mapstructure:"accessControlAllowCredentials,omitempty"`
	AccessControlAllowHeaders     []string `mapstructure:"accessControlAllowHeaders"`
	AccessControlAllowMethods     []string `mapstructure:"accessControlAllowMethods"`
	AccessControlAllowOrigins     []string `mapstructure:"accessControlAllowOrigins"`
	AccessControlExposeHeaders    []string `mapstructure:"accessControlExposeHeaders"`
}

// InterceptEndpoint contains the parameters of endpoint security
type InterceptEndpoint struct {
	Enable          bool
	EndpointCluster EndpointCluster
	ClusterName     string
	ClusterTimeout  time.Duration
	RequestTimeout  time.Duration
	// Level this is an enum allowing only values {api, resource, operation}
	// to indicate from which level interceptor is added
	Level string
	// Includes this is an enum allowing only values in
	// {"request_headers", "request_body", "request_trailer", "response_headers", "response_body", "response_trailer",
	//"invocation_context" }
	Includes *interceptor.RequestInclusions
}

// Certificate contains information of a client certificate
type Certificate struct {
	Alias   string
	Tier    string
	Content []byte
}

// GetCorsConfig returns the CorsConfiguration Object.
func (swagger *MgwSwagger) GetCorsConfig() *CorsConfig {
	return swagger.xWso2Cors
}

// GetAPIType returns the openapi version
func (swagger *MgwSwagger) GetAPIType() string {
	return swagger.apiType
}

// GetVersion returns the API version
func (swagger *MgwSwagger) GetVersion() string {
	return swagger.version
}

// GetTitle returns the API Title
func (swagger *MgwSwagger) GetTitle() string {
	return swagger.title
}

// GetXWso2Basepath returns the basepath set via the vendor extension.
func (swagger *MgwSwagger) GetXWso2Basepath() string {
	return swagger.xWso2Basepath
}

// GetXWso2HTTP2BackendEnabled returns the http2 backend enabled set via the vendor extension.
func (swagger *MgwSwagger) GetXWso2HTTP2BackendEnabled() bool {
	return swagger.xWso2HTTP2BackendEnabled
}

// GetVendorExtensions returns the map of vendor extensions which are defined
// at openAPI's root level.
func (swagger *MgwSwagger) GetVendorExtensions() map[string]interface{} {
	return swagger.vendorExtensions
}

// GetXWso2Endpoints returns the array of x wso2 endpoints.
func (swagger *MgwSwagger) GetXWso2Endpoints() map[string]*EndpointCluster {
	return swagger.xWso2Endpoints
}

// GetResources returns the array of resources (openAPI path level info)
func (swagger *MgwSwagger) GetResources() []*Resource {
	return swagger.resources
}

// GetDescription returns the description of the openapi
func (swagger *MgwSwagger) GetDescription() string {
	return swagger.description
}

// GetXWso2ThrottlingTier returns the Throttling tier via the vendor extension.
func (swagger *MgwSwagger) GetXWso2ThrottlingTier() string {
	return swagger.xWso2ThrottlingTier
}

// GetDisableSecurity returns the authType via the vendor extension.
func (swagger *MgwSwagger) GetDisableSecurity() bool {
	return swagger.disableSecurity
}

// GetID returns the Id of the API
func (swagger *MgwSwagger) GetID() string {
	return swagger.id
}

// GetXWso2RequestBodyPass returns boolean value to indicate
// whether it is allowed to pass request body to the enforcer or not.
func (swagger *MgwSwagger) GetXWso2RequestBodyPass() bool {
	return swagger.xWso2RequestBodyPass
}

// GetClientCerts returns the client certificates of the API
func (swagger *MgwSwagger) GetClientCerts() []Certificate {
	return swagger.clientCertificates
}

// SetClientCerts set the client certificates of the API
func (swagger *MgwSwagger) SetClientCerts(certs []Certificate) {
	swagger.clientCertificates = certs
}

// SetID set the Id of the API
func (swagger *MgwSwagger) SetID(id string) {
	swagger.id = id
}

// SetName sets the name of the API
func (swagger *MgwSwagger) SetName(name string) {
	swagger.title = name
}

// SetSecurityScheme sets the securityschemes of the API
func (swagger *MgwSwagger) SetSecurityScheme(securityScheme []SecurityScheme) {
	swagger.securityScheme = securityScheme
}

// SetSecurity sets the API level security of the API. These refer to the security schemes
// defined for the same API and would have the structure given below,
//
// security:
//   - PetstoreAuth:
//   - 'write:pets'
//   - 'read:pets'
//   - ApiKeyAuth: []
func (swagger *MgwSwagger) SetSecurity(security []map[string][]string) {
	swagger.security = security
}

// SetVersion sets the version of the API
func (swagger *MgwSwagger) SetVersion(version string) {
	swagger.version = version
}

// SetXWso2AuthHeader sets the authHeader of the API
func (swagger *MgwSwagger) SetXWso2AuthHeader(authHeader string) {
	if swagger.xWso2AuthHeader == "" {
		swagger.xWso2AuthHeader = authHeader
	}
}

// GetXWSO2AuthHeader returns the auth header set via the vendor extension.
func (swagger *MgwSwagger) GetXWSO2AuthHeader() string {
	return swagger.xWso2AuthHeader
}

// GetSecurityScheme returns the securitySchemes of the API
func (swagger *MgwSwagger) GetSecurityScheme() []SecurityScheme {
	return swagger.securityScheme
}

// GetSecurity returns the API level security of the API
func (swagger *MgwSwagger) GetSecurity() []map[string][]string {
	return swagger.security
}

// SetXWSO2MutualSSL sets the optional or mandatory mTLS
func (swagger *MgwSwagger) SetXWSO2MutualSSL(mutualSSl string) {
	swagger.xWso2MutualSSL = mutualSSl
}

// GetXWSO2MutualSSL returns the optional or mandatory mTLS
func (swagger *MgwSwagger) GetXWSO2MutualSSL() string {
	return swagger.xWso2MutualSSL
}

// SetXWSO2ApplicationSecurity sets the optional or mandatory application security
func (swagger *MgwSwagger) SetXWSO2ApplicationSecurity(applicationSecurity bool) {
	swagger.xWso2ApplicationSecurity = applicationSecurity
}

// GetXWSO2ApplicationSecurity returns the optional or mandatory application security
func (swagger *MgwSwagger) GetXWSO2ApplicationSecurity() bool {
	return swagger.xWso2ApplicationSecurity
}

// GetOrganizationID returns OrganizationID
func (swagger *MgwSwagger) GetOrganizationID() string {
	return swagger.OrganizationID
}

// Validate method confirms that the mgwSwagger has all required fields in the required format.
// This needs to be checked prior to generate router/enforcer related resources.
func (swagger *MgwSwagger) Validate() error {
	for _, res := range swagger.resources {
		if res.endpoints == nil || len(res.endpoints.Endpoints) == 0 {
			logger.LoggerOasparser.Errorf("No Endpoints are provided for the resources in %s:%s",
				swagger.title, swagger.version)
			return errors.New("no endpoints are provided for the API")
		}
		err := res.endpoints.validateEndpointCluster()
		if err != nil {
			logger.LoggerOasparser.Errorf("Error while parsing the endpoints of the API %s:%s - %v",
				swagger.title, swagger.version, err)
			return err
		}
	}
	return nil
}

func (endpoint *Endpoint) validateEndpoint() error {
	if len(endpoint.ServiceDiscoveryString) > 0 {
		return nil
	}
	if endpoint.Port == 0 || endpoint.Port > 65535 {
		return errors.New("endpoint port value should be between 0 and 65535")
	}
	if len(endpoint.Host) == 0 {
		return errors.New("empty Hostname is provided")
	}
	if strings.HasPrefix(endpoint.Host, "/") {
		return errors.New("relative paths are not supported as endpoint URLs")
	}
	urlString := endpoint.URLType + "://" + endpoint.Host
	_, err := url.ParseRequestURI(urlString)
	return err
}

// GetAuthorityHeader creates the authority header using Host and Port in the form of Host [ ":" Port ]
func (endpoint *Endpoint) GetAuthorityHeader() string {
	return strings.Join([]string{endpoint.Host, strconv.FormatUint(uint64(endpoint.Port), 10)}, ":")
}

func (retryConfig *RetryConfig) validateRetryConfig() {
	conf := config.ReadConfigs()
	maxConfigurableCount := conf.Envoy.Upstream.Retry.MaxRetryCount
	if retryConfig.Count > int32(maxConfigurableCount) || retryConfig.Count < 0 {
		logger.LoggerOasparser.Errorf("Retry count for the API must be within the range 0 - %v."+
			"Reconfiguring retry count as %v", maxConfigurableCount, maxConfigurableCount)
		retryConfig.Count = int32(maxConfigurableCount)
	}
	var validStatusCodes []uint32
	for _, statusCode := range retryConfig.StatusCodes {
		if statusCode > 598 || statusCode < 401 {
			logger.LoggerOasparser.Errorf("Given status code for the API retry config is invalid." +
				"Must be in the range 401 - 598. Dropping the status code.")
		} else {
			validStatusCodes = append(validStatusCodes, statusCode)
		}
	}
	if len(validStatusCodes) < 1 {
		validStatusCodes = append(validStatusCodes, conf.Envoy.Upstream.Retry.StatusCodes...)
	}
	retryConfig.StatusCodes = validStatusCodes
}

func (endpointCluster *EndpointCluster) validateEndpointCluster() error {
	if endpointCluster != nil && len(endpointCluster.Endpoints) > 0 {
		var err error
		for _, endpoint := range endpointCluster.Endpoints {
			err = endpoint.validateEndpoint()
			if err != nil {
				logger.LoggerOasparser.Errorf("Error while parsing the endpoint. %v", err)
				return err
			}
		}

		if endpointCluster.Config != nil {
			// Validate retry
			if endpointCluster.Config.RetryConfig != nil {
				endpointCluster.Config.RetryConfig.validateRetryConfig()
			}
			// Validate timeout
			conf := config.ReadConfigs()
			maxTimeoutInMillis := conf.Envoy.Upstream.Timeouts.MaxRouteTimeoutInSeconds * 1000
			if endpointCluster.Config.TimeoutInMillis > maxTimeoutInMillis {
				endpointCluster.Config.TimeoutInMillis = maxTimeoutInMillis
			}
		}
	}
	return nil
}

func generateEndpointCluster(endpoints []Endpoint, endpointType string) *EndpointCluster {
	if len(endpoints) > 0 {
		endpointCluster := EndpointCluster{
			Endpoints:    endpoints,
			EndpointType: endpointType,
		}
		return &endpointCluster
	}
	return nil
}

// GetOperationInterceptors returns operation interceptors
func (swagger *MgwSwagger) GetOperationInterceptors(apiInterceptor InterceptEndpoint, resourceInterceptor InterceptEndpoint, operations []*Operation, isIn bool) map[string]InterceptEndpoint {
	interceptorOperationMap := make(map[string]InterceptEndpoint)

	for _, op := range operations {
		extensionName := constants.XWso2RequestInterceptor
		// first get operational policies
		operationInterceptor := op.GetCallInterceptorService(isIn)
		// if operational policy interceptor not given check operational level swagger extension
		if !operationInterceptor.Enable {
			if !isIn {
				extensionName = constants.XWso2ResponseInterceptor
			}
			operationInterceptor = swagger.GetInterceptor(op.GetVendorExtensions(), extensionName, constants.OperationLevelInterceptor)
		}
		operationInterceptor.ClusterName = op.iD
		// if operation interceptor not given
		if !operationInterceptor.Enable {
			// assign resource level interceptor
			if resourceInterceptor.Enable {
				operationInterceptor = resourceInterceptor
			} else if apiInterceptor.Enable {
				// if resource interceptor not given add api level interceptor
				operationInterceptor = apiInterceptor
			}
		}
		// add operation to the list only if an interceptor is enabled for the operation
		if operationInterceptor.Enable {
			interceptorOperationMap[strings.ToUpper(op.method)] = operationInterceptor
		}
	}
	return interceptorOperationMap

}

// GetInterceptor returns interceptors
func (swagger *MgwSwagger) GetInterceptor(vendorExtensions map[string]interface{}, extensionName string, level string) InterceptEndpoint {
	var endpointCluster EndpointCluster
	conf := config.ReadConfigs()
	clusterTimeoutV := conf.Envoy.ClusterTimeoutInSeconds
	requestTimeoutV := conf.Envoy.ClusterTimeoutInSeconds
	includesV := &interceptor.RequestInclusions{}

	if x, found := vendorExtensions[extensionName]; found {
		if val, ok := x.(map[string]interface{}); ok {
			//serviceURL mandatory
			if v, found := val[constants.ServiceURL]; found {
				serviceURLV := v.(string)
				endpoint, err := getHTTPEndpoint(serviceURLV)
				if err != nil {
					logger.LoggerOasparser.Error("Error reading interceptors service url value", err)
					return InterceptEndpoint{}
				}
				if endpoint.Basepath != "" {
					logger.LoggerOasparser.Warnf("Interceptor serviceURL basepath is given as %v but it will be ignored",
						endpoint.Basepath)
				}
				endpointCluster.Endpoints = []Endpoint{*endpoint}

			} else {
				logger.LoggerOasparser.Error("Error reading interceptors service url value")
				return InterceptEndpoint{}
			}
			//clusterTimeout optional
			if v, found := val[constants.ClusterTimeout]; found {
				p, err := strconv.ParseInt(fmt.Sprint(v), 0, 0)
				if err == nil {
					clusterTimeoutV = time.Duration(p)
				} else {
					logger.LoggerOasparser.Errorf("Error reading interceptors %v value : %v", constants.ClusterTimeout, err.Error())
				}
			}
			//requestTimeout optional
			if v, found := val[constants.RequestTimeout]; found {
				p, err := strconv.ParseInt(fmt.Sprint(v), 0, 0)
				if err == nil {
					requestTimeoutV = time.Duration(p)
				} else {
					logger.LoggerOasparser.Errorf("Error reading interceptors %v value : %v", constants.RequestTimeout, err.Error())
				}
			}
			//includes optional
			if v, found := val[constants.Includes]; found {
				includes := v.([]interface{})
				if len(includes) > 0 {
					// convert type of includes from "[]interface{}" to "[]string"
					includesStr := make([]string, len(includes))
					for i, v := range includes {
						includesStr[i] = v.(string)
					}
					includesV = GenerateInterceptorIncludes(includesStr)
				}
			}

			return InterceptEndpoint{
				Enable:          true,
				EndpointCluster: endpointCluster,
				ClusterTimeout:  clusterTimeoutV,
				RequestTimeout:  requestTimeoutV,
				Includes:        includesV,
				Level:           level,
			}
		}
		logger.LoggerOasparser.Error("Error parsing response interceptors values to mgwSwagger")
	}
	return InterceptEndpoint{}
}

// GenerateInterceptorIncludes generate includes
func GenerateInterceptorIncludes(includes []string) *interceptor.RequestInclusions {
	includesV := &interceptor.RequestInclusions{}
	for _, include := range includes {
		switch strings.TrimSpace(include) {
		case "request_headers":
			includesV.RequestHeaders = true
		case "request_body":
			includesV.RequestBody = true
		case "request_trailers":
			includesV.RequestTrailer = true
		case "response_headers":
			includesV.ResponseHeaders = true
		case "response_body":
			includesV.ResponseBody = true
		case "response_trailers":
			includesV.ResponseTrailers = true
		case "invocation_context":
			includesV.InvocationContext = true
		}
	}
	return includesV
}
