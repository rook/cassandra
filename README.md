<img alt="Rook" src="Documentation/media/logo.svg" width="50%" height="50%">

[![CNCF Status](https://img.shields.io/badge/cncf%20status-graduated-blue.svg)](https://www.cncf.io/projects)
[![Docker Pulls](https://img.shields.io/docker/pulls/rook/cassandra)](https://hub.docker.com/u/rook)
[![Go Report Card](https://goreportcard.com/badge/github.com/rook/cassandra)](https://goreportcard.com/report/github.com/rook/cassandra)
[![Slack](https://slack.rook.io/badge.svg)](https://slack.rook.io)

# Deprecated

The Rook Cassandra operator has been deprecated due to lack of community support. We recommend deploying Cassandra with the [k8ssandra/cass-operator](https://github.com/k8ssandra/cass-operator/) which is under more active development by the community.

## What is Rook?

Rook is an open source **cloud-native storage orchestrator** for Kubernetes, providing the platform, framework, and support for a diverse set of storage solutions to natively integrate with cloud-native environments.

Rook turns storage software into self-managing, self-scaling, and self-healing storage services. It does this by automating deployment, bootstrapping, configuration, provisioning, scaling, upgrading, migration, disaster recovery, monitoring, and resource management. Rook uses the facilities provided by the underlying cloud-native container management, scheduling and orchestration platform to perform its duties.

Rook integrates deeply into cloud native environments leveraging extension points and providing a seamless experience for scheduling, lifecycle management, resource management, security, monitoring, and user experience.

For more details about the storage solutions currently supported by Rook, please refer to the [project status section](#project-status) below.
We plan to continue adding support for other storage systems and environments based on community demand and engagement in future releases. See our [roadmap](ROADMAP.md) for more details.

Rook is hosted by the [Cloud Native Computing Foundation](https://cncf.io) (CNCF) as a [graduated](https://www.cncf.io/announcements/2020/10/07/cloud-native-computing-foundation-announces-rook-graduation/) level project. If you are a company that wants to help shape the evolution of technologies that are container-packaged, dynamically-scheduled and microservices-oriented, consider joining the CNCF. For details about who's involved and how Rook plays a role, read the CNCF [announcement](https://www.cncf.io/blog/2018/01/29/cncf-host-rook-project-cloud-native-storage-capabilities).

## Contact

Please use the following to reach members of the community:

- Slack: Join our [slack channel](https://slack.rook.io)
- Forums: [rook-dev](https://groups.google.com/forum/#!forum/rook-dev)
- Twitter: [@rook_io](https://twitter.com/rook_io)
- Email (general topics): [cncf-rook-info@lists.cncf.io](mailto:cncf-rook-info@lists.cncf.io)
- Email (security topics): [cncf-rook-security@lists.cncf.io](mailto:cncf-rook-security@lists.cncf.io)

## Project Status

**Deprecated**

The status of each storage provider supported by Rook can be found in the [main Rook repo](https://github.com/rook/rook#project-status).

| Name      | Details                                                                                                                                                                                                                                                                                                                | API Group                  | Status |
| --------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------- | ------ |
| Cassandra | [Cassandra](http://cassandra.apache.org/) is a highly available NoSQL database featuring lightning fast performance, tunable consistency and massive scalability. [Scylla](https://www.scylladb.com) is a close-to-the-hardware rewrite of Cassandra in C++, which enables much lower latencies and higher throughput. | cassandra.rook.io/v1alpha1 | [Deprecated](#deprecated)  |

### Official Releases

Official releases of the Cassandra operator can be found on the [releases page](https://github.com/rook/cassandra/releases).
Please note that it is **strongly recommended** that you use [official releases](https://github.com/rook/cassandra/releases) of Rook, as unreleased versions from the master branch are subject to changes and incompatibilities that will not be supported in the official releases.
Builds from the master branch can have functionality changed and even removed at any time without compatibility support and without prior notice.

Releases of the Cassandra operator prior to v1.7 are found in the main [Rook repo](https://github.com/rook/rook/releases).

## Licensing

Rook is under the Apache 2.0 license.
