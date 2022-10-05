## What's in this folder?

This folder contains atlantis mainfests for deployment in kubernetes.

Reference: https://www.runatlantis.io/docs/deployment.html#kubernetes-kustomize

####

echo -n "" > ghUser
echo -n "" > ghToken
echo -n "" > ghWebhookSecret

kubectl create secret generic atlantis --from-file=ghUser --from-file=ghToken --from-file=ghWebhookSecret
