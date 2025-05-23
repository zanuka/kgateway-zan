// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	gentype "k8s.io/client-go/gentype"

	apiv1alpha1 "github.com/kgateway-dev/kgateway/v2/api/applyconfiguration/api/v1alpha1"
	v1alpha1 "github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
	typedapiv1alpha1 "github.com/kgateway-dev/kgateway/v2/pkg/client/clientset/versioned/typed/api/v1alpha1"
)

// fakeGatewayExtensions implements GatewayExtensionInterface
type fakeGatewayExtensions struct {
	*gentype.FakeClientWithListAndApply[*v1alpha1.GatewayExtension, *v1alpha1.GatewayExtensionList, *apiv1alpha1.GatewayExtensionApplyConfiguration]
	Fake *FakeGatewayV1alpha1
}

func newFakeGatewayExtensions(fake *FakeGatewayV1alpha1, namespace string) typedapiv1alpha1.GatewayExtensionInterface {
	return &fakeGatewayExtensions{
		gentype.NewFakeClientWithListAndApply[*v1alpha1.GatewayExtension, *v1alpha1.GatewayExtensionList, *apiv1alpha1.GatewayExtensionApplyConfiguration](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("gatewayextensions"),
			v1alpha1.SchemeGroupVersion.WithKind("GatewayExtension"),
			func() *v1alpha1.GatewayExtension { return &v1alpha1.GatewayExtension{} },
			func() *v1alpha1.GatewayExtensionList { return &v1alpha1.GatewayExtensionList{} },
			func(dst, src *v1alpha1.GatewayExtensionList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.GatewayExtensionList) []*v1alpha1.GatewayExtension {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.GatewayExtensionList, items []*v1alpha1.GatewayExtension) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
