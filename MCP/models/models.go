package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GcsSource represents the GcsSource schema from the OpenAPI specification
type GcsSource struct {
	Uri string `json:"uri,omitempty"` // Required. Source data URI. (e.g. 'gs://my_bucket/my_object').
}

// ListInstancesResponse represents the ListInstancesResponse schema from the OpenAPI specification
type ListInstancesResponse struct {
	Instances []Instance `json:"instances,omitempty"` // A list of Redis instances in the project in the specified location, or across all locations. If the `location_id` in the parent field of the request is "-", all regions available to the project are queried, and the results aggregated. If in such an aggregated query a location is unavailable, a placeholder Redis entry is included in the response with the `name` field set to a value of the form `projects/{project_id}/locations/{location_id}/instances/`- and the `status` field set to ERROR and `status_message` field set to "location not available for ListInstances".
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token to retrieve the next page of results, or empty if there are no more results in the list.
	Unreachable []string `json:"unreachable,omitempty"` // Locations that could not be reached.
}

// AvailabilityConfiguration represents the AvailabilityConfiguration schema from the OpenAPI specification
type AvailabilityConfiguration struct {
	Availabilitytype string `json:"availabilityType,omitempty"` // Availability type. Potential values: * `ZONAL`: The instance serves data from only one zone. Outages in that zone affect data accessibility. * `REGIONAL`: The instance can serve data from more than one zone in a region (it is highly available).
	Externalreplicaconfigured bool `json:"externalReplicaConfigured,omitempty"`
	Promotablereplicaconfigured bool `json:"promotableReplicaConfigured,omitempty"`
}

// DatabaseResourceFeed represents the DatabaseResourceFeed schema from the OpenAPI specification
type DatabaseResourceFeed struct {
	Feedtimestamp string `json:"feedTimestamp,omitempty"` // Required. Timestamp when feed is generated.
	Feedtype string `json:"feedType,omitempty"` // Required. Type feed to be ingested into condor
	Recommendationsignaldata DatabaseResourceRecommendationSignalData `json:"recommendationSignalData,omitempty"` // Common model for database resource recommendation signal data.
	Resourcehealthsignaldata DatabaseResourceHealthSignalData `json:"resourceHealthSignalData,omitempty"` // Common model for database resource health signal data.
	Resourceid DatabaseResourceId `json:"resourceId,omitempty"` // DatabaseResourceId will serve as primary key for any resource ingestion event.
	Resourcemetadata DatabaseResourceMetadata `json:"resourceMetadata,omitempty"` // Common model for database resource instance metadata.
}

// ListOperationsResponse represents the ListOperationsResponse schema from the OpenAPI specification
type ListOperationsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The standard List next-page token.
	Operations []Operation `json:"operations,omitempty"` // A list of operations that matches the specified filter in the request.
}

// TimeOfDay represents the TimeOfDay schema from the OpenAPI specification
type TimeOfDay struct {
	Seconds int `json:"seconds,omitempty"` // Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
	Hours int `json:"hours,omitempty"` // Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	Minutes int `json:"minutes,omitempty"` // Minutes of hour of day. Must be from 0 to 59.
	Nanos int `json:"nanos,omitempty"` // Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
}

// CertChain represents the CertChain schema from the OpenAPI specification
type CertChain struct {
	Certificates []string `json:"certificates,omitempty"` // The certificates that form the CA chain, from leaf to root order.
}

// GoogleCloudRedisV1OperationMetadata represents the GoogleCloudRedisV1OperationMetadata schema from the OpenAPI specification
type GoogleCloudRedisV1OperationMetadata struct {
	Endtime string `json:"endTime,omitempty"` // End timestamp.
	Statusdetail string `json:"statusDetail,omitempty"` // Operation status details.
	Target string `json:"target,omitempty"` // Operation target.
	Verb string `json:"verb,omitempty"` // Operation verb.
	Apiversion string `json:"apiVersion,omitempty"` // API version.
	Cancelrequested bool `json:"cancelRequested,omitempty"` // Specifies if cancellation was requested for the operation.
	Createtime string `json:"createTime,omitempty"` // Creation timestamp.
}

// DatabaseResourceId represents the DatabaseResourceId schema from the OpenAPI specification
type DatabaseResourceId struct {
	Uniqueid string `json:"uniqueId,omitempty"` // Required. A service-local token that distinguishes this resource from other resources within the same service.
	Provider string `json:"provider,omitempty"` // Required. Cloud provider name. Ex: GCP/AWS/Azure/OnPrem/SelfManaged
	Providerdescription string `json:"providerDescription,omitempty"` // Optional. Needs to be used only when the provider is PROVIDER_OTHER.
	Resourcetype string `json:"resourceType,omitempty"` // Required. The type of resource this ID is identifying. Ex redis.googleapis.com/Instance, redis.googleapis.com/Cluster, alloydb.googleapis.com/Cluster, alloydb.googleapis.com/Instance, spanner.googleapis.com/Instance REQUIRED Please refer go/condor-common-datamodel
}

