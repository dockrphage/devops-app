#!/bin/bash
# This script updates the GitOps repo with the new image

set -e

# Variables passed from GitHub Actions
COMMIT_SHA=${1:-"latest"}
RUN_NUMBER=${2:-"1"}

echo "🔄 Updating GitOps repo with image: $COMMIT_SHA"

# Clone the repo
git clone https://github.com/dockrphage/devops-gitops.git
cd devops-gitops

# Update the deployment file directly
sed -i "s|image: dockrphage/devops-app:.*|image: dockrphage/devops-app:${COMMIT_SHA}|g" app/base/deployment.yaml
sed -i "s|value: \"v.*\"|value: \"v${RUN_NUMBER}\"|g" app/base/deployment.yaml

# Update kustomization
sed -i "s|image: dockrphage/devops-app:.*|image: dockrphage/devops-app:${COMMIT_SHA}|g" app/overlays/dev/kustomization.yaml

# Commit and push
git config user.name "GitHub Actions"
git config user.email "actions@github.com"
git add .
git commit -m "Update image to commit ${COMMIT_SHA} [skip ci]"
git push

echo "✅ GitOps repo updated successfully!"
