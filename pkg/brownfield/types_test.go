// -------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// --------------------------------------------------------------------------------------------

package brownfield

import (
	n "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-03-01/network"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Azure/application-gateway-kubernetes-ingress/pkg/tests/fixtures"
)

var _ = Describe("Test NewExistingResources", func() {

	prohibitedTargets := fixtures.GetAzureIngressProhibitedTargets()

	Context("Test NewExistingResources", func() {
		It("should create the correct struct", func() {
			appGw := n.ApplicationGateway{
				ApplicationGatewayPropertiesFormat: &n.ApplicationGatewayPropertiesFormat{},
			}
			defaultPool := n.ApplicationGatewayBackendAddressPool{}

			actual := NewExistingResources(appGw, prohibitedTargets, &defaultPool)
			expected := ExistingResources{
				ProhibitedTargets:  prohibitedTargets,
				DefaultBackendPool: &n.ApplicationGatewayBackendAddressPool{},
			}
			Expect(actual).To(Equal(expected))
		})
	})

	Context("Test getProhibitedHostNames", func() {
		It("should create a struct", func() {
			appGw := n.ApplicationGateway{
				ApplicationGatewayPropertiesFormat: &n.ApplicationGatewayPropertiesFormat{},
			}
			defaultPool := n.ApplicationGatewayBackendAddressPool{}
			er := NewExistingResources(appGw, prohibitedTargets, &defaultPool)
			actual := er.getProhibitedHostNames()
			expected := map[string]interface{}{
				"bye.com":                 nil,
				"--some-other-hostname--": nil,
				"*.hi.com":                nil,
			}
			Expect(actual).To(Equal(expected))
		})
	})

})
