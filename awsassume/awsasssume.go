package awsassume

import (
	"log"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/alyu/configparser"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func randomSessionName() string {
	const sessionNameLength = 32
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, sessionNameLength)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := sessionNameLength-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

/*
ParseConfig parses the AWS config file.
*/
func ParseConfig() (config *configparser.Configuration) {
	awsConfigPath := path.Join(os.Getenv("HOME"), ".aws/config")

	// check the config file exists
	if _, err := os.Stat(awsConfigPath); err != nil {
		log.Fatal("Unable to find AWS config file at: " + awsConfigPath)
	}

	config, err := configparser.Read(awsConfigPath)

	if err != nil {
		log.Fatal(err)
	}

	return config
}

/*
GetTemporaryCredentials assumes a role and uses Amazon Security Token Service
(STS) to obtain temporary credentials for accessing the AWS account.
*/
func GetTemporaryCredentials(profile string, sessionDuration *int64, externalID *string, roleArn *string, mfaSerial *string, mfaToken *string) {
	// Setup STS connection using Go AWS SDK
	// requires: aws_access_key and aws_secret_access_key of the assumer
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	})
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	stsClient := sts.New(sess)

	// Assume role
	// requires: profile name to assume, the role arn to assume, external id,
	// the mfa serial number of the assumer, and the mfa token.
	var assumeRoleInput *sts.AssumeRoleInput
	assumeRoleInput = &sts.AssumeRoleInput{
		DurationSeconds: sessionDuration,
		ExternalId:      externalID,
		RoleArn:         roleArn,
		SerialNumber:    mfaSerial,
		TokenCode:       mfaToken,
	}

	assumeSession, err := stsClient.AssumeRole(assumeRoleInput)

	if err != nil {
	}

	// return temporary credentials:
	// session_id, sessionkey, sessiontoken
}

/*
GenerateConsoleURL generates the URL to use to access the account using the
temporary credentials provided by STS.
*/
func GenerateConsoleURL() {
}
