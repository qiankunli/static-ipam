## CNI IPAM 插件实现/CNI IPAM plugin

支持指定ip/You can request a specific IP address through Kubernetes annotations

代码拷贝自已有的 cni ipam 实现——static

copy from [CNI static ipam plugin](https://github.com/containernetworking/plugins/blob/master/plugins/ipam/static/main.go)


## 配置/Example configurations

[ipam/static](https://github.com/containernetworking/plugins/tree/master/plugins/ipam/static)

### set kubeconfig in k8s.go

### k8s 网络配置 `/etc/cni/net.d/network.json`


    {
        "cniVersion": "0.3.1",
        "name": "macnet",
        "type": "macvlan",
        "master": "eno3.96",
        "ipam": {
    
            "type": "static-ipam",
                    "addresses": [
                        {
                                "address": "172.31.204.19/16",    // return default ip if there is error
                                "gateway": "172.31.0.1"
                        }
                    ],
            "routes": [
                { "dst": "0.0.0.0/0" }
            ]
        }
    }

  
### k8s deployment.yaml 

    apiVersion: apps/v1beta2
    kind: Deployment
    metadata:
      name: hello-world
    spec:
      replicas: 1
      selector:
        matchLabels:
          appName: hello-world
      strategy:
        rollingUpdate:
          maxSurge: 25%
          maxUnavailable: 25%
        type: RollingUpdate
      template:
        metadata:
          labels:
            appName: hello-world
            type: web
          name: hello-world
          annotations:
            static-ipam.ipv4: "172.31.204.12"
          ...



    
    



