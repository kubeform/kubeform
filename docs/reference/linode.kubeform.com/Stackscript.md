---
title: Stackscript
menu:
  docs_{{ .version }}:
    identifier: stackscript-linode.kubeform.com
    name: Stackscript
    parent: linode.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Stackscript
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `linode.kubeform.com/v1alpha1` |
|    `kind` | string | `Stackscript` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StackscriptSpec](#stackscriptspec)***||
| `status` | ***[StackscriptStatus](#stackscriptstatus)***||
## Phase(`string` alias)

Appears on:[StackscriptStatus](#stackscriptstatus)

## StackscriptSpec

Appears on:[Stackscript](#stackscript), [StackscriptStatus](#stackscriptstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `created` | ***string***| ***(Optional)*** The date this StackScript was created.|
| `deploymentsActive` | ***int64***| ***(Optional)*** Count of currently active, deployed Linodes created from this StackScript.|
| `deploymentsTotal` | ***int64***| ***(Optional)*** The total number of times this StackScript has been deployed.|
| `description` | ***string***|A description for the StackScript.|
| `images` | ***[]string***|An array of Image IDs representing the Images that this StackScript is compatible for deploying with.|
| `isPublic` | ***bool***| ***(Optional)*** This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private.|
| `label` | ***string***|The StackScript's label is for display purposes only.|
| `revNote` | ***string***| ***(Optional)*** This field allows you to add notes for the set of revisions made to this StackScript.|
| `script` | ***string***|The script to execute when provisioning a new Linode with this StackScript.|
| `updated` | ***string***| ***(Optional)*** The date this StackScript was updated.|
| `userDefinedFields` | ***[[]StackscriptSpecUserDefinedFields](#stackscriptspecuserdefinedfields)***| ***(Optional)*** This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment.|
| `userGravatarID` | ***string***| ***(Optional)*** The Gravatar ID for the User who created the StackScript.|
| `username` | ***string***| ***(Optional)*** The User who created the StackScript.|
## StackscriptSpecUserDefinedFields

Appears on:[StackscriptSpec](#stackscriptspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `default` | ***string***| ***(Optional)*** |
| `example` | ***string***| ***(Optional)*** |
| `label` | ***string***| ***(Optional)*** |
| `manyOf` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `oneOf` | ***string***| ***(Optional)*** |
## StackscriptStatus

Appears on:[Stackscript](#stackscript)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StackscriptSpec](#stackscriptspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
