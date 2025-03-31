## Part 1: Installation & Configuration
### 1.2. Install AWS CDK
```
$ nvm use 20.13.1
Now using node v20.13.1 (npm v10.5.2)

$ cdk version
2.1006.0 (build a3b9762)

$ cdklocal version
2.1006.0 (build a3b9762)

$ curl http://localhost:4566/_localstack/health | jq .

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   915  100   915    0     0  39164      0 --:--:-- --:--:-- --:--:-- 39782
{
  "services": {
    "acm": "available",
    "apigateway": "available",
    "cloudformation": "running",
    "cloudwatch": "available",
    "config": "available",
    "dynamodb": "available",
    "dynamodbstreams": "available",
    "ec2": "available",
    "es": "available",
    "events": "running",
    "firehose": "available",
    "iam": "running",
    "kinesis": "available",
    "kms": "available",
    "lambda": "available",
    "logs": "available",
    "opensearch": "available",
    "redshift": "available",
    "resource-groups": "available",
    "resourcegroupstaggingapi": "available",
    "route53": "available",
    "route53resolver": "available",
    "s3": "running",
    "s3control": "available",
    "scheduler": "available",
    "secretsmanager": "available",
    "ses": "available",
    "sns": "available",
    "sqs": "running",
    "ssm": "running",
    "stepfunctions": "available",
    "sts": "running",
    "support": "available",
    "swf": "available",
    "transcribe": "available"
  },
  "edition": "community",
  "version": "4.0.4.dev105"
}
```

## Part 2: Set Environment Variables
```
$ aws sts get-caller-identity --profile localstack --no-cli-pager
{
    "UserId": "AKIAIOSFODNN7EXAMPLE",
    "Account": "000000000000",
    "Arn": "arn:aws:iam::000000000000:root"
}
```

## Part 3: Create the AWS CDK Project
### Step 1: Initialize the CDK Project

