// Code generated by assertions generator; DO NOT EDIT.

package objectassert

import (
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/assert"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/helpers"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/internal/collections"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
)

type ResourceMonitorAssert struct {
	*assert.SnowflakeObjectAssert[sdk.ResourceMonitor, sdk.AccountObjectIdentifier]
}

func ResourceMonitor(t *testing.T, id sdk.AccountObjectIdentifier) *ResourceMonitorAssert {
	t.Helper()
	return &ResourceMonitorAssert{
		assert.NewSnowflakeObjectAssertWithTestClientObjectProvider(sdk.ObjectTypeResourceMonitor, id, func(testClient *helpers.TestClient) assert.ObjectProvider[sdk.ResourceMonitor, sdk.AccountObjectIdentifier] {
			return testClient.ResourceMonitor.Show
		}),
	}
}

func ResourceMonitorFromObject(t *testing.T, resourceMonitor *sdk.ResourceMonitor) *ResourceMonitorAssert {
	t.Helper()
	return &ResourceMonitorAssert{
		assert.NewSnowflakeObjectAssertWithObject(sdk.ObjectTypeResourceMonitor, resourceMonitor.ID(), resourceMonitor),
	}
}

func (r *ResourceMonitorAssert) HasName(expected string) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		if o.Name != expected {
			return fmt.Errorf("expected name: %v; got: %v", expected, o.Name)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasCreditQuota(expected float64) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		if o.CreditQuota != expected {
			return fmt.Errorf("expected credit quota: %v; got: %v", expected, o.CreditQuota)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasUsedCredits(expected float64) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		if o.UsedCredits != expected {
			return fmt.Errorf("expected used credits: %v; got: %v", expected, o.UsedCredits)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasRemainingCredits(expected float64) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		if o.RemainingCredits != expected {
			return fmt.Errorf("expected remaining credits: %v; got: %v", expected, o.RemainingCredits)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasLevel(expected sdk.ResourceMonitorLevel) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		if o.Level == nil {
			return fmt.Errorf("expected level to have value; got: nil")
		}
		if *o.Level != expected {
			return fmt.Errorf("expected level: %v; got: %v", expected, *o.Level)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasFrequency(expected sdk.Frequency) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		if o.Frequency != expected {
			return fmt.Errorf("expected frequency: %v; got: %v", expected, o.Frequency)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasStartTime(expected string) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		if o.StartTime != expected {
			return fmt.Errorf("expected start time: %v; got: %v", expected, o.StartTime)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasEndTime(expected string) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		if o.EndTime != expected {
			return fmt.Errorf("expected end time: %v; got: %v", expected, o.EndTime)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasNotifyAt(expected ...int) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		mapped := collections.Map(o.NotifyAt, func(item int) any { return item })
		mappedExpected := collections.Map(expected, func(item int) any { return item })
		if !slices.Equal(mapped, mappedExpected) {
			return fmt.Errorf("expected notify at: %v; got: %v", expected, o.NotifyAt)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasSuspendAt(expected int) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		if o.SuspendAt == nil {
			return fmt.Errorf("expected suspend at to have value; got: nil")
		}
		if *o.SuspendAt != expected {
			return fmt.Errorf("expected suspend at: %v; got: %v", expected, *o.SuspendAt)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasSuspendImmediateAt(expected int) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		if o.SuspendImmediateAt == nil {
			return fmt.Errorf("expected suspend immediate at to have value; got: nil")
		}
		if *o.SuspendImmediateAt != expected {
			return fmt.Errorf("expected suspend immediate at: %v; got: %v", expected, *o.SuspendImmediateAt)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasCreatedOn(expected time.Time) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		if o.CreatedOn != expected {
			return fmt.Errorf("expected created on: %v; got: %v", expected, o.CreatedOn)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasOwner(expected string) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		if o.Owner != expected {
			return fmt.Errorf("expected owner: %v; got: %v", expected, o.Owner)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasComment(expected string) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		if o.Comment != expected {
			return fmt.Errorf("expected comment: %v; got: %v", expected, o.Comment)
		}
		return nil
	})
	return r
}

func (r *ResourceMonitorAssert) HasNotifyUsers(expected ...string) *ResourceMonitorAssert {
	r.AddAssertion(func(t *testing.T, o *sdk.ResourceMonitor) error {
		t.Helper()
		mapped := collections.Map(o.NotifyUsers, func(item string) any { return item })
		mappedExpected := collections.Map(expected, func(item string) any { return item })
		if !slices.Equal(mapped, mappedExpected) {
			return fmt.Errorf("expected notify users: %v; got: %v", expected, o.NotifyUsers)
		}
		return nil
	})
	return r
}
