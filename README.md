[![Build Status](https://semaphoreci.com/api/v1/partamonov/aws-overview/branches/master/badge.svg)](https://semaphoreci.com/partamonov/aws-overview)

# aws-overview

``aws-overview`` is a small script to get AWS account overview.
Available details are (number of):
* EC2, EC2 Running instances, EC2 Windows running instances
* ELB, ELB without assigned EC2 instances
* Elasticsearch Domains
* RDS, RDS MySQL/MSSQL/Oracle
* CFN
* Lambda functions
* S3 buckets
* Total of all above in all regions

## Installing

* ``go get github.com/tecracer/aws-overview``
* ``go install github.com/tecracer/aws-overview``

For cross platform compilation:
* ``env GOOS=windows GOARCH=amd64 go build``

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
