# Model
model:
  rest_name: firewalllicensingmetrics
  resource_name: firewalllicensingmetrics
  entity_name: FirewallLicensingMetrics
  package: licensingsrv
  group: core/licensing
  description: Represents firewall licensing metrics.
  private: true
  get:
    description: Retrieves the firewall licensing metrics with the given ID.
  delete:
    description: Deletes the firewall licensing metrics with the given ID.
  extends:
  - '@identifiable-stored'
  - '@namespaced'
  - '@base'
  - '@zoned'
  - '@timeable'

# Indexes
indexes:
- - namespace
  - firewallResource

# Attributes
attributes:
  v1:
  - name: EnhancedSupport
    description: Represents enhanced support metric.
    type: integer
    exposed: true
    stored: true

  - name: FWUsageHours
    description: Represents firewall usage hours.
    type: integer
    exposed: true
    stored: true

  - name: TPUsageHours
    description: Represents threat prevention usage hours.
    type: integer
    exposed: true
    stored: true

  - name: TrafficSecured
    description: Represents traffic secured usage.
    type: integer
    exposed: true
    stored: true

  - name: URLFUsageHours
    description: Represents url filtering usage hours.
    type: integer
    exposed: true
    stored: true

  - name: cloudAccountID
    description: Represents the user cloud account ID.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: cloudxyz

  - name: firewallResource
    description: Represents firewall resource.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: firewall1

  - name: tenantCloudAccountID
    description: Represents the cloud ngfw tenant AWS account ID.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 1224212aa

  - name: tenantID
    description: Represents the cloud ngfw tenant ID.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: tenantxyz

  - name: timestamp
    description: The timestamp of this record.
    type: time
    exposed: true
    stored: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
