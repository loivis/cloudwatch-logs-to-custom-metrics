# cloudwatch-logs-to-custom-metrics-with-dimensions

Metric filters could be used to extract logs from CloudWatch log group and create custom metrics. It is all great, except that it's not possible to create metrics with dimension.

# MVP

Put custom metrics with multiple dimensions

## log message format:

```
metrics|namespace|metric_name|value|unit|dimensions

```

Valid units and values are as defined in [AWS API Reference](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html). Values shall be available as const in aws sdk.
