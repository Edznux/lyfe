#!/usr/bin/env bash

# Generate the Strava client
docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate -i https://developers.strava.com/swagger/swagger.json -l go -o /local/strava-client/swagger