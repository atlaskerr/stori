[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/2374/badge)](https://bestpractices.coreinfrastructure.org/projects/2374)

# Stori Registry

Stori is an OCI-compliant image registry that aims to be the successor to the Docker registry.

## Table of Contents
- [Features](#features)
- [Philosophy](#philosophy)

## Features
- __Enhanced API:__ Stori provides a rich API with first-class support for registry administration tasks.

## Philosophy
The Stori project is meant to be a reference implementation of, and serve as a testbed for improvements to, the OCI Distribution Specification.
Stori also solves another problem in the current state of container affairs, the need for an extendable, lightweight, and enterprise-level image registry.

The standard open-source Docker registry provides a core API that only supports the distribution of images with minimal support for maintenance or administration.
The Docker registry also lacks "enterprise" features like auditing, LDAP, and fine-grained access control policies.
The Harbor project is a popular image registry hosted by the CNCF that is built on top of the Docker registry but has many limitations.
Since Harbor uses the Docker registry under the covers, its administrative API is limited to the functionality that the Docker registry API provides.
Also, Harbor is non-trivial to set up.
The only reliable way of installing Harbor on a local machine is to use a custom offline installer that depends on docker-compose and consumes 5GB+ of diskspace when fully installed.
When running Harbor on a small VM like an AWS `t3.micro`, an EBS volume needs to be provisioned for the image storage directory or else there will be no space left on the filesystem.
The Stori project attempts to let registry administrators "have their cake and eat it too" when it comes to having an easily deployable and maintainable enterprise-level image registry with a small footprint.