```
$ mkdir cdk-localstack-demo && cd cdk-localstack-demo                                                                                                          [11:58:31]
cdk init app --language=typescript
Applying project template app for typescript
# Welcome to your CDK TypeScript project

This is a blank project for CDK development with TypeScript.

The `cdk.json` file tells the CDK Toolkit how to execute your app.

## Useful commands

* `npm run build`   compile typescript to js
* `npm run watch`   watch for changes and compile
* `npm run test`    perform the jest unit tests
* `npx cdk deploy`  deploy this stack to your default AWS account/region
* `npx cdk diff`    compare deployed stack with current state
* `npx cdk synth`   emits the synthesized CloudFormation template

Executing npm install...
npm WARN deprecated inflight@1.0.6: This module is not supported, and leaks memory. Do not use it. Check out lru-cache if you want a good and tested way to coalesce async requests by a key value, which is much more comprehensive and powerful.
npm WARN deprecated glob@7.2.3: Glob versions prior to v9 are no longer supported
✅ All done!
```
### Step 2: Install Required CDK Packages
```
$ npm install -ES @aws-cdk/aws-s3 @aws-cdk/aws-iam                                                                                                             [11:59:42]

npm WARN deprecated @aws-cdk/cx-api@1.204.0: AWS CDK v1 has reached End-of-Support on 2023-06-01.
npm WARN deprecated This package is no longer being updated, and users should migrate to AWS CDK v2.
npm WARN deprecated
npm WARN deprecated For more information on how to migrate, see https://docs.aws.amazon.com/cdk/v2/guide/migrating-v2.html
npm WARN deprecated @aws-cdk/core@1.204.0: AWS CDK v1 has reached End-of-Support on 2023-06-01.
npm WARN deprecated This package is no longer being updated, and users should migrate to AWS CDK v2.
npm WARN deprecated
npm WARN deprecated For more information on how to migrate, see https://docs.aws.amazon.com/cdk/v2/guide/migrating-v2.html
npm WARN deprecated @aws-cdk/cloud-assembly-schema@1.204.0: AWS CDK v1 has reached End-of-Support on 2023-06-01.
npm WARN deprecated This package is no longer being updated, and users should migrate to AWS CDK v2.
npm WARN deprecated
npm WARN deprecated For more information on how to migrate, see https://docs.aws.amazon.com/cdk/v2/guide/migrating-v2.html
npm WARN deprecated @aws-cdk/cloud-assembly-schema@1.204.0: AWS CDK v1 has reached End-of-Support on 2023-06-01.
npm WARN deprecated This package is no longer being updated, and users should migrate to AWS CDK v2.
npm WARN deprecated
npm WARN deprecated For more information on how to migrate, see https://docs.aws.amazon.com/cdk/v2/guide/migrating-v2.html
npm WARN deprecated @aws-cdk/cloud-assembly-schema@1.204.0: AWS CDK v1 has reached End-of-Support on 2023-06-01.
npm WARN deprecated This package is no longer being updated, and users should migrate to AWS CDK v2.
npm WARN deprecated
npm WARN deprecated For more information on how to migrate, see https://docs.aws.amazon.com/cdk/v2/guide/migrating-v2.html
npm WARN deprecated @aws-cdk/aws-kms@1.204.0: AWS CDK v1 has reached End-of-Support on 2023-06-01.
npm WARN deprecated This package is no longer being updated, and users should migrate to AWS CDK v2.
npm WARN deprecated
npm WARN deprecated For more information on how to migrate, see https://docs.aws.amazon.com/cdk/v2/guide/migrating-v2.html
npm WARN deprecated @aws-cdk/region-info@1.204.0: AWS CDK v1 has reached End-of-Support on 2023-06-01.
npm WARN deprecated This package is no longer being updated, and users should migrate to AWS CDK v2.
npm WARN deprecated
npm WARN deprecated For more information on how to migrate, see https://docs.aws.amazon.com/cdk/v2/guide/migrating-v2.html
npm WARN deprecated @aws-cdk/aws-events@1.204.0: AWS CDK v1 has reached End-of-Support on 2023-06-01.
npm WARN deprecated This package is no longer being updated, and users should migrate to AWS CDK v2.
npm WARN deprecated
npm WARN deprecated For more information on how to migrate, see https://docs.aws.amazon.com/cdk/v2/guide/migrating-v2.html
npm WARN deprecated @aws-cdk/aws-iam@1.204.0: AWS CDK v1 has reached End-of-Support on 2023-06-01.
npm WARN deprecated This package is no longer being updated, and users should migrate to AWS CDK v2.
npm WARN deprecated
npm WARN deprecated For more information on how to migrate, see https://docs.aws.amazon.com/cdk/v2/guide/migrating-v2.html
npm WARN deprecated @aws-cdk/aws-s3@1.204.0: AWS CDK v1 has reached End-of-Support on 2023-06-01.
npm WARN deprecated This package is no longer being updated, and users should migrate to AWS CDK v2.
npm WARN deprecated
npm WARN deprecated For more information on how to migrate, see https://docs.aws.amazon.com/cdk/v2/guide/migrating-v2.html

added 16 packages, and audited 379 packages in 6s

36 packages are looking for funding
  run `npm fund` for details

8 high severity vulnerabilities

To address issues that do not require attention, run:
  npm audit fix

To address all issues (including breaking changes), run:
  npm audit fix --force

Run `npm audit` for details.
```

## Part 5: Deploy to LocalStack
### Step 1: List Available Stacks

```
$ cdklocal list
StackOne
StackTwo
```

### Step 2: Bootstrap the CDK App

```
$ AWS_PROFILE=localstack cdk bootstrap                                                                                                                         [12:08:51]
 ⏳  Bootstrapping environment aws://000000000000/us-west-2...
Trusted accounts for deployment: (none)
Trusted accounts for lookup: (none)
Using default execution policy of 'arn:aws:iam::aws:policy/AdministratorAccess'. Pass '--cloudformation-execution-policies' to customize.
 ✅  Environment aws://000000000000/us-west-2 bootstrapped (no changes).
```

### Step 3: Check Differences Before Deployment

