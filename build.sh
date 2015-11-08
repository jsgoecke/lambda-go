#!/bin/bash

set -e
set -x

LAMBDA_FUNC="lambda-go"

sed "s/LAMBDA_FUNC/$LAMBDA_FUNC/g" <index.js.template >index.js

GOOS=linux go build -o $LAMBDA_FUNC
zip -r $LAMBDA_FUNC.zip $LAMBDA_FUNC index.js
s3up $LAMBDA_FUNC.zip lambda-go
aws lambda update-function-code --function-name $LAMBDA_FUNC --s3-bucket $LAMBDA_FUNC --s3-key $LAMBDA_FUNC.zip

