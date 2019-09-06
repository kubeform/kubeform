## NodebalancerConfig
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `linode.kubeform.com/v1alpha1` |
|    `kind` | string | `NodebalancerConfig` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NodebalancerConfigSpec](#NodebalancerConfigSpec)***||
| `status` | ***[NodebalancerConfigStatus](#NodebalancerConfigStatus)***||
## NodebalancerConfigSpec
##### (Appears on:[NodebalancerConfig](#NodebalancerConfig), [NodebalancerConfigStatus](#NodebalancerConfigStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `algorithm` | ***string***| ***(Optional)*** What algorithm this NodeBalancer should use for routing traffic to backends: roundrobin, leastconn, source|
| `check` | ***string***| ***(Optional)*** The type of check to perform against backends to ensure they are serving requests. This is used to determine if backends are up or down. If none no check is performed. connection requires only a connection to the backend to succeed. http and http_body rely on the backend serving HTTP, and that the response returned matches what is expected.|
| `checkAttempts` | ***int***| ***(Optional)*** How many times to attempt a check before considering a backend to be down. (1-30)|
| `checkBody` | ***string***| ***(Optional)*** This value must be present in the response body of the check in order for it to pass. If this value is not present in the response body of a check request, the backend is considered to be down|
| `checkInterval` | ***int***| ***(Optional)*** How often, in seconds, to check that backends are up and serving requests.|
| `checkPassive` | ***bool***| ***(Optional)*** If true, any response from this backend with a 5xx status code will be enough for it to be considered unhealthy and taken out of rotation.|
| `checkPath` | ***string***| ***(Optional)*** The URL path to check on each backend. If the backend does not respond to this request it is considered to be down.|
| `checkTimeout` | ***int***| ***(Optional)*** How long, in seconds, to wait for a check attempt before considering it failed. (1-30)|
| `cipherSuite` | ***string***| ***(Optional)*** What ciphers to use for SSL connections served by this NodeBalancer. `legacy` is considered insecure and should only be used if necessary.|
| `nodeStatus` | ***map[string]kubeform.dev/kubeform/apis/linode/v1alpha1.NodebalancerConfigSpecNodeStatus***| ***(Optional)*** |
| `nodebalancerID` | ***int***|The ID of the NodeBalancer to access.|
| `port` | ***int***| ***(Optional)*** The TCP port this Config is for. These values must be unique across configs on a single NodeBalancer (you can't have two configs for port 80, for example). While some ports imply some protocols, no enforcement is done and you may configure your NodeBalancer however is useful to you. For example, while port 443 is generally used for HTTPS, you do not need SSL configured to have a NodeBalancer listening on port 443.|
| `protocol` | ***string***| ***(Optional)*** The protocol this port is configured to serve. If this is set to https you must include an ssl_cert and an ssl_key.|
| `sslCert` | ***string***| ***(Optional)*** The certificate this port is serving. This is not returned. If set, this field will come back as `<REDACTED>`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.|
| `sslCommonname` | ***string***| ***(Optional)*** The common name for the SSL certification this port is serving if this port is not configured to use SSL.|
| `sslFingerprint` | ***string***| ***(Optional)*** The fingerprint for the SSL certification this port is serving if this port is not configured to use SSL.|
| `stickiness` | ***string***| ***(Optional)*** Controls how session stickiness is handled on this port: 'none', 'table', 'http_cookie'|
## NodebalancerConfigStatus
##### (Appears on:[NodebalancerConfig](#NodebalancerConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NodebalancerConfigSpec](#NodebalancerConfigSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `ssl_key` | ***string*** |The private key corresponding to this port's certificate. This is not returned. If set, this field will come back as `<REDACTED>`. Please use the ssl_commonname and ssl_fingerprint to identify the certificate.|
