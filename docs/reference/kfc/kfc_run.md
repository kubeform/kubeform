---
title: Kfc Run
menu:
  docs_{{ .version }}:
    identifier: kfc-run
    name: Kfc Run
    parent: reference-kfc
menu_name: docs_{{ .version }}
section_menu_id: reference
---
## kfc run

Launch Vault operator

### Synopsis

Launch Vault operator

```
kfc run [flags]
```

### Options

```
  -h, --help                help for run
      --kubeconfig string   Path to a kubeconfig. Only required if out-of-cluster.
      --master string       The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.
      --secret-key string   A base64-encoded key, of length 32 bytes when decoded.
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