// GcsDestination represents the GcsDestination schema from the OpenAPI specification
type GcsDestination struct {
	Uri string `json:"uri,omitempty"` // Required. Data destination URI (e.g. 'gs://my_bucket/my_object'). Existing files will be overwritten.
}

// GoogleCloudRedisV1ZoneMetadata represents the GoogleCloudRedisV1ZoneMetadata schema from the OpenAPI specification
type GoogleCloudRedisV1ZoneMetadata struct {
}

// MaintenanceSchedule represents the MaintenanceSchedule schema from the OpenAPI specification
type MaintenanceSchedule struct {
	Canreschedule bool `json:"canReschedule,omitempty"` // If the scheduled maintenance can be rescheduled, default is true.
	Endtime string `json:"endTime,omitempty"` // Output only. The end time of any upcoming scheduled maintenance for this instance.
	Scheduledeadlinetime string `json:"scheduleDeadlineTime,omitempty"` // Output only. The deadline that the maintenance schedule start time can not go beyond, including reschedule.
	Starttime string `json:"startTime,omitempty"` // Output only. The start time of any upcoming scheduled maintenance for this instance.
}

// TlsCertificate represents the TlsCertificate schema from the OpenAPI specification
type TlsCertificate struct {
	Cert string `json:"cert,omitempty"` // PEM representation.
	Createtime string `json:"createTime,omitempty"` // Output only. The time when the certificate was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2020-05-18T00:00:00.094Z`.
	Expiretime string `json:"expireTime,omitempty"` // Output only. The time when the certificate expires in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2020-05-18T00:00:00.094Z`.
	Serialnumber string `json:"serialNumber,omitempty"` // Serial number, as extracted from the certificate.
	Sha1fingerprint string `json:"sha1Fingerprint,omitempty"` // Sha1 Fingerprint of the certificate.
}

// GoogleCloudRedisV1LocationMetadata represents the GoogleCloudRedisV1LocationMetadata schema from the OpenAPI specification
type GoogleCloudRedisV1LocationMetadata struct {
	Availablezones map[string]interface{} `json:"availableZones,omitempty"` // Output only. The set of available zones in the location. The map is keyed by the lowercase ID of each zone, as defined by GCE. These keys can be specified in `location_id` or `alternative_location_id` fields when creating a Redis instance.
}

// RescheduleMaintenanceRequest represents the RescheduleMaintenanceRequest schema from the OpenAPI specification
type RescheduleMaintenanceRequest struct {
	Rescheduletype string `json:"rescheduleType,omitempty"` // Required. If reschedule type is SPECIFIC_TIME, must set up schedule_time as well.
	Scheduletime string `json:"scheduleTime,omitempty"` // Optional. Timestamp when the maintenance shall be rescheduled to if reschedule_type=SPECIFIC_TIME, in RFC 3339 format, for example `2012-11-15T16:19:00.094Z`.
}

// ImportInstanceRequest represents the ImportInstanceRequest schema from the OpenAPI specification
type ImportInstanceRequest struct {
	Inputconfig InputConfig `json:"inputConfig,omitempty"` // The input content
}

// FailoverInstanceRequest represents the FailoverInstanceRequest schema from the OpenAPI specification
type FailoverInstanceRequest struct {
	Dataprotectionmode string `json:"dataProtectionMode,omitempty"` // Optional. Available data protection modes that the user can choose. If it's unspecified, data protection mode will be LIMITED_DATA_LOSS by default.
}

// MaintenancePolicy represents the MaintenancePolicy schema from the OpenAPI specification
type MaintenancePolicy struct {
	Description string `json:"description,omitempty"` // Optional. Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time when the policy was last updated.
	Weeklymaintenancewindow []WeeklyMaintenanceWindow `json:"weeklyMaintenanceWindow,omitempty"` // Optional. Maintenance window that is applied to resources covered by this policy. Minimum 1. For the current version, the maximum number of weekly_window is expected to be one.
	Createtime string `json:"createTime,omitempty"` // Output only. The time when the policy was created.
}

