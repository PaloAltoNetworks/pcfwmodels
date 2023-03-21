# Model
model:
  rest_name: logfieldstat
  resource_name: logfieldstats
  entity_name: LogFieldStat
  package: logging
  group: core/log
  description: Statistics about a NGFW log field.
  private: true
  get:
    description: Retrieves the statistic with the given ID.
  update:
    description: Updates the statistic with the given ID.
  delete:
    description: Deletes the statistic with the given ID.
    global_parameters:
    - $queryable
  extends:
  - '@identifiable-stored'
  - '@named'
  - '@described'
  - '@namespaced'
  - '@base'
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
  - fieldType

# Attributes
attributes:
  v1:
  - name: count
    description: A total count.
    type: integer
    exposed: true
    stored: true

  - name: fieldType
    description: The type of field to get cardinality for.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - SourceIP
    - DestinationIP
    example_value: SourceIP

  - name: firewallName
    description: The NGFW name.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: firewall1

  - name: minutes
    description: Number of minutes in the time range of the statistic.
    type: integer
    exposed: true
    stored: true

  - name: scale
    description: Ratio between top value and another.
    type: integer
    exposed: true
    stored: true

  - name: timestamp
    description: The timestamp of the statistic.
    type: time
    exposed: true
    stored: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"

  - name: top
    description: The top field value.
    type: string
    exposed: true
    stored: true
