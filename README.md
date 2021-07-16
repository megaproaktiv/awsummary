# awsummary

`awsummary` is a script to get a fast (< 3 seconds) AWS account overview of pricey or often used services in all 13 regions.

Available details are (number of):
* EC2, EC2 Running instances, EC2 Windows running instances
* Application LB, Gateway LB, Network LB
* Elasticsearch Domains
* RDS, RDS MySQL/MSSQL/Oracle
* dynamodb reserved (pay while idle!) and on demand capacity
* CFN
* Lambda functions
* S3 buckets
* Total of all above in all regions



## Usage

AWS Credentials expected in ``$HOME/.aws/credentials`` or as environment variables
``AWS_ACCESS_KEY_ID``
``AWS_SECRET_ACCESS_KEY``

```
Usage: aws-overview [-h] [-log-file=path] [-daemon=true] [-repeat-every=60] [other options]
 -h, --help
 -daemon=false/true:           [bool], Run as daemon. Default value is 'false'
 -repeat-every=<INT>:          Repeat period in seconds. Used only in daemon mode. Default value is 180
 -log-file=<PATH>:             Log file location, if skipped logs to STDOUT
 -verbose=true/false:          [bool], if true prints details information about objects
 -machine-readable=true/false: [bool], if true convert output to Logstash format, false print json output
```

## Usage Examples

### Watch out for DynamoDb with reserved capacity 

```bash
2021/07/06 18:03:31 Started
{"Region":"eu-central-1","Stacks":17,"level":"info","msg":"Region specific CFN data","time":"06 Jul 21 18:03 CEST"}
{"Region":"eu-west-1","Stacks":8,"level":"info","msg":"Region specific CFN data","time":"06 Jul 21 18:03 CEST"}
{"Lambda":12,"Region":"eu-west-1","level":"info","msg":"Region specific Lambda data","time":"06 Jul 21 18:03 CEST"}
{"Lambda":4,"Region":"us-east-1","level":"info","msg":"Region specific Lambda data","time":"06 Jul 21 18:03 CEST"}
{"OnDemandCapacity":8,"Region":"eu-central-1","ReservedCapacity":1,"Tables":9,"level":"info","msg":"Region specific DDB data","time":"06 Jul 21 18:03 CEST"}
{"Region":"us-west-2","Stacks":1,"level":"info","msg":"Region specific CFN data","time":"06 Jul 21 18:03 CEST"}
{"Lambda":1,"Region":"us-west-2","level":"info","msg":"Region specific Lambda data","time":"06 Jul 21 18:03 CEST"}
{"ALB":0,"DynamoDB":9,"EC2":0,"EC2Running":0,"EC2RunningWindows":0,"ElasticsearchDomains:":0,"GLB":0,"Lambda":17,"NLB":0,"RDS":0,"RDS_MSSQL":0,"RDS_MySQL":0,"RDS_Oracle":0,"S3Buckets":54,"Stacks":26,"level":"info","msg":"Account overview data","time":"06 Jul 21 18:03 CEST"}
2021/07/06 18:03:32 Exiting!
````

### Pricey MSSQL and Running and stopped instances

```bash
{"Region":"eu-central-1","Stacks":19,"level":"info","msg":"Region specific CFN data","time":"06 Jul 21 18:20 CEST"}
{"ALB":2,"GLB":0,"NLB":2,"Region":"eu-central-1","level":"info","msg":"Region specific ELB data","time":"06 Jul 21 18:20 CEST"}
{"RDS":2,"RDS_MSSQL":2,"RDS_MySQL":0,"RDS_Oracle":0,"Region":"eu-central-1","level":"info","msg":"Region specific RDS data","time":"06 Jul 21 18:20 CEST"}
{"EC2":7,"EC2Running":5,"EC2RunningWindows":3,"Region":"eu-central-1","level":"info","msg":"Region specific EC2 data","time":"06 Jul 21 18:20 CEST"}
{"ALB":2,"DynamoDB":0,"EC2":7,"EC2Running":5,"EC2RunningWindows":3,"ElasticsearchDomains:":0,"GLB":0,"Lambda":0,"NLB":2,"RDS":2,"RDS_MSSQL":2,"RDS_MySQL":0,"RDS_Oracle":0,"S3Buckets":8,"Stacks":19,"level":"info","msg":"Account overview data","time":"06 Jul 21 18:20 CEST"}
2021/07/06 18:20:33 Exiting!
```



## Thanks

Thank you for original script by partamonov.
