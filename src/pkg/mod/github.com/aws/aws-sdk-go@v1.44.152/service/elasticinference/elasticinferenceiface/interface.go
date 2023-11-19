// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package elasticinferenceiface provides an interface to enable mocking the Amazon Elastic  Inference service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package elasticinferenceiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/elasticinference"
)

// ElasticInferenceAPI provides an interface to enable mocking the
// elasticinference.ElasticInference service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Elastic  Inference.
//	func myFunc(svc elasticinferenceiface.ElasticInferenceAPI) bool {
//	    // Make svc.DescribeAcceleratorOfferings request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := elasticinference.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockElasticInferenceClient struct {
//	    elasticinferenceiface.ElasticInferenceAPI
//	}
//	func (m *mockElasticInferenceClient) DescribeAcceleratorOfferings(input *elasticinference.DescribeAcceleratorOfferingsInput) (*elasticinference.DescribeAcceleratorOfferingsOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockElasticInferenceClient{}
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
type ElasticInferenceAPI interface {
	DescribeAcceleratorOfferings(*elasticinference.DescribeAcceleratorOfferingsInput) (*elasticinference.DescribeAcceleratorOfferingsOutput, error)
	DescribeAcceleratorOfferingsWithContext(aws.Context, *elasticinference.DescribeAcceleratorOfferingsInput, ...request.Option) (*elasticinference.DescribeAcceleratorOfferingsOutput, error)
	DescribeAcceleratorOfferingsRequest(*elasticinference.DescribeAcceleratorOfferingsInput) (*request.Request, *elasticinference.DescribeAcceleratorOfferingsOutput)

	DescribeAcceleratorTypes(*elasticinference.DescribeAcceleratorTypesInput) (*elasticinference.DescribeAcceleratorTypesOutput, error)
	DescribeAcceleratorTypesWithContext(aws.Context, *elasticinference.DescribeAcceleratorTypesInput, ...request.Option) (*elasticinference.DescribeAcceleratorTypesOutput, error)
	DescribeAcceleratorTypesRequest(*elasticinference.DescribeAcceleratorTypesInput) (*request.Request, *elasticinference.DescribeAcceleratorTypesOutput)

	DescribeAccelerators(*elasticinference.DescribeAcceleratorsInput) (*elasticinference.DescribeAcceleratorsOutput, error)
	DescribeAcceleratorsWithContext(aws.Context, *elasticinference.DescribeAcceleratorsInput, ...request.Option) (*elasticinference.DescribeAcceleratorsOutput, error)
	DescribeAcceleratorsRequest(*elasticinference.DescribeAcceleratorsInput) (*request.Request, *elasticinference.DescribeAcceleratorsOutput)

	DescribeAcceleratorsPages(*elasticinference.DescribeAcceleratorsInput, func(*elasticinference.DescribeAcceleratorsOutput, bool) bool) error
	DescribeAcceleratorsPagesWithContext(aws.Context, *elasticinference.DescribeAcceleratorsInput, func(*elasticinference.DescribeAcceleratorsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*elasticinference.ListTagsForResourceInput) (*elasticinference.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *elasticinference.ListTagsForResourceInput, ...request.Option) (*elasticinference.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*elasticinference.ListTagsForResourceInput) (*request.Request, *elasticinference.ListTagsForResourceOutput)

	TagResource(*elasticinference.TagResourceInput) (*elasticinference.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *elasticinference.TagResourceInput, ...request.Option) (*elasticinference.TagResourceOutput, error)
	TagResourceRequest(*elasticinference.TagResourceInput) (*request.Request, *elasticinference.TagResourceOutput)

	UntagResource(*elasticinference.UntagResourceInput) (*elasticinference.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *elasticinference.UntagResourceInput, ...request.Option) (*elasticinference.UntagResourceOutput, error)
	UntagResourceRequest(*elasticinference.UntagResourceInput) (*request.Request, *elasticinference.UntagResourceOutput)
}

var _ ElasticInferenceAPI = (*elasticinference.ElasticInference)(nil)
