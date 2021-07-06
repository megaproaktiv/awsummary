package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"github.com/megaproaktiv/awsummary"
	"github.com/megaproaktiv/awsummary/stacks"
	logs "github.com/sirupsen/logrus"
)

var (
	regions = []string{
		"eu-north-1",
		"ap-south-1",
		"eu-west-3",
		"eu-west-2",
		"eu-west-1",
		"ap-northeast-2",
		"ap-northeast-1",
		"sa-east-1",
		"ca-central-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"eu-central-1",
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
	}
	verbose, daemon                                                      bool
	repeat                                                               int
	logfile                                                              string
	s3Number                                                             int
	totalEC2Number, totalEC2RunningNumber, totalEC2RunningWindowsNumber  int
	totalElbNumber, totalElbWithoutEC2Number                             int
	totalRdsNumber, totalOrRdsNumber, totalMyRdsNumber, totalMsRdsNumber int
	totalLambdaNumber                                                    int
	totalCfnNumber                                                       int
	totalESNumber                                                        int
	err                                                                  error
)

var wg sync.WaitGroup

func init() {
	flag.BoolVar(&verbose, "verbose", false, "Show detailed output")
	flag.BoolVar(&daemon, "daemon", false, "Run as daemon")
	flag.IntVar(&repeat, "repeat-every", 180, "Repeat period in seconds")
	flag.StringVar(&logfile, "log-file", "", "Log file location")
	flag.Parse()
	logs.SetFormatter(&logs.JSONFormatter{TimestampFormat: time.RFC822})
}

func main() {
	// Make sure the credentials exists
	awsummary.CheckConfig()

	// Make sure we can create log file
	awsummary.CheckLogFile(logfile)

	go func() {
		log.Print("Started")
		if logfile != "" {
			log.Print("Logs are stored in " + logfile)
		}
		for {
			for _, region := range regions {
				region := region
				wg.Add(1)
				go func() {
					defer wg.Done()
					// Getting EC2 data
					rTotal, rRunning, rWindows := awsummary.ListEC2(region, verbose)
					if rTotal > 0 {
						logs.WithFields(logs.Fields{
							"EC2":               rTotal,
							"EC2Running":        rRunning,
							"EC2RunningWindows": rWindows,
							"Region":            region,
						}).Info(awsummary.Msg("EC2"))
					}
					totalEC2Number += rTotal
					totalEC2RunningNumber += rRunning
					totalEC2RunningWindowsNumber += rWindows
				}()

				wg.Add(1)
				go func() {
					defer wg.Done()
					// Get elasticsearch data
					rEsTotal := awsummary.ListElasticsearchservice(region, verbose)
					totalESNumber += rEsTotal
				}()

				wg.Add(1)
				go func() {
					defer wg.Done()
					// Getting ELB data
					rElbTotal, rElbWithoutEC2Total := awsummary.ListElb(region, verbose)
					if rElbTotal > 0 {
						logs.WithFields(logs.Fields{
							"ELB":      rElbTotal,
							"ELBwoEC2": rElbWithoutEC2Total,
							"Region":   region,
						}).Info(awsummary.Msg("ELB"))
					}
					totalElbNumber += rElbTotal
					totalElbWithoutEC2Number += rElbWithoutEC2Total
				}()

				// Getting RDS data
				wg.Add(1)
				go func() {
					defer wg.Done()
					rRdsTotal, rRdsOTotal, rRdsMyTotal, rRdsMsTotal := awsummary.ListRds(region, verbose)
					if rRdsTotal > 0 {
						logs.WithFields(logs.Fields{
							"RDS":        rRdsTotal,
							"RDS_Oracle": rRdsOTotal,
							"RDS_MySQL":  rRdsMyTotal,
							"RDS_MSSQL":  rRdsMsTotal,
							"Region":     region,
						}).Info(awsummary.Msg("RDS"))
					}
					totalRdsNumber += rRdsTotal
					totalOrRdsNumber += rRdsOTotal
					totalMyRdsNumber += rRdsMyTotal
					totalMsRdsNumber += rRdsMsTotal
				}()

				// Getting Lambda data
				wg.Add(1)
				go func() {
					defer wg.Done()
					rLambdaTotal := awsummary.ListLambda(region, verbose)
					if rLambdaTotal > 0 {
						logs.WithFields(logs.Fields{
							"Lambda": rLambdaTotal,
							"Region": region,
						}).Info(awsummary.Msg("Lambda"))
					}
					totalLambdaNumber += rLambdaTotal
				}()

				// Getting CFN data
				wg.Add(1)
				go func() {
					defer wg.Done()
					rCfnTotal := stacks.List(region, verbose)
					if rCfnTotal > 0 {
						logs.WithFields(logs.Fields{
							"Stacks": rCfnTotal,
							"Region": region,
						}).Info(awsummary.Msg("CFN"))
					}
					totalCfnNumber += rCfnTotal
				}()
			}
			// We do not care about region here, as we will get all
			wg.Add(1)
			go func() {
				defer wg.Done()
				s3Number, err = awsummary.ListS3("eu-west-1", verbose)
				if err != nil {
					log.Fatal("Cannot get S3 data: ", err)
				}
			}()
			
			wg.Wait()
			logs.WithFields(logs.Fields{
				"S3Buckets":         s3Number,
				"EC2":               totalEC2Number,
				"EC2Running":        totalEC2RunningNumber,
				"EC2RunningWindows": totalEC2RunningWindowsNumber,
				"ELB":               totalElbNumber,
				"ELBwithoutEC2":     totalElbWithoutEC2Number,
				"ElasticsearchDomains:": totalESNumber,
				"RDS":               totalRdsNumber,
				"RDS_Oracle":        totalOrRdsNumber,
				"RDS_MySQL":         totalMyRdsNumber,
				"RDS_MSSQL":         totalMsRdsNumber,
				"Lambda":            totalLambdaNumber,
				"Stacks":            totalCfnNumber,
			}).Info("Account overview data")

			if !daemon {
				log.Print("Exiting!")
				os.Exit(0)
			}
			if verbose {
				log.Printf("Sleeping for %d", repeat)
			}
			time.Sleep(time.Duration(repeat) * time.Second)
		}
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case s := <-sig:
			log.Fatalf("Signal (%d) received, stopping", s)
		}
	}
}
