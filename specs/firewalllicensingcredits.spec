# Model
model:
  rest_name: firewalllicensingcredits
  resource_name: firewalllicensingcredits
  entity_name: FirewallLicensingCredits
  package: licensingsrv
  group: core/licensing
  description: Represents firewall licensing credits.
  get:
    description: Retrieves the firewall licensing credits with the given ID.
  delete:
    description: Deletes the firewall licensing credits with the given ID.
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
  - name: EnhancedSupportCredits
    description: Represents enhanced support credits.
    type: integer
    exposed: true
    stored: true

  - name: FWUsageCredits
    description: Represents firewall usage credits.
    type: integer
    exposed: true
    stored: true

  - name: PrismaCloudCredits
    description: Represents prisma cloud management credits.
    type: integer
    exposed: true
    stored: true

  - name: TPUsageCredits
    description: Represents threat prevention usage credits.
    type: integer
    exposed: true
    stored: true

  - name: TrafficSecuredCredits
    description: Represents traffic secured usage credits.
    type: integer
    exposed: true
    stored: true

  - name: URLFUsageCredits
    description: Represents url filtering usage credits.
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
