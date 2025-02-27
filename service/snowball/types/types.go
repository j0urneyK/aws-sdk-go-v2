// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The address that you want the Snow device(s) associated with a specific job to
// be shipped to. Addresses are validated at the time of creation. The address you
// provide must be located within the serviceable area of your region. Although no
// individual elements of the Address are required, if the address is invalid or
// unsupported, then an exception is thrown.
type Address struct {

	// The unique ID for an address.
	AddressId *string

	// The city in an address that a Snow device is to be delivered to.
	City *string

	// The name of the company to receive a Snow device at an address.
	Company *string

	// The country in an address that a Snow device is to be delivered to.
	Country *string

	// If the address you are creating is a primary address, then set this option to
	// true. This field is not supported in most regions.
	IsRestricted bool

	// This field is no longer used and the value is ignored.
	Landmark *string

	// The name of a person to receive a Snow device at an address.
	Name *string

	// The phone number associated with an address that a Snow device is to be
	// delivered to.
	PhoneNumber *string

	// The postal code in an address that a Snow device is to be delivered to.
	PostalCode *string

	// This field is no longer used and the value is ignored.
	PrefectureOrDistrict *string

	// The state or province in an address that a Snow device is to be delivered to.
	StateOrProvince *string

	// The first line in a street address that a Snow device is to be delivered to.
	Street1 *string

	// The second line in a street address that a Snow device is to be delivered to.
	Street2 *string

	// The third line in a street address that a Snow device is to be delivered to.
	Street3 *string

	noSmithyDocumentSerde
}

// Contains a cluster's state, a cluster's ID, and other important information.
type ClusterListEntry struct {

	// The 39-character ID for the cluster that you want to list, for example
	// CID123e4567-e89b-12d3-a456-426655440000.
	ClusterId *string

	// The current state of this cluster. For information about the state of a specific
	// node, see JobListEntry$JobState.
	ClusterState ClusterState

	// The creation date for this cluster.
	CreationDate *time.Time

	// Defines an optional description of the cluster, for example Environmental Data
	// Cluster-01.
	Description *string

	noSmithyDocumentSerde
}

// Contains metadata about a specific cluster.
type ClusterMetadata struct {

	// The automatically generated ID for a specific address.
	AddressId *string

	// The automatically generated ID for a cluster.
	ClusterId *string

	// The current status of the cluster.
	ClusterState ClusterState

	// The creation date for this cluster.
	CreationDate *time.Time

	// The optional description of the cluster.
	Description *string

	// The ID of the address that you want a cluster shipped to, after it will be
	// shipped to its primary address. This field is not supported in most regions.
	ForwardingAddressId *string

	// The type of job for this cluster. Currently, the only job type supported for
	// clusters is LOCAL_USE.
	JobType JobType

	// The KmsKeyARN Amazon Resource Name (ARN) associated with this cluster. This ARN
	// was created using the CreateKey
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateKey.html) API
	// action in Key Management Service (KMS.
	KmsKeyARN *string

	// The Amazon Simple Notification Service (Amazon SNS) notification settings for
	// this cluster.
	Notification *Notification

	// Represents metadata and configuration settings for services on an Amazon Web
	// Services Snow Family device.
	OnDeviceServiceConfiguration *OnDeviceServiceConfiguration

	// The arrays of JobResource objects that can include updated S3Resource objects or
	// LambdaResource objects.
	Resources *JobResource

	// The role ARN associated with this cluster. This ARN was created using the
	// CreateRole
	// (https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) API
	// action in Identity and Access Management (IAM).
	RoleARN *string

	// The shipping speed for each node in this cluster. This speed doesn't dictate how
	// soon you'll get each device, rather it represents how quickly each device moves
	// to its destination while in transit. Regional shipping speeds are as follows:
	//
	// *
	// In Australia, you have access to express shipping. Typically, devices shipped
	// express are delivered in about a day.
	//
	// * In the European Union (EU), you have
	// access to express shipping. Typically, Snow devices shipped express are
	// delivered in about a day. In addition, most countries in the EU have access to
	// standard shipping, which typically takes less than a week, one way.
	//
	// * In India,
	// Snow devices are delivered in one to seven days.
	//
	// * In the US, you have access
	// to one-day shipping and two-day shipping.
	ShippingOption ShippingOption

	// The type of Snowcone device to use for this cluster. For cluster jobs, Amazon
	// Web Services Snow Family currently supports only the EDGE device type.
	SnowballType SnowballType

	// The tax documents required in your Amazon Web Services Region.
	TaxDocuments *TaxDocuments

	noSmithyDocumentSerde
}

