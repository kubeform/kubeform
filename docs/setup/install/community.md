---
title: Install Kubeform Community Edition
description: Installation guide for Kubeform Community edition
menu:
  docs_{{ .version }}:
    identifier: install-kubeform-community
    name: Community Edition
    parent: installation-guide
    weight: 10
product_name: kubeform
menu_name: docs_{{ .version }}
section_menu_id: setup
---

# Install Kubeform Community Edition

Kubeform Community edition is available under [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0) license and free to use for both commercial and non-commercial purposes.

To use Kubeform Community edition, you can grab **1 year** free license from [here](https://license-issuer.appscode.com/).

## Get a License

In this section, we are going to show you how you can get a **1 year** free license for Kubeform Community edition. You can get a license for your Kubernetes cluster by going through the following steps:

- At first, go to [AppsCode License Server](https://license-issuer.appscode.com/) and fill up the form. It will ask for your Name, Email, the product you want to install, and your cluster ID (UID of the `kube-system` namespace).
- Provide your name and email address. You can provide your personal or work email address.
- Then, select `Kubeform Community Edition` in the product field.
- Now, provide your cluster ID. You can get your cluster ID easily by running the following command:

```bash
kubectl get ns kube-system -o=jsonpath='{.metadata.uid}'
```

- Then, you have to agree with the terms and conditions. We recommend reading it before checking the box.
- Now, you can submit the form. After you submit the form, the AppsCode License server will send an email to the provided email address with a link to your license file.
- Navigate to the provided link and save the license into a file. Here, we save the license to a `license.txt` file.

Here is a screenshot of the license form.

<figure align="center">
  <img alt="Kubeform Backend Overview" src="/docs/images/setup/community_license_form.png">
  <figcaption align="center">Fig: Kubeform License Form</figcaption>
</figure>

You can create licenses for as many clusters as you want. You can upgrade your license any time without re-installing Kubeform by following the upgrading guide from [here](/docs/setup/upgrade.md#upgrading-license).

> Kubeform licensing process has been designed to work with CI/CD workflow. You can automatically obtain a license from your CI/CD pipeline by following the guide from [here](https://github.com/appscode/offline-license-server#offline-license-server).

## Install

Kubeform operator can be installed as a Helm chart or simply as Kubernetes manifests.

<ul class="nav nav-tabs" id="installerTab" role="tablist">
  <li class="nav-item">
    <a class="nav-link active" id="helm3-tab" data-toggle="tab" href="#helm3" role="tab" aria-controls="helm3" aria-selected="true">Helm 3 (Recommended)</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" id="helm2-tab" data-toggle="tab" href="#helm2" role="tab" aria-controls="helm2" aria-selected="false">Helm 2</a>
  </li>
  <li class="nav-item">
    <a class="nav-link" id="script-tab" data-toggle="tab" href="#script" role="tab" aria-controls="script" aria-selected="false">YAML</a>
  </li>
</ul>
<div class="tab-content" id="installerTabContent">
  <div class="tab-pane fade show active" id="helm3" role="tabpanel" aria-labelledby="helm3-tab">

## Using Helm 3

Kubeform can be installed via [Helm](https://helm.sh/) using the [chart](https://github.com/kubeform/installer/tree/{{< param "info.installer" >}}/charts/kubeform) from [AppsCode Charts Repository](https://github.com/appscode/charts). To install the chart with the release name `kubeform`:

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubeform --version {{< param "info.community" >}}
NAME            CHART VERSION         APP VERSION         DESCRIPTION
appscode/kubeform  {{< param "info.community" >}}    {{< param "info.community" >}}  Kubeform by AppsCode - Backup your Kubernetes Volumes

$ helm install kubeform appscode/kubeform           \
  --version {{< param "info.community" >}}          \
  --namespace kube-system                           \
  --set-file license=/path/to/the/license.txt
```

To see the detailed configuration options, visit [here](https://github.com/kubeform/installer/tree/{{< param "info.installer" >}}/charts/kubeform).

</div>
<div class="tab-pane fade" id="helm2" role="tabpanel" aria-labelledby="helm2-tab">

## Using Helm 2

Kubeform can be installed via [Helm](https://helm.sh/) using the [chart](https://github.com/kubeform/installer/tree/{{< param "info.installer" >}}/charts/kubeform) from [AppsCode Charts Repository](https://github.com/appscode/charts). To install the chart with the release name `kubeform`:

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search appscode/kubeform --version {{< param "info.community" >}}
NAME            CHART VERSION         APP VERSION         DESCRIPTION
appscode/kubeform  {{< param "info.community" >}}    {{< param "info.community" >}}  Kubeform by AppsCode - Backup your Kubernetes Volumes

$ helm install appscode/kubeform --name kubeform    \
  --version {{< param "info.community" >}}          \
  --namespace kube-system                           \
  --set-file license=/path/to/the/license.txt
```

To see the detailed configuration options, visit [here](https://github.com/kubeform/installer/tree/{{< param "info.installer" >}}/charts/kubeform).

</div>
<div class="tab-pane fade" id="script" role="tabpanel" aria-labelledby="script-tab">

## Using YAML

If you prefer to not use Helm, you can generate YAMLs from Kubeform chart and deploy using `kubectl`. Here we are going to show the prodecure using Helm 3.

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubeform --version {{< param "info.community" >}}
NAME            CHART VERSION         APP VERSION         DESCRIPTION
appscode/kubeform  {{< param "info.community" >}}    {{< param "info.community" >}}  Kubeform by AppsCode - Backup your Kubernetes Volumes

$ helm template kubeform appscode/kubeform          \
  --version {{< param "info.community" >}}          \
  --namespace kube-system                           \
  --set-file license=/path/to/the/license.txt       \
  --no-hooks | kubectl apply -f -
```

To see the detailed configuration options, visit [here](https://github.com/kubeform/installer/tree/{{< param "info.installer" >}}/charts/kubeform).

</div>
</div>

## Verify installation

To check if Kubeform operator pods have started, run the following command:

```bash
$ kubectl get pods --all-namespaces -l app.kubernetes.io/name=kubeform --watch

NAMESPACE     NAME                        READY  STATUS    RESTARTS   AGE
kube-system   kubeform-859d6bdb56-m9br5   1/1    Running   2          5s
```

Once the operator pod is running, you can cancel the above command by typing `Ctrl+C`.
