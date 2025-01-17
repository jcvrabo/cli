package v7

import (
	"code.cloudfoundry.org/cli/v9/actor/actionerror"
	"code.cloudfoundry.org/cli/v9/command/flag"
)

type UnshareRouteCommand struct {
	BaseCommand

	RequireArgs      flag.Domain      `positional-args:"yes"`
	Hostname         string           `long:"hostname" short:"n" description:"Hostname for the HTTP route (required for shared domains)"`
	Path             flag.V7RoutePath `long:"path" description:"Path for the HTTP route"`
	DestinationOrg   string           `short:"o" description:"The org of the destination space (Default: targeted org)"`
	DestinationSpace string           `short:"s" description:"The space to be unshared (Default: targeted space)"`

	relatedCommands interface{} `related_commands:" share-route, delete-route, map-route, unmap-route, routes"`
}

func (cmd UnshareRouteCommand) Usage() string {
	return `
Unshare an existing route from a space:
	CF_NAME unshare-route DOMAIN [--hostname HOSTNAME] [--path PATH] -s OTHER_SPACE [-o OTHER_ORG]`
}

func (cmd UnshareRouteCommand) Examples() string {
	return `
CF_NAME unshare-route example.com --hostname myHost --path foo -s TargetSpace -o TargetOrg        # myhost.example.com/foo
CF_NAME unshare-route example.com --hostname myHost -s TargetSpace                                # myhost.example.com
CF_NAME unshare-route example.com --hostname myHost -s TargetSpace -o TargetOrg                   # myhost.example.com`
}

func (cmd UnshareRouteCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(true, true)
	if err != nil {
		return err
	}

	user, err := cmd.Actor.GetCurrentUser()
	if err != nil {
		return err
	}

	domain, warnings, err := cmd.Actor.GetDomainByName(cmd.RequireArgs.Domain)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}

	path := cmd.Path.Path
	route, warnings, err := cmd.Actor.GetRouteByAttributes(domain, cmd.Hostname, path, 0)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		if _, ok := err.(actionerror.RouteNotFoundError); ok {
			cmd.UI.DisplayText("Can not unshare route:")
			return err
		}
	}

	destinationOrgName := cmd.DestinationOrg

	if destinationOrgName == "" {
		destinationOrgName = cmd.Config.TargetedOrganizationName()
	}

	destinationOrg, warnings, err := cmd.Actor.GetOrganizationByName(destinationOrgName)
	cmd.UI.DisplayWarnings(warnings)

	if err != nil {
		if _, ok := err.(actionerror.OrganizationNotFoundError); ok {
			cmd.UI.DisplayText("Can not unshare route:")
			return err
		}
	}

	targetedSpace, warnings, err := cmd.Actor.GetSpaceByNameAndOrganization(cmd.DestinationSpace, destinationOrg.GUID)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		if _, ok := err.(actionerror.SpaceNotFoundError); ok {
			cmd.UI.DisplayText("Can not unshare route:")
			return err
		}
	}

	url := desiredURL(domain.Name, cmd.Hostname, path, 0)
	cmd.UI.DisplayTextWithFlavor("Unsharing route {{.URL}} from space {{.DestinationSpace}} as {{.User}}",
		map[string]interface{}{
			"URL":              url,
			"DestinationSpace": cmd.DestinationSpace,
			"User":             user.Name,
		})
	warnings, err = cmd.Actor.UnshareRoute(
		route.GUID,
		targetedSpace.GUID,
	)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}
	cmd.UI.DisplayOK()

	return nil
}
