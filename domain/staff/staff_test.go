package staff

import (
    "testing"
    "time"

    "github.com/devit-tel/gotime"
    "github.com/stretchr/testify/require"
)

func freezeTime() time.Time {
    now := time.Now()
    gotime.Freeze(now)
    return now

}

func TestCreate(t *testing.T) {
    now := freezeTime()
    expectedStaff := &Staff{Id: "staff_1", CompanyId: "comp_1", Name: "Tester", Tel: "081-555-4444", CreatedAt: now.Unix(), UpdatedAt: now.Unix()}

    newStaff := Create("staff_1", "comp_1", "Tester", "081-555-4444")
    require.Equal(t, expectedStaff, newStaff)
}

func TestStaff_Update(t *testing.T) {
    now := freezeTime()
    expectedStaff := &Staff{Id: "staff_1", CompanyId: "comp_1", Name: "Tester2", Tel: "081-555-6666", CreatedAt: now.Unix(), UpdatedAt: now.Unix()}

    newStaff := Create("staff_1", "comp_1", "Tester1", "081-555-4444")

    newStaff.Update("Tester2", "081-555-6666")
    require.Equal(t, expectedStaff, newStaff)
}
