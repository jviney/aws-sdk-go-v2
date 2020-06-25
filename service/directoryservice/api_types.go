// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Represents a named directory attribute.
type Attribute struct {
	_ struct{} `type:"structure"`

	// The name of the attribute.
	Name *string `min:"1" type:"string"`

	// The value of the attribute.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Attribute) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Attribute) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Attribute"}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about the certificate.
type Certificate struct {
	_ struct{} `type:"structure"`

	// The identifier of the certificate.
	CertificateId *string `type:"string"`

	// The common name for the certificate.
	CommonName *string `type:"string"`

	// The date and time when the certificate will expire.
	ExpiryDateTime *time.Time `type:"timestamp"`

	// The date and time that the certificate was registered.
	RegisteredDateTime *time.Time `type:"timestamp"`

	// The state of the certificate.
	State CertificateState `type:"string" enum:"true"`

	// Describes a state change for the certificate.
	StateReason *string `type:"string"`
}

// String returns the string representation
func (s Certificate) String() string {
	return awsutil.Prettify(s)
}

// Contains general information about a certificate.
type CertificateInfo struct {
	_ struct{} `type:"structure"`

	// The identifier of the certificate.
	CertificateId *string `type:"string"`

	// The common name for the certificate.
	CommonName *string `type:"string"`

	// The date and time when the certificate will expire.
	ExpiryDateTime *time.Time `type:"timestamp"`

	// The state of the certificate.
	State CertificateState `type:"string" enum:"true"`
}

// String returns the string representation
func (s CertificateInfo) String() string {
	return awsutil.Prettify(s)
}

// Contains information about a computer account in a directory.
type Computer struct {
	_ struct{} `type:"structure"`

	// An array of Attribute objects containing the LDAP attributes that belong
	// to the computer account.
	ComputerAttributes []Attribute `type:"list"`

	// The identifier of the computer.
	ComputerId *string `min:"1" type:"string"`

	// The computer name.
	ComputerName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Computer) String() string {
	return awsutil.Prettify(s)
}

// Points to a remote domain with which you are setting up a trust relationship.
// Conditional forwarders are required in order to set up a trust relationship
// with another domain.
type ConditionalForwarder struct {
	_ struct{} `type:"structure"`

	// The IP addresses of the remote DNS server associated with RemoteDomainName.
	// This is the IP address of the DNS server that your conditional forwarder
	// points to.
	DnsIpAddrs []string `type:"list"`

	// The fully qualified domain name (FQDN) of the remote domains pointed to by
	// the conditional forwarder.
	RemoteDomainName *string `type:"string"`

	// The replication scope of the conditional forwarder. The only allowed value
	// is Domain, which will replicate the conditional forwarder to all of the domain
	// controllers for your AWS directory.
	ReplicationScope ReplicationScope `type:"string" enum:"true"`
}

// String returns the string representation
func (s ConditionalForwarder) String() string {
	return awsutil.Prettify(s)
}

// Contains information for the ConnectDirectory operation when an AD Connector
// directory is being created.
type DirectoryConnectSettings struct {
	_ struct{} `type:"structure"`

	// A list of one or more IP addresses of DNS servers or domain controllers in
	// the on-premises directory.
	//
	// CustomerDnsIps is a required field
	CustomerDnsIps []string `type:"list" required:"true"`

	// The user name of an account in the on-premises directory that is used to
	// connect to the directory. This account must have the following permissions:
	//
	//    * Read users and groups
	//
	//    * Create computer objects
	//
	//    * Join computers to the domain
	//
	// CustomerUserName is a required field
	CustomerUserName *string `min:"1" type:"string" required:"true"`

	// A list of subnet identifiers in the VPC in which the AD Connector is created.
	//
	// SubnetIds is a required field
	SubnetIds []string `type:"list" required:"true"`

	// The identifier of the VPC in which the AD Connector is created.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DirectoryConnectSettings) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DirectoryConnectSettings) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DirectoryConnectSettings"}

	if s.CustomerDnsIps == nil {
		invalidParams.Add(aws.NewErrParamRequired("CustomerDnsIps"))
	}

	if s.CustomerUserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CustomerUserName"))
	}
	if s.CustomerUserName != nil && len(*s.CustomerUserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomerUserName", 1))
	}

	if s.SubnetIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetIds"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains information about an AD Connector directory.
type DirectoryConnectSettingsDescription struct {
	_ struct{} `type:"structure"`

	// A list of the Availability Zones that the directory is in.
	AvailabilityZones []string `type:"list"`

	// The IP addresses of the AD Connector servers.
	ConnectIps []string `type:"list"`

	// The user name of the service account in the on-premises directory.
	CustomerUserName *string `min:"1" type:"string"`

	// The security group identifier for the AD Connector directory.
	SecurityGroupId *string `type:"string"`

	// A list of subnet identifiers in the VPC that the AD Connector is in.
	SubnetIds []string `type:"list"`

	// The identifier of the VPC that the AD Connector is in.
	VpcId *string `type:"string"`
}

// String returns the string representation
func (s DirectoryConnectSettingsDescription) String() string {
	return awsutil.Prettify(s)
}

// Contains information about an AWS Directory Service directory.
type DirectoryDescription struct {
	_ struct{} `type:"structure"`

	// The access URL for the directory, such as http://<alias>.awsapps.com. If
	// no alias has been created for the directory, <alias> is the directory identifier,
	// such as d-XXXXXXXXXX.
	AccessUrl *string `min:"1" type:"string"`

	// The alias for the directory. If no alias has been created for the directory,
	// the alias is the directory identifier, such as d-XXXXXXXXXX.
	Alias *string `min:"1" type:"string"`

	// A DirectoryConnectSettingsDescription object that contains additional information
	// about an AD Connector directory. This member is only present if the directory
	// is an AD Connector directory.
	ConnectSettings *DirectoryConnectSettingsDescription `type:"structure"`

	// The description for the directory.
	Description *string `type:"string"`

	// The desired number of domain controllers in the directory if the directory
	// is Microsoft AD.
	DesiredNumberOfDomainControllers *int64 `min:"2" type:"integer"`

	// The directory identifier.
	DirectoryId *string `type:"string"`

	// The IP addresses of the DNS servers for the directory. For a Simple AD or
	// Microsoft AD directory, these are the IP addresses of the Simple AD or Microsoft
	// AD directory servers. For an AD Connector directory, these are the IP addresses
	// of the DNS servers or domain controllers in the on-premises directory to
	// which the AD Connector is connected.
	DnsIpAddrs []string `type:"list"`

	// The edition associated with this directory.
	Edition DirectoryEdition `type:"string" enum:"true"`

	// Specifies when the directory was created.
	LaunchTime *time.Time `type:"timestamp"`

	// The fully qualified name of the directory.
	Name *string `type:"string"`

	// Describes the AWS Managed Microsoft AD directory in the directory owner account.
	OwnerDirectoryDescription *OwnerDirectoryDescription `type:"structure"`

	// A RadiusSettings object that contains information about the RADIUS server
	// configured for this directory.
	RadiusSettings *RadiusSettings `type:"structure"`

	// The status of the RADIUS MFA server connection.
	RadiusStatus RadiusStatus `type:"string" enum:"true"`

	// The method used when sharing a directory to determine whether the directory
	// should be shared within your AWS organization (ORGANIZATIONS) or with any
	// AWS account by sending a shared directory request (HANDSHAKE).
	ShareMethod ShareMethod `type:"string" enum:"true"`

	// A directory share request that is sent by the directory owner to the directory
	// consumer. The request includes a typed message to help the directory consumer
	// administrator determine whether to approve or reject the share invitation.
	ShareNotes *string `type:"string" sensitive:"true"`

	// Current directory status of the shared AWS Managed Microsoft AD directory.
	ShareStatus ShareStatus `type:"string" enum:"true"`

	// The short name of the directory.
	ShortName *string `type:"string"`

	// The directory size.
	Size DirectorySize `type:"string" enum:"true"`

	// Indicates if single sign-on is enabled for the directory. For more information,
	// see EnableSso and DisableSso.
	SsoEnabled *bool `type:"boolean"`

	// The current stage of the directory.
	Stage DirectoryStage `type:"string" enum:"true"`

	// The date and time that the stage was last updated.
	StageLastUpdatedDateTime *time.Time `type:"timestamp"`

	// Additional information about the directory stage.
	StageReason *string `type:"string"`

	// The directory size.
	Type DirectoryType `type:"string" enum:"true"`

	// A DirectoryVpcSettingsDescription object that contains additional information
	// about a directory. This member is only present if the directory is a Simple
	// AD or Managed AD directory.
	VpcSettings *DirectoryVpcSettingsDescription `type:"structure"`
}

// String returns the string representation
func (s DirectoryDescription) String() string {
	return awsutil.Prettify(s)
}

// Contains directory limit information for a Region.
type DirectoryLimits struct {
	_ struct{} `type:"structure"`

	// The current number of cloud directories in the Region.
	CloudOnlyDirectoriesCurrentCount *int64 `type:"integer"`

	// The maximum number of cloud directories allowed in the Region.
	CloudOnlyDirectoriesLimit *int64 `type:"integer"`

	// Indicates if the cloud directory limit has been reached.
	CloudOnlyDirectoriesLimitReached *bool `type:"boolean"`

	// The current number of AWS Managed Microsoft AD directories in the region.
	CloudOnlyMicrosoftADCurrentCount *int64 `type:"integer"`

	// The maximum number of AWS Managed Microsoft AD directories allowed in the
	// region.
	CloudOnlyMicrosoftADLimit *int64 `type:"integer"`

	// Indicates if the AWS Managed Microsoft AD directory limit has been reached.
	CloudOnlyMicrosoftADLimitReached *bool `type:"boolean"`

	// The current number of connected directories in the Region.
	ConnectedDirectoriesCurrentCount *int64 `type:"integer"`

	// The maximum number of connected directories allowed in the Region.
	ConnectedDirectoriesLimit *int64 `type:"integer"`

	// Indicates if the connected directory limit has been reached.
	ConnectedDirectoriesLimitReached *bool `type:"boolean"`
}

// String returns the string representation
func (s DirectoryLimits) String() string {
	return awsutil.Prettify(s)
}

// Contains VPC information for the CreateDirectory or CreateMicrosoftAD operation.
type DirectoryVpcSettings struct {
	_ struct{} `type:"structure"`

	// The identifiers of the subnets for the directory servers. The two subnets
	// must be in different Availability Zones. AWS Directory Service creates a
	// directory server and a DNS server in each of these subnets.
	//
	// SubnetIds is a required field
	SubnetIds []string `type:"list" required:"true"`

	// The identifier of the VPC in which to create the directory.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DirectoryVpcSettings) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DirectoryVpcSettings) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DirectoryVpcSettings"}

	if s.SubnetIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetIds"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains information about the directory.
type DirectoryVpcSettingsDescription struct {
	_ struct{} `type:"structure"`

	// The list of Availability Zones that the directory is in.
	AvailabilityZones []string `type:"list"`

	// The domain controller security group identifier for the directory.
	SecurityGroupId *string `type:"string"`

	// The identifiers of the subnets for the directory servers.
	SubnetIds []string `type:"list"`

	// The identifier of the VPC that the directory is in.
	VpcId *string `type:"string"`
}

// String returns the string representation
func (s DirectoryVpcSettingsDescription) String() string {
	return awsutil.Prettify(s)
}

// Contains information about the domain controllers for a specified directory.
type DomainController struct {
	_ struct{} `type:"structure"`

	// The Availability Zone where the domain controller is located.
	AvailabilityZone *string `type:"string"`

	// Identifier of the directory where the domain controller resides.
	DirectoryId *string `type:"string"`

	// The IP address of the domain controller.
	DnsIpAddr *string `type:"string"`

	// Identifies a specific domain controller in the directory.
	DomainControllerId *string `type:"string"`

	// Specifies when the domain controller was created.
	LaunchTime *time.Time `type:"timestamp"`

	// The status of the domain controller.
	Status DomainControllerStatus `type:"string" enum:"true"`

	// The date and time that the status was last updated.
	StatusLastUpdatedDateTime *time.Time `type:"timestamp"`

	// A description of the domain controller state.
	StatusReason *string `type:"string"`

	// Identifier of the subnet in the VPC that contains the domain controller.
	SubnetId *string `type:"string"`

	// The identifier of the VPC that contains the domain controller.
	VpcId *string `type:"string"`
}

// String returns the string representation
func (s DomainController) String() string {
	return awsutil.Prettify(s)
}

// Information about SNS topic and AWS Directory Service directory associations.
type EventTopic struct {
	_ struct{} `type:"structure"`

	// The date and time of when you associated your directory with the SNS topic.
	CreatedDateTime *time.Time `type:"timestamp"`

	// The Directory ID of an AWS Directory Service directory that will publish
	// status messages to an SNS topic.
	DirectoryId *string `type:"string"`

	// The topic registration status.
	Status TopicStatus `type:"string" enum:"true"`

	// The SNS topic ARN (Amazon Resource Name).
	TopicArn *string `type:"string"`

	// The name of an AWS SNS topic the receives status messages from the directory.
	TopicName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s EventTopic) String() string {
	return awsutil.Prettify(s)
}

// IP address block. This is often the address block of the DNS server used
// for your on-premises domain.
type IpRoute struct {
	_ struct{} `type:"structure"`

	// IP address block using CIDR format, for example 10.0.0.0/24. This is often
	// the address block of the DNS server used for your on-premises domain. For
	// a single IP address use a CIDR address block with /32. For example 10.0.0.0/32.
	CidrIp *string `type:"string"`

	// Description of the address block.
	Description *string `type:"string"`
}

// String returns the string representation
func (s IpRoute) String() string {
	return awsutil.Prettify(s)
}

// Information about one or more IP address blocks.
type IpRouteInfo struct {
	_ struct{} `type:"structure"`

	// The date and time the address block was added to the directory.
	AddedDateTime *time.Time `type:"timestamp"`

	// IP address block in the IpRoute.
	CidrIp *string `type:"string"`

	// Description of the IpRouteInfo.
	Description *string `type:"string"`

	// Identifier (ID) of the directory associated with the IP addresses.
	DirectoryId *string `type:"string"`

	// The status of the IP address block.
	IpRouteStatusMsg IpRouteStatusMsg `type:"string" enum:"true"`

	// The reason for the IpRouteStatusMsg.
	IpRouteStatusReason *string `type:"string"`
}

// String returns the string representation
func (s IpRouteInfo) String() string {
	return awsutil.Prettify(s)
}

// Contains general information about the LDAPS settings.
type LDAPSSettingInfo struct {
	_ struct{} `type:"structure"`

	// The state of the LDAPS settings.
	LDAPSStatus LDAPSStatus `type:"string" enum:"true"`

	// Describes a state change for LDAPS.
	LDAPSStatusReason *string `type:"string"`

	// The date and time when the LDAPS settings were last updated.
	LastUpdatedDateTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s LDAPSSettingInfo) String() string {
	return awsutil.Prettify(s)
}

// Represents a log subscription, which tracks real-time data from a chosen
// log group to a specified destination.
type LogSubscription struct {
	_ struct{} `type:"structure"`

	// Identifier (ID) of the directory that you want to associate with the log
	// subscription.
	DirectoryId *string `type:"string"`

	// The name of the log group.
	LogGroupName *string `min:"1" type:"string"`

	// The date and time that the log subscription was created.
	SubscriptionCreatedDateTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s LogSubscription) String() string {
	return awsutil.Prettify(s)
}

// Describes the directory owner account details that have been shared to the
// directory consumer account.
type OwnerDirectoryDescription struct {
	_ struct{} `type:"structure"`

	// Identifier of the directory owner account.
	AccountId *string `type:"string"`

	// Identifier of the AWS Managed Microsoft AD directory in the directory owner
	// account.
	DirectoryId *string `type:"string"`

	// IP address of the directory’s domain controllers.
	DnsIpAddrs []string `type:"list"`

	// A RadiusSettings object that contains information about the RADIUS server.
	RadiusSettings *RadiusSettings `type:"structure"`

	// Information about the status of the RADIUS server.
	RadiusStatus RadiusStatus `type:"string" enum:"true"`

	// Information about the VPC settings for the directory.
	VpcSettings *DirectoryVpcSettingsDescription `type:"structure"`
}

// String returns the string representation
func (s OwnerDirectoryDescription) String() string {
	return awsutil.Prettify(s)
}

// Contains information about a Remote Authentication Dial In User Service (RADIUS)
// server.
type RadiusSettings struct {
	_ struct{} `type:"structure"`

	// The protocol specified for your RADIUS endpoints.
	AuthenticationProtocol RadiusAuthenticationProtocol `type:"string" enum:"true"`

	// Not currently used.
	DisplayLabel *string `min:"1" type:"string"`

	// The port that your RADIUS server is using for communications. Your on-premises
	// network must allow inbound traffic over this port from the AWS Directory
	// Service servers.
	RadiusPort *int64 `min:"1025" type:"integer"`

	// The maximum number of times that communication with the RADIUS server is
	// attempted.
	RadiusRetries *int64 `type:"integer"`

	// An array of strings that contains the IP addresses of the RADIUS server endpoints,
	// or the IP addresses of your RADIUS server load balancer.
	RadiusServers []string `type:"list"`

	// The amount of time, in seconds, to wait for the RADIUS server to respond.
	RadiusTimeout *int64 `min:"1" type:"integer"`

	// Required for enabling RADIUS on the directory.
	SharedSecret *string `min:"8" type:"string" sensitive:"true"`

	// Not currently used.
	UseSameUsername *bool `type:"boolean"`
}

// String returns the string representation
func (s RadiusSettings) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RadiusSettings) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RadiusSettings"}
	if s.DisplayLabel != nil && len(*s.DisplayLabel) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DisplayLabel", 1))
	}
	if s.RadiusPort != nil && *s.RadiusPort < 1025 {
		invalidParams.Add(aws.NewErrParamMinValue("RadiusPort", 1025))
	}
	if s.RadiusTimeout != nil && *s.RadiusTimeout < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("RadiusTimeout", 1))
	}
	if s.SharedSecret != nil && len(*s.SharedSecret) < 8 {
		invalidParams.Add(aws.NewErrParamMinLen("SharedSecret", 8))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a schema extension.
type SchemaExtensionInfo struct {
	_ struct{} `type:"structure"`

	// A description of the schema extension.
	Description *string `type:"string"`

	// The identifier of the directory to which the schema extension is applied.
	DirectoryId *string `type:"string"`

	// The date and time that the schema extension was completed.
	EndDateTime *time.Time `type:"timestamp"`

	// The identifier of the schema extension.
	SchemaExtensionId *string `type:"string"`

	// The current status of the schema extension.
	SchemaExtensionStatus SchemaExtensionStatus `type:"string" enum:"true"`

	// The reason for the SchemaExtensionStatus.
	SchemaExtensionStatusReason *string `type:"string"`

	// The date and time that the schema extension started being applied to the
	// directory.
	StartDateTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s SchemaExtensionInfo) String() string {
	return awsutil.Prettify(s)
}

// Identifier that contains details about the directory consumer account.
type ShareTarget struct {
	_ struct{} `type:"structure"`

	// Identifier of the directory consumer account.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// Type of identifier to be used in the Id field.
	//
	// Type is a required field
	Type TargetType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ShareTarget) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ShareTarget) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ShareTarget"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Details about the shared directory in the directory owner account for which