```
$ AWS_PROFILE=localstack cdklocal diff --all                                                                                                                   [12:09:07]
start: Building 057771b8ece806fd99e3c73162c6c9ae2aff2ee2b3e966fff1bc79936d414a8e
success: Built 057771b8ece806fd99e3c73162c6c9ae2aff2ee2b3e966fff1bc79936d414a8e
start: Publishing 057771b8ece806fd99e3c73162c6c9ae2aff2ee2b3e966fff1bc79936d414a8e:current_account-us-west-2
success: Published 057771b8ece806fd99e3c73162c6c9ae2aff2ee2b3e966fff1bc79936d414a8e:current_account-us-west-2
Hold on while we create a read-only change set to get a diff with accurate replacement information (use --no-change-set to use a less accurate but faster template-only diff)

Could not create a change set, will base the diff on template differences (run again with -v to see the reason)

There were no differences

start: Building de279a1d9ceaf490596f1076b7af7f6b3be0a4ae767d0ef527da4a58e49ec055
success: Built de279a1d9ceaf490596f1076b7af7f6b3be0a4ae767d0ef527da4a58e49ec055
start: Publishing de279a1d9ceaf490596f1076b7af7f6b3be0a4ae767d0ef527da4a58e49ec055:current_account-us-west-2
success: Published de279a1d9ceaf490596f1076b7af7f6b3be0a4ae767d0ef527da4a58e49ec055:current_account-us-west-2
Hold on while we create a read-only change set to get a diff with accurate replacement information (use --no-change-set to use a less accurate but faster template-only diff)

Could not create a change set, will base the diff on template differences (run again with -v to see the reason)

There were no differences


✨  Number of stacks with differences: 0
```

### Step 4: Deploy All Stacks
```
$ AWS_PROFILE=localstack cdklocal deploy --all                                                                                                                 [12:09:40]

✨  Synthesis time: 4.34s

StackOne
StackOne: deploying... [1/2]

 ✅  StackOne (no changes)

✨  Deployment time: 0.03s

Stack ARN:
arn:aws:cloudformation:us-west-2:000000000000:stack/StackOne/c62505b6

✨  Total time: 4.37s

StackTwo
StackTwo: deploying... [2/2]

 ✅  StackTwo (no changes)

✨  Deployment time: 0.01s

Stack ARN:
arn:aws:cloudformation:us-west-2:000000000000:stack/StackTwo/783a91d5

✨  Total time: 4.36s
```

## Part 6: Verification
### 1. List Buckets

```
$ aws s3 ls --profile localstack                                                                                                                               [12:11:50]
2025-03-28 07:32:02 cdk-hnb659fds-assets-000000000000-us-west-2
2025-03-28 07:32:44 my-unique-bucket-stackone
2025-03-28 07:32:51 my-unique-bucket-stacktwo
```

