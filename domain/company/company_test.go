package company

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	expectedCompany := &Company{
		Id:   "comp_1",
		Name: "NextComp",
	}

	newCompany := Create("comp_1", "NextComp")
	require.Equal(t, expectedCompany, newCompany)
}
