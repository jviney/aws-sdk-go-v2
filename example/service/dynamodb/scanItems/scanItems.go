// +build example

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/jviney/aws-sdk-go-v2/aws"
	"github.com/jviney/aws-sdk-go-v2/aws/external"
	"github.com/jviney/aws-sdk-go-v2/service/dynamodb"
	"github.com/jviney/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func main() {
	cfg := Config{}
	if err := cfg.Load(); err != nil {
		exitErrorf("failed to load config, %v", err)
	}

	// Create the config that the DynamoDB service will use.
	awscfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		exitErrorf("failed to load config, %v", err)
	}
	if len(cfg.Region) > 0 {
		// The Region for the DynamoDB table. If Config.Region is not set
		// the region must come from the shared config or AWS_REGION
		// environment variable.
		awscfg.Region = cfg.Region
	}

	// Create the DynamoDB service client to make the query request with.
	svc := dynamodb.New(awscfg)

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: aws.String(cfg.Table),
	}
	if cfg.Limit > 0 {
		params.Limit = aws.Int64(cfg.Limit)
	}

	// Make the DynamoDB Query API call
	scanReq := svc.ScanRequest(params)
	result, err := scanReq.Send(context.Background())
	if err != nil {
		exitErrorf("failed to make Query API call, %v", err)
	}

	items := []Item{}

	// Unmarshal the Items field in the result value to the Item Go type.
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		exitErrorf("failed to unmarshal Query result items, %v", err)
	}

	// Print out the items returned
	for i, item := range items {
		fmt.Printf("%d: Key: %d, Desc: %s\n", i, item.Key, item.Desc)
		fmt.Printf("\tNum Data Values: %d\n", len(item.Data))
		for k, v := range item.Data {
			fmt.Printf("\t- %q: %v\n", k, v)
		}
	}
}

type Item struct {
	Key  int
	Desc string
	Data map[string]interface{}
}

type Config struct {
	Table  string // required
	Region string // optional
	Limit  int64  // optional

}

func (c *Config) Load() error {
	flag.Int64Var(&c.Limit, "limit", 0, "Limit is the max items to be returned, 0 is no limit")
	flag.StringVar(&c.Table, "table", "", "Table to Query on")
	flag.StringVar(&c.Region, "region", "", "AWS Region the table is in")
	flag.Parse()

	if len(c.Table) == 0 {
		flag.PrintDefaults()
		return fmt.Errorf("table name is required.")
	}

	return nil
}
