---
title: Release
menu:
  docs_v0.0.1:
    identifier: developer-guide-release
    name: Release Process
    parent: developer-guide
    weight: 20
menu_name: docs_v0.0.1
section_menu_id: setup
---

# Release Process

The following steps must be done from a Linux x64 bit machine.

- Do a global replacement of tags so that docs point to the next release.
- Push changes to the `release-x` branch and apply new tag.
- Push all the changes to remote repo.
- Build and push Kubeform docker image:
```console
$ cd ~/go/src/kubeform.dev/kubeform
$ make push
```

- Now, update the release notes in Github. See previous release notes to get an idea what to include there.
