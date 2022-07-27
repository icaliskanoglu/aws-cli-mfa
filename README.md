# aws-cli-mfa

`aws-cli-mfa` is an open-source, command line application for export temproray credentials to ~/.aws/credentials file. I implemented it with `golang` and it supports multiple platform. You can export temproray credentials and interact aws resources easly.

### Configuring `aws-cli-mfa`

When you enable MFA for an IAM user you can see the MFA serial on console. We need to put that informataion to our aws config file additionally the credentials. Here is the sample configuration: 

#### Installation

You can download the latest prebuild binary from the releases page. [Releases](https://github.com/icaliskanoglu/aws-cli-mfa/releases)

You can clone the go project from the [repository](https://github.com/icaliskanoglu/aws-cli-mfa) and build yourself.

Run these commands to install manually:
```
export VERSION=1.0.0
export PLATFORM=darwin-amd64
wget https://github.com/icaliskanoglu/aws-cli-mfa/releases/download/$VERSION/aws-cli-mfa-$VERSION-$PLATFORM.tar.gz
tar xzvf aws-cli-mfa-$VERSION-$PLATFORM.tar.gz
sudo mv aws-cli-mfa /usr/local/bin/aws-cli-mfa
```

#### ~/.aws/config

```
[my-mfa-profile]
mfa_serial = arn:aws:iam::123454567890:mfa/user-with-mfa
output = json

[profile my-profle]
user_arn = arn:aws:iam::123454567890:user/user-with-mfa
source_profile = my-mfa-profile
```

#### ~/.aws/credentials

```
[my-mfa-profile]
aws_access_key_id = ASDDFDSIOFGJPERFOI
aws_secret_access_key = ldfigjpnsvifgjmsdroifgjmdsorifughjbhehe
```

### Run `aws-cli-mfa` to generate temporary credentials for `my-profile` with mfa:

We need to get MFA code from the selected device. It will export the temprory credentials to `~/.aws/credentials` file.

```
aws-cli-mfa -p my-profle
Exporting credentials for 'my-profle' profile
Enter MFA Code: 123456
Temporary credentials are exported to '/Users/icaliskanoglu/.aws/credentials'.
```

After see the success message, you can see the AWS credential in the `~/.aws/credentials` file.

```
...
[my-profle]
aws_access_key_id = example-access-key-as-in-returned-output
aws_secret_access_key = example-secret-access-key-as-in-returned-output
aws_session_token = example-session-Token-as-in-returned-output
```
