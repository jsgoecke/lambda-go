#!/bin/bash

set -e
set -x

LAMBDA_FUNC="lambda-go"

sed "s/LAMBDA_FUNC/$LAMBDA_FUNC/g" <index.js.template >index.js

GOOS=linux go build -o $LAMBDA_FUNC
zip -r $LAMBDA_FUNC.zip $LAMBDA_FUNC index.js
aws lambda update-function-code --function-name $LAMBDA_FUNC --zip-file fileb://$LAMBDA_FUNC.zip
