## Emissary Config Helm Chart

This folder contains two custom Helm charts for emissary configuration objects
* emissary-config
* emissary-crds

#### Prerequisites
* Helm 3.2.0+

### emissary-config
This chart is for quick install Emissary Hosts/Mapping/Listeners and more configuration objects.

CRDs in templates folder
* hosts.yaml
* listeners.yaml
* mappings.yaml
* tlscontext.yaml

### emissary-crds
This chart is for quick install Emissary CRDs.

CRDs in templates folder
* [getambassador.io_crds.yaml](https://app.getambassador.io/yaml/emissary/3.2.0/emissary-crds.yaml)

### Test
Set values in values.yaml to generate all templates with variables and show the output YAML
```
helm template emissary-config
```

### Package
```
helm package emissary-config

# Successfully packaged chart and saved it to: /Users/wadexu/Downloads/learning_by_doing/Helm/custom-charts/emissary-custom-charts/emissary-config-8.2.0.tgz


helm package emissary-crds

# Successfully packaged chart and saved it to: /Users/wadexu/Downloads/learning_by_doing/Helm/custom-charts/emissary-custom-charts/emissary-crds-8.2.0.tgz
```

### Share Helm Package
* Upload to a bucket
* [Publish to a github repo](https://medium.com/containerum/how-to-make-and-share-your-own-helm-package-50ae40f6c221)

### Example Usage
[Deploy Helm Charts With Terraform](../../../Emissary/terraform_helm_install/dev/emissary.tf)
