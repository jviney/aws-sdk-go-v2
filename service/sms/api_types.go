// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"time"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Information about the application.
type AppSummary struct {
	_ struct{} `type:"structure"`

	// Unique ID of the application.
	AppId *string `locationName:"appId" type:"string"`

	// Time of creation of this application.
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp"`

	// Description of the application.
	Description *string `locationName:"description" type:"string"`

	// Timestamp of the application's creation.
	LastModified *time.Time `locationName:"lastModified" type:"timestamp"`

	// Timestamp of the application's most recent successful replication.
	LatestReplicationTime *time.Time `locationName:"latestReplicationTime" type:"timestamp"`

	// Details about the latest launch of the application.
	LaunchDetails *LaunchDetails `locationName:"launchDetails" type:"structure"`

	// Launch status of the application.
	LaunchStatus AppLaunchStatus `locationName:"launchStatus" type:"string" enum:"true"`

	// A message related to the launch status of the application.
	LaunchStatusMessage *string `locationName:"launchStatusMessage" type:"string"`

	// Name of the application.
	Name *string `locationName:"name" type:"string"`

	// Replication status of the application.
	ReplicationStatus AppReplicationStatus `locationName:"replicationStatus" type:"string" enum:"true"`

	// A message related to the replication status of the application.
	ReplicationStatusMessage *string `locationName:"replicationStatusMessage" type:"string"`

	// Name of the service role in the customer's account used by AWS SMS.
	RoleName *string `locationName:"roleName" type:"string"`

	// Status of the application.
	Status AppStatus `locationName:"status" type:"string" enum:"true"`

	// A message related to the status of the application
	StatusMessage *string `locationName:"statusMessage" type:"string"`

	// Number of server groups present in the application.
	TotalServerGroups *int64 `locationName:"totalServerGroups" type:"integer"`

	// Number of servers present in the application.
	TotalServers *int64 `locationName:"totalServers" type:"integer"`
}

// String returns the string representation
func (s AppSummary) String() string {
	return awsutil.Prettify(s)
}

// Represents a connector.
type Connector struct {
	_ struct{} `type:"structure"`

	// The time the connector was associated.
	AssociatedOn *time.Time `locationName:"associatedOn" type:"timestamp"`

	// The capabilities of the connector.
	CapabilityList []ConnectorCapability `locationName:"capabilityList" type:"list"`

	// The identifier of the connector.
	ConnectorId *string `locationName:"connectorId" type:"string"`

	// The IP address of the connector.
	IpAddress *string `locationName:"ipAddress" type:"string"`

	// The MAC address of the connector.
	MacAddress *string `locationName:"macAddress" type:"string"`

	// The status of the connector.
	Status ConnectorStatus `locationName:"status" type:"string" enum:"true"`

	// The connector version.
	Version *string `locationName:"version" type:"string"`

	// The identifier of the VM manager.
	VmManagerId *string `locationName:"vmManagerId" type:"string"`

	// The name of the VM manager.
	VmManagerName *string `locationName:"vmManagerName" type:"string"`

	// The VM management product.
	VmManagerType VmManagerType `locationName:"vmManagerType" type:"string" enum:"true"`
}

// String returns the string representation
func (s Connector) String() string {
	return awsutil.Prettify(s)
}

// Details about the latest launch of an application.
type LaunchDetails struct {
	_ struct{} `type:"structure"`

	// Latest time this application was launched successfully.
	LatestLaunchTime *time.Time `locationName:"latestLaunchTime" type:"timestamp"`

	// Identifier of the latest stack launched for this application.
	StackId *string `locationName:"stackId" type:"string"`

	// Name of the latest stack launched for this application.
	StackName *string `locationName:"stackName" type:"string"`
}

// String returns the string representation
func (s LaunchDetails) String() string {
	return awsutil.Prettify(s)
}

