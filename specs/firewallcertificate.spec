# Model
model:
  rest_name: firewallcertificate
  resource_name: firewallcertificates
  entity_name: FirewallCertificate
  package: ngfw
  group: core/ngfw
  description: Represents a cloud certificate.
  get:
    description: Retrieves the certificate with the given ID.
  update:
    description: Updates the certificate with the given ID.
  delete:
    description: Deletes the certificate with the given ID.
  extends:
  - '@identifiable-stored'
  - '@named'
  - '@described'
  - '@namespaced'
  - '@base'
  - '@zoned'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: certificate
    description: Depends on the Type. For AWS, it is the AWS Certificate ARN.
    type: string
    exposed: true
    stored: true

  - name: selfSigned
    description: A self-signed root CA certificate.
    type: boolean
    exposed: true
    stored: true
