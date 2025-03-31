## Part 1: Installation & Configuration
### Step 1: Install CUE
```
$ go install cuelang.org/go/cmd/cue@latest
go: downloading cuelang.org/go v0.12.0
go: downloading cuelabs.dev/go/oci/ociregistry v0.0.0-20241125120445-2c00c104c6e1
go: downloading github.com/opencontainers/go-digest v1.0.0
go: downloading github.com/opencontainers/image-spec v1.1.0
go: downloading github.com/rogpeppe/go-internal v1.13.2-0.20241226121412-a5dc8ff20d0a
go: downloading github.com/spf13/cobra v1.8.1
go: downloading github.com/spf13/pflag v1.0.5
go: downloading golang.org/x/oauth2 v0.25.0
go: downloading golang.org/x/text v0.21.0
go: downloading golang.org/x/tools v0.29.0
go: downloading github.com/cockroachdb/apd/v3 v3.2.1
go: downloading github.com/emicklei/proto v1.13.4
go: downloading golang.org/x/mod v0.22.0
go: downloading github.com/protocolbuffers/txtpbfmt v0.0.0-20241112170944-20d2c9ebc01d
go: downloading github.com/pelletier/go-toml/v2 v2.2.3
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading golang.org/x/sync v0.10.0
go: downloading golang.org/x/net v0.34.0
go: downloading github.com/google/uuid v1.6.0
go: downloading github.com/mitchellh/go-wordwrap v1.0.1
```
### Step 2: Verify Installation
```
$ cue version
cue version v0.12.0

go version go1.23.7
      -buildmode exe
       -compiler gc
  DefaultGODEBUG asynctimerchan=1,gotypesalias=0,httpservecontentkeepheaders=1,tls3des=1,tlskyber=0,x509keypairleaf=0,x509negativeserial=1
     CGO_ENABLED 1
          GOARCH arm64
            GOOS darwin
         GOARM64 v8.0
cue.lang.version v0.12.0
```

## Part 4: Format Conversion
```
$ cue export example/deployment.cue --out yaml

deployment:
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: nginx-deployment
    namespace: default
  spec:
    replicas: 3
    selector:
      matchLabels:
        app: nginx
    template:
      metadata:
        labels:
          app: nginx
      spec:
        containers:
          - name: nginx
            image: nginx:1.14.2
            ports:
              - containerPort: 80
cue export example/person.cue --out yaml
person:
  name: John Doe
  age: 30
  email: john@example.com
  hobbies:
    - reading
    - coding
```
## Part 5: Advanced Tasks
```
$ cue export main.cue --out json
{
    "resources": {
        "deployment": {
            "apiVersion": "apps/v1",
            "kind": "Deployment",
            "metadata": {
                "name": "nginx-deployment",
                "namespace": "default",
                "labels": {
                    "app": "nginx"
                }
            },
            "spec": {
                "replicas": 3,
                "selector": {
                    "matchLabels": {
                        "app": "nginx"
                    }
                },
                "template": {
                    "metadata": {
                        "labels": {
                            "app": "nginx"
                        }
                    },
                    "spec": {
                        "containers": [
                            {
                                "name": "nginx",
                                "image": "nginx:1.14.2",
                                "resources": {
                                    "limits": {
                                        "cpu": "500m",
                                        "memory": "128Mi"
                                    },
                                    "requests": {
                                        "cpu": "250m",
                                        "memory": "64Mi"
                                    }
                                },
                                "ports": [
                                    {
                                        "containerPort": 80
                                    }
                                ]
                            }
                        ]
                    }
                }
            }
        },
        "service": {
            "apiVersion": "v1",
            "kind": "Service",
            "metadata": {
                "name": "nginx-service",
                "namespace": "default",
                "labels": {
                    "app": "nginx"
                }
            },
            "spec": {
                "selector": {
                    "app": "nginx"
                },
                "ports": [
                    {
                        "port": 80,
                        "targetPort": 80
                    }
                ],
                "type": "ClusterIP"
            }
        },
        "configMap": {
            "apiVersion": "v1",
            "kind": "ConfigMap",
            "metadata": {
                "name": "nginx-config",
                "namespace": "default"
            },
            "data": {
                "config-key": "config-value"
            }
        }
    }
}
```