### 2. Check IAM Roles
```
$ aws iam list-roles --profile localstack --no-cli-pager
{
    "Roles": [
        {
            "Path": "/",
            "RoleName": "StackTwo-MyRoleF48FFE04-ab48ed92",
            "RoleId": "AROAQAAAAAAACQQY5OMRU",
            "Arn": "arn:aws:iam::000000000000:role/StackTwo-MyRoleF48FFE04-ab48ed92",
            "CreateDate": "2025-03-28T05:32:51.822541+00:00",
            "AssumeRolePolicyDocument": {
                "Statement": [
                    {
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "s3.amazonaws.com"
                        }
                    }
                ],
                "Version": "2012-10-17"
            },
            "MaxSessionDuration": 3600
        },
        {
            "Path": "/",
            "RoleName": "cdk-hnb659fds-lookup-role-000000000000-us-west-2",
            "RoleId": "AROAQAAAAAAACR6WCNF3U",
            "Arn": "arn:aws:iam::000000000000:role/cdk-hnb659fds-lookup-role-000000000000-us-west-2",
            "CreateDate": "2025-03-28T05:32:02.927576+00:00",
            "AssumeRolePolicyDocument": {
                "Statement": [
                    {
                        "Action": "sts:TagSession",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "000000000000"
                        }
                    },
                    {
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "000000000000"
                        }
                    },
                    "__aws_no_value__"
                ]
            },
            "MaxSessionDuration": 3600,
            "Tags": [
                {
                    "Key": "aws-cdk:bootstrap-role",
                    "Value": "{'Key': 'aws-cdk:bootstrap-role', 'Value': 'lookup'}"
                }
            ]
        },
        {
            "Path": "/",
            "RoleName": "cdk-hnb659fds-image-publishing-role-000000000000-us-west-2",
            "RoleId": "AROAQAAAAAAAD5RV7L6P2",
            "Arn": "arn:aws:iam::000000000000:role/cdk-hnb659fds-image-publishing-role-000000000000-us-west-2",
            "CreateDate": "2025-03-28T05:32:02.774944+00:00",
            "AssumeRolePolicyDocument": {
                "Statement": [
                    {
                        "Action": "sts:TagSession",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "000000000000"
                        }
                    },
                    {
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "000000000000"
                        }
                    }
                ]
            },
            "MaxSessionDuration": 3600,
            "Tags": [
                {
                    "Key": "aws-cdk:bootstrap-role",
                    "Value": "{'Key': 'aws-cdk:bootstrap-role', 'Value': 'image-publishing'}"
                }
            ]
        },
        {
            "Path": "/",
            "RoleName": "StackOne-MyRoleF48FFE04-48737083",
            "RoleId": "AROAQAAAAAAAIKKP7HBZA",
            "Arn": "arn:aws:iam::000000000000:role/StackOne-MyRoleF48FFE04-48737083",
            "CreateDate": "2025-03-28T05:32:44.571339+00:00",
            "AssumeRolePolicyDocument": {
                "Statement": [
                    {
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "s3.amazonaws.com"
                        }
                    }
                ],
                "Version": "2012-10-17"
            },
            "MaxSessionDuration": 3600
        },
        {
            "Path": "/",
            "RoleName": "cdk-hnb659fds-file-publishing-role-000000000000-us-west-2",
            "RoleId": "AROAQAAAAAAAJXK3CAAIE",
            "Arn": "arn:aws:iam::000000000000:role/cdk-hnb659fds-file-publishing-role-000000000000-us-west-2",
            "CreateDate": "2025-03-28T05:32:02.143188+00:00",
            "AssumeRolePolicyDocument": {
                "Statement": [
                    {
                        "Action": "sts:TagSession",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "000000000000"
                        }
                    },
                    {
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "000000000000"
                        }
                    }
                ]
            },
            "MaxSessionDuration": 3600,
            "Tags": [
                {
                    "Key": "aws-cdk:bootstrap-role",
                    "Value": "{'Key': 'aws-cdk:bootstrap-role', 'Value': 'file-publishing'}"
                }
            ]
        },
        {
            "Path": "/",
            "RoleName": "cdk-hnb659fds-cfn-exec-role-000000000000-us-west-2",
            "RoleId": "AROAQAAAAAAAN36ZGADBY",
            "Arn": "arn:aws:iam::000000000000:role/cdk-hnb659fds-cfn-exec-role-000000000000-us-west-2",
            "CreateDate": "2025-03-28T05:32:03.017326+00:00",
            "AssumeRolePolicyDocument": {
                "Statement": [
                    {
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "cloudformation.amazonaws.com"
                        }
                    }
                ],
                "Version": "2012-10-17"
            },
            "MaxSessionDuration": 3600
        },
        {
            "Path": "/",
            "RoleName": "cdk-hnb659fds-deploy-role-000000000000-us-west-2",
            "RoleId": "AROAQAAAAAAAOLZUSPLRF",
            "Arn": "arn:aws:iam::000000000000:role/cdk-hnb659fds-deploy-role-000000000000-us-west-2",
            "CreateDate": "2025-03-28T05:32:05.436865+00:00",
            "AssumeRolePolicyDocument": {
                "Statement": [
                    {
                        "Action": "sts:TagSession",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "000000000000"
                        }
                    },
                    {
                        "Action": "sts:AssumeRole",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "000000000000"
                        }
                    }
                ]
            },
            "MaxSessionDuration": 3600,
            "Tags": [
                {
                    "Key": "aws-cdk:bootstrap-role",
                    "Value": "{'Key': 'aws-cdk:bootstrap-role', 'Value': 'deploy'}"
                }
            ]
        }
    ]
}
```
