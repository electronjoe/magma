---
id: heap_profile
title: Heap Profiling MME
hide_title: true
---
# Heap Profiling

## Tooling

We will use the Gperftools [heap profiler](https://gperftools.github.io/gperftools/heapprofile.html).

## Environment Setup

```shell script
vagrant up magma
vagrant ssh magma
sudo apt-get update
sudo apt-get install google-perftools
sudo ln -s /usr/lib/libprofiler.so.0 /usr/lib/libprofiler.so
sudo ln -s /usr/lib/libtcmalloc.so.4 /usr/lib/libtcmalloc.so
```

## Building

```shell script
make
```
