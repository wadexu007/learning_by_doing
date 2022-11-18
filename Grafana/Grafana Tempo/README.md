## What is this folder?
The folder contains all resources to show how to automate instrument Golang app (OpenTelemetry SDK) to send `Trace` data --> Grafanan agent --> Grafana Tempo in Grafana Cloud.

## About Grafana Tempo
Grafana Tempo is an open source, easy-to-use, and high-scale distributed tracing backend. Tempo is cost-efficient, requiring only object storage to operate, and is deeply integrated with Grafana, Prometheus, and Loki. Tempo can ingest common open source tracing protocols, including Jaeger, Zipkin and OpenTelemetry.

## Sending Data to Tempo
In order to send spans to Tempo we will be using the [Grafana Agent](https://github.com/grafana/agent).

### Install Grafana agent
Manifests: Trace collection (Deployment): [agent-trace.yaml](https://github.com/grafana/agent/blob/main/production/kubernetes/agent-traces.yaml)

Replace ${NAMESPACE} in [yaml](agent-trace.yaml) to yours if you need. 
```
kubectl create ns trace

kubectl apply -f agent-trace.yaml
```

## Configure Grafana Agent
Replace `YOUR_TEMPO_ENDPOINT`, `YOUR_TEMPO_USER` and `YOUR_TEMPO_PASSWORD` in [configmap.yaml](./configmap.yaml)

The Agent supports consuming traces in a number of common formats. The following example config could be used to start all possible receivers with default ports. It is recommended to only start the ones you need.
```
receivers:
  jaeger:
    protocols:
      grpc:
      thrift_binary:
      thrift_compact:
      thrift_http:
  zipkin:
  otlp:
    protocols:
      http:
      grpc:
  opencensus:
```

Here we only choose OpenTelemetry `otlp` for receivers.

```
kubectl -n trace apply -f configmap.yaml
```

## Deploy a demo App
Send traces to Grafana Cloud's Tempo service with OpenTelemetry SDK in your demo app code (Golang).

Refer to [instrument your golang code](https://opentelemetry.io/docs/instrumentation/go/getting-started/#trace-instrumentation) to write your golang app.

Then containerized it with docker and deploy to k8s, my k8s manifests [deployment yaml](./demo-app.yaml) and [source code](../../Golang/demo_app_with_instrumentation/) for reference, replace `OTEL_EXPORTER_OTLP_ENDPOINT` if need, my grafana agent is `grafana-agent-traces.trace.svc.cluster.local:4317`. 
```
kubectl apply -f demo-app.yaml                
```

Kubectl port-forwarding can be used to connect to demo app API without exposing the service.
```
kubectl port-forward svc/bookstore -n exercise 8080:8080
```

Call demo app API
```
curl -X POST 'http://localhost:8080/books' -d '{"title":"Cloud Native","author":"wadexu"}' | jq
curl -X POST 'http://localhost:8080/books' -d '{"title":"Linux","author":"wadexu"}' | jq
curl -X GET http://localhost:8080/books | jq
```

## View Trace
Search for the trace in Grafana Cloud by navigating to Explore and choosing your traces data source.

![alt text](../Images/grafana_cloud_trace.jpg "This is Trace data in Grafana Cloud - Tempo.")

## Conclusion
OpenTelemetry is the future for setting up observability for cloud-native apps. It is backed by a huge community and covers a wide variety of technology and frameworks. Using OpenTelemetry, engineering teams can instrument polyglot and distributed applications with peace of mind.

## Userful Links
https://uptrace.dev/opentelemetry/go-tracing.html

<br>
