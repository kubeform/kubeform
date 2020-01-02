---
title: Install | Kubeform
menu:
  docs_{{ .version }}:
    identifier: install-kubeform
    name: Install
    parent: setup
    weight: 10
menu_name: docs_{{ .version }}
section_menu_id: setup
---

# Installation Guide

Kubeform controller can be installed via Helm.

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

Kubeform can be installed via [Helm](https://helm.sh/) using the [chart](https://github.com/kubeform/installer/tree/{{< param "info.version" >}}/charts/kubeform) from [AppsCode Charts Repository](https://github.com/appscode/charts). To install the chart with the release name `my-release`:

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubeform
NAME                CHART VERSION  APP VERSION  DESCRIPTION
appscode/kubeform   {{< param "info.version" >}}         {{< param "info.version" >}}       Kubeform by AppsCode - Build Cloud Infrastructure from Kubernetes

$ helm install kfc appscode/kubeform --version {{< param "info.version" >}} --namespace kube-system
```

To see the detailed configuration options, visit [here](https://github.com/kubeform/installer/tree/{{< param "info.version" >}}/charts/kubeform).

</div>
<div class="tab-pane fade" id="helm2" role="tabpanel" aria-labelledby="helm2-tab">

## Using Helm 2

Kubeform can be installed via [Helm](https://helm.sh/) using the [chart](https://github.com/kubeform/installer/tree/{{< param "info.version" >}}/charts/kubeform) from [AppsCode Charts Repository](https://github.com/appscode/charts). To install the chart with the release name `my-release`:

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search appscode/kubeform
NAME                CHART VERSION  APP VERSION  DESCRIPTION
appscode/kubeform   {{< param "info.version" >}}         {{< param "info.version" >}}       Kubeform by AppsCode - Build Cloud Infrastructure from Kubernetes

$ helm install appscode/kubeform --name kfc --version {{< param "info.version" >}} --namespace kube-system
```

To see the detailed configuration options, visit [here](https://github.com/kubeform/installer/tree/{{< param "info.version" >}}/charts/kubeform).

</div>
<div class="tab-pane fade show active" id="script" role="tabpanel" aria-labelledby="script-tab">

## Using YAML

If you prefer to not use Helm, you can generate YAMLs from Kubeform chart and deploy using `kubectl`. Here we are going to show the prodecure using Helm 3.

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubeform
NAME                CHART VERSION  APP VERSION  DESCRIPTION
appscode/kubeform   {{< param "info.version" >}}         {{< param "info.version" >}}       Kubeform by AppsCode - Build Cloud Infrastructure from Kubernetes

$ helm install kfc appscode/kubeform \
  --version {{< param "info.version" >}} \
  --namespace kube-system \
  --no-hooks | kubectl apply -f -
```

To see the detailed configuration options, visit [here](https://github.com/kubeform/installer/tree/{{< param "info.version" >}}/charts/kubeform).

</div>
</div>

### Installing in GKE Cluster

If you are installing Kubeform on a GKE cluster, you will need cluster admin permissions to install Kubeform operator. Run the following command to grant admin permision to the cluster.

```console
$ kubectl create clusterrolebinding "cluster-admin-$(whoami)" \
  --clusterrole=cluster-admin \
  --user="$(gcloud config get-value core/account)"
```

In addition, if your GKE cluster is a [private cluster](https://cloud.google.com/kubernetes-engine/docs/how-to/private-clusters), you will need to either add an additional firewall rule that allows master nodes access port `8443/tcp` on worker nodes, or change the existing rule that allows access to ports `443/tcp` and `10250/tcp` to also allow access to port `8443/tcp`. The procedure to add or modify firewall rules is described in the official GKE documentation for private clusters mentioned before.

## Verify installation
To check if Kubeform operator pods have started, run the following command:
```console
$ kubectl get pods --all-namespaces -l app=kubeform --watch

NAMESPACE     NAME                              READY     STATUS    RESTARTS   AGE
kube-system   stash-operator-859d6bdb56-m9br5   2/2       Running   2          5s
```

Once the operator pods are running, you can cancel the above command by typing `Ctrl+C`.