// Cluster represents the Cluster schema from the OpenAPI specification
type Cluster struct {
	Pscconfigs []PscConfig `json:"pscConfigs,omitempty"` // Required. Each PscConfig configures the consumer network where IPs will be designated to the cluster for client access through Private Service Connect Automation. Currently, only one PscConfig is supported.
	Sizegb int `json:"sizeGb,omitempty"` // Output only. Redis memory size in GB for the entire cluster rounded up to the next integer.
	Uid string `json:"uid,omitempty"` // Output only. System assigned, unique identifier for the cluster.
	Replicacount int `json:"replicaCount,omitempty"` // Optional. The number of replica nodes per shard.
	Transitencryptionmode string `json:"transitEncryptionMode,omitempty"` // Optional. The in-transit encryption for the Redis cluster. If not provided, encryption is disabled for the cluster.
	Pscconnections []PscConnection `json:"pscConnections,omitempty"` // Output only. PSC connections for discovery of the cluster topology and accessing the cluster.
	Stateinfo StateInfo `json:"stateInfo,omitempty"` // Represents additional information about the state of the cluster.
	Createtime string `json:"createTime,omitempty"` // Output only. The timestamp associated with the cluster creation request.
	Name string `json:"name,omitempty"` // Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/clusters/{cluster_id}`
	Shardcount int `json:"shardCount,omitempty"` // Required. Number of shards for the Redis cluster.
	State string `json:"state,omitempty"` // Output only. The current state of this cluster. Can be CREATING, READY, UPDATING, DELETING and SUSPENDED
	Authorizationmode string `json:"authorizationMode,omitempty"` // Optional. The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster.
	Discoveryendpoints []DiscoveryEndpoint `json:"discoveryEndpoints,omitempty"` // Output only. Endpoints created on each given network, for Redis clients to connect to the cluster. Currently only one discovery endpoint is supported.
}

// Compliance represents the Compliance schema from the OpenAPI specification
type Compliance struct {
	Standard string `json:"standard,omitempty"` // Industry-wide compliance standards or benchmarks, such as CIS, PCI, and OWASP.
	Version string `json:"version,omitempty"` // Version of the standard or benchmark, for example, 1.1
}

// CustomMetadataData represents the CustomMetadataData schema from the OpenAPI specification
type CustomMetadataData struct {
	Databasemetadata []DatabaseMetadata `json:"databaseMetadata,omitempty"`
}

// WeeklyMaintenanceWindow represents the WeeklyMaintenanceWindow schema from the OpenAPI specification
type WeeklyMaintenanceWindow struct {
	Day string `json:"day,omitempty"` // Required. The day of week that maintenance updates occur.
	Duration string `json:"duration,omitempty"` // Output only. Duration of the maintenance window. The current window is fixed at 1 hour.
	Starttime TimeOfDay `json:"startTime,omitempty"` // Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
}

// OutputConfig represents the OutputConfig schema from the OpenAPI specification
type OutputConfig struct {
	Gcsdestination GcsDestination `json:"gcsDestination,omitempty"` // The Cloud Storage location for the output content
}

// ManagedCertificateAuthority represents the ManagedCertificateAuthority schema from the OpenAPI specification
type ManagedCertificateAuthority struct {
	Cacerts []CertChain `json:"caCerts,omitempty"` // The PEM encoded CA certificate chains for redis managed server authentication
}

// InputConfig represents the InputConfig schema from the OpenAPI specification
type InputConfig struct {
	Gcssource GcsSource `json:"gcsSource,omitempty"` // The Cloud Storage location for the input content
}

// PscConnection represents the PscConnection schema from the OpenAPI specification
type PscConnection struct {
	Address string `json:"address,omitempty"` // Output only. The IP allocated on the consumer network for the PSC forwarding rule.
	Forwardingrule string `json:"forwardingRule,omitempty"` // Output only. The URI of the consumer side forwarding rule. Example: projects/{projectNumOrId}/regions/us-east1/forwardingRules/{resourceId}.
	Network string `json:"network,omitempty"` // The consumer network where the IP address resides, in the form of projects/{project_id}/global/networks/{network_id}.
	Projectid string `json:"projectId,omitempty"` // Output only. The consumer project_id where the forwarding rule is created from.
	Pscconnectionid string `json:"pscConnectionId,omitempty"` // Output only. The PSC connection id of the forwarding rule connected to the service attachment.
}

// UpgradeInstanceRequest represents the UpgradeInstanceRequest schema from the OpenAPI specification
type UpgradeInstanceRequest struct {
	Redisversion string `json:"redisVersion,omitempty"` // Required. Specifies the target version of Redis software to upgrade to.
}

// ListLocationsResponse represents the ListLocationsResponse schema from the OpenAPI specification
type ListLocationsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The standard List next-page token.
	Locations []Location `json:"locations,omitempty"` // A list of locations that matches the specified filter in the request.
}

// Product represents the Product schema from the OpenAPI specification
type Product struct {
	Engine string `json:"engine,omitempty"` // The specific engine that the underlying database is running.
	TypeField string `json:"type,omitempty"` // Type of specific database product. It could be CloudSQL, AlloyDB etc..
	Version string `json:"version,omitempty"` // Version of the underlying database engine. Example values: For MySQL, it could be "8.0", "5.7" etc.. For Postgres, it could be "14", "15" etc..
}

