## Loki
[Grafana Loki](https://grafana.com/docs/loki/latest/) is a set of components that can be composed into a fully featured logging stack.

[Promtail](https://grafana.com/docs/loki/latest/clients/promtail/) is an agent which ships the contents of local logs to a private Grafana Loki instance or Grafana Cloud. It is usually deployed to every machine that has applications needed to be monitored.

## Install Promtail
* Install agent [Promtail](https://grafana.com/docs/loki/latest/clients/promtail/) for a kubernetes cluster

Replace <User ID> and <Your Grafana.com API Key> with a Grafana.com API key with the MetricsPublisher role.

```
kubectl create ns loki

sh promtail.sh <User ID> <API Key> logs-prod3.grafana.net loki | kubectl apply --namespace=loki -f  -
```

now logs start to sending to Loki services in Grafana Cloud

## Check Result
Go back to Grafana Cloud Console to Manage your Grafana Cloud Stack.

Click `Launch` Grafana, go to `Explore` select datasource (Grafana logs datasource automatical added to Grafana Cloud)

![alt text.](../Images/grafana_cloud_logs.jpg "This is test result image.")

## Clean Up
```
kubectl delete ds promtail -n loki

daemonset.apps "promtail" deleted
```