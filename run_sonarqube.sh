#!/bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

set -a
source "${SCRIPT_DIR}/.env"
set +a

go test -coverprofile="${SCRIPT_DIR}/docs/coverage.out" "${SCRIPT_DIR}/src/tests/..."
go test -json "${SCRIPT_DIR}/src/tests/..." > "${SCRIPT_DIR}/docs/test_report.json"
sonar-scanner -Dsonar.sources="${SCRIPT_DIR}/." -Dsonar.token=$SONARQUBE_TOKEN
