// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the longer ID format settings for all resource types in a specific
// Region. This request is useful for performing a quick audit to determine whether
// a specific Region is fully opted in for longer IDs (17-character IDs). This
// request only returns information about resource types that support longer IDs.
// The following resource types support longer IDs: bundle | conversion-task |
// customer-gateway | dhcp-options | elastic-ip-allocation | elastic-ip-association
// | export-task | flow-log | image | import-task | instance | internet-gateway |
// network-acl | network-acl-association | network-interface |
// network-interface-attachment | prefix-list | reservation | route-table |
// route-table-association | security-group | snapshot | subnet |
// subnet-cidr-block-association | volume | vpc | vpc-cidr-block-association |
// vpc-endpoint | vpc-peering-connection | vpn-connection | vpn-gateway.
func (c *Client) DescribeAggregateIdFormat(ctx context.Context, params *DescribeAggregateIdFormatInput, optFns ...func(*Options)) (*DescribeAggregateIdFormatOutput, error) {
	if params == nil {
		params = &DescribeAggregateIdFormatInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAggregateIdFormat", params, optFns, addOperationDescribeAggregateIdFormatMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAggregateIdFormatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAggregateIdFormatInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DescribeAggregateIdFormatOutput struct {

	// Information about each resource's ID format.
	Statuses []types.IdFormat

	// Indicates whether all resource types in the Region are configured to use longer
	// IDs. This value is only true if all users are configured to use longer IDs for
	// all resources types in the Region.
	UseLongIdsAggregated *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAggregateIdFormatMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeAggregateIdFormat{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeAggregateIdFormat{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAggregateIdFormat(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeAggregateIdFormat(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeAggregateIdFormat",
	}
}
