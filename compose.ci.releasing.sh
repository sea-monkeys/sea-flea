#!/bin/bash
docker compose --env-file credentials.env --env-file compose.ci.releasing.env  --file compose.ci.releasing.yml up --build

