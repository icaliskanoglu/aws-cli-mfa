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

## Run script to generate temproray credentials for my-profle with mfa:

```
python refresh_credentials.py my-profle
OTP from device: 123456
```


[Original Post: Multi-factor Authentication (MFA) for AWS CLI](https://www.redpill-linpro.com/techblog/2020/02/18/awscli.html)
