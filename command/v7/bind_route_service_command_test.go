package v7_test

import (
	"errors"
	"fmt"

	"code.cloudfoundry.org/cli/v9/actor/actionerror"
	"code.cloudfoundry.org/cli/v9/actor/v7action"
	"code.cloudfoundry.org/cli/v9/command/commandfakes"
	v7 "code.cloudfoundry.org/cli/v9/command/v7"
	"code.cloudfoundry.org/cli/v9/command/v7/v7fakes"
	"code.cloudfoundry.org/cli/v9/types"
	"code.cloudfoundry.org/cli/v9/util/configv3"
	"code.cloudfoundry.org/cli/v9/util/ui"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("bind-route-service Command", func() {
	var (
		cmd             v7.BindRouteServiceCommand
		testUI          *ui.UI
		fakeConfig      *commandfakes.FakeConfig
		fakeSharedActor *commandfakes.FakeSharedActor
		executeErr      error
		fakeActor       *v7fakes.FakeActor
	)

	const (
		fakeUserName            = "fake-user-name"
		fakeServiceInstanceName = "fake-service-instance-name"
		fakeDomainName          = "fake-domain-name"
		fakeHostname            = "fake-hostname"
		fakePath                = "fake-path"
		fakeOrgName             = "fake-org-name"
		fakeSpaceName           = "fake-space-name"
		fakeSpaceGUID           = "fake-space-guid"
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(NewBuffer(), NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeActor = new(v7fakes.FakeActor)

		cmd = v7.BindRouteServiceCommand{
			BaseCommand: v7.BaseCommand{
				UI:          testUI,
				Config:      fakeConfig,
				SharedActor: fakeSharedActor,
				Actor:       fakeActor,
			},
		}

		fakeConfig.TargetedSpaceReturns(configv3.Space{
			Name: fakeSpaceName,
			GUID: fakeSpaceGUID,
		})

		fakeConfig.TargetedOrganizationReturns(configv3.Organization{Name: fakeOrgName})

		fakeActor.GetCurrentUserReturns(configv3.User{Name: fakeUserName}, nil)

		fakeActor.CreateRouteBindingReturns(
			nil,
			v7action.Warnings{"fake warning"},
			nil,
		)

		setPositionalFlags(&cmd, fakeDomainName, fakeServiceInstanceName)
		setFlag(&cmd, "--hostname", fakeHostname)
		setFlag(&cmd, "--path", fakePath)
		setFlag(&cmd, "-c", `{"foo": "bar"}`)
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	It("checks the user is logged in, and targeting an org and space", func() {
		Expect(executeErr).NotTo(HaveOccurred())

		Expect(fakeSharedActor.CheckTargetCallCount()).To(Equal(1))
		actualOrg, actualSpace := fakeSharedActor.CheckTargetArgsForCall(0)
		Expect(actualOrg).To(BeTrue())
		Expect(actualSpace).To(BeTrue())
	})

	It("delegates to the actor", func() {
		Expect(fakeActor.CreateRouteBindingCallCount()).To(Equal(1))
		Expect(fakeActor.CreateRouteBindingArgsForCall(0)).To(Equal(v7action.CreateRouteBindingParams{
			SpaceGUID:           fakeSpaceGUID,
			ServiceInstanceName: fakeServiceInstanceName,
			DomainName:          fakeDomainName,
			Hostname:            fakeHostname,
			Path:                fmt.Sprintf("/%s", fakePath),
			Parameters:          types.NewOptionalObject(map[string]interface{}{"foo": "bar"}),
		}))
	})

	Describe("intro message", func() {
		It("prints an intro and warnings", func() {
			Expect(executeErr).NotTo(HaveOccurred())
			Expect(testUI.Err).To(Say("fake warning"))
			Expect(testUI.Out).To(Say(
				`Binding route %s.%s/%s to service instance %s in org %s / space %s as %s\.\.\.\n`,
				fakeHostname,
				fakeDomainName,
				fakePath,
				fakeServiceInstanceName,
				fakeOrgName,
				fakeSpaceName,
				fakeUserName,
			))
		})

		When("no hostname", func() {
			BeforeEach(func() {
				setFlag(&cmd, "--hostname", "")
			})

			It("prints an appropriate intro", func() {
				Expect(testUI.Out).To(Say(
					`Binding route %s/%s to service instance`,
					fakeDomainName,
					fakePath,
				))
			})
		})

		When("no path", func() {
			BeforeEach(func() {
				setFlag(&cmd, "--path", "")
			})

			It("prints an appropriate intro", func() {
				Expect(testUI.Out).To(Say(
					`Binding route %s.%s to service instance`,
					fakeHostname,
					fakeDomainName,
				))
			})
		})
	})

	When("binding already exists", func() {
		BeforeEach(func() {
			fakeActor.CreateRouteBindingReturns(
				nil,
				v7action.Warnings{"fake warning"},
				actionerror.ResourceAlreadyExistsError{},
			)
		})

		It("prints a message and warnings", func() {
			Expect(testUI.Out).To(SatisfyAll(
				Say(
					`Route %s.%s/%s is already bound to service instance %s\.\n`,
					fakeHostname,
					fakeDomainName,
					fakePath,
					fakeServiceInstanceName,
				),
				Say(`OK\n`),
			))

			Expect(testUI.Err).To(Say("fake warning"))
		})
	})

	Describe("processing the response stream", func() {
		Context("nil stream", func() {
			It("prints a message and warnings", func() {
				Expect(testUI.Out).To(Say(`OK\n`))
				Expect(testUI.Err).To(Say("fake warning"))
			})
		})

		Context("stream goes to complete", func() {
			BeforeEach(func() {
				eventStream := make(chan v7action.PollJobEvent)
				go func() {
					eventStream <- v7action.PollJobEvent{
						State:    v7action.JobProcessing,
						Warnings: v7action.Warnings{"job processing warning"},
					}
					eventStream <- v7action.PollJobEvent{
						State:    v7action.JobComplete,
						Warnings: v7action.Warnings{"job complete warning"},
					}
					close(eventStream)
				}()

				fakeActor.CreateRouteBindingReturns(
					eventStream,
					v7action.Warnings{"fake warning"},
					nil,
				)
			})

			It("prints a message and warnings", func() {
				Expect(testUI.Out).To(Say(`OK\n`))

				Expect(testUI.Err).To(SatisfyAll(
					Say("fake warning"),
					Say("job processing warning"),
					Say("job complete warning"),
				))
			})
		})

		Context("stream goes to polling", func() {
			BeforeEach(func() {
				eventStream := make(chan v7action.PollJobEvent)
				go func() {
					eventStream <- v7action.PollJobEvent{
						State:    v7action.JobProcessing,
						Warnings: v7action.Warnings{"job processing warning"},
					}
					eventStream <- v7action.PollJobEvent{
						State:    v7action.JobPolling,
						Warnings: v7action.Warnings{"job polling warning"},
					}
				}()

				fakeActor.CreateRouteBindingReturns(
					eventStream,
					v7action.Warnings{"fake warning"},
					nil,
				)
			})

			It("prints a message and warnings", func() {
				Expect(testUI.Out).To(SatisfyAll(
					Say(`OK\n`),
					Say(`\n`),
					Say(`Binding in progress\.\n`),
				))

				Expect(testUI.Err).To(SatisfyAll(
					Say("fake warning"),
					Say("job processing warning"),
					Say("job polling warning"),
				))
			})
		})

		Context("stream goes to error", func() {
			BeforeEach(func() {
				eventStream := make(chan v7action.PollJobEvent)
				go func() {
					eventStream <- v7action.PollJobEvent{
						State:    v7action.JobFailed,
						Warnings: v7action.Warnings{"job failed warning"},
						Err:      errors.New("boom"),
					}
				}()

				fakeActor.CreateRouteBindingReturns(
					eventStream,
					v7action.Warnings{"fake warning"},
					nil,
				)
			})

			It("prints warnings and returns the error", func() {
				Expect(executeErr).To(MatchError("boom"))

				Expect(testUI.Err).To(SatisfyAll(
					Say("fake warning"),
					Say("job failed warning"),
				))
			})
		})

		When("--wait flag specified", func() {
			BeforeEach(func() {
				eventStream := make(chan v7action.PollJobEvent)
				go func() {
					eventStream <- v7action.PollJobEvent{
						State:    v7action.JobProcessing,
						Warnings: v7action.Warnings{"job processing warning"},
					}
					eventStream <- v7action.PollJobEvent{
						State:    v7action.JobPolling,
						Warnings: v7action.Warnings{"job polling warning"},
					}
					eventStream <- v7action.PollJobEvent{
						State:    v7action.JobComplete,
						Warnings: v7action.Warnings{"job complete warning"},
					}
					close(eventStream)
				}()

				fakeActor.CreateRouteBindingReturns(
					eventStream,
					v7action.Warnings{"fake warning"},
					nil,
				)

				setFlag(&cmd, "--wait")
			})

			It("waits for the event stream to complete", func() {
				Expect(testUI.Out).To(SatisfyAll(
					Say(`Waiting for the operation to complete\.\.\.\n`),
					Say(`\n`),
					Say(`OK\n`),
				))

				Expect(testUI.Err).To(SatisfyAll(
					Say("fake warning"),
					Say("job processing warning"),
					Say("job polling warning"),
					Say("job complete warning"),
				))
			})
		})
	})

	When("checking the target returns an error", func() {
		BeforeEach(func() {
			fakeSharedActor.CheckTargetReturns(errors.New("explode"))
		})

		It("returns the error", func() {
			Expect(executeErr).To(MatchError("explode"))
		})
	})

	When("actor returns error", func() {
		BeforeEach(func() {
			fakeActor.CreateRouteBindingReturns(
				nil,
				v7action.Warnings{"fake warning"},
				errors.New("boom"),
			)
		})

		It("prints warnings and returns the error", func() {
			Expect(testUI.Err).To(Say("fake warning"))
			Expect(executeErr).To(MatchError("boom"))
		})
	})

	When("getting the username returns an error", func() {
		BeforeEach(func() {
			fakeActor.GetCurrentUserReturns(configv3.User{}, errors.New("bad thing"))
		})

		It("returns the error", func() {
			Expect(executeErr).To(MatchError("bad thing"))
		})
	})
})
