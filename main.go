package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	awsevents "github.com/aws/aws-lambda-go/events"
	awslambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	awssession "github.com/aws/aws-sdk-go/aws/session"
	awscloudwatch "github.com/aws/aws-sdk-go/service/cloudwatch"
)

var (
	client *awscloudwatch.CloudWatch
)

func init() {
	client = awscloudwatch.New(awssession.Must(awssession.NewSession()))
}

func main() {
	awslambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, event awsevents.CloudwatchLogsEvent) error {
	fmt.Println(event.AWSLogs.Data)
	data, err := event.AWSLogs.Parse()
	if err != nil {
		return err
	}

	var errs []string

	for i := range data.LogEvents {
		metrics := convert(&data.LogEvents[i])

		in := &awscloudwatch.PutMetricDataInput{
			Namespace:  aws.String(data.LogGroup),
			MetricData: metrics,
		}
		_, err = client.PutMetricData(in)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) != 0 {
		return errors.New(strings.Join(errs, ","))
	}

	return nil
}

func convert(event *events.CloudwatchLogsLogEvent) (metrics []*awscloudwatch.MetricDatum) {
	// REPORT RequestId: 654e4b71-ddaf-4ce3-be71-abfef7187d4f Duration: 241.77 ms Billed Duration: 300 ms Memory Size: 512 MB Max Memory Used: 45 MB
	switch {
	case strings.HasPrefix(event.Message, "REPORT RequestId:"):
	case true:
		// metrics|namespace|metric_name|value|unit|dimensions
	}

	return nil
}

// namespace = "test-namespace"
// metrics = []*awscloudwatch.MetricDatum{
// 	{
// 		Dimensions: []*awscloudwatch.Dimension{
// 			{Name: aws.String("dimension1-name"), Value: aws.String("dimension1-value")},
// 			{Name: aws.String("dimension2-name"), Value: aws.String("dimension2-value")},
// 		},
// 		MetricName: aws.String("test-metric"),
// 		Value:      aws.Float64(1),
// 		Unit:       aws.String(awscloudwatch.StandardUnitCount),
// 	},
// }
