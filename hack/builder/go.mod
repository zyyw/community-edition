module github.com/vmware-tanzu/community-edition/hack/builder

go 1.16

// -- Used to build TCE plugins with local codebase --
// replace github.com/vmware-tanzu/tanzu-framework => ../../../tanzu-framework

require (
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/spf13/cobra v1.2.1
	github.com/vmware-labs/yaml-jsonpath v0.3.2 // indirect
	github.com/vmware-tanzu/carvel-secretgen-controller v0.5.0 // indirect
	github.com/vmware-tanzu/tanzu-framework v0.2.1
	golang.org/x/sys v0.0.0-20210910150752-751e447fb3d0 // indirect
	k8s.io/klog/v2 v2.9.0 // indirect
	sigs.k8s.io/controller-runtime v0.9.0 // indirect
)
