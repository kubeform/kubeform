[![Build Status](https://github.com/kubeform/kubeform/workflows/CI/badge.svg)](https://github.com/kubeform/kubeform/actions?workflow=CI)
[![Docker Pulls](https://img.shields.io/docker/pulls/kubeform/kfc.svg)](https://hub.docker.com/r/kubeform/kfc/)
[![Slack](https://slack.appscode.com/badge.svg)](https://slack.appscode.com)
[![Twitter](https://img.shields.io/twitter/follow/kubeform.svg?style=social&logo=twitter&label=Follow)](https://twitter.com/intent/follow?screen_name=Kubeform)

# Kubeform

Kubeform by AppsCode is a Kubernetes operator for [Terraform](https://www.terraform.io/). Kubeform provides Kubernetes CRDs for Terraform resources and modules so that you can manage any cloud infrastructure in a Kubernetes native way. You just write a CRD for a cloud infrastructure, apply it and Kubeform will create it for you! Kubeform currently supports 5 top cloud platforms. These are AWS, Google Cloud, Azure, Digitalocean and Linode.

## Features

- Native kubernetes support
- Built on Terraform
- Supports Terraform resources and modules
- Use cloud infrastructures as code
- Define & Manage cloud infrastructures as Kubernetes `CRD` (Custom Resource Definition)
- Supports multiple cloud platform
- 100% open source

## Installation

To install Kubeform, please follow the guide [here](https://kubeform.com/docs/latest/setup/install/).

## Using Kubeform

Want to learn how to use Kubeform? Please start [here](https://kubeform.com/docs/latest/guides/).

## Contribution guidelines

Want to help improve Kubeform? Please start [here](https://kubeform.com/docs/latest/welcome/contributing/).

## Acknowledgement

- Many thanks to [Hashicorp](https://www.hashicorp.com/) for [Terraform](https://www.terraform.io/) project.

## Support

We use Slack for public discussions. To chit chat with us or the rest of the community, join us in the [AppsCode Slack team](https://appscode.slack.com/messages/C8NCX6N23/details/) channel `#kubeform`. To sign up, use our [Slack inviter](https://slack.appscode.com/).

If you have found a bug with Kubeform or want to request for new features, please [file an issue](https://github.com/kubeform/project/issues/new)
