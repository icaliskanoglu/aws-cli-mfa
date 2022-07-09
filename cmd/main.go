package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/go-ini/ini"
	log "github.com/sirupsen/logrus"
)

func exportCredentials(credentials *sts.Credentials, profile string, credentialsFile *ini.File, path string) {

	sect, err := credentialsFile.NewSection(profile)

	if err != nil {
		log.WithError(err).WithField("Section", profile).Panicf("Could not create section!")
	}

	sect.NewKey("aws_access_key_id", *credentials.AccessKeyId)
	sect.NewKey("aws_secret_access_key", *credentials.SecretAccessKey)
	sect.NewKey("aws_session_token", *credentials.SessionToken)

	err = credentialsFile.SaveTo(path)
	if err != nil {
		log.WithError(err).Panicf("Could not save credentials file!")
	}
}

func mustLoadFile(path string) *ini.File {
	confFile, err := ini.Load(path)
	if err != nil {
		log.WithError(err).WithField("Path", path).Panicf("Could not load credentials file!")
	}
	return confFile
}

func mustCreateTempCredentials(sourceProfile string, mfaArn string) *sts.Credentials {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter MFA Code: ")
	tokenString, err := reader.ReadString('\n')
	if err != nil {
		log.WithError(err).Panicf("Could not read token!")
	}
	tokenString = strings.Replace(tokenString, "\n", "", -1)

	_, err = strconv.Atoi(tokenString)
	if err != nil {
		log.WithError(err).Panicf("Could not read token!")
	}
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: sourceProfile,
	})
	if err != nil {
		log.WithError(err).WithField("Profile", sourceProfile).Panicf("Could not load profile!")
	}

	svc := sts.New(sess)
	output, err := svc.GetSessionToken(&sts.GetSessionTokenInput{
		SerialNumber: &mfaArn,
		TokenCode:    &tokenString,
	})

	if err != nil {
		log.WithError(err).WithField("Profile", sourceProfile).Panicf("Could not get session token!")
	}
	return output.Credentials
}

func main() {
	//ctx := context.TODO()

	var profile string

	flag.StringVar(&profile, "profile", "default", "aws profile")
	flag.StringVar(&profile, "p", "default", "aws profile")

	flag.Parse()

	fmt.Printf("Exporting credentials for '%s'", profile)
	usr, err := user.Current()
	if err != nil {
		log.WithError(err).Panicf("Could not get current User!")
	}
	credentialsPath := usr.HomeDir + "/.aws/credentials"
	credentialsFile := mustLoadFile(credentialsPath)

	configPath := usr.HomeDir + "/.aws/config"
	configFile := mustLoadFile(configPath)

	sourceProfile := configFile.Section("profile " + profile).Key("source_profile").Value()
	mfaArn := configFile.Section(sourceProfile).Key("mfa_serial").Value()

	tempCredentials := mustCreateTempCredentials(sourceProfile, mfaArn)

	exportCredentials(tempCredentials, profile, credentialsFile, credentialsPath)

	fmt.Printf("Temporary credentials are exported to '%s'!", credentialsPath)
}
