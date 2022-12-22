// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLicenseFeaturesToMap(t *testing.T) {
	l := License{}
	l.Features.SetDefaults()

	m := l.FeaturesToMap()

	CheckTrue(t, m["ldap"].(bool))
	CheckTrue(t, m["ldap_groups"].(bool))
	CheckTrue(t, m["mfa"].(bool))
	CheckTrue(t, m["google"].(bool))
	CheckTrue(t, m["office365"].(bool))
	CheckTrue(t, m["compliance"].(bool))
	CheckTrue(t, m["cluster"].(bool))
	CheckTrue(t, m["metrics"].(bool))
	CheckTrue(t, m["mhpns"].(bool))
	CheckTrue(t, m["saml"].(bool))
	CheckTrue(t, m["elastic_search"].(bool))
	CheckTrue(t, m["email_notification_contents"].(bool))
	CheckTrue(t, m["data_retention"].(bool))
	CheckTrue(t, m["message_export"].(bool))
	CheckTrue(t, m["custom_permissions_schemes"].(bool))
	CheckTrue(t, m["id_loaded"].(bool))
	CheckTrue(t, m["future"].(bool))
	CheckTrue(t, m["shared_channels"].(bool))
	CheckTrue(t, m["remote_cluster_service"].(bool))
}

func TestLicenseFeaturesSetDefaults(t *testing.T) {
	f := Features{}
	f.SetDefaults()

	CheckInt(t, *f.Users, 0)
	CheckTrue(t, *f.LDAP)
	CheckTrue(t, *f.LDAPGroups)
	CheckTrue(t, *f.MFA)
	CheckTrue(t, *f.GoogleOAuth)
	CheckTrue(t, *f.Office365OAuth)
	CheckTrue(t, *f.Compliance)
	CheckTrue(t, *f.Cluster)
	CheckTrue(t, *f.Metrics)
	CheckTrue(t, *f.MHPNS)
	CheckTrue(t, *f.SAML)
	CheckTrue(t, *f.Elasticsearch)
	CheckTrue(t, *f.EmailNotificationContents)
	CheckTrue(t, *f.DataRetention)
	CheckTrue(t, *f.MessageExport)
	CheckTrue(t, *f.CustomPermissionsSchemes)
	CheckTrue(t, *f.GuestAccountsPermissions)
	CheckTrue(t, *f.IDLoadedPushNotifications)
	CheckTrue(t, *f.SharedChannels)
	CheckTrue(t, *f.RemoteClusterService)
	CheckTrue(t, *f.FutureFeatures)

	f = Features{}
	f.SetDefaults()

	*f.Users = 300
	*f.FutureFeatures = false
	*f.LDAP = true
	*f.LDAPGroups = true
	*f.MFA = true
	*f.GoogleOAuth = true
	*f.Office365OAuth = true
	*f.Compliance = true
	*f.Cluster = true
	*f.Metrics = true
	*f.MHPNS = true
	*f.SAML = true
	*f.Elasticsearch = true
	*f.DataRetention = true
	*f.MessageExport = true
	*f.CustomPermissionsSchemes = true
	*f.GuestAccounts = true
	*f.GuestAccountsPermissions = true
	*f.EmailNotificationContents = true
	*f.IDLoadedPushNotifications = true
	*f.SharedChannels = true

	f.SetDefaults()

	CheckInt(t, *f.Users, 300)
	CheckTrue(t, *f.LDAP)
	CheckTrue(t, *f.LDAPGroups)
	CheckTrue(t, *f.MFA)
	CheckTrue(t, *f.GoogleOAuth)
	CheckTrue(t, *f.Office365OAuth)
	CheckTrue(t, *f.Compliance)
	CheckTrue(t, *f.Cluster)
	CheckTrue(t, *f.Metrics)
	CheckTrue(t, *f.MHPNS)
	CheckTrue(t, *f.SAML)
	CheckTrue(t, *f.Elasticsearch)
	CheckTrue(t, *f.EmailNotificationContents)
	CheckTrue(t, *f.DataRetention)
	CheckTrue(t, *f.MessageExport)
	CheckTrue(t, *f.CustomPermissionsSchemes)
	CheckTrue(t, *f.GuestAccounts)
	CheckTrue(t, *f.GuestAccountsPermissions)
	CheckTrue(t, *f.IDLoadedPushNotifications)
	CheckTrue(t, *f.SharedChannels)
	CheckTrue(t, *f.RemoteClusterService)
	CheckFalse(t, *f.FutureFeatures)
}

func TestIsCloud(t *testing.T) {
	l1 := License{}
	l1.Features = &Features{}
	l1.Features.SetDefaults()
	assert.False(t, l1.IsCloud())

	boolTrue := true
	l1.Features.Cloud = &boolTrue
	assert.True(t, l1.IsCloud())

	var license *License
	assert.False(t, license.IsCloud())

	l1.Features = nil
	assert.False(t, l1.IsCloud())

	t.Run("false if license is nil", func(t *testing.T) {
		var license *License
		assert.False(t, license.IsCloud())
	})
}

func TestHasLDAP(t *testing.T) {
	boolTrue := true
	t.Run("false if license is nil", func(t *testing.T) {
		var license *License
		assert.False(t, license.HasLDAP())
	})

	l1 := License{}
	assert.False(t, l1.HasLDAP())
	l1.Features = &Features{}
	assert.False(t, l1.HasLDAP())
	l1.Features.SetDefaults()
	assert.False(t, l1.HasLDAP())
	l1.Features.LDAP = &boolTrue
	assert.True(t, l1.HasLDAP())

	assert.False(t, (&License{SkuShortName: LicenseShortSkuE10}).HasLDAP())
	assert.False(t, (&License{SkuShortName: LicenseShortSkuE20}).HasLDAP())
	assert.False(t, (&License{SkuShortName: LicenseShortSkuProfessional}).HasLDAP())
	assert.False(t, (&License{SkuShortName: LicenseShortSkuEnterprise}).HasLDAP())
}
