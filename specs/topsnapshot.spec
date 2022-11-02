# Model
model:
  rest_name: topsnapshot
  resource_name: topsnapshots
  entity_name: TopSnapshot
  package: logserv
  group: core/log
  description: Snapshot of top items for a log field.
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
  - topType

# Attributes
attributes:
  v1:
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
    - SixHours
    - TwentyFourHours
    - Other
    default_value: Other

  - name: timestamp
    description: The timestamp of the snapshot.
    type: time
    exposed: true
    stored: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"

  - name: topItems
    description: The top items.
    type: refList
    exposed: true
    subtype: topitem
    stored: true
    extensions:
      refMode: pointer

  - name: topType
    description: The type of field to get top results for.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - SourceIP
    - DestinationIP
    - SourceCountry
    - DestinationCountry
    - Application
    - Attacker
    - AttackVictim
    - Attack
    - AttackType
    - AttackedApplication
    - URLCategory
    example_value: SourceIP
