# Docker Compose Developer Info

## Building all Containers

For e.g. integ-test.

```shell
cd cwf/gateway/docker
docker-compose --file docker-compose.yml --file docker-compose.override.yml --file docker-compose.integ-test.yml -- build --parallel
```

You can validate images have been build.

```shell
➜  docker git:(pr-docker-compose-docs) ✗ docker image ls                                                                                                                     
REPOSITORY                                   TAG       IMAGE ID       CREATED          SIZE
cwf_gateway_pipelined                        latest    0597a986c606   18 minutes ago   1.5GB
cwf_gateway_python                           latest    c45fa7d642a9   19 minutes ago   1.21GB
cwf_gateway_go                               latest    d4584e3e9435   20 minutes ago   811MB
cwf_gateway_sessiond                         latest    c2ff545ae959   21 minutes ago   936MB
cwf_cwag_go                                  latest    326e80190b69   21 minutes ago   195MB
```

## Spin up just pipelined

```shell
docker run -i -t cwf_gateway_pipelined:latest /bin/bash
```

### Generate OVS DB

```shell
ovsdb-tool create /etc/openvswitch/conf.db /usr/share/openvswitch/vswitch.ovsschema
mkdir /var/run/openvswitch
ovsdb-server --unixctl=/var/run/openvswitch/db.sock --detach
```

### Setting up ovs

```shell
set bridge cwag_br0 protocols=protocols=OpenFlow10,OpenFlow13,OpenFlow14 other-config:disable-in-band=true
/usr/bin/ovs-vsctl set-controller cwag_br0 tcp:127.0.0.1:6633
```

## TODO

**Quick Q - looks like pipelined is broken out from the general gateway_python, any idea why?**

```
REPOSITORY                                   TAG       IMAGE ID       CREATED          SIZE
cwf_gateway_pipelined                        latest    0597a986c606   18 minutes ago   1.5GB
cwf_gateway_python                           latest    c45fa7d642a9   19 minutes ago   1.21GB
```

> Yeah we should probably unify these. The reasoning is that orc8r/gateway/python services were containerized first for the feg. cwf_gateway_python actually is created via the dockerfile at magma/feg/gateway/docker/python and contains only the services at magma/orc8r/gateway/python. The cwf_gateway_pipelined contains all of the services in lte/gateway/python. It shouldn't be named pipelined but probably something more generic as it actually can run more than just the pipelined service.

**Also - in exploring this a bit by hand - looks like docker-compose.integ-test.yml when running pipelined executes some ovs-vsctl commands inside the container. We running OVS inside of here? How does that interact with the parent environment?**

> Yeah this is part of the reason the image is so bloated. We have to install ovs inside the container so that the ovs-vsctl commands work. But we mount the host ovs db into the container so that it works properly

**Does the host spin up ONLY the database, or does the host also run both database and switch?**

> The db and the switch

Looks like openflow mabye `MME` <-> `pipelined` and then `ovs-vsctl` for `pipelined` <-> `ovs db` which then gets applied

## Resources

- [Great Article on ovsdb](https://relaxdiego.com/2014/09/ovsdb.html)
