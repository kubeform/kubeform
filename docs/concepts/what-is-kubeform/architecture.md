---
title: Kubeform Architecture
menu:
  docs_{{ .version }}:
    identifier: kubeform-architecture
    name: Kubeform Architecture
    parent: what-is-kubeform
    weight: 20
menu_name: docs_{{ .version }}
section_menu_id: concepts
---

## Kubeform Architecture

The following diagram shows how `Kubeform` creates a resource on a Cloud Provider (GCP, AWS, etc.).

<figure align="center">
 <img alt="Kubeform Architecture" src="/docs/images/concepts/what-is-kubeform/architecture.jpg">
 <figcaption align="center">Fig: Kubeform Architecture</figcaption>
</figure>

The Resource Creation Process of Kubeform consists of the following steps:

1. At first, a user creates a provider secret with access credentials of the Cloud provider where the resource will be created.

2. Then, he creates a sensitive secret with sensitive fields of the cloud resource that he wants to create. This is optional, if a user do not create a sensitive secret then kfc will create a sensitive secret if the cloud resource has any sensitive field.

3. Then, he creates a CRD of the resource that specifies the information of the Cloud Resource. The CRD also points to the secrets that he created.

4. The KubeForm Controller (KFC) watches the created CRD and also the sensitive secret continuously.

5. If the KubeForm Controller (KFC) get any new changes in sensitive secret or in the CRD it starts reconciling and the resource enters into `InProgress` phase.
   
6. The KubeForm Controller (KFC) Create, Update or Delete the respective cloud resource through the Resource CRUD API.  

7. If the resource is being Created or Updated the resource enters into `InProgress` phase. When the recociling process ends and resource successfully created or updated theen the phase is `Current`.

8. The KubeForm Controller (KFC) update the resource `spec.state` after successfully creating or updating the resource.

9. If the resource is being Deleted then the resource phase is `Terminating`.

10. The KubeForm Controller (KFC) deletes the resource after the respective cloud resource get destroyed.