package v7_test

import (
	"code.cloudfoundry.org/cli/v9/command/flag"
	v7 "code.cloudfoundry.org/cli/v9/command/v7"
	"code.cloudfoundry.org/cli/v9/command/v7/v7fakes"
	"code.cloudfoundry.org/cli/v9/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("unset-label command", func() {
	var (
		cmd             v7.UnsetLabelCommand
		resourceName    string
		fakeLabelSetter *v7fakes.FakeLabelUnsetter

		executeErr error
	)

	BeforeEach(func() {
		fakeLabelSetter = new(v7fakes.FakeLabelUnsetter)
		cmd = v7.UnsetLabelCommand{
			LabelUnsetter: fakeLabelSetter,
		}

		cmd.RequiredArgs = flag.UnsetLabelArgs{
			ResourceType: "anything",
			ResourceName: resourceName,
			LabelKeys:    []string{"FOO", "ENV"},
		}
		cmd.BuildpackStack = "some-stack"
		cmd.ServiceBroker = "some-service-broker"
		cmd.ServiceOffering = "some-service-offering"
	})

	It("calls execute with the right parameters", func() {
		executeErr = cmd.Execute(nil)

		Expect(executeErr).ToNot(HaveOccurred())
		Expect(fakeLabelSetter.ExecuteCallCount()).To(Equal(1))
		targetResource, keys := fakeLabelSetter.ExecuteArgsForCall(0)
		Expect(targetResource.ResourceType).To(Equal(cmd.RequiredArgs.ResourceType))
		Expect(targetResource.ResourceName).To(Equal(cmd.RequiredArgs.ResourceName))
		Expect(targetResource.BuildpackStack).To(Equal(cmd.BuildpackStack))
		Expect(targetResource.ServiceBroker).To(Equal(cmd.ServiceBroker))
		Expect(targetResource.ServiceOffering).To(Equal(cmd.ServiceOffering))
		Expect(keys).To(Equal(map[string]types.NullString{
			"FOO": types.NewNullString(),
			"ENV": types.NewNullString(),
		}))
	})
})
