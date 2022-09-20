//go:build go1.9
// +build go1.9

package endpoints_test

import (
	"fmt"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/endpoints"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
)

// ***************************************************************************
// All endpoint metadata is sourced from the testdata/endpoints.json file at
// test startup. Not the live endpoints model file. Update the testdata file
// for the tests to use the latest live model.
// ***************************************************************************

func ExampleEnumPartitions() {
	resolver := endpoints.DefaultResolver()
	partitions := resolver.(endpoints.EnumPartitions).Partitions()

	for _, p := range partitions {
		fmt.Println("Regions for", p.ID())
		for id := range p.Regions() {
			fmt.Println("*", id)
		}

		fmt.Println("Services for", p.ID())
		for id := range p.Services() {
			fmt.Println("*", id)
		}
	}
}

func ExampleResolverFunc() {
	myCustomResolver := func(service, region string, optFns ...func(*endpoints.Options)) (
		endpoints.ResolvedEndpoint, error,
	) {
		if service == endpoints.S3ServiceID {
			return endpoints.ResolvedEndpoint{
				URL:           "s3.custom.endpoint.com",
				SigningRegion: "custom-signing-region",
			}, nil
		}

		return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String("us-west-2"),
		EndpointResolver: endpoints.ResolverFunc(myCustomResolver),
	}))

	// Create the S3 service client with the shared session. This will
	// automatically use the S3 custom endpoint configured in the custom
	// endpoint resolver wrapping the default endpoint resolver.
	s3Svc := s3.New(sess)
	// Operation calls will be made to the custom endpoint.
	s3Svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("myBucket"),
		Key:    aws.String("myObjectKey"),
	})
}
