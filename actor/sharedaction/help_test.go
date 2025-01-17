package sharedaction_test

import (
	"code.cloudfoundry.org/cli/v9/actor/actionerror"
	. "code.cloudfoundry.org/cli/v9/actor/sharedaction"
	"code.cloudfoundry.org/cli/v9/actor/sharedaction/sharedactionfakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type commandList struct {
	App     appCommand         `command:"app" description:"Display health and status for an app"`
	Restage restageCommand     `command:"restage" alias:"rg" description:"Restage an app"`
	Help    helpCommand        `command:"help" alias:"h" description:"Show help"`
	Usage   usageMethodCommand `command:"fancy"`
}

type appCommand struct {
	GUID            bool        `long:"guid" description:"Retrieve and display the given app's guid.  All other health and status output for the app is suppressed." default:"some-default"`
	usage           interface{} `usage:"CF_NAME app APP_NAME"`
	relatedCommands interface{} `related_commands:"apps, events, logs, map-route, unmap-route, push"`
}

type restageCommand struct {
	envCFStagingTimeout interface{} `environmentName:"CF_STAGING_TIMEOUT" environmentDescription:"Max wait time for buildpack staging, in minutes" environmentDefault:"15"`
	envCFStartupTimeout interface{} `environmentName:"CF_STARTUP_TIMEOUT" environmentDescription:"Max wait time for app instance startup, in minutes" environmentDefault:"5"`
}

type helpCommand struct {
	AllCommands bool        `short:"a" description:"All available CLI commands"`
	usage       interface{} `usage:"CF_NAME help [COMMAND]"`
}

type usageMethodCommand struct {
}

func (c usageMethodCommand) Usage() string {
	return `
Usage line 1
Usage line 2
`
}

func (c usageMethodCommand) Examples() string {
	return `
Examples line 1
Examples line 2
`
}

func (c usageMethodCommand) Resources() string {
	return `
Resource 1
Resource 2
`
}

var _ = Describe("Help Actions", func() {
	var actor *Actor

	BeforeEach(func() {
		actor = NewActor(&sharedactionfakes.FakeConfig{})
	})

	Describe("CommandInfoByName", func() {
		When("the command exists", func() {
			When("passed the command name", func() {
				It("returns command info", func() {
					commandInfo, err := actor.CommandInfoByName(commandList{}, "app")
					Expect(err).NotTo(HaveOccurred())

					Expect(commandInfo.Name).To(Equal("app"))
					Expect(commandInfo.Description).To(Equal("Display health and status for an app"))
					Expect(commandInfo.Alias).To(BeEmpty())
					Expect(commandInfo.Usage).To(Equal("CF_NAME app APP_NAME"))
					Expect(commandInfo.Flags).To(HaveLen(1))
					Expect(commandInfo.Flags).To(ContainElement(CommandFlag{
						Short:       "",
						Long:        "guid",
						Description: "Retrieve and display the given app's guid.  All other health and status output for the app is suppressed.",
						Default:     "some-default",
					}))
					Expect(commandInfo.RelatedCommands).To(Equal([]string{
						"apps", "events", "logs", "map-route", "push", "unmap-route",
					}))
				})

				When("the command uses timeout environment variables", func() {
					It("has timeout environment variables", func() {
						commandInfo, err := actor.CommandInfoByName(commandList{}, "restage")
						Expect(err).NotTo(HaveOccurred())

						Expect(commandInfo.Environment).To(ConsistOf(
							EnvironmentVariable{
								Name:         "CF_STAGING_TIMEOUT",
								Description:  "Max wait time for buildpack staging, in minutes",
								DefaultValue: "15",
							},
							EnvironmentVariable{
								Name:         "CF_STARTUP_TIMEOUT",
								Description:  "Max wait time for app instance startup, in minutes",
								DefaultValue: "5",
							}))
					})
				})

				When("the command does not use environment variables", func() {
					It("does not have environment variables", func() {
						commandInfo, err := actor.CommandInfoByName(commandList{}, "app")
						Expect(err).NotTo(HaveOccurred())

						Expect(commandInfo.Environment).To(BeEmpty())
					})
				})

				When("the command has a Usage() method", func() {
					It("retrieves the usage text from the method", func() {
						commandInfo, err := actor.CommandInfoByName(commandList{}, "fancy")
						Expect(err).NotTo(HaveOccurred())

						Expect(commandInfo.Usage).To(Equal("Usage line 1\n   Usage line 2"))
					})
				})

				When("the command has a Examples() method", func() {
					It("retrieves the examples text from the method", func() {
						commandInfo, err := actor.CommandInfoByName(commandList{}, "fancy")
						Expect(err).NotTo(HaveOccurred())

						Expect(commandInfo.Examples).To(Equal("Examples line 1\n   Examples line 2"))
					})
				})

				When("the command has a Resources() method", func() {
					It("retrieves the resources text from the method", func() {
						commandInfo, err := actor.CommandInfoByName(commandList{}, "fancy")
						Expect(err).NotTo(HaveOccurred())

						Expect(commandInfo.Resources).To(Equal("Resource 1\n   Resource 2"))
					})
				})
			})

			When("passed the command alias", func() {
				It("returns command info", func() {
					commandInfo, err := actor.CommandInfoByName(commandList{}, "h")
					Expect(err).NotTo(HaveOccurred())

					Expect(commandInfo.Name).To(Equal("help"))
					Expect(commandInfo.Description).To(Equal("Show help"))
					Expect(commandInfo.Alias).To(Equal("h"))
					Expect(commandInfo.Usage).To(Equal("CF_NAME help [COMMAND]"))
					Expect(commandInfo.Flags).To(ConsistOf(
						CommandFlag{
							Short:       "a",
							Long:        "",
							Description: "All available CLI commands",
						},
					))
				})
			})
		})

		When("the command does not exist", func() {
			It("returns err", func() {
				_, err := actor.CommandInfoByName(commandList{}, "does-not-exist")

				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(actionerror.InvalidCommandError{CommandName: "does-not-exist"}))
			})
		})
	})

	Describe("CommandInfos", func() {
		It("returns back all the command's names and descriptions", func() {
			commands := actor.CommandInfos(commandList{})

			Expect(commands["app"]).To(Equal(CommandInfo{
				Name:        "app",
				Description: "Display health and status for an app",
			}))
			Expect(commands["help"]).To(Equal(CommandInfo{
				Name:        "help",
				Description: "Show help",
				Alias:       "h",
			}))
			Expect(commands["restage"]).To(Equal(CommandInfo{
				Name:        "restage",
				Description: "Restage an app",
				Alias:       "rg",
			}))
		})
	})
})
