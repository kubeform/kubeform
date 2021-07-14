---
title: How Kubeform handles Terraform State
menu:
  docs_{{ .version }}:
    identifier: kubeform-tfstate
    name: Terraform State
    parent: what-is-kubeform
    weight: 30
menu_name: docs_{{ .version }}
section_menu_id: concepts
---

# How Kubeform handles Terraform State

## Terraform Resource

For individual resources, Kubeform doesn't store the full `tfstate` file in the CRD. The `tfstate` json file contains `version`, `terraform_version`, `serial`, `lineage` and `resources` fields.

Kubeform stores the `tfstate` file in the CRD using the following way:

1. The `version`, `terraform_version`, `serial` and, `lineage` fields are stored in `status.state` field in the custom resources.
2. The data from the `resources` field that are not sensitive are stored in the `output` field of the CRD.
3. The data from the `resources` field that are sensitive are stored in a secret. The secret name is specified in a `secretRef` field of the CRD. If the `secretRef` is not set, then KFC creates a secret and stores the sensitive outputs there.

To obtain the `tfstate` from the state field:

1. The `version`, `terraform_version`, `serial` and, `lineage` fields are retrieves from the `status.state` field of a Kubeform resource.
2. The non-sensitive data from the `output` field of `status` and the sensitive data from the secret are merged and then set on the `resources` field.
