# Model
model:
  rest_name: firewalllog
  resource_name: firewalllogs
  entity_name: FirewallLog
  package: logging
  group: core/log
  description: Represents a log line in a logquery result.
  detached: true

# Attributes
attributes:
  v1:
  - name: accountID
    description: AWS account id that generated the log.
    type: string
    exposed: true

  - name: action
    description: Action the NGFW took.
    type: string
    exposed: true

  - name: app
    description: Application.
    type: string
    exposed: true

  - name: destinationIP
    description: Destination ip.
    type: string
    exposed: true

  - name: destinationPort
    description: Destination port.
    type: integer
    exposed: true

  - name: firewallName
    description: Name of firewall that generated the log.
    type: string
    exposed: true

  - name: protocol
    description: Network protocol.
    type: string
    exposed: true

  - name: region
    description: AWS region that generated the log.
    type: string
    exposed: true

  - name: repeatCount
    description: Repeat count.
    type: integer
    exposed: true

  - name: rule
    description: NGFW rule.
    type: string
    exposed: true

  - name: sessionID
    description: Session id.
    type: string
    exposed: true

  - name: sourceIP
    description: Source ip.
    type: string
    exposed: true

  - name: sourcePort
    description: Source port.
    type: integer
    exposed: true

  - name: threatCategory
    description: Threat category.
    type: string
    exposed: true

  - name: threatContentName
    description: Threat content name.
    type: string
    exposed: true

  - name: threatContentver
    description: Contentver.
    type: string
    exposed: true

  - name: threatDataFilterReason
    description: Data filter reason.
    type: string
    exposed: true

  - name: threatDirection
    description: Traffic direction.
    type: string
    exposed: true

  - name: threatFileType
    description: Filetype.
    type: string
    exposed: true

  - name: threatSeverity
    description: Severity.
    type: string
    exposed: true

  - name: threatSubType
    description: Sub type.
    type: string
    exposed: true

  - name: timeGenerated
    description: Time generated.
    type: time
    exposed: true
    example_value: "2023-03-06T20:50:56Z"

  - name: trafficBytesReceived
    description: Bytes received.
    type: string
    exposed: true

  - name: trafficBytesSent
    description: Bytes sent.
    type: string
    exposed: true

  - name: trafficCategory
    description: Category.
    type: string
    exposed: true

  - name: trafficDestinationCountry
    description: Traffic destination country.
    type: string
    exposed: true

  - name: trafficElapsedTime
    description: Elapsed time.
    type: integer
    exposed: true

  - name: trafficPacketsReceived
    description: Bytes received.
    type: string
    exposed: true

  - name: trafficPacketsSent
    description: Bytes received.
    type: string
    exposed: true

  - name: trafficSessionEndReason
    description: Session end reason.
    type: string
    exposed: true

  - name: trafficSourceCountry
    description: Traffic source country.
    type: string
    exposed: true

  - name: trafficStartTime
    description: Start time.
    type: time
    exposed: true
    example_value: "2023-03-06T20:50:56Z"

  - name: type
    description: Type of the log.
    type: enum
    exposed: true
    allowed_choices:
    - Traffic
    - Threat
    - URLFiltering

  - name: urlFilteringCategory
    description: Url filtering Category.
    type: string
    exposed: true

  - name: urlFilteringStartTime
    description: Start time.
    type: time
    exposed: true
    example_value: "2023-03-06T20:50:56Z"

  - name: xffIP
    description: Xff ip.
    type: string
    exposed: true
