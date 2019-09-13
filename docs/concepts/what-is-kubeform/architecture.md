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
 <img alt="Kubeform Architecture" src="/docs/images/concepts/what-is-kubeform/architecture.svg">
 <figcaption align="center">Fig: Kubeform Architecture</figcaption>
</figure>

The Resource Creation Process of Kubeform consists of the following steps:

1. At first, a user creates a secret with access credentials of the Cloud provider where the resource will be created.

2. Then, he creates a CRD of the resource that specifies the information of the Cloud Resource. The CRD also holds the credential information.

3. The KubeForm Controller (KFC) watches the created CRD.

4. Then, KFC creates `.tf` files from the CRD `Spec` and the provider secret.

5. If the `.tfstate` file doesn't exist then KFC creates the `.tfstate` file from the `status.output` & `status.state` fields of the CRD.

6. Then KFC runs `terraform apply` commands on the `.tfstate` file and `.tf` files that were created in the previous steps.

7. After successful execution of `terraform apply` command, it creates the specified resource on the specified Cloud Provider.

8. Then `terraform apply` command updates the `tfstate` file if necessary.

9. If the `.tfstate` file is updated, KFC also updates the `status.output` & `status.state` fields of the CRD.