// A JSON-formatted object that describes a compatible Amazon Machine Image (AMI),
// including the ID and name for a Snow device AMI. This AMI is compatible with the
// device's physical hardware requirements, and it should be able to be run in an
// SBE1 instance on the device.
type CompatibleImage struct {

	// The unique identifier for an individual Snow device AMI.
	AmiId *string

	// The optional name of a compatible image.
	Name *string

	noSmithyDocumentSerde
}

// Defines the real-time status of a Snow device's data transfer while the device
// is at Amazon Web Services. This data is only available while a job has a
// JobState value of InProgress, for both import and export jobs.
type DataTransfer struct {

	// The number of bytes transferred between a Snow device and Amazon S3.
	BytesTransferred int64

	// The number of objects transferred between a Snow device and Amazon S3.
	ObjectsTransferred int64

	// The total bytes of data for a transfer between a Snow device and Amazon S3. This
	// value is set to 0 (zero) until all the keys that will be transferred have been
	// listed.
	TotalBytes int64

	// The total number of objects for a transfer between a Snow device and Amazon S3.
	// This value is set to 0 (zero) until all the keys that will be transferred have
	// been listed.
	TotalObjects int64

	noSmithyDocumentSerde
}

// The name and version of the service dependant on the requested service.
type DependentService struct {

	// The name of the dependent service.
	ServiceName ServiceName

	// The version of the dependent service.
	ServiceVersion *ServiceVersion

	noSmithyDocumentSerde
}

// The container for SnowconeDeviceConfiguration.
type DeviceConfiguration struct {

	// Returns information about the device configuration for an Snowcone job.
	SnowconeDeviceConfiguration *SnowconeDeviceConfiguration

	noSmithyDocumentSerde
}

// A JSON-formatted object that contains the IDs for an Amazon Machine Image (AMI),
// including the Amazon EC2 AMI ID and the Snow device AMI ID. Each AMI has these
// two IDs to simplify identifying the AMI in both the Amazon Web Services Cloud
// and on the device.
type Ec2AmiResource struct {

	// The ID of the AMI in Amazon EC2.
	//
	// This member is required.
	AmiId *string

	// The ID of the AMI on the Snow device.
	SnowballAmiId *string

	noSmithyDocumentSerde
}

// An object representing the metadata and configuration settings of EKS Anywhere
// on the Snow Family device.
type EKSOnDeviceServiceConfiguration struct {

	// The version of EKS Anywhere on the Snow Family device.
	EKSAnywhereVersion *string

	// The Kubernetes version for EKS Anywhere on the Snow Family device.
	KubernetesVersion *string

	noSmithyDocumentSerde
}

// The container for the EventTriggerDefinition$EventResourceARN.
type EventTriggerDefinition struct {

	// The Amazon Resource Name (ARN) for any local Amazon S3 resource that is an
	// Lambda function's event trigger associated with this job.
	EventResourceARN *string

	noSmithyDocumentSerde
}

// The tax documents required in Amazon Web Services Region in India.
type INDTaxDocuments struct {

	// The Goods and Services Tax (GST) documents required in Amazon Web Services
	// Region in India.
	GSTIN *string

	noSmithyDocumentSerde
}

// Each JobListEntry object contains a job's state, a job's ID, and a value that
// indicates whether the job is a job part, in the case of an export job.
type JobListEntry struct {

	// The creation date for this job.
	CreationDate *time.Time

	// The optional description of this specific job, for example Important Photos
	// 2016-08-11.
	Description *string

	// A value that indicates that this job is a main job. A main job represents a
	// successful request to create an export job. Main jobs aren't associated with any
	// Snowballs. Instead, each main job will have at least one job part, and each job
	// part is associated with a Snowball. It might take some time before the job parts
	// associated with a particular main job are listed, because they are created after
	// the main job is created.
	IsMaster bool

	// The automatically generated ID for a job, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	JobId *string

	// The current state of this job.
	JobState JobState

	// The type of job.
	JobType JobType

	// The type of device used with this job.
	SnowballType SnowballType

	noSmithyDocumentSerde
}

