package datablocks

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"text/template"
)

func TestDatablocksSyncer(t *testing.T) {
	temp := template.Must(template.New("sync-template").Parse("https://sync.v5prebid.datablocks.net/s2ssync?gdpr={{.GDPR}}&gdpr_consent={{.GDPRConsent}}&r=https%3A%2F%2Flocalhost%3A8888%2Fsetuid%3Fbidder%3Ddatablocks%26gdpr%3D{{.GDPR}}%26gdpr_consent%3D{{.GDPRConsent}}%26uid%3D%24%7Buid%7D"))
	syncer := NewDatablocksSyncer(temp)
	syncInfo, err := syncer.GetUsersyncInfo("1", "BONciguONcjGKADACHENAOLS1rAHDAFAAEAASABQAMwAeACEAFw")
	assert.NoError(t, err)
	assert.Equal(t, "https://sync.v5prebid.datablocks.net/s2ssync?gdpr=1&gdpr_consent=BONciguONcjGKADACHENAOLS1rAHDAFAAEAASABQAMwAeACEAFw&r=https%3A%2F%2Flocalhost%3A8888%2Fsetuid%3Fbidder%3Ddatablocks%26gdpr%3D1%26gdpr_consent%3DBONciguONcjGKADACHENAOLS1rAHDAFAAEAASABQAMwAeACEAFw%26uid%3D%24%7Buid%7D", syncInfo.URL)
	assert.Equal(t, "redirect", syncInfo.Type)
	assert.EqualValues(t, datablocksGDPRVendorID, syncer.GDPRVendorID())
	assert.Equal(t, false, syncInfo.SupportCORS)
}
