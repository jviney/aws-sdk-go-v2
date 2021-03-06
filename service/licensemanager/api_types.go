// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package licensemanager

import (
	"fmt"
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Describes automated discovery.
type AutomatedDiscoveryInformation struct {
	_ struct{} `type:"structure"`

	// Time that automated discovery last ran.
	LastRunTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s AutomatedDiscoveryInformation) String() string {
	return awsutil.Prettify(s)
}

// Details about license consumption.
type ConsumedLicenseSummary struct {
	_ struct{} `type:"structure"`

	// Number of licenses consumed by the resource.
	ConsumedLicenses *int64 `type:"long"`

	// Resource type of the resource consuming a license.
	ResourceType ResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ConsumedLicenseSummary) String() string {
	return awsutil.Prettify(s)
}

// A filter name and value pair that is used to return more specific results
// from a describe operation. Filters can be used to match a set of resources
// by specific criteria, such as tags, attributes, or IDs.
type Filter struct {
	_ struct{} `type:"structure"`

	// Name of the filter. Filter names are case-sensitive.
	Name *string `type:"string"`

	// Filter values. Filter values are case-sensitive.
	Values []string `type:"list"`
}

// String returns the string representation
func (s Filter) String() string {
	return awsutil.Prettify(s)
}

// An inventory filter.
type InventoryFilter struct {
	_ struct{} `type:"structure"`

	// Condition of the filter.
	//
	// Condition is a required field
	Condition InventoryFilterCondition `type:"string" required:"true" enum:"true"`

	// Name of the filter.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// Value of the filter.
	Value *string `type:"string"`
}

// String returns the string representation
func (s InventoryFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InventoryFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InventoryFilter"}
	if len(s.Condition) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Condition"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A license configuration is an abstraction of a customer license agreement
// that can be consumed and enforced by License Manager. Components include
// specifications for the license type (licensing by instance, socket, CPU,
// or vCPU), allowed tenancy (shared tenancy, Dedicated Instance, Dedicated
// Host, or all of these), host affinity (how long a VM must be associated with
// a host), and the number of licenses purchased and used.
type LicenseConfiguration struct {
	_ struct{} `type:"structure"`

	// Automated discovery information.
	AutomatedDiscoveryInformation *AutomatedDiscoveryInformation `type:"structure"`

	// Summaries for licenses consumed by various resources.
	ConsumedLicenseSummaryList []ConsumedLicenseSummary `type:"list"`

	// Number of licenses consumed.
	ConsumedLicenses *int64 `type:"long"`

	// Description of the license configuration.
	Description *string `type:"string"`

	// Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string `type:"string"`

	// Unique ID of the license configuration.
	LicenseConfigurationId *string `type:"string"`

	// Number of licenses managed by the license configuration.
	LicenseCount *int64 `type:"long"`

	// Number of available licenses as a hard limit.
	LicenseCountHardLimit *bool `type:"boolean"`

	// Dimension to use to track the license inventory.
	LicenseCountingType LicenseCountingType `type:"string" enum:"true"`

	// License rules.
	LicenseRules []string `type:"list"`

	// Summaries for managed resources.
	ManagedResourceSummaryList []ManagedResourceSummary `type:"list"`

	// Name of the license configuration.
	Name *string `type:"string"`

	// Account ID of the license configuration's owner.
	OwnerAccountId *string `type:"string"`

	// Product information.
	ProductInformationList []ProductInformation `type:"list"`

	// Status of the license configuration.
	Status *string `type:"string"`
}

// String returns the string representation
func (s LicenseConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Describes an association with a license configuration.
type LicenseConfigurationAssociation struct {
	_ struct{} `type:"structure"`

	// Time when the license configuration was associated with the resource.
	AssociationTime *time.Time `type:"timestamp"`

	// Amazon Resource Name (ARN) of the resource.
	ResourceArn *string `type:"string"`

	// ID of the AWS account that owns the resource consuming licenses.
	ResourceOwnerId *string `type:"string"`

	// Type of server resource.
	ResourceType ResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s LicenseConfigurationAssociation) String() string {
	return awsutil.Prettify(s)
}

// Details about the usage of a resource associated with a license configuration.
type LicenseConfigurationUsage struct {
	_ struct{} `type:"structure"`

	// Time when the license configuration was initially associated with the resource.
	AssociationTime *time.Time `type:"timestamp"`

	// Number of licenses consumed by the resource.
	ConsumedLicenses *int64 `type:"long"`

	// Amazon Resource Name (ARN) of the resource.
	ResourceArn *string `type:"string"`

	// ID of the account that owns the resource.
	ResourceOwnerId *string `type:"string"`

	// Status of the resource.
	ResourceStatus *string `type:"string"`

	// Type of resource.
	ResourceType ResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s LicenseConfigurationUsage) String() string {
	return awsutil.Prettify(s)
}

// Describes the failure of a license operation.
type LicenseOperationFailure struct {
	_ struct{} `type:"structure"`

	// Error message.
	ErrorMessage *string `type:"string"`

	// Failure time.
	FailureTime *time.Time `type:"timestamp"`

	// Reserved.
	MetadataList []Metadata `type:"list"`

	// Name of the operation.
	OperationName *string `type:"string"`

	// The requester is "License Manager Automated Discovery".
	OperationRequestedBy *string `type:"string"`

	// Amazon Resource Name (ARN) of the resource.
	ResourceArn *string `type:"string"`

	// ID of the AWS account that owns the resource.
	ResourceOwnerId *string `type:"string"`

	// Resource type.
	ResourceType ResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s LicenseOperationFailure) String() string {
	return awsutil.Prettify(s)
}

// Details for associating a license configuration with a resource.
type LicenseSpecification struct {
	_ struct{} `type:"structure"`

	// Amazon Resource Name (ARN) of the license configuration.
	//
	// LicenseConfigurationArn is a required field
	LicenseConfigurationArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s LicenseSpecification) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *LicenseSpecification) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "LicenseSpecification"}

	if s.LicenseConfigurationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LicenseConfigurationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Summary information about a managed resource.
type ManagedResourceSummary struct {
	_ struct{} `type:"structure"`

	// Number of resources associated with licenses.
	AssociationCount *int64 `type:"long"`

	// Type of resource associated with a license.
	ResourceType ResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ManagedResourceSummary) String() string {
	return awsutil.Prettify(s)
}

// Reserved.
type Metadata struct {
	_ struct{} `type:"structure"`

	// Reserved.
	Name *string `type:"string"`

	// Reserved.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Metadata) String() string {
	return awsutil.Prettify(s)
}

// Configuration information for AWS Organizations.
type OrganizationConfiguration struct {
	_ struct{} `type:"structure"`

	// Enables AWS Organization integration.
	//
	// EnableIntegration is a required field
	EnableIntegration *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s OrganizationConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *OrganizationConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "OrganizationConfiguration"}

	if s.EnableIntegration == nil {
		invalidParams.Add(aws.NewErrParamRequired("EnableIntegration"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes product information for a license configuration.
type ProductInformation struct {
	_ struct{} `type:"structure"`

	// Product information filters. The following filters and logical operators
	// are supported:
	//
	//    * Application Name - The name of the application. Logical operator is
	//    EQUALS.
	//
	//    * Application Publisher - The publisher of the application. Logical operator
	//    is EQUALS.
	//
	//    * Application Version - The version of the application. Logical operator
	//    is EQUALS.
	//
	//    * Platform Name - The name of the platform. Logical operator is EQUALS.
	//
	//    * Platform Type - The platform type. Logical operator is EQUALS.
	//
	//    * License Included - The type of license included. Logical operators are
	//    EQUALS and NOT_EQUALS. Possible values are sql-server-enterprise | sql-server-standard
	//    | sql-server-web | windows-server-datacenter.
	//
	// ProductInformationFilterList is a required field
	ProductInformationFilterList []ProductInformationFilter `type:"list" required:"true"`

	// Resource type. The value is SSM_MANAGED.
	//
	// ResourceType is a required field
	ResourceType *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ProductInformation) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ProductInformation) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ProductInformation"}

	if s.ProductInformationFilterList == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductInformationFilterList"))
	}

	if s.ResourceType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}
	if s.ProductInformationFilterList != nil {
		for i, v := range s.ProductInformationFilterList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ProductInformationFilterList", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes product information filters.
type ProductInformationFilter struct {
	_ struct{} `type:"structure"`

	// Logical operator.
	//
	// ProductInformationFilterComparator is a required field
	ProductInformationFilterComparator *string `type:"string" required:"true"`

	// Filter name.
	//
	// ProductInformationFilterName is a required field
	ProductInformationFilterName *string `type:"string" required:"true"`

	// Filter value.
	//
	// ProductInformationFilterValue is a required field
	ProductInformationFilterValue []string `type:"list" required:"true"`
}

// String returns the string representation
func (s ProductInformationFilter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ProductInformationFilter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ProductInformationFilter"}

	if s.ProductInformationFilterComparator == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductInformationFilterComparator"))
	}

	if s.ProductInformationFilterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductInformationFilterName"))
	}

	if s.ProductInformationFilterValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductInformationFilterValue"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Details about a resource.
type ResourceInventory struct {
	_ struct{} `type:"structure"`

	// Platform of the resource.
	Platform *string `type:"string"`

	// Platform version of the resource in the inventory.
	PlatformVersion *string `type:"string"`

	// Amazon Resource Name (ARN) of the resource.
	ResourceArn *string `type:"string"`

	// ID of the resource.
	ResourceId *string `type:"string"`

	// ID of the account that owns the resource.
	ResourceOwningAccountId *string `type:"string"`

	// Type of resource.
	ResourceType ResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ResourceInventory) String() string {
	return awsutil.Prettify(s)
}

// Details about a tag for a license configuration.
type Tag struct {
	_ struct{} `type:"structure"`

	// Tag key.
	Key *string `type:"string"`

	// Tag value.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}