// Contains job logs. Whenever a Snow device is used to import data into or export
// data out of Amazon S3, you'll have the option of downloading a PDF job report.
// Job logs are returned as a part of the response syntax of the DescribeJob action
// in the JobMetadata data type. The job logs can be accessed for up to 60 minutes
// after this request has been made. To access any of the job logs after 60 minutes
// have passed, you'll have to make another call to the DescribeJob action. For
// import jobs, the PDF job report becomes available at the end of the import
// process. For export jobs, your job report typically becomes available while the
// Snow device for your job part is being delivered to you. The job report provides
// you insight into the state of your Amazon S3 data transfer. The report includes
// details about your job or job part for your records. For deeper visibility into
// the status of your transferred objects, you can look at the two associated logs:
// a success log and a failure log. The logs are saved in comma-separated value
// (CSV) format, and the name of each log includes the ID of the job or job part
// that the log describes.
type JobLogs struct {

	// A link to an Amazon S3 presigned URL where the job completion report is located.
	JobCompletionReportURI *string

	// A link to an Amazon S3 presigned URL where the job failure log is located.
	JobFailureLogURI *string

	// A link to an Amazon S3 presigned URL where the job success log is located.
	JobSuccessLogURI *string

	noSmithyDocumentSerde
}

// Contains information about a specific job including shipping information, job
// status, and other important metadata. This information is returned as a part of
// the response syntax of the DescribeJob action.
type JobMetadata struct {

	// The ID for the address that you want the Snow device shipped to.
	AddressId *string

	// The 39-character ID for the cluster, for example
	// CID123e4567-e89b-12d3-a456-426655440000.
	ClusterId *string

	// The creation date for this job.
	CreationDate *time.Time

	// A value that defines the real-time status of a Snow device's data transfer while
	// the device is at Amazon Web Services. This data is only available while a job
	// has a JobState value of InProgress, for both import and export jobs.
	DataTransferProgress *DataTransfer

	// The description of the job, provided at job creation.
	Description *string

	// The container for SnowconeDeviceConfiguration.
	DeviceConfiguration *DeviceConfiguration

	// The ID of the address that you want a job shipped to, after it will be shipped
	// to its primary address. This field is not supported in most regions.
	ForwardingAddressId *string

	// The automatically generated ID for a job, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	JobId *string

	// Links to Amazon S3 presigned URLs for the job report and logs. For import jobs,
	// the PDF job report becomes available at the end of the import process. For
	// export jobs, your job report typically becomes available while the Snow device
	// for your job part is being delivered to you.
	JobLogInfo *JobLogs

	// The current status of the jobs.
	JobState JobState

	// The type of job.
	JobType JobType

	// The Amazon Resource Name (ARN) for the Key Management Service (KMS) key
	// associated with this job. This ARN was created using the CreateKey
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateKey.html) API
	// action in KMS.
	KmsKeyARN *string

	// The ID of the long-term pricing type for the device.
	LongTermPricingId *string

	// The Amazon Simple Notification Service (Amazon SNS) notification settings
	// associated with a specific job. The Notification object is returned as a part of
	// the response syntax of the DescribeJob action in the JobMetadata data type.
	Notification *Notification

	// Represents metadata and configuration settings for services on an Amazon Web
	// Services Snow Family device.
	OnDeviceServiceConfiguration *OnDeviceServiceConfiguration

	// Allows you to securely operate and manage Snowcone devices remotely from outside
	// of your internal network. When set to INSTALLED_AUTOSTART, remote management
	// will automatically be available when the device arrives at your location.
	// Otherwise, you need to use the Snowball Client to manage the device.
	RemoteManagement RemoteManagement

	// An array of S3Resource objects. Each S3Resource object represents an Amazon S3
	// bucket that your transferred data will be exported from or imported into.
	Resources *JobResource

	// The role ARN associated with this job. This ARN was created using the CreateRole
	// (https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) API
	// action in Identity and Access Management.
	RoleARN *string

	// A job's shipping information, including inbound and outbound tracking numbers
	// and shipping speed options.
	ShippingDetails *ShippingDetails

	// The Snow device capacity preference for this job, specified at job creation. In
	// US regions, you can choose between 50 TB and 80 TB Snowballs. All other regions
	// use 80 TB capacity Snowballs. For more information, see
	// "https://docs.aws.amazon.com/snowball/latest/snowcone-guide/snow-device-types.html"
	// (Snow Family Devices and Capacity) in the Snowcone User Guide or
	// "https://docs.aws.amazon.com/snowball/latest/developer-guide/snow-device-types.html"
	// (Snow Family Devices and Capacity) in the Snowcone User Guide.
	SnowballCapacityPreference SnowballCapacity

	// The type of device used with this job.
	SnowballType SnowballType

	// The metadata associated with the tax documents required in your Amazon Web
	// Services Region.
	TaxDocuments *TaxDocuments

	noSmithyDocumentSerde
}

