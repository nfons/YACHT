# YACHT
Yet Another Config Handling Tool

Using Go Templating engine to quickly generate env specific yaml files
### Why:
Kubernetes Config management was a pain, especially when dealing with multiple environments
that might contain some small changes.

Suppose we have a dev and Prod clusters

    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: MY_APP
      label: [THIS VALUE MIGHT CHANGE]
    spec:
      replicas: 1
        spec:
          containers:
          - name: MY_APP
            image: SOME_IMAGE
            ports:
            - containerPort: 80
            env:
            - name: app_env
              value: [ THIS VALUE WILL CHANGE ]

The above example would require 2x YAMLS for dev/prod, by using yacht, you can reduce down to 1


---


### How:

`go get -u github.com/nfons/yacht`

Write your yaml files as you normally do, but define each variable as go template variables:


    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: {{.HTTPSVC_APPNAME}}
    spec:
      replicas: 1
      selector:
        matchLabels:
          app: {{.HTTPSVC_APPNAME}}
      template:
        metadata:
          labels:
            app: {{.HTTPSVC_APPNAME}}
        spec:
          containers:
          - name: app
            image: {{.HTTPSVC_IMAGE}}
            ports:
            - containerPort: 8080



create an env file (i.e dev.conf) with the specific variables you want:

    HTTPSVC_IMAGE=gcr.io/google_containers/echoserver:1.8
    HTTPSVC_APPNAME=http-svc

next run yacht:

`yacht -e dev -f YOUR_YAML`

To Test/examples, checkout:
https://github.com/nfons/yacht-test

## License
GNU GENERAL PUBLIC LICENSE
                       Version 3, 29 June 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.