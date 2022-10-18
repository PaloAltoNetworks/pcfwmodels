# Model
model:
  rest_name: trendqueryitem
  resource_name: trendqueryitems
  entity_name: TrendQueryItem
  package: logserv
  group: core/log
  description: Represents an item in a trendquery result.
  detached: true

# Attributes
attributes:
  v1:
  - name: count
    description: The value of the subject at the given timestamp.
    type: integer
    exposed: true

  - name: minutes
    description: Number of minutes represented by the trend.
    type: integer
    exposed: true

  - name: timestamp
    description: The timestamp of the trend.
    type: time
    exposed: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