// Contains an array of Amazon Web Services resource objects. Each object
// represents an Amazon S3 bucket, an Lambda function, or an Amazon Machine Image
// (AMI) based on Amazon EC2 that is associated with a particular job.
type JobResource struct {

	// The Amazon Machine Images (AMIs) associated with this job.
	Ec2AmiResources []Ec2AmiResource

	// The Python-language Lambda functions for this job.
	LambdaResources []LambdaResource

	// An array of S3Resource objects.
	S3Resources []S3Resource

	noSmithyDocumentSerde
}

// Contains a key range. For export jobs, a S3Resource object can have an optional
// KeyRange value. The length of the range is defined at job creation, and has
// either an inclusive BeginMarker, an inclusive EndMarker, or both. Ranges are
// UTF-8 binary sorted.
type KeyRange struct {

	// The key that starts an optional key range for an export job. Ranges are
	// inclusive and UTF-8 binary sorted.
	BeginMarker *string

	// The key that ends an optional key range for an export job. Ranges are inclusive
	// and UTF-8 binary sorted.
	EndMarker *string

	noSmithyDocumentSerde
}

// Identifies
type LambdaResource struct {

	// The array of ARNs for S3Resource objects to trigger the LambdaResource objects
	// associated with this job.
	EventTriggers []EventTriggerDefinition

	// An Amazon Resource Name (ARN) that represents an Lambda function to be triggered
	// by PUT object actions on the associated local Amazon S3 resource.
	LambdaArn *string

	noSmithyDocumentSerde
}

// Each LongTermPricingListEntry object contains information about a long-term
// pricing type.
type LongTermPricingListEntry struct {

	// The current active jobs on the device the long-term pricing type.
	CurrentActiveJob *string

	// If set to true, specifies that the current long-term pricing type for the device
	// should be automatically renewed before the long-term pricing contract expires.
	IsLongTermPricingAutoRenew *bool

	// The IDs of the jobs that are associated with a long-term pricing type.
	JobIds []string

	// The end date the long-term pricing contract.
	LongTermPricingEndDate *time.Time

	// The ID of the long-term pricing type for the device.
	LongTermPricingId *string

	// The start date of the long-term pricing contract.
	LongTermPricingStartDate *time.Time

	// The status of the long-term pricing type.
	LongTermPricingStatus *string

	// The type of long-term pricing that was selected for the device.
	LongTermPricingType LongTermPricingType

	// A new device that replaces a device that is ordered with long-term pricing.
	ReplacementJob *string

	// The type of Snow Family devices associated with this long-term pricing job.
	SnowballType SnowballType

	noSmithyDocumentSerde
}

// An object that represents the metadata and configuration settings for the NFS
// (Network File System) service on an Amazon Web Services Snow Family device.
type NFSOnDeviceServiceConfiguration struct {

	// The maximum NFS storage for one Snow Family device.
	StorageLimit int32

	// The scale unit of the NFS storage on the device. Valid values: TB.
	StorageUnit StorageUnit

	noSmithyDocumentSerde
}

// The Amazon Simple Notification Service (Amazon SNS) notification settings
// associated with a specific job. The Notification object is returned as a part of
// the response syntax of the DescribeJob action in the JobMetadata data type. When
// the notification settings are defined during job creation, you can choose to
// notify based on a specific set of job states using the JobStatesToNotify array
// of strings, or you can specify that you want to have Amazon SNS notifications
// sent out for all job states with NotifyAll set to true.
type Notification struct {

	// The list of job states that will trigger a notification for this job.
	JobStatesToNotify []JobState

	// Any change in job state will trigger a notification for this job.
	NotifyAll bool

	// The new SNS TopicArn that you want to associate with this job. You can create
	// Amazon Resource Names (ARNs) for topics by using the CreateTopic
	// (https://docs.aws.amazon.com/sns/latest/api/API_CreateTopic.html) Amazon SNS API
	// action. You can subscribe email addresses to an Amazon SNS topic through the
	// Amazon Web Services Management Console, or by using the Subscribe
	// (https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html) Amazon Simple
	// Notification Service (Amazon SNS) API action.
	SnsTopicARN *string

	noSmithyDocumentSerde
}

// An object that represents the metadata and configuration settings for services
// on an Amazon Web Services Snow Family device.
type OnDeviceServiceConfiguration struct {

	// The configuration of EKS Anywhere on the Snow Family device.
	EKSOnDeviceService *EKSOnDeviceServiceConfiguration

	// Represents the NFS (Network File System) service on a Snow Family device.
	NFSOnDeviceService *NFSOnDeviceServiceConfiguration

	// Represents the Storage Gateway service Tape Gateway type on a Snow Family
	// device.
	TGWOnDeviceService *TGWOnDeviceServiceConfiguration

	noSmithyDocumentSerde
}

