## Datadog Application Performance Monitoring (APM) 

## Sending Traces to Datadog

### Java
Dockerfile
```
...

ENV DD_ENV xxx
ENV SERVICE_ENV xxx

...

RUN mkdir -p /dd-java-agent

...

COPY ./dd-java-agent.jar /dd-java-agent/dd-java-agent.jar


CMD java -javaagent:/dd-java-agent/dd-java-agent.jar -Ddd.profiling.enabled=true -Ddd.logs.injection=true \
    -Ddd.trace.analytics.enabled=true -Ddd.service=msg-api -Ddd.env=$DD_ENV -Ddd.agent.host=172.17.0.1 \
    -Ddd.tags=env:$SERVICE_ENV -jar audiencecuration.jar
```


### Nodejs
https://docs.datadoghq.com/tracing/setup/nodejs/

Dockerfile
```
...

ENV DD_ENV xxx
ENV DD_SERVICE xxx
ENV DD_TAGS env:xxx

...
```

server/server.js
```
...
const tracer = require('dd-trace').init();
...
const app = express();
```

package.json
```
    "dd-trace": "0.26.0"
```

### Golang
https://docs.datadoghq.com/tracing/setup/go/

```
```

## Datadog Agent Configuration
* Datadog agent in VM Instannce 
Datadog.yaml
```
apm_config:
  ## @param enabled - boolean - optional - default: true
  ## Set to true to enable the APM Agent.
  #
  enabled: true
  
  apm_non_local_traffic: true
```

* Datadog agent in Kubernetes
https://docs.datadoghq.com/agent/kubernetes/apm/?tab=helm


##  View Result
Datadog Console -> APM -> Services