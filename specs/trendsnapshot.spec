# Model
model:
  rest_name: trendsnapshot
  resource_name: trendsnapshots
  entity_name: TrendSnapshot
  package: logging
  group: core/log
  description: Snapshot of a trend item for a log field.
  private: true
  get:
    description: Retrieves the snapshot with the given ID.
  update:
    description: Updates the snapshot with the given ID.
  delete:
    description: Deletes the snapshot with the given ID.
    global_parameters:
    - $queryable
  extends:
  - '@identifiable-stored'
  - '@named'
  - '@described'
  - '@namespaced'
  - '@zoned'

# Ordering
default_order:
- :no-inherit
- timestamp

# Indexes
indexes:
- - namespace
  - timestamp
- - namespace
  - firewallName
- - namespace
  - rangeType
- - namespace
  - trendType

# Attributes
attributes:
  v1:
  - name: count
    description: The count of the given item over the given time range.
    type: integer
    exposed: true
    stored: true

  - name: firewallName
    description: The NGFW name.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: firewall1

  - name: minutes
    description: Number of minutes represented by snapshot.
    type: integer
    exposed: true
    stored: true

  - name: rangeType
    description: The time window of the snapshot.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - OneHour
    - Other
    default_value: Other

  - name: timestamp
    description: The timestamp of the snapshot.
    type: time
    exposed: true
    stored: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"

  - name: trendType
    description: The type of field for this trend.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - IngressBytes
    - EgressBytes
    - ThreatsDetected
    - ThreatsBlocked
    example_value: IngressBytes