// the share request in the directory consumer account has been accepted.
type SharedDirectory struct {
	_ struct{} `type:"structure"`

	// The date and time that the shared directory was created.
	CreatedDateTime *time.Time `type:"timestamp"`

	// The date and time that the shared directory was last updated.
	LastUpdatedDateTime *time.Time `type:"timestamp"`

	// Identifier of the directory owner account, which contains the directory that
	// has been shared to the consumer account.
	OwnerAccountId *string `type:"string"`

	// Identifier of the directory in the directory owner account.
	OwnerDirectoryId *string `type:"string"`

	// The method used when sharing a directory to determine whether the directory
	// should be shared within your AWS organization (ORGANIZATIONS) or with any
	// AWS account by sending a shared directory request (HANDSHAKE).
	ShareMethod ShareMethod `type:"string" enum:"true"`

	// A directory share request that is sent by the directory owner to the directory
	// consumer. The request includes a typed message to help the directory consumer
	// administrator determine whether to approve or reject the share invitation.
	ShareNotes *string `type:"string" sensitive:"true"`

	// Current directory status of the shared AWS Managed Microsoft AD directory.
	ShareStatus ShareStatus `type:"string" enum:"true"`

	// Identifier of the directory consumer account that has access to the shared
	// directory (OwnerDirectoryId) in the directory owner account.
	SharedAccountId *string `type:"string"`

	// Identifier of the shared directory in the directory consumer account. This
	// identifier is different for each directory owner account.
	SharedDirectoryId *string `type:"string"`
}

