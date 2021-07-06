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

Thank you for original script by partamonov.