// Each S3Resource object represents an Amazon S3 bucket that your transferred data
// will be exported from or imported into. For export jobs, this object can have an
// optional KeyRange value. The length of the range is defined at job creation, and
// has either an inclusive BeginMarker, an inclusive EndMarker, or both. Ranges are
// UTF-8 binary sorted.
type S3Resource struct {

	// The Amazon Resource Name (ARN) of an Amazon S3 bucket.
	BucketArn *string

	// For export jobs, you can provide an optional KeyRange within a specific Amazon
	// S3 bucket. The length of the range is defined at job creation, and has either an
	// inclusive BeginMarker, an inclusive EndMarker, or both. Ranges are UTF-8 binary
	// sorted.
	KeyRange *KeyRange

	// Specifies the service or services on the Snow Family device that your
	// transferred data will be exported from or imported into. Amazon Web Services
	// Snow Family supports Amazon S3 and NFS (Network File System).
	TargetOnDeviceServices []TargetOnDeviceService

	noSmithyDocumentSerde
}

// The version of the requested service.
type ServiceVersion struct {

	// The version number of the requested service.
	Version *string

	noSmithyDocumentSerde
}

// The Status and TrackingNumber information for an inbound or outbound shipment.
type Shipment struct {

	// Status information for a shipment.
	Status *string

	// The tracking number for this job. Using this tracking number with your region's
	// carrier's website, you can track a Snow device as the carrier transports it. For
	// India, the carrier is Amazon Logistics. For all other regions, UPS is the
	// carrier.
	TrackingNumber *string

	noSmithyDocumentSerde
}

// A job's shipping information, including inbound and outbound tracking numbers
// and shipping speed options.
type ShippingDetails struct {

	// The Status and TrackingNumber values for a Snow device being returned to Amazon
	// Web Services for a particular job.
	InboundShipment *Shipment

	// The Status and TrackingNumber values for a Snow device being delivered to the
	// address that you specified for a particular job.
	OutboundShipment *Shipment

	// The shipping speed for a particular job. This speed doesn't dictate how soon
	// you'll get the Snow device from the job's creation date. This speed represents
	// how quickly it moves to its destination while in transit. Regional shipping
	// speeds are as follows:
	//
	// * In Australia, you have access to express shipping.
	// Typically, Snow devices shipped express are delivered in about a day.
	//
	// * In the
	// European Union (EU), you have access to express shipping. Typically, Snow
	// devices shipped express are delivered in about a day. In addition, most
	// countries in the EU have access to standard shipping, which typically takes less
	// than a week, one way.
	//
	// * In India, Snow devices are delivered in one to seven
	// days.
	//
	// * In the United States of America (US), you have access to one-day
	// shipping and two-day shipping.
	ShippingOption ShippingOption

	noSmithyDocumentSerde
}

// Specifies the device configuration for an Snowcone job.
type SnowconeDeviceConfiguration struct {

	// Configures the wireless connection for the Snowcone device.
	WirelessConnection *WirelessConnection

	noSmithyDocumentSerde
}

// An object that represents the service or services on the Snow Family device that
// your transferred data will be exported from or imported into. Amazon Web
// Services Snow Family supports Amazon S3 and NFS (Network File System).
type TargetOnDeviceService struct {

	// Specifies the name of the service on the Snow Family device that your
	// transferred data will be exported from or imported into.
	ServiceName DeviceServiceName

	// Specifies whether the data is being imported or exported. You can import or
	// export the data, or use it locally on the device.
	TransferOption TransferOption

	noSmithyDocumentSerde
}

// The tax documents required in your Amazon Web Services Region.
type TaxDocuments struct {

	// The tax documents required in Amazon Web Services Region in India.
	IND *INDTaxDocuments

	noSmithyDocumentSerde
}

// An object that represents the metadata and configuration settings for the
// Storage Gateway service Tape Gateway type on an Amazon Web Services Snow Family
// device.
type TGWOnDeviceServiceConfiguration struct {

	// The maximum number of virtual tapes to store on one Snow Family device. Due to
	// physical resource limitations, this value must be set to 80 for Snowball Edge.
	StorageLimit int32

	// The scale unit of the virtual tapes on the device.
	StorageUnit StorageUnit

	noSmithyDocumentSerde
}

// Configures the wireless connection on an Snowcone device.
type WirelessConnection struct {

	// Enables the Wi-Fi adapter on an Snowcone device.
	IsWifiEnabled bool

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
