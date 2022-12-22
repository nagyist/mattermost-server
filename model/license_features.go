// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

// Reading features from the struct is deprecated.
// Instead we prefer each product vertical owning the definition
// of what feature is enabled, gated by checking against the license.SkuShortName
// See https://docs.google.com/document/d/1PhlIiGXbAl-sTRFXLRT5R71eSzUTA0HXAVNXPu1aAVU/edit#
// for more information.

// Where possible, we can define access to features in this file to prevent littering
// the codebases with repetitive checks and possible drift between repeated checks.

type Features struct {
	Users                     *int  `json:"users"`
	LDAP                      *bool `json:"ldap"`
	LDAPGroups                *bool `json:"ldap_groups"`
	MFA                       *bool `json:"mfa"`
	GoogleOAuth               *bool `json:"google_oauth"`
	Office365OAuth            *bool `json:"office365_oauth"`
	OpenId                    *bool `json:"openid"`
	Compliance                *bool `json:"compliance"`
	Cluster                   *bool `json:"cluster"`
	Metrics                   *bool `json:"metrics"`
	MHPNS                     *bool `json:"mhpns"`
	SAML                      *bool `json:"saml"`
	Elasticsearch             *bool `json:"elastic_search"`
	Announcement              *bool `json:"announcement"`
	ThemeManagement           *bool `json:"theme_management"`
	EmailNotificationContents *bool `json:"email_notification_contents"`
	DataRetention             *bool `json:"data_retention"`
	MessageExport             *bool `json:"message_export"`
	CustomPermissionsSchemes  *bool `json:"custom_permissions_schemes"`
	CustomTermsOfService      *bool `json:"custom_terms_of_service"`
	GuestAccounts             *bool `json:"guest_accounts"`
	GuestAccountsPermissions  *bool `json:"guest_accounts_permissions"`
	IDLoadedPushNotifications *bool `json:"id_loaded"`
	LockTeammateNameDisplay   *bool `json:"lock_teammate_name_display"`
	EnterprisePlugins         *bool `json:"enterprise_plugins"`
	AdvancedLogging           *bool `json:"advanced_logging"`
	Cloud                     *bool `json:"cloud"`
	SharedChannels            *bool `json:"shared_channels"`
	RemoteClusterService      *bool `json:"remote_cluster_service"`

	// after we enabled more features we'll need to control them with this
	FutureFeatures *bool `json:"future_features"`
}

func (l *License) FeaturesToMap() map[string]any {
	return map[string]any{
		"ldap":                        l.HasLDAP(),
		"ldap_groups":                 l != nil && *l.Features.LDAPGroups,
		"mfa":                         l != nil && *l.Features.MFA,
		"google":                      l != nil && *l.Features.GoogleOAuth,
		"office365":                   l != nil && *l.Features.Office365OAuth,
		"openid":                      l != nil && *l.Features.OpenId,
		"compliance":                  l != nil && *l.Features.Compliance,
		"cluster":                     l != nil && *l.Features.Cluster,
		"metrics":                     l != nil && *l.Features.Metrics,
		"mhpns":                       l != nil && *l.Features.MHPNS,
		"saml":                        l != nil && *l.Features.SAML,
		"elastic_search":              l != nil && *l.Features.Elasticsearch,
		"email_notification_contents": l != nil && *l.Features.EmailNotificationContents,
		"data_retention":              l != nil && *l.Features.DataRetention,
		"message_export":              l != nil && *l.Features.MessageExport,
		"custom_permissions_schemes":  l != nil && *l.Features.CustomPermissionsSchemes,
		"guest_accounts":              l != nil && *l.Features.GuestAccounts,
		"guest_accounts_permissions":  l != nil && *l.Features.GuestAccountsPermissions,
		"id_loaded":                   l != nil && *l.Features.IDLoadedPushNotifications,
		"lock_teammate_name_display":  l != nil && *l.Features.LockTeammateNameDisplay,
		"enterprise_plugins":          l.HasEnterpriseMarketplacePlugins(),
		"advanced_logging":            l != nil && *l.Features.AdvancedLogging,
		"cloud":                       l != nil && *l.Features.Cloud,
		"shared_channels":             l != nil && *l.Features.SharedChannels,
		"remote_cluster_service":      l != nil && *l.Features.RemoteClusterService,
		"future":                      l != nil && *l.Features.FutureFeatures,
	}
}

