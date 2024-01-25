#!/bin/bash

# Export environment variables
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

set -a
source "${SCRIPT_DIR}/../.env"
set +a

# Run Docker Compose
docker-compose -f "${SCRIPT_DIR}/docker-compose.yml" up -d