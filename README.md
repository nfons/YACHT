# YACHT
Yet Another Config Handling Tool:

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

The above example would require 2x YAMLS for dev/prod

---


### How:
Coming soon

## License
GNU GENERAL PUBLIC LICENSE
                       Version 3, 29 June 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.