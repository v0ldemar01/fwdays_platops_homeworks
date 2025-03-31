#KubernetesDeployment: {
    apiVersion: "apps/v1"
    kind:       "Deployment"
    metadata: {
        name:      string
        namespace: string | *"default"
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
                image: string
                ports?: [...{
                    containerPort: int & >=1 & <=65535
                }]
            }]
        }
    }
}

deployment: #KubernetesDeployment & {
    metadata: {
        name: "nginx-deployment"
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
                ports: [{
                    containerPort: 80
                }]
            }]
        }
    }
}