// String returns the string representation
func (s SharedDirectory) String() string {
	return awsutil.Prettify(s)
}

// Describes a directory snapshot.
type Snapshot struct {
	_ struct{} `type:"structure"`

	// The directory identifier.
	DirectoryId *string `type:"string"`

	// The descriptive name of the snapshot.
	Name *string `type:"string"`

	// The snapshot identifier.
	SnapshotId *string `type:"string"`

	// The date and time that the snapshot was taken.
	StartTime *time.Time `type:"timestamp"`

	// The snapshot status.
	Status SnapshotStatus `type:"string" enum:"true"`

	// The snapshot type.
	Type SnapshotType `type:"string" enum:"true"`
}

// String returns the string representation
func (s Snapshot) String() string {
	return awsutil.Prettify(s)
}

// Contains manual snapshot limit information for a directory.
type SnapshotLimits struct {
	_ struct{} `type:"structure"`

	// The current number of manual snapshots of the directory.
	ManualSnapshotsCurrentCount *int64 `type:"integer"`

	// The maximum number of manual snapshots allowed.
	ManualSnapshotsLimit *int64 `type:"integer"`

	// Indicates if the manual snapshot limit has been reached.
	ManualSnapshotsLimitReached *bool `type:"boolean"`
}

// String returns the string representation
func (s SnapshotLimits) String() string {
	return awsutil.Prettify(s)
}