// Entitlement represents the Entitlement schema from the OpenAPI specification
type Entitlement struct {
	Entitlementstate string `json:"entitlementState,omitempty"` // The current state of user's accessibility to a feature/benefit.
	TypeField string `json:"type,omitempty"` // An enum that represents the type of this entitlement.
}

// BackupRun represents the BackupRun schema from the OpenAPI specification
type BackupRun struct {
	Endtime string `json:"endTime,omitempty"` // The time the backup operation completed. REQUIRED
	ErrorField OperationError `json:"error,omitempty"` // An error that occurred during a backup creation operation.
	Starttime string `json:"startTime,omitempty"` // The time the backup operation started. REQUIRED
	Status string `json:"status,omitempty"` // The status of this run. REQUIRED
}

// ExportInstanceRequest represents the ExportInstanceRequest schema from the OpenAPI specification
type ExportInstanceRequest struct {
	Outputconfig OutputConfig `json:"outputConfig,omitempty"` // The output content
}

// Empty represents the Empty schema from the OpenAPI specification
type Empty struct {
}

// OperationMetadata represents the OperationMetadata schema from the OpenAPI specification
type OperationMetadata struct {
	Requestedcancellation bool `json:"requestedCancellation,omitempty"` // Output only. Identifies whether the user has requested cancellation of the operation. Operations that have successfully been cancelled have Operation.error value with a google.rpc.Status.code of 1, corresponding to `Code.CANCELLED`.
	Statusmessage string `json:"statusMessage,omitempty"` // Output only. Human-readable status of the operation, if any.
	Target string `json:"target,omitempty"` // Output only. Server-defined resource path for the target of the operation.
	Verb string `json:"verb,omitempty"` // Output only. Name of the verb executed by the operation.
	Apiversion string `json:"apiVersion,omitempty"` // Output only. API version used to start the operation.
	Createtime string `json:"createTime,omitempty"` // Output only. The time the operation was created.
	Endtime string `json:"endTime,omitempty"` // Output only. The time the operation finished running.
}

// PersistenceConfig represents the PersistenceConfig schema from the OpenAPI specification
type PersistenceConfig struct {
	Persistencemode string `json:"persistenceMode,omitempty"` // Optional. Controls whether Persistence features are enabled. If not provided, the existing value will be used.
	Rdbnextsnapshottime string `json:"rdbNextSnapshotTime,omitempty"` // Output only. The next time that a snapshot attempt is scheduled to occur.
	Rdbsnapshotperiod string `json:"rdbSnapshotPeriod,omitempty"` // Optional. Period between RDB snapshots. Snapshots will be attempted every period starting from the provided snapshot start time. For example, a start time of 01/01/2033 06:45 and SIX_HOURS snapshot period will do nothing until 01/01/2033, and then trigger snapshots every day at 06:45, 12:45, 18:45, and 00:45 the next day, and so on. If not provided, TWENTY_FOUR_HOURS will be used as default.
	Rdbsnapshotstarttime string `json:"rdbSnapshotStartTime,omitempty"` // Optional. Date and time that the first snapshot was/will be attempted, and to which future snapshots will be aligned. If not provided, the current time will be used.
}

// ReconciliationOperationMetadata represents the ReconciliationOperationMetadata schema from the OpenAPI specification
type ReconciliationOperationMetadata struct {
	Deleteresource bool `json:"deleteResource,omitempty"` // DEPRECATED. Use exclusive_action instead.
	Exclusiveaction string `json:"exclusiveAction,omitempty"` // Excluisive action returned by the CLH.
}

// StateInfo represents the StateInfo schema from the OpenAPI specification
type StateInfo struct {
	Updateinfo UpdateInfo `json:"updateInfo,omitempty"` // Represents information about an updating cluster.
}

// Location represents the Location schema from the OpenAPI specification
type Location struct {
	Name string `json:"name,omitempty"` // Full resource name for the region. For example: "projects/example-project/locations/us-east1".
	Displayname string `json:"displayName,omitempty"` // The friendly name for this location, typically a nearby city name. For example, "Tokyo".
	Labels map[string]interface{} `json:"labels,omitempty"` // Cross-service attributes for the location. For example {"cloud.googleapis.com/region": "us-east1"}
	Locationid string `json:"locationId,omitempty"` // Resource ID for the region. For example: "us-east1".
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Output only. The set of available zones in the location. The map is keyed by the lowercase ID of each zone, as defined by Compute Engine. These keys can be specified in `location_id` or `alternative_location_id` fields when creating a Redis instance.
}

// DatabaseMetadata represents the DatabaseMetadata schema from the OpenAPI specification
type DatabaseMetadata struct {
	Backuprun BackupRun `json:"backupRun,omitempty"` // A backup run.
	Product Product `json:"product,omitempty"` // Product specification for Condor resources.
	Resourceid DatabaseResourceId `json:"resourceId,omitempty"` // DatabaseResourceId will serve as primary key for any resource ingestion event.
	Resourcename string `json:"resourceName,omitempty"` // Required. Database name. Resource name to follow CAIS resource_name format as noted here go/condor-common-datamodel
	Backupconfiguration BackupConfiguration `json:"backupConfiguration,omitempty"` // Configuration for automatic backups
}