func (f *Features) SetDefaults() {
	if f.FutureFeatures == nil {
		f.FutureFeatures = NewBool(true)
	}

	if f.Users == nil {
		f.Users = NewInt(0)
	}

	if f.LDAP == nil {
		f.LDAP = NewBool(*f.FutureFeatures)
	}

	if f.LDAPGroups == nil {
		f.LDAPGroups = NewBool(*f.FutureFeatures)
	}

	if f.MFA == nil {
		f.MFA = NewBool(*f.FutureFeatures)
	}

	if f.GoogleOAuth == nil {
		f.GoogleOAuth = NewBool(*f.FutureFeatures)
	}

	if f.Office365OAuth == nil {
		f.Office365OAuth = NewBool(*f.FutureFeatures)
	}

	if f.OpenId == nil {
		f.OpenId = NewBool(*f.FutureFeatures)
	}

	if f.Compliance == nil {
		f.Compliance = NewBool(*f.FutureFeatures)
	}

	if f.Cluster == nil {
		f.Cluster = NewBool(*f.FutureFeatures)
	}

	if f.Metrics == nil {
		f.Metrics = NewBool(*f.FutureFeatures)
	}

	if f.MHPNS == nil {
		f.MHPNS = NewBool(*f.FutureFeatures)
	}

	if f.SAML == nil {
		f.SAML = NewBool(*f.FutureFeatures)
	}

	if f.Elasticsearch == nil {
		f.Elasticsearch = NewBool(*f.FutureFeatures)
	}

	if f.Announcement == nil {
		f.Announcement = NewBool(true)
	}

	if f.ThemeManagement == nil {
		f.ThemeManagement = NewBool(true)
	}

	if f.EmailNotificationContents == nil {
		f.EmailNotificationContents = NewBool(*f.FutureFeatures)
	}

	if f.DataRetention == nil {
		f.DataRetention = NewBool(*f.FutureFeatures)
	}

	if f.MessageExport == nil {
		f.MessageExport = NewBool(*f.FutureFeatures)
	}

	if f.CustomPermissionsSchemes == nil {
		f.CustomPermissionsSchemes = NewBool(*f.FutureFeatures)
	}

	if f.GuestAccounts == nil {
		f.GuestAccounts = NewBool(*f.FutureFeatures)
	}

	if f.GuestAccountsPermissions == nil {
		f.GuestAccountsPermissions = NewBool(*f.FutureFeatures)
	}

	if f.CustomTermsOfService == nil {
		f.CustomTermsOfService = NewBool(*f.FutureFeatures)
	}

	if f.IDLoadedPushNotifications == nil {
		f.IDLoadedPushNotifications = NewBool(*f.FutureFeatures)
	}

	if f.LockTeammateNameDisplay == nil {
		f.LockTeammateNameDisplay = NewBool(*f.FutureFeatures)
	}

	if f.EnterprisePlugins == nil {
		f.EnterprisePlugins = NewBool(*f.FutureFeatures)
	}

	if f.AdvancedLogging == nil {
		f.AdvancedLogging = NewBool(*f.FutureFeatures)
	}

	if f.Cloud == nil {
		f.Cloud = NewBool(false)
	}

	if f.SharedChannels == nil {
		f.SharedChannels = NewBool(*f.FutureFeatures)
	}

	if f.RemoteClusterService == nil {
		f.RemoteClusterService = NewBool(*f.FutureFeatures)
	}
}

func (l *License) IsCloud() bool {
	return l != nil && l.Features != nil && l.Features.Cloud != nil && *l.Features.Cloud
}

func (l *License) HasEnterpriseMarketplacePlugins() bool {
	if l == nil {
		return false
	}
	return (l.Features != nil && l.Features.EnterprisePlugins != nil && *l.Features.EnterprisePlugins) ||
		l.SkuShortName == LicenseShortSkuE20 ||
		l.SkuShortName == LicenseShortSkuProfessional ||
		l.SkuShortName == LicenseShortSkuEnterprise
}

func (l *License) HasLDAP() bool {
	if l == nil {
		return false
	}
	return (l.Features != nil && l.Features.LDAP != nil && *l.Features.LDAP) ||
		l.SkuShortName == LicenseShortSkuE10 ||
		l.SkuShortName == LicenseShortSkuE20 ||
		l.SkuShortName == LicenseShortSkuProfessional ||
		l.SkuShortName == LicenseShortSkuEnterprise
}
