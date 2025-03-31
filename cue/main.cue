#KubernetesDeployment: {
    apiVersion: "apps/v1"
    kind:       "Deployment"
    metadata: {
        name:      =~"^[a-z0-9-]+$" // Enforce naming convention: lowercase, numbers, and dashes
        namespace: string | *"default"
        labels: {
            app: string
        }
    }
    spec: {
        replicas: int & >=1 & <=10
        selector: matchLabels: {
            app: string
        }
        template: {
            metadata: labels: {
                app: string
            }
            spec: containers: [...{
                name:  string
                image: =~"^([a-zA-Z0-9]+(?:[._/-][a-zA-Z0-9]+)*)(:[a-zA-Z0-9._-]+)?$"
                resources?: {
                    limits: {
                        cpu:    string
                        memory: string
                    }
                    requests: {
                        cpu:    string
                        memory: string
                    }
                }
                ports?: [...{
                    containerPort: int & >=1 & <=65535
                }]
            }]
        }
    }
}

#KubernetesService: {
    apiVersion: "v1"
    kind:       "Service"
    metadata: {
        name:      =~"^[a-z0-9-]+$" // Enforce naming convention
        namespace: string | *"default"
        labels: {
            app: string
        }
    }
    spec: {
        selector: {
            app: string
        }
        ports: [...{
            port:       int & >=1 & <=65535
            targetPort: int & >=1 & <=65535
        }]
        type: *"ClusterIP" | "NodePort" | "LoadBalancer"
    }
}

#KubernetesConfigMap: {
    apiVersion: "v1"
    kind:       "ConfigMap"
    metadata: {
        name:      =~"^[a-z0-9-]+$" // Enforce naming convention
        namespace: string | *"default"
    }
    data: [string]: string
}

// Combine resources into a single configuration
resources: {
    deployment: #KubernetesDeployment & {
        metadata: {
            name: "nginx-deployment"
            labels: {
                app: "nginx"
            }
        }
        spec: {
            replicas: 3
            selector: matchLabels: {
                app: "nginx"
            }
            template: {
                metadata: labels: {
                    app: "nginx"
                }
                spec: containers: [{
                    name:  "nginx"
                    image: "nginx:1.14.2"
                    resources: {
                        limits: {
                            cpu:    "500m"
                            memory: "128Mi"
                        }
                        requests: {
                            cpu:    "250m"
                            memory: "64Mi"
                        }
                    }
                    ports: [{
                        containerPort: 80
                    }]
                }]
            }
        }
    }

    service: #KubernetesService & {
        metadata: {
            name: "nginx-service"
            labels: {
                app: "nginx"
            }
        }
        spec: {
            selector: {
                app: "nginx"
            }
            ports: [{
                port:       80
                targetPort: 80
            }]
            type: "ClusterIP"
        }
    }

    configMap: #KubernetesConfigMap & {
        metadata: {
            name: "nginx-config"
        }
        data: {
            "config-key": "config-value"
        }
    }
}
