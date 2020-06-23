---
title: Medium Kubernetes Common Library
category: library
---
# Medium Kubernetes Common Library

This library contains common code used across multiple kubernetes related projects.

# Matcher-Gen

`matcher-gen` is a kubernetes code generator that creates match asserters for
kubernetes resources that ignore attributes that aren't important when testing.
Here's an example of creating match assertions for the kubernetes core/v1 types
in the go.medium.engineering/kubernetes/test/core/v1 package:

```!bash
$ go get go.medium.engineering/kubernetes/cmd/matcher-gen
$ matcher-gen -i k8s.io/api/core/v1 -p go.medium.engeering/kubernetes/test/core/v1
```

An `AssertMatch` function in go.medium.engineering/kubernetes/test is used to
check if a resource is available and matches via a client. Custom typed Match
and UnMatch asserts are registered with a comparator. A default comparator is
provided and the matcher-generated code automatically registers generated
assertions with it, you just have to import the generated package somewhere. If
you'd like to manage your own comparater, you can pass it to the
RegisterComparators method created by the code generator.
