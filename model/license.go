// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	DayInSeconds      = 24 * 60 * 60
	DayInMilliseconds = DayInSeconds * 1000

	ExpiredLicenseError = "api.license.add_license.expired.app_error"
	InvalidLicenseError = "api.license.add_license.invalid.app_error"
	LicenseGracePeriod  = DayInMilliseconds * 10 //10 days
	LicenseRenewalLink  = "https://mattermost.com/renew/"

	LicenseShortSkuE10          = "E10"
	LicenseShortSkuE20          = "E20"
	LicenseShortSkuProfessional = "professional"
	LicenseShortSkuEnterprise   = "enterprise"
)

const (
	LicenseUpForRenewalEmailSent = "LicenseUpForRenewalEmailSent"
)

var (
	trialDuration      = 30*(time.Hour*24) + (time.Hour * 8)                                            // 720 hours (30 days) + 8 hours is trial license duration
	adminTrialDuration = 30*(time.Hour*24) + (time.Hour * 23) + (time.Minute * 59) + (time.Second * 59) // 720 hours (30 days) + 23 hours, 59 mins and 59 seconds

	// a sanctioned trial's duration is either more than the upper bound,
	// or less than the lower bound
	sanctionedTrialDurationLowerBound = 31*(time.Hour*24) + (time.Hour * 23) + (time.Minute * 59) + (time.Second * 59) // 744 hours (31 days) + 23 hours, 59 mins and 59 seconds
	sanctionedTrialDurationUpperBound = 29*(time.Hour*24) + (time.Hour * 23) + (time.Minute * 59) + (time.Second * 59) // 696 hours (29 days) + 23 hours, 59 mins and 59 seconds
)

type LicenseRecord struct {
	Id       string `json:"id"`
	CreateAt int64  `json:"create_at"`
	Bytes    string `json:"-"`
}

type License struct {
	Id           string    `json:"id"`
	IssuedAt     int64     `json:"issued_at"`
	StartsAt     int64     `json:"starts_at"`
	ExpiresAt    int64     `json:"expires_at"`
	Customer     *Customer `json:"customer"`
	Features     *Features `json:"features"`
	SkuName      string    `json:"sku_name"`
	SkuShortName string    `json:"sku_short_name"`
	IsTrial      bool      `json:"is_trial"`
	IsGovSku     bool      `json:"is_gov_sku"`
	SignupJWT    *string   `json:"signup_jwt"`
}

type Customer struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

type TrialLicenseRequest struct {
	ServerID              string `json:"server_id"`
	Email                 string `json:"email"`
	Name                  string `json:"name"`
	SiteURL               string `json:"site_url"`
	SiteName              string `json:"site_name"`
	Users                 int    `json:"users"`
	TermsAccepted         bool   `json:"terms_accepted"`
	ReceiveEmailsAccepted bool   `json:"receive_emails_accepted"`
}

func (l *License) IsExpired() bool {
	return l.ExpiresAt < GetMillis()
}

func (l *License) IsPastGracePeriod() bool {
	timeDiff := GetMillis() - l.ExpiresAt
	return timeDiff > LicenseGracePeriod
}

func (l *License) IsWithinExpirationPeriod() bool {
	days := l.DaysToExpiration()
	return days <= 60 && days >= 58
}

func (l *License) DaysToExpiration() int {
	dif := l.ExpiresAt - GetMillis()
	d, _ := time.ParseDuration(fmt.Sprint(dif) + "ms")
	days := d.Hours() / 24
	return int(days)
}

func (l *License) IsStarted() bool {
	return l.StartsAt < GetMillis()
}

func (l *License) IsTrialLicense() bool {
	return l.IsTrial || (l.ExpiresAt-l.StartsAt) == trialDuration.Milliseconds() || (l.ExpiresAt-l.StartsAt) == adminTrialDuration.Milliseconds()
}

func (l *License) IsSanctionedTrial() bool {
	duration := l.ExpiresAt - l.StartsAt

	return l.IsTrialLicense() &&
		(duration >= sanctionedTrialDurationLowerBound.Milliseconds() || duration <= sanctionedTrialDurationUpperBound.Milliseconds())
}

// NewTestLicense returns a license that expires in the future and has the given features.
func NewTestLicense(features ...string) *License {
	ret := &License{
		ExpiresAt: GetMillis() + 90*DayInMilliseconds,
		Customer:  &Customer{},
		Features:  &Features{},
	}
	ret.Features.SetDefaults()

	featureMap := map[string]bool{}
	for _, feature := range features {
		featureMap[feature] = true
	}
	featureJson, _ := json.Marshal(featureMap)
	json.Unmarshal(featureJson, &ret.Features)

	return ret
}

// NewTestLicense returns a license that expires in the future and set as false the given features.
func NewTestLicenseWithFalseDefaults(features ...string) *License {
	ret := &License{
		ExpiresAt: GetMillis() + 90*DayInMilliseconds,
		Customer:  &Customer{},
		Features:  &Features{},
	}
	ret.Features.SetDefaults()

	featureMap := map[string]bool{}
	for _, feature := range features {
		featureMap[feature] = false
	}
	featureJson, _ := json.Marshal(featureMap)
	json.Unmarshal(featureJson, &ret.Features)

	return ret
}

func NewTestLicenseSKU(skuShortName string, features ...string) *License {
	lic := NewTestLicense(features...)
	lic.SkuShortName = skuShortName
	return lic
}

func (lr *LicenseRecord) IsValid() *AppError {
	if !IsValidId(lr.Id) {
		return NewAppError("LicenseRecord.IsValid", "model.license_record.is_valid.id.app_error", nil, "", http.StatusBadRequest)
	}

	if lr.CreateAt == 0 {
		return NewAppError("LicenseRecord.IsValid", "model.license_record.is_valid.create_at.app_error", nil, "", http.StatusBadRequest)
	}

	if lr.Bytes == "" || len(lr.Bytes) > 10000 {
		return NewAppError("LicenseRecord.IsValid", "model.license_record.is_valid.create_at.app_error", nil, "", http.StatusBadRequest)
	}

	return nil
}

func (lr *LicenseRecord) PreSave() {
	lr.CreateAt = GetMillis()
}
