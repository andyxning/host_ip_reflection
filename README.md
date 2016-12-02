![](https://img.shields.io/badge/LICENSE-AGPL-blue.svg)
[![Build Status](https://travis-ci.org/andyxning/host_ip_reflection.svg?branch=master)](https://travis-ci.org/andyxning/host_ip_reflection)

### host_ip_reflection
--
`host_ip_reflection` can answer the question "What Is My IP Address?".
[ifconfig.me](ifconfig.me) is the well-known public service that can be used to query
the local client or outward ip in a NAT environment.

`host_ip_reflection` supports HTTP API to query the client ip or outward ip in a NAT
environment.

`host_ip_reflection` also supports `X-Real-Ip` HTTP Header. This makes `host_ip_reflection`
can be deployed behind a reverse proxy such as `nginx` or `haproxy`.

### API
--
|API|Description|HTTP Method|Response|
|---|-----------|-----------|--------|
|/|query local or outward ip in a NAT environment|GET|Success wtih response code `200` and json response body {"IP":client_node_ip}. Failure with none 200 response code.|

### Example
--
```
>> curl http://127.0.0.1:3087
{"ip":"127.0.0.1"}
```

### Usage
--
* In Kubernetes, if we use `NodePort` service type. We can query `host_ip_reflection` to learn
the node ip where the container is deployed.
  * Note that Kubernetes 1.4 added the ability to learn the nodename for container using downward
    in [PR 27880](https://github.com/kubernetes/kubernetes/pull/27880). However, if you can not upgrade to this version and still need to learn the node ip, you can achieve this by curl `host_ip_reflection` api in container.

