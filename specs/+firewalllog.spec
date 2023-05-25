# Model
model:
  rest_name: firewalllog
  resource_name: firewalllogs
  entity_name: FirewallLog
  package: logging
  group: core/log
  description: Represents a log line in a log query result.
  detached: true

# Attributes
attributes:
  v1:
  - name: XFFIP
    description: |-
      The IP address of the user who requested the web page or the IP address of the
      next to last device that the request traversed. If the request goes through one
      or more proxies, load balancers, or other upstream devices, the firewall
      displays the IP address of the most recent device.
    type: string
    exposed: true
    read_only: true

  - name: accountID
    description: AWS account ID that generated the log.
    type: string
    exposed: true
    read_only: true

  - name: action
    description: Action taken for the session.
    type: string
    exposed: true
    read_only: true

  - name: app
    description: Application associated with the session.
    type: string
    exposed: true
    read_only: true

  - name: destinationDAG
    description: Original destination source dynamic address group.
    type: string
    exposed: true
    read_only: true

  - name: destinationIP
    description: Original session destination IP address.
    type: string
    exposed: true
    read_only: true

  - name: destinationPort
    description: Destination port utilized by the session.
    type: integer
    exposed: true
    read_only: true

  - name: firewallName
    description: Name of firewall that generated the log.
    type: string
    exposed: true
    read_only: true

  - name: protocol
    description: IP protocol associated with the session.
    type: string
    exposed: true
    read_only: true

  - name: region
    description: AWS region that generated the log.
    type: string
    exposed: true
    read_only: true

  - name: repeatCount
    description: |-
      Number of sessions with same Source IP, Destination IP, Application, and Subtype
      seen within 5 seconds.
    type: integer
    exposed: true
    read_only: true

  - name: rule
    description: Name of the rule that the session matched.
    type: string
    exposed: true
    read_only: true

  - name: sessionID
    description: An internal numerical identifier applied to each session.
    type: string
    exposed: true
    read_only: true

  - name: sourceDAG
    description: Original session source dynamic address group.
    type: string
    exposed: true
    read_only: true

  - name: sourceIP
    description: Original session source IP address.
    type: string
    exposed: true
    read_only: true

  - name: sourcePort
    description: Source port utilized by the session.
    type: integer
    exposed: true
    read_only: true

  - name: threatCategory
    description: |-
      Describes threat categories used to classify different types of threat
      signatures.
    type: string
    exposed: true
    read_only: true

  - name: threatContentName
    description: |-
      Palo Alto Networks identifier for known and custom threats. It is a description
      string followed by a 64-bit numerical identifier in parentheses for some
      Subtypes.
    type: string
    exposed: true
    read_only: true

  - name: threatContentver
    description: Applications and Threats version on your firewall when the log was
      generated.
    type: string
    exposed: true
    read_only: true

  - name: threatDataFilterReason
    description: Reason for Data Filtering action.
    type: string
    exposed: true
    read_only: true

  - name: threatDirection
    description: Indicates the direction of the attack, client-to-server or server-to-client.
    type: string
    exposed: true
    read_only: true

  - name: threatFileType
    description: File type associated with the threat.
    type: string
    exposed: true
    read_only: true

  - name: threatSeverity
    description: |-
      Severity associated with the threat; values are informational, low, medium,
      high, critical.
    type: string
    exposed: true
    read_only: true

  - name: threatSubType
    description: Subtype of threat log.
    type: string
    exposed: true
    read_only: true

  - name: timeGenerated
    description: Time the log was generated on the dataplane.
    type: time
    exposed: true
    read_only: true
    example_value: "2023-03-06T20:50:56Z"

  - name: trafficBytesReceived
    description: Number of bytes in the server-to-client direction of the session.
    type: string
    exposed: true
    read_only: true

  - name: trafficBytesSent
    description: Number of bytes in the client-to-server direction of the session.
    type: string
    exposed: true
    read_only: true

  - name: trafficCategory
    description: URL category associated with the session (if applicable).
    type: string
    exposed: true
    read_only: true

  - name: trafficDestinationCountry
    description: |-
      Destination country or Internal region for private addresses. Maximum length is
      32 bytes.
    type: string
    exposed: true
    read_only: true

  - name: trafficElapsedTime
    description: Elapsed time of the session.
    type: integer
    exposed: true
    read_only: true

  - name: trafficPacketsReceived
    description: Number of server-to-client packets for the session.
    type: string
    exposed: true
    read_only: true

  - name: trafficPacketsSent
    description: Number of client-to-server packets for the session.
    type: string
    exposed: true
    read_only: true

  - name: trafficSessionEndReason
    description: |-
      The reason a session terminated. If the termination had multiple causes, this
      field displays only the highest priority reason.
    type: string
    exposed: true
    read_only: true

  - name: trafficSourceCountry
    description: |-
      Source country or Internal region for private addresses; maximum length is 32
      bytes.
    type: string
    exposed: true
    read_only: true

  - name: trafficStartTime
    description: Time of session start.
    type: time
    exposed: true
    read_only: true
    example_value: "2023-03-06T20:50:56Z"

  - name: type
    description: Type of the log.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - Traffic
    - Threat
    - URLFiltering

  - name: urlFilteringCategories
    description: Lists the URL filtering categories that the firewall used to enforce
      policy.
    type: list
    exposed: true
    subtype: string
    read_only: true
    example_value:
    - news
    - low-risk

  - name: urlFilteringContentType
    description: Content type of the HTTP response data. Maximum length 32 bytes.
    type: string
    exposed: true
    read_only: true

  - name: urlFilteringDirection
    description: Indicates the direction of the attack, client-to-server or server-to-client.
    type: string
    exposed: true
    read_only: true

  - name: urlFilteringHTTP2Connection
    description: |-
      Identifies if traffic used an HTTP/2 connection by displaying one
      of the following values: TCP connection session ID—session is HTTP/2 0—session
      is not HTTP/2.
    type: string
    exposed: true
    read_only: true
    example_value: 22749

  - name: urlFilteringHTTPMethod
    description: |-
      Describes the HTTP Method used in the web request. Only the following
      methods are logged: Connect, Delete, Get, Head, Options, Post, Put.
    type: string
    exposed: true
    read_only: true
    example_value: get

  - name: urlFilteringReferer
    description: |-
      The Referer field in the HTTP header contains the URL of the web
      page that linked the user to another web page; it is the source that redirected
      (referred) the user to the web page that is being requested.
    type: string
    exposed: true
    read_only: true
    example_value: https://example.com/

  - name: urlFilteringURLFilename
    description: The actual URI of the request.
    type: string
    exposed: true
    read_only: true

  - name: urlFilteringURLIdx
    description: |-
      When an application uses TCP keep-alives to keep a connection open
      for a length of time, all the log entries for that session have a single session
      ID. In such cases, when you have a single threat log (and session ID) that
      includes multiple URL entries, the url_idx is a counter that allows you to
      correlate the order of each log entry within the single session.
    type: integer
    exposed: true
    read_only: true
    example_value: 1
