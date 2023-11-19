// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package backupgatewayiface provides an interface to enable mocking the AWS Backup Gateway service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package backupgatewayiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/backupgateway"
)

// BackupGatewayAPI provides an interface to enable mocking the
// backupgateway.BackupGateway service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Backup Gateway.
//	func myFunc(svc backupgatewayiface.BackupGatewayAPI) bool {
//	    // Make svc.AssociateGatewayToServer request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := backupgateway.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockBackupGatewayClient struct {
//	    backupgatewayiface.BackupGatewayAPI
//	}
//	func (m *mockBackupGatewayClient) AssociateGatewayToServer(input *backupgateway.AssociateGatewayToServerInput) (*backupgateway.AssociateGatewayToServerOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockBackupGatewayClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type BackupGatewayAPI interface {
	AssociateGatewayToServer(*backupgateway.AssociateGatewayToServerInput) (*backupgateway.AssociateGatewayToServerOutput, error)
	AssociateGatewayToServerWithContext(aws.Context, *backupgateway.AssociateGatewayToServerInput, ...request.Option) (*backupgateway.AssociateGatewayToServerOutput, error)
	AssociateGatewayToServerRequest(*backupgateway.AssociateGatewayToServerInput) (*request.Request, *backupgateway.AssociateGatewayToServerOutput)

	CreateGateway(*backupgateway.CreateGatewayInput) (*backupgateway.CreateGatewayOutput, error)
	CreateGatewayWithContext(aws.Context, *backupgateway.CreateGatewayInput, ...request.Option) (*backupgateway.CreateGatewayOutput, error)
	CreateGatewayRequest(*backupgateway.CreateGatewayInput) (*request.Request, *backupgateway.CreateGatewayOutput)

	DeleteGateway(*backupgateway.DeleteGatewayInput) (*backupgateway.DeleteGatewayOutput, error)
	DeleteGatewayWithContext(aws.Context, *backupgateway.DeleteGatewayInput, ...request.Option) (*backupgateway.DeleteGatewayOutput, error)
	DeleteGatewayRequest(*backupgateway.DeleteGatewayInput) (*request.Request, *backupgateway.DeleteGatewayOutput)

	DeleteHypervisor(*backupgateway.DeleteHypervisorInput) (*backupgateway.DeleteHypervisorOutput, error)
	DeleteHypervisorWithContext(aws.Context, *backupgateway.DeleteHypervisorInput, ...request.Option) (*backupgateway.DeleteHypervisorOutput, error)
	DeleteHypervisorRequest(*backupgateway.DeleteHypervisorInput) (*request.Request, *backupgateway.DeleteHypervisorOutput)

	DisassociateGatewayFromServer(*backupgateway.DisassociateGatewayFromServerInput) (*backupgateway.DisassociateGatewayFromServerOutput, error)
	DisassociateGatewayFromServerWithContext(aws.Context, *backupgateway.DisassociateGatewayFromServerInput, ...request.Option) (*backupgateway.DisassociateGatewayFromServerOutput, error)
	DisassociateGatewayFromServerRequest(*backupgateway.DisassociateGatewayFromServerInput) (*request.Request, *backupgateway.DisassociateGatewayFromServerOutput)

	GetGateway(*backupgateway.GetGatewayInput) (*backupgateway.GetGatewayOutput, error)
	GetGatewayWithContext(aws.Context, *backupgateway.GetGatewayInput, ...request.Option) (*backupgateway.GetGatewayOutput, error)
	GetGatewayRequest(*backupgateway.GetGatewayInput) (*request.Request, *backupgateway.GetGatewayOutput)

	GetVirtualMachine(*backupgateway.GetVirtualMachineInput) (*backupgateway.GetVirtualMachineOutput, error)
	GetVirtualMachineWithContext(aws.Context, *backupgateway.GetVirtualMachineInput, ...request.Option) (*backupgateway.GetVirtualMachineOutput, error)
	GetVirtualMachineRequest(*backupgateway.GetVirtualMachineInput) (*request.Request, *backupgateway.GetVirtualMachineOutput)

	ImportHypervisorConfiguration(*backupgateway.ImportHypervisorConfigurationInput) (*backupgateway.ImportHypervisorConfigurationOutput, error)
	ImportHypervisorConfigurationWithContext(aws.Context, *backupgateway.ImportHypervisorConfigurationInput, ...request.Option) (*backupgateway.ImportHypervisorConfigurationOutput, error)
	ImportHypervisorConfigurationRequest(*backupgateway.ImportHypervisorConfigurationInput) (*request.Request, *backupgateway.ImportHypervisorConfigurationOutput)

	ListGateways(*backupgateway.ListGatewaysInput) (*backupgateway.ListGatewaysOutput, error)
	ListGatewaysWithContext(aws.Context, *backupgateway.ListGatewaysInput, ...request.Option) (*backupgateway.ListGatewaysOutput, error)
	ListGatewaysRequest(*backupgateway.ListGatewaysInput) (*request.Request, *backupgateway.ListGatewaysOutput)

	ListGatewaysPages(*backupgateway.ListGatewaysInput, func(*backupgateway.ListGatewaysOutput, bool) bool) error
	ListGatewaysPagesWithContext(aws.Context, *backupgateway.ListGatewaysInput, func(*backupgateway.ListGatewaysOutput, bool) bool, ...request.Option) error

	ListHypervisors(*backupgateway.ListHypervisorsInput) (*backupgateway.ListHypervisorsOutput, error)
	ListHypervisorsWithContext(aws.Context, *backupgateway.ListHypervisorsInput, ...request.Option) (*backupgateway.ListHypervisorsOutput, error)
	ListHypervisorsRequest(*backupgateway.ListHypervisorsInput) (*request.Request, *backupgateway.ListHypervisorsOutput)

	ListHypervisorsPages(*backupgateway.ListHypervisorsInput, func(*backupgateway.ListHypervisorsOutput, bool) bool) error
	ListHypervisorsPagesWithContext(aws.Context, *backupgateway.ListHypervisorsInput, func(*backupgateway.ListHypervisorsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*backupgateway.ListTagsForResourceInput) (*backupgateway.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *backupgateway.ListTagsForResourceInput, ...request.Option) (*backupgateway.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*backupgateway.ListTagsForResourceInput) (*request.Request, *backupgateway.ListTagsForResourceOutput)

	ListVirtualMachines(*backupgateway.ListVirtualMachinesInput) (*backupgateway.ListVirtualMachinesOutput, error)
	ListVirtualMachinesWithContext(aws.Context, *backupgateway.ListVirtualMachinesInput, ...request.Option) (*backupgateway.ListVirtualMachinesOutput, error)
	ListVirtualMachinesRequest(*backupgateway.ListVirtualMachinesInput) (*request.Request, *backupgateway.ListVirtualMachinesOutput)

	ListVirtualMachinesPages(*backupgateway.ListVirtualMachinesInput, func(*backupgateway.ListVirtualMachinesOutput, bool) bool) error
	ListVirtualMachinesPagesWithContext(aws.Context, *backupgateway.ListVirtualMachinesInput, func(*backupgateway.ListVirtualMachinesOutput, bool) bool, ...request.Option) error

	PutMaintenanceStartTime(*backupgateway.PutMaintenanceStartTimeInput) (*backupgateway.PutMaintenanceStartTimeOutput, error)
	PutMaintenanceStartTimeWithContext(aws.Context, *backupgateway.PutMaintenanceStartTimeInput, ...request.Option) (*backupgateway.PutMaintenanceStartTimeOutput, error)
	PutMaintenanceStartTimeRequest(*backupgateway.PutMaintenanceStartTimeInput) (*request.Request, *backupgateway.PutMaintenanceStartTimeOutput)

	TagResource(*backupgateway.TagResourceInput) (*backupgateway.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *backupgateway.TagResourceInput, ...request.Option) (*backupgateway.TagResourceOutput, error)
	TagResourceRequest(*backupgateway.TagResourceInput) (*request.Request, *backupgateway.TagResourceOutput)

	TestHypervisorConfiguration(*backupgateway.TestHypervisorConfigurationInput) (*backupgateway.TestHypervisorConfigurationOutput, error)
	TestHypervisorConfigurationWithContext(aws.Context, *backupgateway.TestHypervisorConfigurationInput, ...request.Option) (*backupgateway.TestHypervisorConfigurationOutput, error)
	TestHypervisorConfigurationRequest(*backupgateway.TestHypervisorConfigurationInput) (*request.Request, *backupgateway.TestHypervisorConfigurationOutput)

	UntagResource(*backupgateway.UntagResourceInput) (*backupgateway.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *backupgateway.UntagResourceInput, ...request.Option) (*backupgateway.UntagResourceOutput, error)
	UntagResourceRequest(*backupgateway.UntagResourceInput) (*request.Request, *backupgateway.UntagResourceOutput)

	UpdateGatewayInformation(*backupgateway.UpdateGatewayInformationInput) (*backupgateway.UpdateGatewayInformationOutput, error)
	UpdateGatewayInformationWithContext(aws.Context, *backupgateway.UpdateGatewayInformationInput, ...request.Option) (*backupgateway.UpdateGatewayInformationOutput, error)
	UpdateGatewayInformationRequest(*backupgateway.UpdateGatewayInformationInput) (*request.Request, *backupgateway.UpdateGatewayInformationOutput)

	UpdateGatewaySoftwareNow(*backupgateway.UpdateGatewaySoftwareNowInput) (*backupgateway.UpdateGatewaySoftwareNowOutput, error)
	UpdateGatewaySoftwareNowWithContext(aws.Context, *backupgateway.UpdateGatewaySoftwareNowInput, ...request.Option) (*backupgateway.UpdateGatewaySoftwareNowOutput, error)
	UpdateGatewaySoftwareNowRequest(*backupgateway.UpdateGatewaySoftwareNowInput) (*request.Request, *backupgateway.UpdateGatewaySoftwareNowOutput)

	UpdateHypervisor(*backupgateway.UpdateHypervisorInput) (*backupgateway.UpdateHypervisorOutput, error)
	UpdateHypervisorWithContext(aws.Context, *backupgateway.UpdateHypervisorInput, ...request.Option) (*backupgateway.UpdateHypervisorOutput, error)
	UpdateHypervisorRequest(*backupgateway.UpdateHypervisorInput) (*request.Request, *backupgateway.UpdateHypervisorOutput)
}

var _ BackupGatewayAPI = (*backupgateway.BackupGateway)(nil)
