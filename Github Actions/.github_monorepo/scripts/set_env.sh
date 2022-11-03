#!/bin/bash
echo $GITHUB_BASE_REF
echo $GITHUB_REF
if [[ $GITHUB_BASE_REF == "main" || $GITHUB_REF == "refs/heads/main" ]]; then
    echo "::set-output name=deploy_env::prod"
fi

if [[ $GITHUB_BASE_REF == "develop" || $GITHUB_REF == "refs/heads/develop" ]]; then
    echo "::set-output name=deploy_env::dev"
    echo "::set-output name=gcp_project::xperiences-eng-cn-dev"
    echo "::set-output name=gke_cluster::xpe-dev-gke"
    echo "::set-output name=gke_region::asia-east2"
fi

if [[ $GITHUB_BASE_REF == "hackathon" || $GITHUB_REF == "refs/heads/hackathon" ]]; then
    echo "::set-output name=deploy_env::dev"
    echo "::set-output name=gcp_project::xperiences-eng-cn-dev"
    echo "::set-output name=gke_cluster::xpe-dev-gke"
    echo "::set-output name=gke_region::asia-east2"
fi
