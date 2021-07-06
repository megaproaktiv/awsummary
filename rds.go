package awsummary

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"log"
	"strings"
)

func ListRds(region string, verbose bool) (rdsNumber, rdsOracleNumber, rdsMysqlNumber, rdsMSsqlNumber int) {
	svc := rds.New(session.New(&aws.Config{Region: aws.String(region)}))

	params := &rds.DescribeDBInstancesInput{
		MaxRecords: aws.Int64(100),
	}

	resp, err := svc.DescribeDBInstances(params)
	if err != nil {
		log.Fatal("Cannot get RDS data: ", err)
	}

	for _, name := range resp.DBInstances {
		switch {
		case strings.Contains(*name.Engine, "oracle"):
			rdsOracleNumber++
		case strings.Contains(*name.Engine, "mysql"):
			rdsMysqlNumber++
		case strings.Contains(*name.Engine, "sqlserver"):
			rdsMSsqlNumber++
		}
		if verbose {
			log.Println("RDS Name: ", *name.DBInstanceIdentifier, " created: ", *name.InstanceCreateTime)
			log.Println("RDS Size: ", *name.DBInstanceClass)
			log.Println("RDS Engine: ", *name.Engine)
		}
	}
	rdsNumber = len(resp.DBInstances)
	return rdsNumber, rdsOracleNumber, rdsMysqlNumber, rdsMSsqlNumber
}
