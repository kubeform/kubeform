---
title: Kubeform Overview
description: Kubeform Overview
menu:
  docs_{{ .version }}:
    identifier: overview-concepts
    name: Overview
    parent: what-is-kubeform
    weight: 10
menu_name: docs_{{ .version }}
section_menu_id: concepts
---

# Kubeform

 Kubeform by AppsCode is a Kubernetes operator provisioning cloud or on-prem resources using [Terraform](https://terraform.io) providers. If you need to create a cloud infrastructure from kubernetes to use it with any kubernetes workloads or CRD you might want to create it from kubernetes. Kubeform allows you to create cloud infrastructure from kubernetes. You can just write a CRD for a cloud infrastructure, apply it and Kubeform will create it!

## Features

- **Native Kubernetes Support**
  Standard Kubernetes is all you need. If you can run Kubernetes, you can provision and manage cloud infrastructure using Kubeform. Use standard Kubernetes CLI and API to provision and manage different cloud infrastructures. This native integration with Kubernetes means your cloud infrastructures are ready for any higher level deployment tooling like GitOps, Open Policy Agent (OPA), etc.

- **Built on Terraform**
  Terraform is a popular tool for provisioning cloud infrastructure. Kubeform is built on top of Terraform. You can think it like a kubernetes operator for terraform. You can easily convert any resource of terraform using a CRD. You can just write a yaml for the resource which is almost same as a `tf` configuration. After applying the yaml using kubernetes API, Kubeform will provision the cloud infrastructure for you!

- **Infrastructure as Code**:
  Using Kubeform, you can just write a yaml and provision a cloud infrastructure in popular cloud platforms. You can easily provision and manage cloud infrastructures using codes. You can automate the lifecycle of cloud infrastructure using Kubeform.

- **Supports Multiple Cloud Platform**:
  Kubeform supports 5 top cloud platforms. These are AWS, Google Cloud, Azure, DigitalOcean & Linode. Using Kubeform, Infrastructure of any of these clouds can be provisioned using a simple CRD instance.

- **Unleash Developer Velocity**:
  With Kubeform, you can launch a cloud infrastructure via a simple CRD instance and start developing in minutes. Kubeform eliminates lengthy deployment and management processes with on-demand provisioning, scaling, patching and updating cloud infrastructure and shortens time-to-value. Kubeform gives each developer on your team their own development instance, so your team can focus on what matters most — velocity and not have to worry about stepping on each other’s toes.