// UpdateInfo represents the UpdateInfo schema from the OpenAPI specification
type UpdateInfo struct {
	Targetreplicacount int `json:"targetReplicaCount,omitempty"` // Target number of replica nodes per shard.
	Targetshardcount int `json:"targetShardCount,omitempty"` // Target number of shards for redis cluster
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Message string `json:"message,omitempty"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
	Code int `json:"code,omitempty"` // The status code, which should be an enum value of google.rpc.Code.
	Details []map[string]interface{} `json:"details,omitempty"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.
}

// DatabaseResourceHealthSignalData represents the DatabaseResourceHealthSignalData schema from the OpenAPI specification
type DatabaseResourceHealthSignalData struct {
	Resourcecontainer string `json:"resourceContainer,omitempty"` // Closest parent container of this resource. In GCP, 'container' refers to a Cloud Resource Manager project. It must be resource name of a Cloud Resource Manager project with the format of "provider//", such as "projects/123". For GCP provided resources, number should be project number.
	State string `json:"state,omitempty"`
	Description string `json:"description,omitempty"` // Description associated with signal
	Name string `json:"name,omitempty"` // Required. The name of the signal, ex: PUBLIC_SQL_INSTANCE, SQL_LOG_ERROR_VERBOSITY etc.
	Provider string `json:"provider,omitempty"` // Cloud provider name. Ex: GCP/AWS/Azure/OnPrem/SelfManaged
	Signalid string `json:"signalId,omitempty"` // Required. Unique identifier for the signal. This is an unique id which would be mainatined by partner to identify a signal.
	Compliance []Compliance `json:"compliance,omitempty"` // Industry standards associated with this signal; if this signal is an issue, that could be a violation of the associated industry standard(s). For example, AUTO_BACKUP_DISABLED signal is associated with CIS GCP 1.1, CIS GCP 1.2, CIS GCP 1.3, NIST 800-53 and ISO-27001 compliance standards. If a database resource does not have automated backup enable, it will violate these following industry standards.
	Externaluri string `json:"externalUri,omitempty"` // The external-uri of the signal, using which more information about this signal can be obtained. In GCP, this will take user to SCC page to get more details about signals.
	Resourcename string `json:"resourceName,omitempty"` // Required. Database resource name associated with the signal. Resource name to follow CAIS resource_name format as noted here go/condor-common-datamodel
	Signalclass string `json:"signalClass,omitempty"` // Required. The class of the signal, such as if it's a THREAT or VULNERABILITY.
	Additionalmetadata map[string]interface{} `json:"additionalMetadata,omitempty"` // Any other additional metadata
	Signaltype string `json:"signalType,omitempty"` // Required. Type of signal, for example, `AVAILABLE_IN_MULTIPLE_ZONES`, `LOGGING_MOST_ERRORS`, etc.
	Eventtime string `json:"eventTime,omitempty"` // Required. The last time at which the event described by this signal took place
}

// NodeInfo represents the NodeInfo schema from the OpenAPI specification
type NodeInfo struct {
	Id string `json:"id,omitempty"` // Output only. Node identifying string. e.g. 'node-0', 'node-1'
	Zone string `json:"zone,omitempty"` // Output only. Location of the node.
}

// OperationError represents the OperationError schema from the OpenAPI specification
type OperationError struct {
	Code string `json:"code,omitempty"` // Identifies the specific error that occurred. REQUIRED
	Errortype string `json:"errorType,omitempty"`
	Message string `json:"message,omitempty"` // Additional information about the error encountered. REQUIRED
}

// PscConfig represents the PscConfig schema from the OpenAPI specification
type PscConfig struct {
	Network string `json:"network,omitempty"` // Required. The network where the IP address of the discovery endpoint will be reserved, in the form of projects/{network_project}/global/networks/{network_id}.
}

