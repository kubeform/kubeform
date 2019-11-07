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

## Terraform Module

For Terraform modules, Kubeform stores the full `tfstate` file in the CRD status object. As the `tfstate` file may contain sensitive information, KFC doesn't store the `tfstate` file directly.

Before storing the `tfstate` file in the CRD, the `tfstate` is transformed in the following steps:

1. First, the content of the tfstate file is compressed using `gzip`.
2. Then the compressed data is encrypted using a secret key. The secret key is a base64 encoded key which is 32 bytes when decoded. The secret key is provided by the user to the Kubeform operator(`kfc`).
3. The encrypted data is `basE91` encoded. `basE91` is an advanced method for encoding binary data as ASCII characters. It is similar to base64, but is more efficient. The overhead produced by basE91 depends on the input data. It amounts at most to 23% (versus 33% for base64).
4. Finally, the base91 encoded data is stored in the `state` field of `status` object in the CRD.

The process is shown in the following figure:

<figure align="center">
 <img alt="Kubeform Architecture" src="/docs/images/concepts/what-is-kubeform/tfstate-to-state.svg">
 <figcaption align="center">Fig: `tfstate` to `state` string in CRD</figcaption>
</figure>

To obtain the tfstate file from the `status.state` field of a custom resource, the steps is followed in reverse:

1. First, the `status.state` field of the Kubeform module resource is base91 decoded.
2. The base91 decoded data is then decrypted using the secret key.
3. Then the decrypted data is decompressed using `gzip`.
4. Finally, decompressed data is saved as a `tfstate` file.

The process is shown in the following figure:

<figure align="center">
 <img alt="Kubeform Architecture" src="/docs/images/concepts/what-is-kubeform/state-to-tfstate.svg">
 <figcaption align="center">Fig: `state` string in CRD to `tfstate`</figcaption>
</figure>
