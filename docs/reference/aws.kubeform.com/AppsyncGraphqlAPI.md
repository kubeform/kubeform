---
title: AppsyncGraphqlAPI
menu:
  docs_{{ .version }}:
    identifier: appsyncgraphqlapi-aws.kubeform.com
    name: AppsyncGraphqlAPI
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppsyncGraphqlAPI
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AppsyncGraphqlAPI` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppsyncGraphqlAPISpec](#appsyncgraphqlapispec)***||
| `status` | ***[AppsyncGraphqlAPIStatus](#appsyncgraphqlapistatus)***||
## AppsyncGraphqlAPISpec

Appears on:[AppsyncGraphqlAPI](#appsyncgraphqlapi), [AppsyncGraphqlAPIStatus](#appsyncgraphqlapistatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `additionalAuthenticationProvider` | ***[[]AppsyncGraphqlAPISpecAdditionalAuthenticationProvider](#appsyncgraphqlapispecadditionalauthenticationprovider)***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `authenticationType` | ***string***||
| `logConfig` | ***[[]AppsyncGraphqlAPISpecLogConfig](#appsyncgraphqlapispeclogconfig)***| ***(Optional)*** |
| `name` | ***string***||
| `openidConnectConfig` | ***[[]AppsyncGraphqlAPISpecOpenidConnectConfig](#appsyncgraphqlapispecopenidconnectconfig)***| ***(Optional)*** |
| `schema` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `uris` | ***map[string]string***| ***(Optional)*** |
| `userPoolConfig` | ***[[]AppsyncGraphqlAPISpecUserPoolConfig](#appsyncgraphqlapispecuserpoolconfig)***| ***(Optional)*** |
## AppsyncGraphqlAPISpecAdditionalAuthenticationProvider

Appears on:[AppsyncGraphqlAPISpec](#appsyncgraphqlapispec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticationType` | ***string***||
| `openidConnectConfig` | ***[[]AppsyncGraphqlAPISpecAdditionalAuthenticationProviderOpenidConnectConfig](#appsyncgraphqlapispecadditionalauthenticationprovideropenidconnectconfig)***| ***(Optional)*** |
| `userPoolConfig` | ***[[]AppsyncGraphqlAPISpecAdditionalAuthenticationProviderUserPoolConfig](#appsyncgraphqlapispecadditionalauthenticationprovideruserpoolconfig)***| ***(Optional)*** |
## AppsyncGraphqlAPISpecAdditionalAuthenticationProviderOpenidConnectConfig

Appears on:[AppsyncGraphqlAPISpecAdditionalAuthenticationProvider](#appsyncgraphqlapispecadditionalauthenticationprovider)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authTtl` | ***int64***| ***(Optional)*** |
| `clientID` | ***string***| ***(Optional)*** |
| `iatTtl` | ***int64***| ***(Optional)*** |
| `issuer` | ***string***||
## AppsyncGraphqlAPISpecAdditionalAuthenticationProviderUserPoolConfig

Appears on:[AppsyncGraphqlAPISpecAdditionalAuthenticationProvider](#appsyncgraphqlapispecadditionalauthenticationprovider)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `appIDClientRegex` | ***string***| ***(Optional)*** |
| `awsRegion` | ***string***| ***(Optional)*** |
| `userPoolID` | ***string***||
## AppsyncGraphqlAPISpecLogConfig

Appears on:[AppsyncGraphqlAPISpec](#appsyncgraphqlapispec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cloudwatchLogsRoleArn` | ***string***||
| `fieldLogLevel` | ***string***||
## AppsyncGraphqlAPISpecOpenidConnectConfig

Appears on:[AppsyncGraphqlAPISpec](#appsyncgraphqlapispec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authTtl` | ***int64***| ***(Optional)*** |
| `clientID` | ***string***| ***(Optional)*** |
| `iatTtl` | ***int64***| ***(Optional)*** |
| `issuer` | ***string***||
## AppsyncGraphqlAPISpecUserPoolConfig

Appears on:[AppsyncGraphqlAPISpec](#appsyncgraphqlapispec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `appIDClientRegex` | ***string***| ***(Optional)*** |
| `awsRegion` | ***string***| ***(Optional)*** |
| `defaultAction` | ***string***||
| `userPoolID` | ***string***||
## AppsyncGraphqlAPIStatus

Appears on:[AppsyncGraphqlAPI](#appsyncgraphqlapi)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppsyncGraphqlAPISpec](#appsyncgraphqlapispec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AppsyncGraphqlAPIStatus](#appsyncgraphqlapistatus)

---
