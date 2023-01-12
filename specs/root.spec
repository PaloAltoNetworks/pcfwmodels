# Model
model:
  rest_name: root
  resource_name: root
  entity_name: Root
  package: root
  group: core
  description: root object.
  get:
    description: gets the object.
  root: true

# Relations
relations:
- rest_name: awsfirewall
  get:
    description: Retrieves the list of awsfirewalls.
    global_parameters:
    - $queryable
  create:
    description: Creates a new awsfirewall.

- rest_name: awslogdefinition
  get:
    description: Retrieves the list of awslogdefinitions.
    global_parameters:
    - $queryable
  create:
    description: Creates a new awslogdefinition.

- rest_name: deploymentadvisorsubnethelper
  create:
    description: Creates a list of free subnets per availability zone.
    global_parameters:
    - $queryable

- rest_name: deploymentadvisorterraform
  create:
    description: Creates a new terraform plan.
    global_parameters:
    - $queryable

- rest_name: firewallappidlist
  get:
    description: Retrieves a firewallappidlist.
    global_parameters:
    - $queryable

- rest_name: firewallapplicationgroup
  get:
    description: Retrieves the list of firewallapplicationgroups.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewallapplicationgroup.

- rest_name: firewallcertificate
  get:
    description: Retrieves the list of firewallcertificates.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewallcertificate.

- rest_name: firewallcidrlist
  get:
    description: Retrieves the list of firewallcidrlists.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewallcidrlist.

- rest_name: firewallcommitaction
  create:
    description: Creates a new firewallcommitaction.

- rest_name: firewallcommittedruleset
  get:
    description: Retrieves the list of firewallcommittedrulesets.
    global_parameters:
    - $queryable

- rest_name: firewallcountrycodelist
  get:
    description: Retrieves a firewallcountrycodelist.
    global_parameters:
    - $queryable

- rest_name: firewallcustomexternalfeed
  get:
    description: Retrieves the list of firewallcustomexternalfeeds.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewallcustomexternalfeed.

- rest_name: firewallcustomurlcategory
  get:
    description: Retrieves the list of firewallcustomurlcategorys.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewallcustomurlcategory.

- rest_name: firewallexternalfeedlist
  get:
    description: Retrieves a firewallexternalfeedlist.
    global_parameters:
    - $queryable

- rest_name: firewallfqdnlist
  get:
    description: Retrieves a the list of firewallfqdnlists.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewallfqdnlist.

- rest_name: firewalllicensingcredits
  get:
    description: Retrieves the list of firewall licensing credits.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewalllicensingcredits.

- rest_name: firewalllicensingmetrics
  create:
    description: Creates a new firewall licensing metrics.

- rest_name: firewallrollbackaction
  create:
    description: Creates a new firewallrollbackaction.

- rest_name: firewallrule
  get:
    description: Retrieves the list of firewallrules.
    global_parameters:
    - $queryable

- rest_name: firewallruleset
  get:
    description: Retrieves the list of firewallrulesets.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewallruleset.

- rest_name: firewallsecurityprofile
  get:
    description: Retrieves the list of firewallsecurityprofiles.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewallsecurityprofile.

- rest_name: firewalltemplate
  get:
    description: Retrieves the list of firewalltemplates.
    global_parameters:
    - $queryable
  create:
    description: Creates a new firewalltemplate.

- rest_name: firewallurlcategorylist
  get:
    description: Retrieves a firewallurlcategorylist.
    global_parameters:
    - $queryable

- rest_name: logfieldstat
  get:
    description: Retrieves statistics for a log field.
    global_parameters:
    - $timewindow
    - $queryable
  create:
    description: Creates a statistic for a log field.

- rest_name: logquery
  create:
    description: Creates a new log query.
    global_parameters:
    - $timewindow
    - $queryable

- rest_name: pcfwaccount
  get:
    description: Retrieves the status of cloud account on Prisma firewall service.
    global_parameters:
    - $queryable
  create:
    description: Onboards Cloud Account on the Prisma Prisma Cloud service.

- rest_name: pcfwtenant
  get:
    description: Retrieves the status of Prisma Cloud Tenant.
    global_parameters:
    - $queryable
  create:
    description: Onboards Prisma Cloud tenant on PC Firewall.

- rest_name: testmodel
  get:
    description: Retrieves the list of testmodels.
    global_parameters:
    - $queryable
  create:
    description: Creates a new testmodel.

- rest_name: topquery
  create:
    description: Creates a new top query.
    global_parameters:
    - $timewindow
    - $queryable

- rest_name: topsnapshot
  get:
    description: Retrieves top snapshots.
    global_parameters:
    - $timewindow
    - $queryable
  create:
    description: Creates a new top snapshot.

- rest_name: totalquery
  create:
    description: Creates a new total query.
    global_parameters:
    - $timewindow
    - $queryable

- rest_name: trendquery
  create:
    description: Creates a new trend query.
    global_parameters:
    - $timewindow
    - $queryable

- rest_name: trendsnapshot
  get:
    description: Retrieves trend snapshots.
    global_parameters:
    - $timewindow
    - $queryable
  create:
    description: Creates a new trend snapshot.
