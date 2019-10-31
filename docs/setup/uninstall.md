---
title: Uninstall | Kubeform
menu:
  docs_{{ .version }}:
    identifier: uninstall-kubeform
    name: Uninstall
    parent: setup
    weight: 20
menu_name: docs_{{ .version }}
section_menu_id: setup
---

# Uninstall Kubeform

To uninstall Kubeform installed using helm, use the following command.

```console
$ helm delete kfc
```

If you want to remove the release from the store and make its name free for later use, Then provide the `--purge` flag:

```console
$ helm delete kfc --purge
```