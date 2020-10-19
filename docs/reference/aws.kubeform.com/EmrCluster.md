---
title: EmrCluster
menu:
  docs_{{ .version }}:
    identifier: emrcluster-aws.kubeform.com
    name: EmrCluster
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EmrCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EmrCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EmrClusterSpec](#emrclusterspec)***||
| `status` | ***[EmrClusterStatus](#emrclusterstatus)***||
## EmrClusterSpec

Appears on:[EmrCluster](#emrcluster), [EmrClusterStatus](#emrclusterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `additionalInfo` | ***string***| ***(Optional)*** |
| `applications` | ***[]string***| ***(Optional)*** |
| `autoscalingRole` | ***string***| ***(Optional)*** |
| `bootstrapAction` | ***[[]EmrClusterSpecBootstrapAction](#emrclusterspecbootstrapaction)***| ***(Optional)*** |
| `clusterState` | ***string***| ***(Optional)*** |
| `configurations` | ***string***| ***(Optional)*** |
| `configurationsJSON` | ***string***| ***(Optional)*** |
| `coreInstanceCount` | ***int64***| ***(Optional)*** Deprecated|
| `coreInstanceGroup` | ***[[]EmrClusterSpecCoreInstanceGroup](#emrclusterspeccoreinstancegroup)***| ***(Optional)*** |
| `coreInstanceType` | ***string***| ***(Optional)*** Deprecated|
| `customAmiID` | ***string***| ***(Optional)*** |
| `ebsRootVolumeSize` | ***int64***| ***(Optional)*** |
| `ec2Attributes` | ***[[]EmrClusterSpecEc2Attributes](#emrclusterspecec2attributes)***| ***(Optional)*** |
| `instanceGroup` | ***[[]EmrClusterSpecInstanceGroup](#emrclusterspecinstancegroup)***| ***(Optional)*** Deprecated|
| `keepJobFlowAliveWhenNoSteps` | ***bool***| ***(Optional)*** |
| `kerberosAttributes` | ***[[]EmrClusterSpecKerberosAttributes](#emrclusterspeckerberosattributes)***| ***(Optional)*** |
| `logURI` | ***string***| ***(Optional)*** |
| `masterInstanceGroup` | ***[[]EmrClusterSpecMasterInstanceGroup](#emrclusterspecmasterinstancegroup)***| ***(Optional)*** |
| `masterInstanceType` | ***string***| ***(Optional)*** Deprecated|
| `masterPublicDNS` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `releaseLabel` | ***string***||
| `scaleDownBehavior` | ***string***| ***(Optional)*** |
| `securityConfiguration` | ***string***| ***(Optional)*** |
| `serviceRole` | ***string***||
| `step` | ***[[]EmrClusterSpecStep](#emrclusterspecstep)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `terminationProtection` | ***bool***| ***(Optional)*** |
| `visibleToAllUsers` | ***bool***| ***(Optional)*** |
## EmrClusterSpecBootstrapAction

Appears on:[EmrClusterSpec](#emrclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `args` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `path` | ***string***||
## EmrClusterSpecCoreInstanceGroup

Appears on:[EmrClusterSpec](#emrclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoscalingPolicy` | ***string***| ***(Optional)*** |
| `bidPrice` | ***string***| ***(Optional)*** |
| `ebsConfig` | ***[[]EmrClusterSpecCoreInstanceGroupEbsConfig](#emrclusterspeccoreinstancegroupebsconfig)***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `instanceCount` | ***int64***| ***(Optional)*** |
| `instanceType` | ***string***||
| `name` | ***string***| ***(Optional)*** |
## EmrClusterSpecCoreInstanceGroupEbsConfig

Appears on:[EmrClusterSpecCoreInstanceGroup](#emrclusterspeccoreinstancegroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `iops` | ***int64***| ***(Optional)*** |
| `size` | ***int64***||
| `type` | ***string***||
| `volumesPerInstance` | ***int64***| ***(Optional)*** |
## EmrClusterSpecEc2Attributes

Appears on:[EmrClusterSpec](#emrclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `additionalMasterSecurityGroups` | ***string***| ***(Optional)*** |
| `additionalSlaveSecurityGroups` | ***string***| ***(Optional)*** |
| `emrManagedMasterSecurityGroup` | ***string***| ***(Optional)*** |
| `emrManagedSlaveSecurityGroup` | ***string***| ***(Optional)*** |
| `instanceProfile` | ***string***||
| `keyName` | ***string***| ***(Optional)*** |
| `serviceAccessSecurityGroup` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
## EmrClusterSpecInstanceGroup

Appears on:[EmrClusterSpec](#emrclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoscalingPolicy` | ***string***| ***(Optional)*** |
| `bidPrice` | ***string***| ***(Optional)*** |
| `ebsConfig` | ***[[]EmrClusterSpecInstanceGroupEbsConfig](#emrclusterspecinstancegroupebsconfig)***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `instanceCount` | ***int64***| ***(Optional)*** |
| `instanceRole` | ***string***||
| `instanceType` | ***string***||
| `name` | ***string***| ***(Optional)*** |
## EmrClusterSpecInstanceGroupEbsConfig

Appears on:[EmrClusterSpecInstanceGroup](#emrclusterspecinstancegroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `iops` | ***int64***| ***(Optional)*** |
| `size` | ***int64***||
| `type` | ***string***||
| `volumesPerInstance` | ***int64***| ***(Optional)*** |
## EmrClusterSpecKerberosAttributes

Appears on:[EmrClusterSpec](#emrclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `adDomainJoinUser` | ***string***| ***(Optional)*** |
| `realm` | ***string***||
## EmrClusterSpecMasterInstanceGroup

Appears on:[EmrClusterSpec](#emrclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bidPrice` | ***string***| ***(Optional)*** |
| `ebsConfig` | ***[[]EmrClusterSpecMasterInstanceGroupEbsConfig](#emrclusterspecmasterinstancegroupebsconfig)***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `instanceCount` | ***int64***| ***(Optional)*** |
| `instanceType` | ***string***||
| `name` | ***string***| ***(Optional)*** |
## EmrClusterSpecMasterInstanceGroupEbsConfig

Appears on:[EmrClusterSpecMasterInstanceGroup](#emrclusterspecmasterinstancegroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `iops` | ***int64***| ***(Optional)*** |
| `size` | ***int64***||
| `type` | ***string***||
| `volumesPerInstance` | ***int64***| ***(Optional)*** |
## EmrClusterSpecStep

Appears on:[EmrClusterSpec](#emrclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `actionOnFailure` | ***string***||
| `hadoopJarStep` | ***[[]EmrClusterSpecStepHadoopJarStep](#emrclusterspecstephadoopjarstep)***||
| `name` | ***string***||
## EmrClusterSpecStepHadoopJarStep

Appears on:[EmrClusterSpecStep](#emrclusterspecstep)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `args` | ***[]string***| ***(Optional)*** |
| `jar` | ***string***||
| `mainClass` | ***string***| ***(Optional)*** |
| `properties` | ***map[string]string***| ***(Optional)*** |
## EmrClusterStatus

Appears on:[EmrCluster](#emrcluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EmrClusterSpec](#emrclusterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[EmrClusterStatus](#emrclusterstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `kerberos_attributes.<index>.ad_domain_join_password` | ***string*** ||
| `kerberos_attributes.<index>.cross_realm_trust_principal_password` | ***string*** ||
| `kerberos_attributes.<index>.kdc_admin_password` | ***string*** ||
