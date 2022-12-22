// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLicenseIsExpired(t *testing.T) {
	l1 := License{}
	l1.ExpiresAt = GetMillis() - 1000
	assert.True(t, l1.IsExpired())

	l1.ExpiresAt = GetMillis() + 10000
	assert.False(t, l1.IsExpired())
}

func TestLicenseIsPastGracePeriod(t *testing.T) {
	l1 := License{}
	l1.ExpiresAt = GetMillis() - LicenseGracePeriod - 1000
	assert.True(t, l1.IsPastGracePeriod())

	l1.ExpiresAt = GetMillis() + 1000
	assert.False(t, l1.IsPastGracePeriod())
}

func TestLicenseIsStarted(t *testing.T) {
	l1 := License{}
	l1.StartsAt = GetMillis() - 1000

	assert.True(t, l1.IsStarted())

	l1.StartsAt = GetMillis() + 10000
	assert.False(t, l1.IsStarted())
}

func TestLicenseRecordIsValid(t *testing.T) {
	lr := LicenseRecord{
		CreateAt: GetMillis(),
		Bytes:    "asdfghjkl;",
	}

	appErr := lr.IsValid()
	assert.NotNil(t, appErr)

	lr.Id = NewId()
	lr.CreateAt = 0
	appErr = lr.IsValid()
	assert.NotNil(t, appErr)

	lr.CreateAt = GetMillis()
	lr.Bytes = ""
	appErr = lr.IsValid()
	assert.NotNil(t, appErr)

	lr.Bytes = strings.Repeat("0123456789", 1001)
	appErr = lr.IsValid()
	assert.NotNil(t, appErr)

	lr.Bytes = "ASDFGHJKL;"
	appErr = lr.IsValid()
	assert.Nil(t, appErr)
}

func TestLicenseRecordPreSave(t *testing.T) {
	lr := LicenseRecord{}
	lr.PreSave()

	assert.NotZero(t, lr.CreateAt)
}

func TestLicense_IsTrialLicense(t *testing.T) {
	t.Run("detect trial license directly from the flag", func(t *testing.T) {
		license := &License{
			IsTrial: true,
		}
		assert.True(t, license.IsTrial)

		license.IsTrial = false
		assert.False(t, license.IsTrialLicense())
	})

	t.Run("detect trial license form duration", func(t *testing.T) {
		startDate, err := time.Parse(time.RFC822, "01 Jan 21 00:00 UTC")
		assert.NoError(t, err)

		endDate, err := time.Parse(time.RFC822, "31 Jan 21 08:00 UTC")
		assert.NoError(t, err)

		license := &License{
			StartsAt:  startDate.UnixNano() / int64(time.Millisecond),
			ExpiresAt: endDate.UnixNano() / int64(time.Millisecond),
		}
		assert.True(t, license.IsTrialLicense())

		endDate, err = time.Parse(time.RFC822, "01 Feb 21 08:00 UTC")
		assert.NoError(t, err)

		license.ExpiresAt = endDate.UnixNano() / int64(time.Millisecond)
		assert.False(t, license.IsTrialLicense())

		// 30 days + 23 hours 59 mins 59 seconds
		endDate, err = time.Parse("02 Jan 06 15:04:05 MST", "31 Jan 21 23:59:59 UTC")
		assert.NoError(t, err)
		license.ExpiresAt = endDate.UnixNano() / int64(time.Millisecond)
		assert.True(t, license.IsTrialLicense())
	})

	t.Run("detect trial with both flag and duration", func(t *testing.T) {
		startDate, err := time.Parse(time.RFC822, "01 Jan 21 00:00 UTC")
		assert.NoError(t, err)

		endDate, err := time.Parse(time.RFC822, "31 Jan 21 08:00 UTC")
		assert.NoError(t, err)

		license := &License{
			IsTrial:   true,
			StartsAt:  startDate.UnixNano() / int64(time.Millisecond),
			ExpiresAt: endDate.UnixNano() / int64(time.Millisecond),
		}

		assert.True(t, license.IsTrialLicense())
		license.IsTrial = false

		// detecting trial from duration
		assert.True(t, license.IsTrialLicense())

		endDate, _ = time.Parse(time.RFC822, "1 Feb 2021 08:00 UTC")
		license.ExpiresAt = endDate.UnixNano() / int64(time.Millisecond)
		assert.False(t, license.IsTrialLicense())

		license.IsTrial = true
		assert.True(t, license.IsTrialLicense())
	})
}

func TestLicense_IsSanctionedTrial(t *testing.T) {
	t.Run("short duration sanctioned trial", func(t *testing.T) {
		startDate, err := time.Parse(time.RFC822, "01 Jan 21 00:00 UTC")
		assert.NoError(t, err)

		endDate, err := time.Parse(time.RFC822, "08 Jan 21 08:00 UTC")
		assert.NoError(t, err)

		license := &License{
			IsTrial:   true,
			StartsAt:  startDate.UnixNano() / int64(time.Millisecond),
			ExpiresAt: endDate.UnixNano() / int64(time.Millisecond),
		}

		assert.True(t, license.IsSanctionedTrial())

		license.IsTrial = false
		assert.False(t, license.IsSanctionedTrial())
	})

	t.Run("long duration sanctioned trial", func(t *testing.T) {
		startDate, err := time.Parse(time.RFC822, "01 Jan 21 00:00 UTC")
		assert.NoError(t, err)

		endDate, err := time.Parse(time.RFC822, "02 Feb 21 08:00 UTC")
		assert.NoError(t, err)

		license := &License{
			IsTrial:   true,
			StartsAt:  startDate.UnixNano() / int64(time.Millisecond),
			ExpiresAt: endDate.UnixNano() / int64(time.Millisecond),
		}

		assert.True(t, license.IsSanctionedTrial())

		license.IsTrial = false
		assert.False(t, license.IsSanctionedTrial())
	})

	t.Run("invalid duration for sanctioned trial", func(t *testing.T) {
		startDate, err := time.Parse(time.RFC822, "01 Jan 21 00:00 UTC")
		assert.NoError(t, err)

		endDate, err := time.Parse(time.RFC822, "31 Jan 21 08:00 UTC")
		assert.NoError(t, err)

		license := &License{
			IsTrial:   true,
			StartsAt:  startDate.UnixNano() / int64(time.Millisecond),
			ExpiresAt: endDate.UnixNano() / int64(time.Millisecond),
		}

		assert.False(t, license.IsSanctionedTrial())
	})

	t.Run("boundary conditions for sanctioned trial", func(t *testing.T) {
		startDate, err := time.Parse(time.RFC822, "01 Jan 21 00:00 UTC")
		assert.NoError(t, err)

		// 29 days + 23 hours 59 mins 59 seconds
		endDate, err := time.Parse("02 Jan 06 15:04:05 MST", "30 Jan 21 23:59:59 UTC")
		assert.NoError(t, err)

		license := &License{
			IsTrial:   true,
			StartsAt:  startDate.UnixNano() / int64(time.Millisecond),
			ExpiresAt: endDate.UnixNano() / int64(time.Millisecond),
		}

		assert.True(t, license.IsSanctionedTrial())

		// 31 days + 23 hours 59 mins 59 seconds
		endDate, err = time.Parse("02 Jan 06 15:04:05 MST", "01 Feb 21 23:59:59 UTC")
		assert.NoError(t, err)
		license.ExpiresAt = endDate.UnixNano() / int64(time.Millisecond)
		assert.True(t, license.IsSanctionedTrial())
	})
}
