module github.com/rook/cassandra

go 1.16

require (
	github.com/banzaicloud/k8s-objectmatcher v1.1.0
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f
	github.com/csi-addons/volume-replication-operator v0.1.1-0.20210525040814-ab575a2879fb
	github.com/davecgh/go-spew v1.1.1
	github.com/elazarl/goproxy v0.0.0-20191011121108-aa519ddbe484 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/k8snetworkplumbingwg/network-attachment-definition-client v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/tevino/abool v1.2.0
	github.com/yanniszark/go-nodetool v0.0.0-20191206125106-cd8f91fa16be
	golang.org/x/oauth2 v0.0.0-20201208152858-08078c50e5b5 // indirect
	google.golang.org/genproto v0.0.0-20201201144952-b05cb90ed32e // indirect
	google.golang.org/grpc v1.33.2 // indirect
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.21.2
	k8s.io/apiextensions-apiserver v0.21.1
	k8s.io/apimachinery v0.21.2
	k8s.io/apiserver v0.21.1
	k8s.io/client-go v0.21.2
	k8s.io/cloud-provider v0.21.1
	k8s.io/utils v0.0.0-20210527160623-6fdb442a123b
	sigs.k8s.io/controller-runtime v0.9.0
)

replace (
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.1
	github.com/kubernetes-incubator/external-storage => github.com/libopenstorage/external-storage v0.20.4-openstorage-rc3
	github.com/portworx/sched-ops => github.com/portworx/sched-ops v0.20.4-openstorage-rc3
	google.golang.org/grpc => google.golang.org/grpc v1.27.1
)

exclude (
	// This tag doesn't exist, but is imported by github.com/portworx/sched-ops.
	github.com/kubernetes-incubator/external-storage v0.20.4-openstorage-rc2
	// Exclude pre-go-mod kubernetes tags, because they are older
	// than v0.x releases but are picked when updating dependencies.
	k8s.io/client-go v1.4.0
	k8s.io/client-go v1.5.0
	k8s.io/client-go v1.5.1
	k8s.io/client-go v1.5.2
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/client-go v2.0.0+incompatible
	k8s.io/client-go v2.0.0-alpha.1+incompatible
	k8s.io/client-go v3.0.0+incompatible
	k8s.io/client-go v3.0.0-beta.0+incompatible
	k8s.io/client-go v4.0.0+incompatible
	k8s.io/client-go v4.0.0-beta.0+incompatible
	k8s.io/client-go v5.0.0+incompatible
	k8s.io/client-go v5.0.1+incompatible
	k8s.io/client-go v6.0.0+incompatible
	k8s.io/client-go v7.0.0+incompatible
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/client-go v9.0.0+incompatible
	k8s.io/client-go v9.0.0-invalid+incompatible
)
