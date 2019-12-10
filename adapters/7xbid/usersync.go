package bid7x

import (
	"text/template"

	"github.com/prebid/prebid-server/adapters"
	"github.com/prebid/prebid-server/usersync"
)

func New7xbidSyncer(temp *template.Template) usersync.Usersyncer {
	//TODO: Fix it
	return adapters.NewSyncer("7xBid", 58, temp, adapters.SyncTypeIframe)
}
