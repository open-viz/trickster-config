module go.openviz.dev/trickster-config

require (
	github.com/cert-manager/cert-manager v1.13.3
	github.com/elastic/go-elasticsearch/v7 v7.15.1
	go.bytebuilders.dev/audit 3ff33160c6f02f6151e59cdd44dd50a347c02ba0
	go.bytebuilders.dev/license-proxyserver 31122ab825027d2495c9320b63d99660f1ca56be
	go.bytebuilders.dev/license-verifier v0.13.4
	go.bytebuilders.dev/license-verifier/kubernetes v0.13.4
	go.mongodb.org/mongo-driver v1.10.2
	gomodules.xyz/logs v0.0.7
	gomodules.xyz/password-generator v0.2.9
	kmodules.xyz/client-go v0.29.4
	kmodules.xyz/go-containerregistry v0.0.12
	kmodules.xyz/resource-metadata v0.18.1
	kubedb.dev/apimachinery a1d475ceb73e12977cce84eb7564393b9ae9b6e3
	kubedb.dev/db-client-go v0.0.8
	kubestash.dev/apimachinery cc46ddfd674a760d87ec2fe4122f7816296654c8
)

replace github.com/Masterminds/sprig/v3 => github.com/gomodules/sprig/v3 v3.2.3-0.20220405051441-0a8a99bac1b8

replace sigs.k8s.io/controller-runtime => github.com/kmodules/controller-runtime ac-0.17.0

replace github.com/imdario/mergo => github.com/imdario/mergo v0.3.6

replace k8s.io/apiserver => github.com/kmodules/apiserver ac-1.29.0

replace k8s.io/kubernetes => github.com/kmodules/kubernetes ac-1.29.0
