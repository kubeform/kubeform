---
title: Welcome | Kubeform
description: Welcome to Kubeform
menu:
  docs_{{ .version }}:
    identifier: readme
    name: Readme
    parent: welcome
    weight: -1
menu_name: docs_{{ .version }}
section_menu_id: welcome
url: /docs/{{ .version }}/welcome/
aliases:
  - /docs/{{ .version }}/
  - /docs/{{ .version }}/README/
---

# Kubeform

Kubeform by AppsCode is a Kubernetes operator provisioning cloud or on-prem resources using [Terraform](https://terraform.io) providers. Kubeform provides auto-generated Kubernetes CRDs for Terraform resources so that you can manage any cloud infrastructure in a Kubernetes native way. You just write a CRD for a cloud infrastructure, apply it and Kubeform will create it for you! Kubeform currently supports 5 top cloud platforms. These are AWS, Google Cloud, Azure, DigitalOcean and Linode.

The key features of Kubeform are:

- Native kubernetes support
- Built on Terraform
- Supports Terraform resources and modules
- Use cloud infrastructures as code
- Define & Manage cloud infrastructures as Kubernetes `CRD` (Custom Resource Definition)
- Supports multiple cloud platform

From here you can learn all about Kubeform's architecture and how to deploy and use Kubeform.

- [Concepts](/docs/concepts/). Concepts explain some significant aspect of Kubeform. This is where you can learn about what Kubeform does and how it does it.

- [Setup](/docs/setup/). Setup contains instructions for installing
  the Kubeform in various cloud providers.

- [Guides](/docs/guides). Guides show you how to perform tasks with Kubeform.

- [Reference](/docs/reference/). Detailed exhaustive lists of
command-line options, configuration options, API definitions, and procedures.

We're always looking for help improving our documentation, so please don't hesitate to [file an issue](https://github.com/kubeform/kubeform/issues/new) if you see some problem. Or better yet, submit your own [contributions](/docs/CONTRIBUTING.md) to help make our docs better.
