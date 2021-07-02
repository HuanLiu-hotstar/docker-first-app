# deploy your server to k8s

- we assume there is an image in local docker engine,name=golang-test:v3

## create a k8s deployment

```sh
# generate a deployment yaml files  
kubectl create deployment first-app --image="golang-test:v3" -o yaml --dry-run=client > first-app.yaml

# the reason generate deployment yaml is that ,we will modify imagePullPolicy in yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: first-app
  name: first-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: first-app
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: first-app
    spec:
      containers:
      - image: golang-test:v3
        name: golang-test
        imagePullPolicy: Never ## Never: use local docker image other than pull docker image from docker-hub 
        resources: {}
status: {}
# now apply deployment
kubectl apply -f first-app.yaml

# find the pod in k8s
kubectl get pods
# find the deployment
kubectl get deployment
# now container is running
docker ps
# output like this:
# CONTAINER ID   IMAGE               COMMAND      CREATED STATUS     PORTS                NAMES
# 0f3afe30a460   5a4d02edae48        "/app/main"  About a minute ago Up About a minute                                                       k8s_golang-test_first-app-6cdc85f58-hk8pv_default_8932e5bd-0d33-494f-b6b0-42e772186b89_0


# expose service outside k8s cluster
kubectl expose deployment/first-app --type="NodePort" --port 9090
# - type:expose type 
# - port: is the server listen port in container

# lookup the service
kubectl get services
# output like this
# NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
# first-app    NodePort    10.105.179.76   <none>        9090:32158/TCP   4s


# now you can access server in your local machine 
#  32158 maps the request to 9090 in pod
curl localhost:32158/world
# server output for you request
# welcome 1 guest world






```

## Upgrade Server

```sh

## build the new image 
#  docker build . -t golang-test:v5
## update the images
## describe in kubectl apply yaml files where images name 
## kubectl set image deployment/deployment-name imagename=new_imagename --record 
kubectl set image deployment/first-app golang-test=golang-test:v5 --record 


# upgrade server 
kubectl rollout status deployments/first-app

```

