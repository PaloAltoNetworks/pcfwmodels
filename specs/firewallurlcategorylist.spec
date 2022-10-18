# Model
model:
  rest_name: firewallurlcategorylist
  resource_name: firewallurlcategorylists
  entity_name: FirewallURLCategoryList
  package: ngfw
  group: core/ngfw
  description: This a readonly list that returns the default NGFW URLCategories.
  extends:
  - '@namespaced'

# Attributes
attributes:
  v1:
  - name: categories
    description: A list of NGFW URLCategory objects.
    type: refList
    exposed: true
    subtype: firewallurlcategory
    read_only: true
    extensions:
      refMode: pointer
