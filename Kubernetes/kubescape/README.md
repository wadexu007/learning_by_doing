## Kubernetes Static Analyzer

Let's use an open-source tool called [kubescape](https://github.com/kubescape/kubescape) which scans yaml files for misconfigurations. Scanning your k8s manifest files is a useful and quick way to search for common vulnerabilities that may be lurking in your code.

We will find a vulnerability and fix below configuration.

```
cat sample-configmap.yaml
```
You should see the username and password values in plaintext.

Remember that ConfigMaps are used for nonsensitive information. Instead we use Secrets to pass credentials into k8s clusters.

### Installation
**On macOS**
```
brew tap kubescape/tap

brew install kubescape-cli
```

### Scan

Run the following command to scan the manifest files:
```
kubescape scan framework nsa sample-configmap.yaml
```

You should see the following results from running the scan. The scan has caught the credentials used in the ConfigMap file noted by the control `Applications credentials in configuration files`.

```
[info] Kubescape scanner starting
[info] Downloading/Loading policy definitions
[success] Downloaded/Loaded policy
[info] Accessing local objects
[success] Done accessing local objects
[info] Scanning GitLocal
[success] Done scanning GitLocal

^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Controls: 1 (Failed: 1, Excluded: 0, Skipped: 0)
Failed Resources by Severity: Critical — 0, High — 1, Medium — 0, Low — 0

+----------+-------------------------------------------------+------------------+--------------------+---------------+--------------+
| SEVERITY |                  CONTROL NAME                   | FAILED RESOURCES | EXCLUDED RESOURCES | ALL RESOURCES | % RISK-SCORE |
+----------+-------------------------------------------------+------------------+--------------------+---------------+--------------+
| High     | Applications credentials in configuration files |        1         |         0          |       1       |     100%     |
+----------+-------------------------------------------------+------------------+--------------------+---------------+--------------+
|          |                RESOURCE SUMMARY                 |        1         |         0          |       1       |   100.00%    |
+----------+-------------------------------------------------+------------------+--------------------+---------------+--------------+
FRAMEWORK NSA


~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
Scan results have not been submitted: run kubescape with the '--submit' flag
Sign up for free: https://cloud.armosec.io/account/sign-up?utm_source=GitHub&utm_medium=CLI&utm_campaign=no_submit
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
```

### More Examples
Kubescape support both Helm and Kustomize.

Example for [Kustomize manifests](https://github.com/wadexu007/learning_by_doing/tree/main/Kustomize/demo-manifests/services/demo-app/dev).
```
cd Kustomize/demo-manifests/services/demo-app/dev

kubescape scan framework nsa .
```

Scan a running Kubernetes cluster (current context)
```
kubescape scan --verbose
```

Scan a running Kubernetes cluster (current context) with [host scaning enabled](https://hub.armosec.io/docs/host-sensor?utm_source=github&utm_medium=repository)
```
kubescape scan --enable-host-scan  --verbose
```

<br>