# Model
model:
  rest_name: logincident
  resource_name: logincidents
  entity_name: LogIncident
  package: logging
  group: core/log
  description: Represents a security incident.
  detached: true

# Attributes
attributes:
  v1:
  - name: attackerAlternateIPs
    description: Other IP addresses for the attacker instance.
    type: list
    exposed: true
    subtype: string

  - name: attackerID
    description: The attacker ID (EC2 instance ID).
    type: string
    exposed: true
    example_value: i-450224347700

  - name: attackerIP
    description: The attacker IP address.
    type: string
    exposed: true
    example_value: 212.44.7.91

  - name: attackerName
    description: The attacker name.
    type: string
    exposed: true

  - name: attackerPort
    description: The attacker port.
    type: integer
    exposed: true

  - name: firewallName
    description: The NGFW name.
    type: string
    exposed: true

  - name: threatCategory
    description: The threat category.
    type: string
    exposed: true

  - name: threatName
    description: The threat name.
    type: string
    exposed: true

  - name: timestamp
    description: The timestamp of the incident.
    type: time
    exposed: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"

  - name: victimAlternateIPs
    description: Other IP addresses for the victim instance.
    type: list
    exposed: true
    subtype: string

  - name: victimID
    description: The victim ID (EC2 instance ID).
    type: string
    exposed: true
    example_value: i-450224347700

  - name: victimIP
    description: The victim IP address.
    type: string
    exposed: true
    example_value: 212.3.41.8

  - name: victimName
    description: The victim name.
    type: string
    exposed: true

  - name: victimPort
    description: The victim port.
    type: integer
    exposed: true
    example_value: 80
