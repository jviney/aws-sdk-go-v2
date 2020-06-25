// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a Query operation.
type QueryInput struct {
	_ struct{} `type:"structure"`

	// This is a legacy parameter. Use ProjectionExpression instead. For more information,
	// see AttributesToGet (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributesToGet.html)
	// in the Amazon DynamoDB Developer Guide.
	AttributesToGet []string `min:"1" type:"list"`

	// This is a legacy parameter. Use FilterExpression instead. For more information,
	// see ConditionalOperator (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html)
	// in the Amazon DynamoDB Developer Guide.
	ConditionalOperator ConditionalOperator `type:"string" enum:"true"`

	// Determines the read consistency model: If set to true, then the operation
	// uses strongly consistent reads; otherwise, the operation uses eventually
	// consistent reads.
	//
	// Strongly consistent reads are not supported on global secondary indexes.
	// If you query a global secondary index with ConsistentRead set to true, you
	// will receive a ValidationException.
	ConsistentRead *bool `type:"boolean"`

	// The primary key of the first item that this operation will evaluate. Use
	// the value that was returned for LastEvaluatedKey in the previous operation.
	//
	// The data type for ExclusiveStartKey must be String, Number, or Binary. No
	// set data types are allowed.
	ExclusiveStartKey map[string]AttributeValue `type:"map"`

	// One or more substitution tokens for attribute names in an expression. The
	// following are some use cases for using ExpressionAttributeNames:
	//
	//    * To access an attribute whose name conflicts with a DynamoDB reserved
	//    word.
	//
	//    * To create a placeholder for repeating occurrences of an attribute name
	//    in an expression.
	//
	//    * To prevent special characters in an attribute name from being misinterpreted
	//    in an expression.
	//
	// Use the # character in an expression to dereference an attribute name. For
	// example, consider the following attribute name:
	//
	//    * Percentile
	//
	// The name of this attribute conflicts with a reserved word, so it cannot be
	// used directly in an expression. (For the complete list of reserved words,
	// see Reserved Words (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html)
	// in the Amazon DynamoDB Developer Guide). To work around this, you could specify
	// the following for ExpressionAttributeNames:
	//
	//    * {"#P":"Percentile"}
	//
	// You could then use this substitution in an expression, as in this example:
	//
	//    * #P = :val
	//
	// Tokens that begin with the : character are expression attribute values, which
	// are placeholders for the actual value at runtime.
	//
	// For more information on expression attribute names, see Specifying Item Attributes
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeNames map[string]string `type:"map"`

	// One or more values that can be substituted in an expression.
	//
	// Use the : (colon) character in an expression to dereference an attribute
	// value. For example, suppose that you wanted to check whether the value of
	// the ProductStatus attribute was one of the following:
	//
	// Available | Backordered | Discontinued
	//
	// You would first need to specify ExpressionAttributeValues as follows:
	//
	// { ":avail":{"S":"Available"}, ":back":{"S":"Backordered"}, ":disc":{"S":"Discontinued"}
	// }
	//
	// You could then use these values in an expression, such as this:
	//
	// ProductStatus IN (:avail, :back, :disc)
	//
	// For more information on expression attribute values, see Specifying Conditions
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeValues map[string]AttributeValue `type:"map"`

	// A string that contains conditions that DynamoDB applies after the Query operation,
	// but before the data is returned to you. Items that do not satisfy the FilterExpression
	// criteria are not returned.
	//
	// A FilterExpression does not allow key attributes. You cannot define a filter
	// expression based on a partition key or a sort key.
	//
	// A FilterExpression is applied after the items have already been read; the
	// process of filtering does not consume any additional read capacity units.
	//
	// For more information, see Filter Expressions (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/QueryAndScan.html#FilteringResults)
	// in the Amazon DynamoDB Developer Guide.
	FilterExpression *string `type:"string"`

	// The name of an index to query. This index can be any local secondary index
	// or global secondary index on the table. Note that if you use the IndexName
	// parameter, you must also provide TableName.
	IndexName *string `min:"3" type:"string"`

	// The condition that specifies the key values for items to be retrieved by
	// the Query action.
	//
	// The condition must perform an equality test on a single partition key value.
	//
	// The condition can optionally perform one of several comparison tests on a
	// single sort key value. This allows Query to retrieve one item with a given
	// partition key value and sort key value, or several items that have the same
	// partition key value but different sort key values.
	//
	// The partition key equality test is required, and must be specified in the
	// following format:
	//
	// partitionKeyName = :partitionkeyval
	//
	// If you also want to provide a condition for the sort key, it must be combined
	// using AND with the condition for the sort key. Following is an example, using
	// the = comparison operator for the sort key:
	//
	// partitionKeyName = :partitionkeyval AND sortKeyName = :sortkeyval
	//
	// Valid comparisons for the sort key condition are as follows:
	//
	//    * sortKeyName = :sortkeyval - true if the sort key value is equal to :sortkeyval.
	//
	//    * sortKeyName < :sortkeyval - true if the sort key value is less than
	//    :sortkeyval.
	//
	//    * sortKeyName <= :sortkeyval - true if the sort key value is less than
	//    or equal to :sortkeyval.
	//
	//    * sortKeyName > :sortkeyval - true if the sort key value is greater than
	//    :sortkeyval.
	//
	//    * sortKeyName >= :sortkeyval - true if the sort key value is greater than
	//    or equal to :sortkeyval.
	//
	//    * sortKeyName BETWEEN :sortkeyval1 AND :sortkeyval2 - true if the sort
	//    key value is greater than or equal to :sortkeyval1, and less than or equal
	//    to :sortkeyval2.
	//
	//    * begins_with ( sortKeyName, :sortkeyval ) - true if the sort key value
	//    begins with a particular operand. (You cannot use this function with a
	//    sort key that is of type Number.) Note that the function name begins_with
	//    is case-sensitive.
	//
	// Use the ExpressionAttributeValues parameter to replace tokens such as :partitionval
	// and :sortval with actual values at runtime.
	//
	// You can optionally use the ExpressionAttributeNames parameter to replace
	// the names of the partition key and sort key with placeholder tokens. This
	// option might be necessary if an attribute name conflicts with a DynamoDB
	// reserved word. For example, the following KeyConditionExpression parameter
	// causes an error because Size is a reserved word:
	//
	//    * Size = :myval
	//
	// To work around this, define a placeholder (such a #S) to represent the attribute
	// name Size. KeyConditionExpression then is as follows:
	//
	//    * #S = :myval
	//
	// For a list of reserved words, see Reserved Words (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html)
	// in the Amazon DynamoDB Developer Guide.
	//
	// For more information on ExpressionAttributeNames and ExpressionAttributeValues,
	// see Using Placeholders for Attribute Names and Values (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ExpressionPlaceholders.html)
	// in the Amazon DynamoDB Developer Guide.
	KeyConditionExpression *string `type:"string"`

	// This is a legacy parameter. Use KeyConditionExpression instead. For more
	// information, see KeyConditions (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.KeyConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	KeyConditions map[string]Condition `type:"map"`

	// The maximum number of items to evaluate (not necessarily the number of matching
	// items). If DynamoDB processes the number of items up to the limit while processing
	// the results, it stops the operation and returns the matching values up to
	// that point, and a key in LastEvaluatedKey to apply in a subsequent operation,
	// so that you can pick up where you left off. Also, if the processed dataset
	// size exceeds 1 MB before DynamoDB reaches this limit, it stops the operation
	// and returns the matching values up to the limit, and a key in LastEvaluatedKey
	// to apply in a subsequent operation to continue the operation. For more information,
	// see Query and Scan (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/QueryAndScan.html)
	// in the Amazon DynamoDB Developer Guide.
	Limit *int64 `min:"1" type:"integer"`

	// A string that identifies one or more attributes to retrieve from the table.
	// These attributes can include scalars, sets, or elements of a JSON document.
	// The attributes in the expression must be separated by commas.
	//
	// If no attribute names are specified, then all attributes will be returned.
	// If any of the requested attributes are not found, they will not appear in
	// the result.
	//
	// For more information, see Accessing Item Attributes (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html)
	// in the Amazon DynamoDB Developer Guide.
	ProjectionExpression *string `type:"string"`

	// This is a legacy parameter. Use FilterExpression instead. For more information,
	// see QueryFilter (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.QueryFilter.html)
	// in the Amazon DynamoDB Developer Guide.
	QueryFilter map[string]Condition `type:"map"`

	// Determines the level of detail about provisioned throughput consumption that
	// is returned in the response:
	//
	//    * INDEXES - The response includes the aggregate ConsumedCapacity for the
	//    operation, together with ConsumedCapacity for each table and secondary
	//    index that was accessed. Note that some operations, such as GetItem and
	//    BatchGetItem, do not access any indexes at all. In these cases, specifying
	//    INDEXES will only return ConsumedCapacity information for table(s).
	//
	//    * TOTAL - The response includes only the aggregate ConsumedCapacity for
	//    the operation.
	//
	//    * NONE - No ConsumedCapacity details are included in the response.
	ReturnConsumedCapacity ReturnConsumedCapacity `type:"string" enum:"true"`

	// Specifies the order for index traversal: If true (default), the traversal
	// is performed in ascending order; if false, the traversal is performed in
	// descending order.
	//
	// Items with the same partition key value are stored in sorted order by sort
	// key. If the sort key data type is Number, the results are stored in numeric
	// order. For type String, the results are stored in order of UTF-8 bytes. For
	// type Binary, DynamoDB treats each byte of the binary data as unsigned.
	//
	// If ScanIndexForward is true, DynamoDB returns the results in the order in
	// which they are stored (by sort key value). This is the default behavior.
	// If ScanIndexForward is false, DynamoDB reads the results in reverse order
	// by sort key value, and then returns the results to the client.
	ScanIndexForward *bool `type:"boolean"`

	// The attributes to be returned in the result. You can retrieve all item attributes,
	// specific item attributes, the count of matching items, or in the case of
	// an index, some or all of the attributes projected into the index.
	//
	//    * ALL_ATTRIBUTES - Returns all of the item attributes from the specified
	//    table or index. If you query a local secondary index, then for each matching
	//    item in the index, DynamoDB fetches the entire item from the parent table.
	//    If the index is configured to project all item attributes, then all of
	//    the data can be obtained from the local secondary index, and no fetching
	//    is required.
	//
	//    * ALL_PROJECTED_ATTRIBUTES - Allowed only when querying an index. Retrieves
	//    all attributes that have been projected into the index. If the index is
	//    configured to project all attributes, this return value is equivalent
	//    to specifying ALL_ATTRIBUTES.
	//
	//    * COUNT - Returns the number of matching items, rather than the matching
	//    items themselves.
	//
	//    * SPECIFIC_ATTRIBUTES - Returns only the attributes listed in AttributesToGet.
	//    This return value is equivalent to specifying AttributesToGet without
	//    specifying any value for Select. If you query or scan a local secondary
	//    index and request only attributes that are projected into that index,
	//    the operation will read only the index and not the table. If any of the
	//    requested attributes are not projected into the local secondary index,
	//    DynamoDB fetches each of these attributes from the parent table. This
	//    extra fetching incurs additional throughput cost and latency. If you query
	//    or scan a global secondary index, you can only request attributes that
	//    are projected into the index. Global secondary index queries cannot fetch
	//    attributes from the parent table.
	//
	// If neither Select nor AttributesToGet are specified, DynamoDB defaults to
	// ALL_ATTRIBUTES when accessing a table, and ALL_PROJECTED_ATTRIBUTES when
	// accessing an index. You cannot use both Select and AttributesToGet together
	// in a single request, unless the value for Select is SPECIFIC_ATTRIBUTES.
	// (This usage is equivalent to specifying AttributesToGet without any value
	// for Select.)
	//
	// If you use the ProjectionExpression parameter, then the value for Select
	// can only be SPECIFIC_ATTRIBUTES. Any other value for Select will return an
	// error.
	Select Select `type:"string" enum:"true"`

	// The name of the table containing the requested items.
	//
	// TableName is a required field
	TableName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s QueryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *QueryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "QueryInput"}
	if s.AttributesToGet != nil && len(s.AttributesToGet) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AttributesToGet", 1))
	}
	if s.IndexName != nil && len(*s.IndexName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("IndexName", 3))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 3))
	}
	if s.KeyConditions != nil {
		for i, v := range s.KeyConditions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "KeyConditions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.QueryFilter != nil {
		for i, v := range s.QueryFilter {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "QueryFilter", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a Query operation.
type QueryOutput struct {
	_ struct{} `type:"structure"`

	// The capacity units consumed by the Query operation. The data returned includes
	// the total provisioned throughput consumed, along with statistics for the
	// table and any indexes involved in the operation. ConsumedCapacity is only
	// returned if the ReturnConsumedCapacity parameter was specified. For more
	// information, see Provisioned Throughput (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html)
	// in the Amazon DynamoDB Developer Guide.
	ConsumedCapacity *ConsumedCapacity `type:"structure"`

	// The number of items in the response.
	//
	// If you used a QueryFilter in the request, then Count is the number of items
	// returned after the filter was applied, and ScannedCount is the number of
	// matching items before the filter was applied.
	//
	// If you did not use a filter in the request, then Count and ScannedCount are
	// the same.
	Count *int64 `type:"integer"`

	// An array of item attributes that match the query criteria. Each element in
	// this array consists of an attribute name and the value for that attribute.
	Items []map[string]AttributeValue `type:"list"`

	// The primary key of the item where the operation stopped, inclusive of the
	// previous result set. Use this value to start a new operation, excluding this
	// value in the new request.
	//
	// If LastEvaluatedKey is empty, then the "last page" of results has been processed
	// and there is no more data to be retrieved.
	//
	// If LastEvaluatedKey is not empty, it does not necessarily mean that there
	// is more data in the result set. The only way to know when you have reached
	// the end of the result set is when LastEvaluatedKey is empty.
	LastEvaluatedKey map[string]AttributeValue `type:"map"`

	// The number of items evaluated, before any QueryFilter is applied. A high
	// ScannedCount value with few, or no, Count results indicates an inefficient
	// Query operation. For more information, see Count and ScannedCount (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/QueryAndScan.html#Count)
	// in the Amazon DynamoDB Developer Guide.
	//
	// If you did not use a filter in the request, then ScannedCount is the same
	// as Count.
	ScannedCount *int64 `type:"integer"`
}

// String returns the string representation
func (s QueryOutput) String() string {
	return awsutil.Prettify(s)
}

const opQuery = "Query"

// QueryRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// The Query operation finds items based on primary key values. You can query
// any table or secondary index that has a composite primary key (a partition
// key and a sort key).
//
// Use the KeyConditionExpression parameter to provide a specific value for
// the partition key. The Query operation will return all of the items from
// the table or index with that partition key value. You can optionally narrow
// the scope of the Query operation by specifying a sort key value and a comparison
// operator in KeyConditionExpression. To further refine the Query results,
// you can optionally provide a FilterExpression. A FilterExpression determines
// which items within the results should be returned to you. All of the other
// results are discarded.
//
// A Query operation always returns a result set. If no matching items are found,
// the result set will be empty. Queries that do not return results consume
// the minimum number of read capacity units for that type of read operation.
//
// DynamoDB calculates the number of read capacity units consumed based on item
// size, not on the amount of data that is returned to an application. The number
// of capacity units consumed will be the same whether you request all of the
// attributes (the default behavior) or just some of them (using a projection
// expression). The number will also be the same whether or not you use a FilterExpression.
//
// Query results are always sorted by the sort key value. If the data type of
// the sort key is Number, the results are returned in numeric order; otherwise,
// the results are returned in order of UTF-8 bytes. By default, the sort order
// is ascending. To reverse the order, set the ScanIndexForward parameter to
// false.
//
// A single Query operation will read up to the maximum number of items set
// (if using the Limit parameter) or a maximum of 1 MB of data and then apply
// any filtering to the results using FilterExpression. If LastEvaluatedKey
// is present in the response, you will need to paginate the result set. For
// more information, see Paginating the Results (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Query.html#Query.Pagination)
// in the Amazon DynamoDB Developer Guide.
//
// FilterExpression is applied after a Query finishes, but before the results
// are returned. A FilterExpression cannot contain partition key or sort key
// attributes. You need to specify those attributes in the KeyConditionExpression.
//
// A Query operation can return an empty result set and a LastEvaluatedKey if
// all the items read for the page of results are filtered out.
//
// You can query a table, a local secondary index, or a global secondary index.
// For a query on a table or on a local secondary index, you can set the ConsistentRead
// parameter to true and obtain a strongly consistent result. Global secondary
// indexes support eventually consistent reads only, so do not specify ConsistentRead
// when querying a global secondary index.
//
//    // Example sending a request using QueryRequest.
//    req := client.QueryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/Query
func (c *Client) QueryRequest(input *QueryInput) QueryRequest {
	op := &aws.Operation{
		Name:       opQuery,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"ExclusiveStartKey"},
			OutputTokens:    []string{"LastEvaluatedKey"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &QueryInput{}
	}

	req := c.newRequest(op, input, &QueryOutput{})

	if req.Config.EnableEndpointDiscovery {
		de := discovererDescribeEndpoints{
			Client:        c,
			Required:      false,
			EndpointCache: c.endpointCache,
			Params: map[string]*string{
				"op": &req.Operation.Name,
			},
		}

		for k, v := range de.Params {
			if v == nil {
				delete(de.Params, k)
			}
		}

		req.Handlers.Build.PushFrontNamed(aws.NamedHandler{
			Name: "crr.endpointdiscovery",
			Fn:   de.Handler,
		})
	}

	return QueryRequest{Request: req, Input: input, Copy: c.QueryRequest}
}

// QueryRequest is the request type for the
// Query API operation.
type QueryRequest struct {
	*aws.Request
	Input *QueryInput
	Copy  func(*QueryInput) QueryRequest
}

// Send marshals and sends the Query API request.
func (r QueryRequest) Send(ctx context.Context) (*QueryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &QueryResponse{
		QueryOutput: r.Request.Data.(*QueryOutput),
		response:    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewQueryRequestPaginator returns a paginator for Query.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.QueryRequest(input)
//   p := dynamodb.NewQueryRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewQueryPaginator(req QueryRequest) QueryPaginator {
	return QueryPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *QueryInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// QueryPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type QueryPaginator struct {
	aws.Pager
}

func (p *QueryPaginator) CurrentPage() *QueryOutput {
	return p.Pager.CurrentPage().(*QueryOutput)
}

// QueryResponse is the response type for the
// Query API operation.
type QueryResponse struct {
	*QueryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// Query request.
func (r *QueryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
