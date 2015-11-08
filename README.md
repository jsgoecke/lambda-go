# lambda-go

## Overview

A working example for developing [Alexa Skills Kit](http://developer.amazon.com/public/solutions/alexa/alexa-skills-kit) for your [Amazon Echo](http://amazon.com/echo) with [AWS Lambda](https://aws.amazon.com/lambda/). This example wraps a Go executable/process in an Node.js wrapper. Yes, Go works just fine on AWS.

## Requirements

1. [Golang](https://golang.org/) v1.4+
  * Linux compilation must be installed, as AWS Lambda is Linux based
2. go get github.com/jasonmoo/lambda_proc
3. [s3up](https://labix.org/s3up)
4. [awscli](https://aws.amazon.com/cli/)
  * On OSX: brew install awscli

## Installation

1. git clone https://github.com/jsgoecke/lambda-go.git
  * Change the LAMBDA_PROC variable in the file 'build.sh' to the name of your Lambda function on AWS Lambda
2. Setup an AWS Lambda function with the name set for LAMBDA_PROC in build.sh
3. Setup an AWS S3 Bucket with the same name set for LAMBDA_PROC in build.sh
4. Setup an Alex Skill with the 'Alexa Skills Settings' and pointing to your AWS Lambda function, its recommended to use [ARN](http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)

## Deploy

After configuring your s3up and awscli credentials for S3/Lambda, then:

```
./build.sh
```

### Result

```json
+ LAMBDA_FUNC=lambda-go
+ cp index.js.template index.js
+ sed s/LAMBDA_FUNC/lambda-go/g
+ GOOS=linux
+ go build -o lambda-go
+ zip -r lambda-go.zip lambda-go index.js
updating: lambda-go (deflated 70%)
updating: index.js (deflated 63%)
+ s3up lambda-go.zip lambda-go
+ aws lambda update-function-code --function-name lambda-go --s3-bucket lambda-go --s3-key lambda-go.zip
{
    "CodeSha256": "foo", 
    "FunctionName": "lambda-go", 
    "CodeSize": 905999, 
    "MemorySize": 128, 
    "FunctionArn": "bar", 
    "Version": "$LATEST", 
    "Role": "baz", 
    "Timeout": 3, 
    "LastModified": "2015-11-08T15:06:44.854+0000", 
    "Handler": "index.handler", 
    "Runtime": "nodejs", 
    "Description": "Lambda Go Example"
}
```

## Alexa Skill Settings

### Intent Schema 

```json
{
  "intents": [
    {
      "intent": "Command",
      "slots": [
        {
          "name": "Action",
          "type": "ACTION"
        },
        {
          "name": "Name",
          "type": "NAMES"
        }
      ]
    }
  ]
}
```

### Custom Slot Types

```
ACTION 	walk | talk | bike | run 	
NAMES 	fred | thelma | shaggy
```

### Sample Utterance

```
Command {Name} go {Action}
```

## Test

### Question

```
shaggy go run
```

### Request from Alexa

```json
{
  "session": {
    "sessionId": "SessionId.cb4e8a06-58ff-4d6a-befd-1e7a16be0a32",
    "application": {
      "applicationId": "foo"
    },
    "attributes": null,
    "user": {
      "userId": "amzn1.account.bar",
      "accessToken": null
    },
    "new": true
  },
  "request": {
    "type": "IntentRequest",
    "requestId": "EdwRequestId.58a4d6d4-f656-4e7e-8ea0-81ad6e66aa82",
    "timestamp": 1446994642447,
    "intent": {
      "name": "Command",
      "slots": {
        "Action": {
          "name": "Action",
          "value": "run"
        },
        "Name": {
          "name": "Name",
          "value": "shaggy"
        }
      }
    },
    "reason": null
  }
}
```

### Response to Alexa

```json
{
  "version": "1.0",
  "response": {
    "outputSpeech": {
      "type": "PlainText",
      "text": "You told shaggy to run"
    },
    "reprompt": null,
    "shouldEndSession": true
  },
  "sessionAttributes": {}
}
```