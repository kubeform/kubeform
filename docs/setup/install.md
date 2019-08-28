---
title: Install
description: Kubeform Install
menu:
  product_kubeform_0.0.1:
    identifier: install-Kubeform
    name: Install
    parent: setup
    weight: 10
product_name: Kubeform
menu_name: product_kubeform_0.0.1
section_menu_id: setup
---

# Installation Guide

Kubeform controller can be installed via Helm.

## Using Helm
Kubeform can be installed via [Helm](https://helm.sh/) using the [chart](https://github.com/kubeform/installer/tree/v0.1.0/chart/kubeform) from [AppsCode Charts Repository](https://github.com/appscode/charts). To install the chart with the release name `my-release`:

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search appscode/kubeform
NAME                CHART VERSION  APP VERSION  DESCRIPTION
appscode/kubeform   v0.1.0         v0.1.0       Kubeform by AppsCode - Build Cloud Infrastructure from Kubernetes

$ helm install appscode/kubeform --name kfc --version v0.1.0 --namespace kube-system
```

To see the detailed configuration options, visit [here](https://github.com/kubeform/installer/tree/v0.1.0/chart/kubeform).