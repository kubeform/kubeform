---
title: Developer Guide
menu:
  docs_{{ .version }}:
    identifier: developer-guide-readme
    name: Overview
    parent: developer-guide
    weight: 10
menu_name: docs_{{ .version }}
section_menu_id: setup
aliases:
  - /docs/{{ .version }}/setup/developer-guide/
---

> New to Kubeform? Please start [here](/docs/concepts/README.md).

## Development Guide
This document is intended to be the canonical source of truth for things like supported toolchain versions for building Kubeform.
If you find a requirement that this doc does not capture, please submit an issue on github.

This document is intended to be relative to the branch in which it is found. It is guaranteed that requirements will change over time
for the development branch, but release branches of Kubeform should not change.

### Build Kubeform
Some of the Kubeform development helper scripts rely on a fairly up-to-date GNU tools environment, so most recent Linux distros should
work just fine out-of-the-box.

#### Setup GO
Kubeform is written in Google's GO programming language. Currently, Kubeform is developed and tested on **go 1.9.2**. If you haven't set up a GO
development environment, please follow [these instructions](https://golang.org/doc/code.html) to install GO.
