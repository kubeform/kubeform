---
title: Kfc Version
menu:
  docs_{{ .version }}:
    identifier: kfc-version
    name: Kfc Version
    parent: reference-kfc
menu_name: docs_{{ .version }}
section_menu_id: reference
---
## kfc version

Prints binary version number.

### Synopsis

Prints binary version number.

```
kfc version [flags]
```

### Options

```
      --check string   Check version constraint
  -h, --help           help for version
      --short          Print just the version number.
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --enable-analytics                 Send analytical events to Google Analytics (default true)
      --log-flush-frequency duration     Maximum number of seconds between log flushes (default 5s)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [kfc](/docs/reference/kfc/kfc.md)	 - Kubeform controller by AppsCode - HashiCorp Terraform Operator for Kubernetes

