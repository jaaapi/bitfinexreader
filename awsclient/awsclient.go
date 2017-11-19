package awsclient

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
    "fmt"
)

func CreateClient() func {
	 sess, err = createAwsSession()
	 if err != nil {
		fmt.Println("Error creating session ", err)
		return
	}
	svc := createFirehoseService(sess)
	return svc
}

func createAwsSession() (*Session, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-1")})
	return sess, err
}

func createFirehoseService(*Session session) *Firehose {
	svc := firehose.New(sess)
	return svc
}

func PutToFirehose(svc Firehose,  data String, deliveryStreamName String) Error {
	record := &firehose.Record{Data: append(data, '\n')}
	_, err := svc.PutRecord(
		&firehose.PutRecordInput{
			DeliveryStreamName: aws.String(deliveryStreamName),
			Records:            record
		}
	)
	return err
}
