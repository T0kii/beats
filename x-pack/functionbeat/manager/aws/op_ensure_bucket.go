// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package aws

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/smithy-go"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/T0kii/beats/v7/x-pack/functionbeat/manager/executor"
	"github.com/elastic/elastic-agent-libs/logp"
)

// This error is not provided by the S3 error package.
const notFound = "NotFound"

type opEnsureBucket struct {
	log        *logp.Logger
	svc        *s3.Client
	bucketName string
}

func newOpEnsureBucket(log *logp.Logger, cfg aws.Config, bucketName string) *opEnsureBucket {
	return &opEnsureBucket{log: log, svc: s3.NewFromConfig(cfg), bucketName: bucketName}
}

func (o *opEnsureBucket) Execute(_ executor.Context) error {
	o.log.Debugf("Verifying presence of S3 bucket: %s", o.bucketName)

	check := &s3.HeadBucketInput{Bucket: aws.String(o.bucketName)}
	_, err := o.svc.HeadBucket(context.TODO(), check)
	if err == nil {
		// The bucket exists and we have permission to access it.
		return nil
	}

	var apiError smithy.APIError
	if errors.As(err, &apiError) {
		if apiError.ErrorCode() == notFound {
			// bucket do not exist let's create it.
			input := &s3.CreateBucketInput{Bucket: aws.String(o.bucketName)}
			resp, err := o.svc.CreateBucket(context.TODO(), input)
			if err != nil {
				o.log.Debugf("Could not create bucket, resp: %v", resp)
				return err
			}
			// bucket created successfully
			return nil
		}
	}

	// Catchall for unauthorized access.
	return fmt.Errorf("bucket '%s' already exist and you don't have permission to access it: %w", o.bucketName, err)
}