// DatabaseResourceMetadata represents the DatabaseResourceMetadata schema from the OpenAPI specification
type DatabaseResourceMetadata struct {
	Primaryresourceid DatabaseResourceId `json:"primaryResourceId,omitempty"` // DatabaseResourceId will serve as primary key for any resource ingestion event.
	Resourcename string `json:"resourceName,omitempty"` // Required. Different from DatabaseResourceId.unique_id, a resource name can be reused over time. That is, after a resource named "ABC" is deleted, the name "ABC" can be used to to create a new resource within the same source. Resource name to follow CAIS resource_name format as noted here go/condor-common-datamodel
	Availabilityconfiguration AvailabilityConfiguration `json:"availabilityConfiguration,omitempty"` // Configuration for availability of database instance
	Userlabels map[string]interface{} `json:"userLabels,omitempty"` // User-provided labels, represented as a dictionary where each label is a single key value pair.
	Custommetadata CustomMetadataData `json:"customMetadata,omitempty"` // Any custom metadata associated with the resource. i.e. A spanner instance can have multiple databases with its own unique metadata. Information for these individual databases can be captured in custom metadata data
	Location string `json:"location,omitempty"` // The resource location. REQUIRED
	Backupconfiguration BackupConfiguration `json:"backupConfiguration,omitempty"` // Configuration for automatic backups
	Creationtime string `json:"creationTime,omitempty"` // The creation time of the resource, i.e. the time when resource is created and recorded in partner service.
	Expectedstate string `json:"expectedState,omitempty"` // The state that the instance is expected to be in. For example, an instance state can transition to UNHEALTHY due to wrong patch update, while the expected state will remain at the HEALTHY.
	Product Product `json:"product,omitempty"` // Product specification for Condor resources.
	Resourcecontainer string `json:"resourceContainer,omitempty"` // Closest parent Cloud Resource Manager container of this resource. It must be resource name of a Cloud Resource Manager project with the format of "/", such as "projects/123". For GCP provided resources, number should be project number.
	Backuprun BackupRun `json:"backupRun,omitempty"` // A backup run.
	Currentstate string `json:"currentState,omitempty"` // Current state of the instance.
	Id DatabaseResourceId `json:"id,omitempty"` // DatabaseResourceId will serve as primary key for any resource ingestion event.
	Instancetype string `json:"instanceType,omitempty"` // The type of the instance. Specified at creation time.
	Entitlements []Entitlement `json:"entitlements,omitempty"` // Entitlements associated with the resource
	Updationtime string `json:"updationTime,omitempty"` // The time at which the resource was updated and recorded at partner service.
}

