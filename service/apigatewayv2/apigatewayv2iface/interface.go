// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package apigatewayv2iface provides an interface to enable mocking the AmazonApiGatewayV2 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package apigatewayv2iface

import (
	"github.com/jviney/aws-sdk-go-v2/service/apigatewayv2"
)

// ClientAPI provides an interface to enable mocking the
// apigatewayv2.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AmazonApiGatewayV2.
//    func myFunc(svc apigatewayv2iface.ClientAPI) bool {
//        // Make svc.CreateApi request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := apigatewayv2.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        apigatewayv2iface.ClientPI
//    }
//    func (m *mockClientClient) CreateApi(input *apigatewayv2.CreateApiInput) (*apigatewayv2.CreateApiOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	CreateApiRequest(*apigatewayv2.CreateApiInput) apigatewayv2.CreateApiRequest

	CreateApiMappingRequest(*apigatewayv2.CreateApiMappingInput) apigatewayv2.CreateApiMappingRequest

	CreateAuthorizerRequest(*apigatewayv2.CreateAuthorizerInput) apigatewayv2.CreateAuthorizerRequest

	CreateDeploymentRequest(*apigatewayv2.CreateDeploymentInput) apigatewayv2.CreateDeploymentRequest

	CreateDomainNameRequest(*apigatewayv2.CreateDomainNameInput) apigatewayv2.CreateDomainNameRequest

	CreateIntegrationRequest(*apigatewayv2.CreateIntegrationInput) apigatewayv2.CreateIntegrationRequest

	CreateIntegrationResponseRequest(*apigatewayv2.CreateIntegrationResponseInput) apigatewayv2.CreateIntegrationResponseRequest

	CreateModelRequest(*apigatewayv2.CreateModelInput) apigatewayv2.CreateModelRequest

	CreateRouteRequest(*apigatewayv2.CreateRouteInput) apigatewayv2.CreateRouteRequest

	CreateRouteResponseRequest(*apigatewayv2.CreateRouteResponseInput) apigatewayv2.CreateRouteResponseRequest

	CreateStageRequest(*apigatewayv2.CreateStageInput) apigatewayv2.CreateStageRequest

	CreateVpcLinkRequest(*apigatewayv2.CreateVpcLinkInput) apigatewayv2.CreateVpcLinkRequest

	DeleteAccessLogSettingsRequest(*apigatewayv2.DeleteAccessLogSettingsInput) apigatewayv2.DeleteAccessLogSettingsRequest

	DeleteApiRequest(*apigatewayv2.DeleteApiInput) apigatewayv2.DeleteApiRequest

	DeleteApiMappingRequest(*apigatewayv2.DeleteApiMappingInput) apigatewayv2.DeleteApiMappingRequest

	DeleteAuthorizerRequest(*apigatewayv2.DeleteAuthorizerInput) apigatewayv2.DeleteAuthorizerRequest

	DeleteCorsConfigurationRequest(*apigatewayv2.DeleteCorsConfigurationInput) apigatewayv2.DeleteCorsConfigurationRequest

	DeleteDeploymentRequest(*apigatewayv2.DeleteDeploymentInput) apigatewayv2.DeleteDeploymentRequest

	DeleteDomainNameRequest(*apigatewayv2.DeleteDomainNameInput) apigatewayv2.DeleteDomainNameRequest

	DeleteIntegrationRequest(*apigatewayv2.DeleteIntegrationInput) apigatewayv2.DeleteIntegrationRequest

	DeleteIntegrationResponseRequest(*apigatewayv2.DeleteIntegrationResponseInput) apigatewayv2.DeleteIntegrationResponseRequest

	DeleteModelRequest(*apigatewayv2.DeleteModelInput) apigatewayv2.DeleteModelRequest

	DeleteRouteRequest(*apigatewayv2.DeleteRouteInput) apigatewayv2.DeleteRouteRequest

	DeleteRouteRequestParameterRequest(*apigatewayv2.DeleteRouteRequestParameterInput) apigatewayv2.DeleteRouteRequestParameterRequest

	DeleteRouteResponseRequest(*apigatewayv2.DeleteRouteResponseInput) apigatewayv2.DeleteRouteResponseRequest

	DeleteRouteSettingsRequest(*apigatewayv2.DeleteRouteSettingsInput) apigatewayv2.DeleteRouteSettingsRequest

	DeleteStageRequest(*apigatewayv2.DeleteStageInput) apigatewayv2.DeleteStageRequest

	DeleteVpcLinkRequest(*apigatewayv2.DeleteVpcLinkInput) apigatewayv2.DeleteVpcLinkRequest

	ExportApiRequest(*apigatewayv2.ExportApiInput) apigatewayv2.ExportApiRequest

	GetApiRequest(*apigatewayv2.GetApiInput) apigatewayv2.GetApiRequest

	GetApiMappingRequest(*apigatewayv2.GetApiMappingInput) apigatewayv2.GetApiMappingRequest

	GetApiMappingsRequest(*apigatewayv2.GetApiMappingsInput) apigatewayv2.GetApiMappingsRequest

	GetApisRequest(*apigatewayv2.GetApisInput) apigatewayv2.GetApisRequest

	GetAuthorizerRequest(*apigatewayv2.GetAuthorizerInput) apigatewayv2.GetAuthorizerRequest

	GetAuthorizersRequest(*apigatewayv2.GetAuthorizersInput) apigatewayv2.GetAuthorizersRequest

	GetDeploymentRequest(*apigatewayv2.GetDeploymentInput) apigatewayv2.GetDeploymentRequest

	GetDeploymentsRequest(*apigatewayv2.GetDeploymentsInput) apigatewayv2.GetDeploymentsRequest

	GetDomainNameRequest(*apigatewayv2.GetDomainNameInput) apigatewayv2.GetDomainNameRequest

	GetDomainNamesRequest(*apigatewayv2.GetDomainNamesInput) apigatewayv2.GetDomainNamesRequest

	GetIntegrationRequest(*apigatewayv2.GetIntegrationInput) apigatewayv2.GetIntegrationRequest

	GetIntegrationResponseRequest(*apigatewayv2.GetIntegrationResponseInput) apigatewayv2.GetIntegrationResponseRequest

	GetIntegrationResponsesRequest(*apigatewayv2.GetIntegrationResponsesInput) apigatewayv2.GetIntegrationResponsesRequest

	GetIntegrationsRequest(*apigatewayv2.GetIntegrationsInput) apigatewayv2.GetIntegrationsRequest

	GetModelRequest(*apigatewayv2.GetModelInput) apigatewayv2.GetModelRequest

	GetModelTemplateRequest(*apigatewayv2.GetModelTemplateInput) apigatewayv2.GetModelTemplateRequest

	GetModelsRequest(*apigatewayv2.GetModelsInput) apigatewayv2.GetModelsRequest

	GetRouteRequest(*apigatewayv2.GetRouteInput) apigatewayv2.GetRouteRequest

	GetRouteResponseRequest(*apigatewayv2.GetRouteResponseInput) apigatewayv2.GetRouteResponseRequest

	GetRouteResponsesRequest(*apigatewayv2.GetRouteResponsesInput) apigatewayv2.GetRouteResponsesRequest

	GetRoutesRequest(*apigatewayv2.GetRoutesInput) apigatewayv2.GetRoutesRequest

	GetStageRequest(*apigatewayv2.GetStageInput) apigatewayv2.GetStageRequest

	GetStagesRequest(*apigatewayv2.GetStagesInput) apigatewayv2.GetStagesRequest

	GetTagsRequest(*apigatewayv2.GetTagsInput) apigatewayv2.GetTagsRequest

	GetVpcLinkRequest(*apigatewayv2.GetVpcLinkInput) apigatewayv2.GetVpcLinkRequest

	GetVpcLinksRequest(*apigatewayv2.GetVpcLinksInput) apigatewayv2.GetVpcLinksRequest

	ImportApiRequest(*apigatewayv2.ImportApiInput) apigatewayv2.ImportApiRequest

	ReimportApiRequest(*apigatewayv2.ReimportApiInput) apigatewayv2.ReimportApiRequest

	TagResourceRequest(*apigatewayv2.TagResourceInput) apigatewayv2.TagResourceRequest

	UntagResourceRequest(*apigatewayv2.UntagResourceInput) apigatewayv2.UntagResourceRequest

	UpdateApiRequest(*apigatewayv2.UpdateApiInput) apigatewayv2.UpdateApiRequest

	UpdateApiMappingRequest(*apigatewayv2.UpdateApiMappingInput) apigatewayv2.UpdateApiMappingRequest

	UpdateAuthorizerRequest(*apigatewayv2.UpdateAuthorizerInput) apigatewayv2.UpdateAuthorizerRequest

	UpdateDeploymentRequest(*apigatewayv2.UpdateDeploymentInput) apigatewayv2.UpdateDeploymentRequest

	UpdateDomainNameRequest(*apigatewayv2.UpdateDomainNameInput) apigatewayv2.UpdateDomainNameRequest

	UpdateIntegrationRequest(*apigatewayv2.UpdateIntegrationInput) apigatewayv2.UpdateIntegrationRequest

	UpdateIntegrationResponseRequest(*apigatewayv2.UpdateIntegrationResponseInput) apigatewayv2.UpdateIntegrationResponseRequest

	UpdateModelRequest(*apigatewayv2.UpdateModelInput) apigatewayv2.UpdateModelRequest

	UpdateRouteRequest(*apigatewayv2.UpdateRouteInput) apigatewayv2.UpdateRouteRequest

	UpdateRouteResponseRequest(*apigatewayv2.UpdateRouteResponseInput) apigatewayv2.UpdateRouteResponseRequest

	UpdateStageRequest(*apigatewayv2.UpdateStageInput) apigatewayv2.UpdateStageRequest

	UpdateVpcLinkRequest(*apigatewayv2.UpdateVpcLinkInput) apigatewayv2.UpdateVpcLinkRequest
}

var _ ClientAPI = (*apigatewayv2.Client)(nil)
