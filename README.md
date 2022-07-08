# aws-cli-mfa

## Prepare mfa settings

### ~/.aws/config

```
[my-mfa-profile]
mfa_serial = arn:aws:iam::123454567890:mfa/user-with-mfa
output = json

[profile my-profle]
user_arn = arn:aws:iam::123454567890:user/user-with-mfa
source_profile = my-mfa-profile
```


### ~/.aws/credentials

```
[my-mfa-profile]
aws_access_key_id = ASDDFDSIOFGJPERFOI
aws_secret_access_key = ldfigjpnsvifgjmsdroifgjmdsorifughjbhehe
```

## Run `aws-cli-mfa` to generate temporary credentials for `my-profile` with mfa:

```
aws-cli-mfa my-profle
Enter MFA Code: 123456
```