// Represents a replication job.
type ReplicationJob struct {
	_ struct{} `type:"structure"`

	// The description of the replication job.
	Description *string `locationName:"description" type:"string"`

	// Whether the replication job should produce encrypted AMIs or not. See also
	// KmsKeyId below.
	Encrypted *bool `locationName:"encrypted" type:"boolean"`

	// The time between consecutive replication runs, in hours.
	Frequency *int64 `locationName:"frequency" type:"integer"`

	// KMS key ID for replication jobs that produce encrypted AMIs. Can be any of
	// the following:
	//
	//    * KMS key ID
	//
	//    * KMS key alias
	//
	//    * ARN referring to KMS key ID
	//
	//    * ARN referring to KMS key alias
	//
	// If encrypted is true but a KMS key id is not specified, the customer's default
	// KMS key for EBS is used.
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`

	// The ID of the latest Amazon Machine Image (AMI).
	LatestAmiId *string `locationName:"latestAmiId" type:"string"`

	// The license type to be used for the AMI created by a successful replication
	// run.
	LicenseType LicenseType `locationName:"licenseType" type:"string" enum:"true"`

	// The start time of the next replication run.
	NextReplicationRunStartTime *time.Time `locationName:"nextReplicationRunStartTime" type:"timestamp"`

	// Number of recent AMIs to keep in the customer's account for a replication
	// job. By default the value is set to zero, meaning that all AMIs are kept.
	NumberOfRecentAmisToKeep *int64 `locationName:"numberOfRecentAmisToKeep" type:"integer"`

	// The identifier of the replication job.
	ReplicationJobId *string `locationName:"replicationJobId" type:"string"`

	// Information about the replication runs.
	ReplicationRunList []ReplicationRun `locationName:"replicationRunList" type:"list"`

	// The name of the IAM role to be used by the Server Migration Service.
	RoleName *string `locationName:"roleName" type:"string"`

	RunOnce *bool `locationName:"runOnce" type:"boolean"`

	// The seed replication time.
	SeedReplicationTime *time.Time `locationName:"seedReplicationTime" type:"timestamp"`

	// The identifier of the server.
	ServerId *string `locationName:"serverId" type:"string"`

	// The type of server.
	ServerType ServerType `locationName:"serverType" type:"string" enum:"true"`

	// The state of the replication job.
	State ReplicationJobState `locationName:"state" type:"string" enum:"true"`

	// The description of the current status of the replication job.
	StatusMessage *string `locationName:"statusMessage" type:"string"`

	// Information about the VM server.
	VmServer *VmServer `locationName:"vmServer" type:"structure"`
}

// String returns the string representation
func (s ReplicationJob) String() string {
	return awsutil.Prettify(s)
}

// Represents a replication run.
type ReplicationRun struct {
	_ struct{} `type:"structure"`

	// The identifier of the Amazon Machine Image (AMI) from the replication run.
	AmiId *string `locationName:"amiId" type:"string"`

	// The completion time of the last replication run.
	CompletedTime *time.Time `locationName:"completedTime" type:"timestamp"`

	// The description of the replication run.
	Description *string `locationName:"description" type:"string"`

	// Whether the replication run should produce encrypted AMI or not. See also
	// KmsKeyId below.
	Encrypted *bool `locationName:"encrypted" type:"boolean"`

	// KMS key ID for replication jobs that produce encrypted AMIs. Can be any of
	// the following:
	//
	//    * KMS key ID
	//
	//    * KMS key alias
	//
	//    * ARN referring to KMS key ID
	//
	//    * ARN referring to KMS key alias
	//
	// If encrypted is true but a KMS key id is not specified, the customer's default
	// KMS key for EBS is used.
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`

	// The identifier of the replication run.
	ReplicationRunId *string `locationName:"replicationRunId" type:"string"`

	// The start time of the next replication run.
	ScheduledStartTime *time.Time `locationName:"scheduledStartTime" type:"timestamp"`

	// Details of the current stage of the replication run.
	StageDetails *ReplicationRunStageDetails `locationName:"stageDetails" type:"structure"`

	// The state of the replication run.
	State ReplicationRunState `locationName:"state" type:"string" enum:"true"`

	// The description of the current status of the replication job.
	StatusMessage *string `locationName:"statusMessage" type:"string"`

	// The type of replication run.
	Type ReplicationRunType `locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s ReplicationRun) String() string {
	return awsutil.Prettify(s)
}

// Details of the current stage of a replication run.
type ReplicationRunStageDetails struct {
	_ struct{} `type:"structure"`

	// String describing the current stage of a replication run.
	Stage *string `locationName:"stage" type:"string"`

	// String describing the progress of the current stage of a replication run.
	StageProgress *string `locationName:"stageProgress" type:"string"`
}

// String returns the string representation
func (s ReplicationRunStageDetails) String() string {
	return awsutil.Prettify(s)
}

// Location of the Amazon S3 object in the customer's account.
type S3Location struct {
	_ struct{} `type:"structure"`

	// Amazon S3 bucket name.
	Bucket *string `locationName:"bucket" type:"string"`

	// Amazon S3 bucket key.
	Key *string `locationName:"key" type:"string"`
}

// String returns the string representation
func (s S3Location) String() string {
	return awsutil.Prettify(s)
}

// Represents a server.
type Server struct {
	_ struct{} `type:"structure"`

	// The identifier of the replication job.
	ReplicationJobId *string `locationName:"replicationJobId" type:"string"`

	// Indicates whether the replication job is deleted or failed.
	ReplicationJobTerminated *bool `locationName:"replicationJobTerminated" type:"boolean"`

	// The identifier of the server.
	ServerId *string `locationName:"serverId" type:"string"`

	// The type of server.
	ServerType ServerType `locationName:"serverType" type:"string" enum:"true"`

	// Information about the VM server.
	VmServer *VmServer `locationName:"vmServer" type:"structure"`
}

// String returns the string representation
func (s Server) String() string {
	return awsutil.Prettify(s)
}

// A logical grouping of servers.
type ServerGroup struct {
	_ struct{} `type:"structure"`

	// Name of a server group.
	Name *string `locationName:"name" type:"string"`

	// Identifier of a server group.
	ServerGroupId *string `locationName:"serverGroupId" type:"string"`

	// List of servers belonging to a server group.
	ServerList []Server `locationName:"serverList" type:"list"`
}

// String returns the string representation
func (s ServerGroup) String() string {
	return awsutil.Prettify(s)
}

// Launch configuration for a server group.
type ServerGroupLaunchConfiguration struct {
	_ struct{} `type:"structure"`

	// Launch order of servers in the server group.
	LaunchOrder *int64 `locationName:"launchOrder" type:"integer"`

	// Identifier of the server group the launch configuration is associated with.
	ServerGroupId *string `locationName:"serverGroupId" type:"string"`

	// Launch configuration for servers in the server group.
	ServerLaunchConfigurations []ServerLaunchConfiguration `locationName:"serverLaunchConfigurations" type:"list"`
}

// String returns the string representation
func (s ServerGroupLaunchConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Replication configuration for a server group.
type ServerGroupReplicationConfiguration struct {
	_ struct{} `type:"structure"`

	// Identifier of the server group this replication configuration is associated
	// with.
	ServerGroupId *string `locationName:"serverGroupId" type:"string"`

	// Replication configuration for servers in the server group.
	ServerReplicationConfigurations []ServerReplicationConfiguration `locationName:"serverReplicationConfigurations" type:"list"`
}

// String returns the string representation
func (s ServerGroupReplicationConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Launch configuration for a server.
type ServerLaunchConfiguration struct {
	_ struct{} `type:"structure"`

	// If true, a publicly accessible IP address is created when launching the server.
	AssociatePublicIpAddress *bool `locationName:"associatePublicIpAddress" type:"boolean"`

	// Name of the EC2 SSH Key to be used for connecting to the launched server.
	Ec2KeyName *string `locationName:"ec2KeyName" type:"string"`

	// Instance type to be used for launching the server.
	InstanceType *string `locationName:"instanceType" type:"string"`

	// Logical ID of the server in the Amazon CloudFormation template.
	LogicalId *string `locationName:"logicalId" type:"string"`

	// Identifier of the security group that applies to the launched server.
	SecurityGroup *string `locationName:"securityGroup" type:"string"`

	// Identifier of the server the launch configuration is associated with.
	Server *Server `locationName:"server" type:"structure"`

	// Identifier of the subnet the server should be launched into.
	Subnet *string `locationName:"subnet" type:"string"`

	// Location of the user-data script to be executed when launching the server.
	UserData *UserData `locationName:"userData" type:"structure"`

	// Identifier of the VPC the server should be launched into.
	Vpc *string `locationName:"vpc" type:"string"`
}

// String returns the string representation
func (s ServerLaunchConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Replication configuration of a server.
type ServerReplicationConfiguration struct {
	_ struct{} `type:"structure"`

	// Identifier of the server this replication configuration is associated with.
	Server *Server `locationName:"server" type:"structure"`

	// Parameters for replicating the server.
	ServerReplicationParameters *ServerReplicationParameters `locationName:"serverReplicationParameters" type:"structure"`
}

// String returns the string representation
func (s ServerReplicationConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Replication parameters for replicating a server.
type ServerReplicationParameters struct {
	_ struct{} `type:"structure"`

	// When true, the replication job produces encrypted AMIs. See also KmsKeyId
	// below.
	Encrypted *bool `locationName:"encrypted" type:"boolean"`

	// Frequency of creating replication jobs for the server.
	Frequency *int64 `locationName:"frequency" type:"integer"`

	// KMS key ID for replication jobs that produce encrypted AMIs. Can be any of
	// the following:
	//
	//    * KMS key ID
	//
	//    * KMS key alias
	//
	//    * ARN referring to KMS key ID
	//
	//    * ARN referring to KMS key alias
	//
	// If encrypted is true but a KMS key id is not specified, the customer's default
	// KMS key for EBS is used.
	KmsKeyId *string `locationName:"kmsKeyId" type:"string"`

	// License type for creating a replication job for the server.
	LicenseType LicenseType `locationName:"licenseType" type:"string" enum:"true"`

	// Number of recent AMIs to keep when creating a replication job for this server.
	NumberOfRecentAmisToKeep *int64 `locationName:"numberOfRecentAmisToKeep" type:"integer"`

	RunOnce *bool `locationName:"runOnce" type:"boolean"`

	// Seed time for creating a replication job for the server.
	SeedTime *time.Time `locationName:"seedTime" type:"timestamp"`
}

// String returns the string representation
func (s ServerReplicationParameters) String() string {
	return awsutil.Prettify(s)
}

// A label that can be assigned to an application.
type Tag struct {
	_ struct{} `type:"structure"`

	// Tag key.
	Key *string `locationName:"key" type:"string"`

	// Tag value.
	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// A script that runs on first launch of an Amazon EC2 instance. Used for configuring
// the server during launch.
type UserData struct {
	_ struct{} `type:"structure"`

	// Amazon S3 location of the user-data script.
	S3Location *S3Location `locationName:"s3Location" type:"structure"`
}

// String returns the string representation
func (s UserData) String() string {
	return awsutil.Prettify(s)
}

// Represents a VM server.
type VmServer struct {
	_ struct{} `type:"structure"`

	// The name of the VM manager.
	VmManagerName *string `locationName:"vmManagerName" type:"string"`

	// The type of VM management product.
	VmManagerType VmManagerType `locationName:"vmManagerType" type:"string" enum:"true"`

	// The name of the VM.
	VmName *string `locationName:"vmName" type:"string"`

	// The VM folder path in the vCenter Server virtual machine inventory tree.
	VmPath *string `locationName:"vmPath" type:"string"`

	// Information about the VM server location.
	VmServerAddress *VmServerAddress `locationName:"vmServerAddress" type:"structure"`
}

// String returns the string representation
func (s VmServer) String() string {
	return awsutil.Prettify(s)
}

// Represents a VM server location.
type VmServerAddress struct {
	_ struct{} `type:"structure"`

	// The identifier of the VM.
	VmId *string `locationName:"vmId" type:"string"`

	// The identifier of the VM manager.
	VmManagerId *string `locationName:"vmManagerId" type:"string"`
}

// String returns the string representation
func (s VmServerAddress) String() string {
	return awsutil.Prettify(s)
}