// Operation represents the Operation schema from the OpenAPI specification
type Operation struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"` // { `createTime`: The time the operation was created. `endTime`: The time the operation finished running. `target`: Server-defined resource path for the target of the operation. `verb`: Name of the verb executed by the operation. `statusDetail`: Human-readable status of the operation, if any. `cancelRequested`: Identifies whether the user has requested cancellation of the operation. Operations that have successfully been cancelled have Operation.error value with a google.rpc.Status.code of 1, corresponding to `Code.CANCELLED`. `apiVersion`: API version used to start the operation. }
	Name string `json:"name,omitempty"` // The server-assigned name, which is only unique within the same service that originally returns it. If you use the default HTTP mapping, the `name` should be a resource name ending with `operations/{unique_id}`.
	Response map[string]interface{} `json:"response,omitempty"` // The normal, successful response of the operation. If the original method returns no data on success, such as `Delete`, the response is `google.protobuf.Empty`. If the original method is standard `Get`/`Create`/`Update`, the response should be the resource. For other methods, the response should have the type `XxxResponse`, where `Xxx` is the original method name. For example, if the original method name is `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.
	Done bool `json:"done,omitempty"` // If the value is `false`, it means the operation is still in progress. If `true`, the operation is completed, and either `error` or `response` is available.
	ErrorField Status `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
}

// RetentionSettings represents the RetentionSettings schema from the OpenAPI specification
type RetentionSettings struct {
	Retentionunit string `json:"retentionUnit,omitempty"` // The unit that 'retained_backups' represents.
	Timebasedretention string `json:"timeBasedRetention,omitempty"`
	Quantitybasedretention int `json:"quantityBasedRetention,omitempty"`
}

// InstanceAuthString represents the InstanceAuthString schema from the OpenAPI specification
type InstanceAuthString struct {
	Authstring string `json:"authString,omitempty"` // AUTH string set on the instance.
}

// DiscoveryEndpoint represents the DiscoveryEndpoint schema from the OpenAPI specification
type DiscoveryEndpoint struct {
	Address string `json:"address,omitempty"` // Output only. Address of the exposed Redis endpoint used by clients to connect to the service. The address could be either IP or hostname.
	Port int `json:"port,omitempty"` // Output only. The port number of the exposed Redis endpoint.
	Pscconfig PscConfig `json:"pscConfig,omitempty"`
}

// Instance represents the Instance schema from the OpenAPI specification
type Instance struct {
	Readreplicasmode string `json:"readReplicasMode,omitempty"` // Optional. Read replicas mode for the instance. Defaults to READ_REPLICAS_DISABLED.
	Readendpoint string `json:"readEndpoint,omitempty"` // Output only. Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only. Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes will exhibit some lag behind the primary. Write requests must target 'host'.
	Secondaryiprange string `json:"secondaryIpRange,omitempty"` // Optional. Additional IP range for node placement. Required when enabling read replicas on an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address range associated with the private service access connection, or "auto".
	Maintenancepolicy MaintenancePolicy `json:"maintenancePolicy,omitempty"` // Maintenance policy for an instance.
	Connectmode string `json:"connectMode,omitempty"` // Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
	Authorizednetwork string `json:"authorizedNetwork,omitempty"` // Optional. The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
	Displayname string `json:"displayName,omitempty"` // An arbitrary and optional user-provided name for the instance.
	Customermanagedkey string `json:"customerManagedKey,omitempty"` // Optional. The KMS key reference that the customer provides when trying to create the instance.
	Name string `json:"name,omitempty"` // Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Redis instances are managed and addressed at regional level so location_id here refers to a GCP region; however, users may choose which specific zone (or collection of zones for cross-zone instances) an instance should be provisioned in. Refer to location_id and alternative_location_id fields for more details.
	Maintenanceversion string `json:"maintenanceVersion,omitempty"` // Optional. The self service update maintenance version. The version is date based such as "20210712_00_00".
	Persistenceiamidentity string `json:"persistenceIamIdentity,omitempty"` // Output only. Cloud IAM identity used by import / export operations to transfer data to/from Cloud Storage. Format is "serviceAccount:". The value may change over time for a given instance so should be checked before each import/export operation.
	Createtime string `json:"createTime,omitempty"` // Output only. The time the instance was created.
	Servercacerts []TlsCertificate `json:"serverCaCerts,omitempty"` // Output only. List of server CA certificates for the instance.
	Port int `json:"port,omitempty"` // Output only. The port number of the exposed Redis endpoint.
	Satisfiespzs bool `json:"satisfiesPzs,omitempty"` // Optional. Output only. Reserved for future use.
	Authenabled bool `json:"authEnabled,omitempty"` // Optional. Indicates whether OSS Redis AUTH is enabled for the instance. If set to "true" AUTH is enabled on the instance. Default value is "false" meaning AUTH is disabled.
	Labels map[string]interface{} `json:"labels,omitempty"` // Resource labels to represent user provided metadata
	Replicacount int `json:"replicaCount,omitempty"` // Optional. The number of replica nodes. The valid range for the Standard Tier with read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled for a Standard Tier instance, the only valid value is 1 and the default is 1. The valid value for basic tier is 0 and the default is also 0.
	Redisconfigs map[string]interface{} `json:"redisConfigs,omitempty"` // Optional. Redis configuration parameters, according to http://redis.io/topics/config. Currently, the only supported parameters are: Redis version 3.2 and newer: * maxmemory-policy * notify-keyspace-events Redis version 4.0 and newer: * activedefrag * lfu-decay-time * lfu-log-factor * maxmemory-gb Redis version 5.0 and newer: * stream-node-max-bytes * stream-node-max-entries
	Nodes []NodeInfo `json:"nodes,omitempty"` // Output only. Info per node.
	Alternativelocationid string `json:"alternativeLocationId,omitempty"` // Optional. If specified, at least one node will be provisioned in this zone in addition to the zone specified in location_id. Only applicable to standard tier. If provided, it must be a different zone from the one provided in [location_id]. Additional nodes beyond the first 2 will be placed in zones selected by the service.
	Redisversion string `json:"redisVersion,omitempty"` // Optional. The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: * `REDIS_3_2` for Redis 3.2 compatibility * `REDIS_4_0` for Redis 4.0 compatibility (default) * `REDIS_5_0` for Redis 5.0 compatibility * `REDIS_6_X` for Redis 6.x compatibility * `REDIS_7_0` for Redis 7.0 compatibility
	Transitencryptionmode string `json:"transitEncryptionMode,omitempty"` // Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
	Availablemaintenanceversions []string `json:"availableMaintenanceVersions,omitempty"` // Optional. The available maintenance versions that an instance could update to.
	Readendpointport int `json:"readEndpointPort,omitempty"` // Output only. The port number of the exposed readonly redis endpoint. Standard tier only. Write requests should target 'port'.
	Satisfiespzi bool `json:"satisfiesPzi,omitempty"` // Optional. Output only. Reserved for future use.
	Persistenceconfig PersistenceConfig `json:"persistenceConfig,omitempty"` // Configuration of the persistence functionality.
	Currentlocationid string `json:"currentLocationId,omitempty"` // Output only. The current zone where the Redis primary node is located. In basic tier, this will always be the same as [location_id]. In standard tier, this can be the zone of any node in the instance.
	Statusmessage string `json:"statusMessage,omitempty"` // Output only. Additional information about the current status of this instance, if available.
	State string `json:"state,omitempty"` // Output only. The current state of this instance.
	Locationid string `json:"locationId,omitempty"` // Optional. The zone where the instance will be provisioned. If not provided, the service will choose a zone from the specified region for the instance. For standard tier, additional nodes will be added across multiple zones for protection against zonal failures. If specified, at least one node will be provisioned in this zone.
	Tier string `json:"tier,omitempty"` // Required. The service tier of the instance.
	Reservediprange string `json:"reservedIpRange,omitempty"` // Optional. For DIRECT_PEERING mode, the CIDR range of internal addresses that are reserved for this instance. Range must be unique and non-overlapping with existing subnets in an authorized network. For PRIVATE_SERVICE_ACCESS mode, the name of one allocated IP address ranges associated with this private service access connection. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29. For READ_REPLICAS_ENABLED the default block size is /28.
	Host string `json:"host,omitempty"` // Output only. Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.
	Maintenanceschedule MaintenanceSchedule `json:"maintenanceSchedule,omitempty"` // Upcoming maintenance schedule. If no maintenance is scheduled, fields are not populated.
	Memorysizegb int `json:"memorySizeGb,omitempty"` // Required. Redis memory size in GiB.
	Suspensionreasons []string `json:"suspensionReasons,omitempty"` // Optional. reasons that causes instance in "SUSPENDED" state.
}

// DatabaseResourceRecommendationSignalData represents the DatabaseResourceRecommendationSignalData schema from the OpenAPI specification
type DatabaseResourceRecommendationSignalData struct {
	Additionalmetadata map[string]interface{} `json:"additionalMetadata,omitempty"` // Optional. Any other additional metadata specific to recommendation
	Lastrefreshtime string `json:"lastRefreshTime,omitempty"` // Required. last time recommendationw as refreshed
	Recommendationstate string `json:"recommendationState,omitempty"` // Required. Recommendation state
	Recommender string `json:"recommender,omitempty"` // Required. Name of recommendation. Examples: organizations/1234/locations/us-central1/recommenders/google.cloudsql.instance.PerformanceRecommender/recommendations/9876
	Recommenderid string `json:"recommenderId,omitempty"` // Required. ID of recommender. Examples: "google.cloudsql.instance.PerformanceRecommender"
	Recommendersubtype string `json:"recommenderSubtype,omitempty"` // Required. Contains an identifier for a subtype of recommendations produced for the same recommender. Subtype is a function of content and impact, meaning a new subtype might be added when significant changes to `content` or `primary_impact.category` are introduced. See the Recommenders section to see a list of subtypes for a given Recommender. Examples: For recommender = "google.cloudsql.instance.PerformanceRecommender", recommender_subtype can be "MYSQL_HIGH_NUMBER_OF_OPEN_TABLES_BEST_PRACTICE"/"POSTGRES_HIGH_TRANSACTION_ID_UTILIZATION_BEST_PRACTICE"
	Resourcename string `json:"resourceName,omitempty"` // Required. Database resource name associated with the signal. Resource name to follow CAIS resource_name format as noted here go/condor-common-datamodel
	Signaltype string `json:"signalType,omitempty"` // Required. Type of signal, for example, `SIGNAL_TYPE_IDLE`, `SIGNAL_TYPE_HIGH_NUMBER_OF_TABLES`, etc.
}

// BackupConfiguration represents the BackupConfiguration schema from the OpenAPI specification
type BackupConfiguration struct {
	Automatedbackupenabled bool `json:"automatedBackupEnabled,omitempty"` // Whether customer visible automated backups are enabled on the instance.
	Backupretentionsettings RetentionSettings `json:"backupRetentionSettings,omitempty"`
	Pointintimerecoveryenabled bool `json:"pointInTimeRecoveryEnabled,omitempty"` // Whether point-in-time recovery is enabled. This is optional field, if the database service does not have this feature or metadata is not available in control plane, this can be omitted.
}

// CertificateAuthority represents the CertificateAuthority schema from the OpenAPI specification
type CertificateAuthority struct {
	Managedserverca ManagedCertificateAuthority `json:"managedServerCa,omitempty"`
	Name string `json:"name,omitempty"` // Identifier. Unique name of the resource in this scope including project, location and cluster using the form: `projects/{project}/locations/{location}/clusters/{cluster}/certificateAuthority`
}

// ListClustersResponse represents the ListClustersResponse schema from the OpenAPI specification
type ListClustersResponse struct {
	Clusters []Cluster `json:"clusters,omitempty"` // A list of Redis clusters in the project in the specified location, or across all locations. If the `location_id` in the parent field of the request is "-", all regions available to the project are queried, and the results aggregated. If in such an aggregated query a location is unavailable, a placeholder Redis entry is included in the response with the `name` field set to a value of the form `projects/{project_id}/locations/{location_id}/clusters/`- and the `status` field set to ERROR and `status_message` field set to "location not available for ListClusters".
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token to retrieve the next page of results, or empty if there are no more results in the list.
	Unreachable []string `json:"unreachable,omitempty"` // Locations that could not be reached.
}
