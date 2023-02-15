# Model
model:
  rest_name: topitem
  resource_name: topitems
  entity_name: TopItem
  package: logging
  group: core/log
  description: A field value and its count.
  private: true
  detached: true

# Attributes
attributes:
  v1:
  - name: count
    description: The count of the specified item.
    type: integer
    exposed: true
    stored: true

  - name: value
    description: The value of the log field.
    type: string
    exposed: true
    stored: true
