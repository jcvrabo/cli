package v7action_test

import (
	"errors"

	"code.cloudfoundry.org/cli/v9/actor/actionerror"
	. "code.cloudfoundry.org/cli/v9/actor/v7action"
	"code.cloudfoundry.org/cli/v9/actor/v7action/v7actionfakes"
	"code.cloudfoundry.org/cli/v9/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/v9/api/cloudcontroller/ccv3"
	"code.cloudfoundry.org/cli/v9/resources"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Space Manifest Actions", func() {
	var (
		actor                     *Actor
		fakeCloudControllerClient *v7actionfakes.FakeCloudControllerClient
	)

	BeforeEach(func() {
		fakeCloudControllerClient = new(v7actionfakes.FakeCloudControllerClient)
		actor = NewActor(fakeCloudControllerClient, nil, nil, nil, nil, nil)
	})

	Describe("DiffSpaceManifest", func() {
		var (
			spaceGUID   string
			rawManifest []byte

			diff       resources.ManifestDiff
			warnings   Warnings
			executeErr error
		)

		BeforeEach(func() {
			spaceGUID = "some-space-guid"
			rawManifest = []byte("---\n- applications:\n name: my-app")
		})

		JustBeforeEach(func() {
			diff, warnings, executeErr = actor.DiffSpaceManifest(spaceGUID, rawManifest)
		})

		When("getting the diff succeeds", func() {
			BeforeEach(func() {
				returnedDiff := resources.ManifestDiff{
					Diffs: []resources.Diff{
						{Op: resources.AddOperation, Path: "/some/path", Value: "wow"},
					},
				}

				fakeCloudControllerClient.GetSpaceManifestDiffReturns(
					returnedDiff,
					ccv3.Warnings{"diff-manifest-warning"},
					nil,
				)
			})

			It("returns the diff and warnings", func() {
				Expect(executeErr).NotTo(HaveOccurred())
				Expect(diff).To(Equal(resources.ManifestDiff{
					Diffs: []resources.Diff{
						{Op: resources.AddOperation, Path: "/some/path", Value: "wow"},
					},
				}))
				Expect(warnings).To(ConsistOf("diff-manifest-warning"))
			})
		})

		When("getting the diff errors", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.GetSpaceManifestDiffReturns(
					resources.ManifestDiff{},
					ccv3.Warnings{"diff-manifest-warning"},
					errors.New("diff-manifest-error"),
				)
			})

			It("returns the error and warnings", func() {
				Expect(executeErr).To(MatchError("diff-manifest-error"))
				Expect(warnings).To(ConsistOf("diff-manifest-warning"))
			})
		})
	})

	Describe("SetSpaceManifest", func() {
		var (
			spaceGUID   string
			rawManifest []byte

			warnings   Warnings
			executeErr error
		)

		BeforeEach(func() {
			spaceGUID = "some-space-guid"
			rawManifest = []byte("---\n- applications:\n name: my-app")
		})

		JustBeforeEach(func() {
			warnings, executeErr = actor.SetSpaceManifest(spaceGUID, rawManifest)
		})

		When("applying the manifest succeeds", func() {
			BeforeEach(func() {
				fakeCloudControllerClient.UpdateSpaceApplyManifestReturns(
					"some-job-url",
					ccv3.Warnings{"apply-manifest-1-warning"},
					nil,
				)
			})

			When("polling finishes successfully", func() {
				BeforeEach(func() {
					fakeCloudControllerClient.PollJobReturns(
						ccv3.Warnings{"poll-1-warning"},
						nil,
					)
				})

				It("uploads the app manifest", func() {
					Expect(executeErr).ToNot(HaveOccurred())
					Expect(warnings).To(ConsistOf("apply-manifest-1-warning", "poll-1-warning"))

					Expect(fakeCloudControllerClient.UpdateSpaceApplyManifestCallCount()).To(Equal(1))
					guidInCall, appManifest := fakeCloudControllerClient.UpdateSpaceApplyManifestArgsForCall(0)
					Expect(guidInCall).To(Equal("some-space-guid"))
					Expect(appManifest).To(Equal(rawManifest))

					Expect(fakeCloudControllerClient.PollJobCallCount()).To(Equal(1))
					jobURL := fakeCloudControllerClient.PollJobArgsForCall(0)
					Expect(jobURL).To(Equal(ccv3.JobURL("some-job-url")))
				})
			})

			When("polling returns a generic error", func() {
				var expectedErr error

				BeforeEach(func() {
					expectedErr = errors.New("poll-1-error")
					fakeCloudControllerClient.PollJobReturns(
						ccv3.Warnings{"poll-1-warning"},
						expectedErr,
					)
				})

				It("reports a polling error", func() {
					Expect(executeErr).To(Equal(expectedErr))
					Expect(warnings).To(ConsistOf("apply-manifest-1-warning", "poll-1-warning"))
				})
			})

			When("polling returns an job failed error", func() {
				var expectedErr error

				BeforeEach(func() {
					expectedErr = ccerror.V3JobFailedError{Detail: "some-job-failed"}
					fakeCloudControllerClient.PollJobReturns(
						ccv3.Warnings{"poll-1-warning"},
						expectedErr,
					)
				})

				It("reports a polling error", func() {
					Expect(executeErr).To(Equal(actionerror.ApplicationManifestError{Message: "some-job-failed"}))
					Expect(warnings).To(ConsistOf("apply-manifest-1-warning", "poll-1-warning"))
				})
			})
		})

		When("applying the manifest errors", func() {
			var applyErr error

			BeforeEach(func() {
				applyErr = errors.New("some-apply-manifest-error")
				fakeCloudControllerClient.UpdateSpaceApplyManifestReturns(
					"",
					ccv3.Warnings{"apply-manifest-1-warning"},
					applyErr,
				)
			})

			It("reports a error trying to apply the manifest", func() {
				Expect(executeErr).To(Equal(applyErr))
				Expect(warnings).To(ConsistOf("apply-manifest-1-warning"))
			})
		})
	})
})
