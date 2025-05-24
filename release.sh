#!/bin/bash
set -o allexport; source compose.ci.releasing.env; set +o allexport

: <<'COMMENT'
Todo:
- 🛠️ update VERSION in compose.ci.releasing.env
- 📝 update documents and README.md if necessary
- 📦 run compose.ci.releasing.sh to generate the binaries and publish the Docker images
- 🏷️ run the current script to create a git tag and push it to the repository

Remark: delete a tag: git tag -d v0.0.1
COMMENT

echo "Generating release: ${TAG} ${ABOUT}"

find . -name '.DS_Store' -type f -delete

git add .
git commit -m "📦 ${ABOUT}"
git push

git tag -a ${TAG} -m "${ABOUT}"
git push origin ${TAG}