// Metadata assigned to a directory consisting of a key-value pair.
type Tag struct {
	_ struct{} `type:"structure"`

	// Required name of the tag. The string value can be Unicode characters and
	// cannot be prefixed with "aws:". The string can contain only the set of Unicode
	// letters, digits, white-space, '_', '.', '/', '=', '+', '-' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// The optional value of the tag. The string value can be Unicode characters.
	// The string can contain only the set of Unicode letters, digits, white-space,
	// '_', '.', '/', '=', '+', '-' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes a trust relationship between an AWS Managed Microsoft AD directory
// and an external domain.
type Trust struct {
	_ struct{} `type:"structure"`

	// The date and time that the trust relationship was created.
	CreatedDateTime *time.Time `type:"timestamp"`

	// The Directory ID of the AWS directory involved in the trust relationship.
	DirectoryId *string `type:"string"`

	// The date and time that the trust relationship was last updated.
	LastUpdatedDateTime *time.Time `type:"timestamp"`

	// The Fully Qualified Domain Name (FQDN) of the external domain involved in
	// the trust relationship.
	RemoteDomainName *string `type:"string"`

	// Current state of selective authentication for the trust.
	SelectiveAuth SelectiveAuth `type:"string" enum:"true"`

	// The date and time that the TrustState was last updated.
	StateLastUpdatedDateTime *time.Time `type:"timestamp"`

	// The trust relationship direction.
	TrustDirection TrustDirection `type:"string" enum:"true"`

	// The unique ID of the trust relationship.
	TrustId *string `type:"string"`

	// The trust relationship state.
	TrustState TrustState `type:"string" enum:"true"`

	// The reason for the TrustState.
	TrustStateReason *string `type:"string"`

	// The trust relationship type. Forest is the default.
	TrustType TrustType `type:"string" enum:"true"`
}

// String returns the string representation
func (s Trust) String() string {
	return awsutil.Prettify(s)
}

// Identifier that contains details about the directory consumer account with
// whom the directory is being unshared.
type UnshareTarget struct {
	_ struct{} `type:"structure"`

	// Identifier of the directory consumer account.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// Type of identifier to be used in the Id field.
	//
	// Type is a required field
	Type TargetType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s UnshareTarget) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UnshareTarget) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UnshareTarget"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
