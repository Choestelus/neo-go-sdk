package utility_test

import (
	"testing"

	"github.com/Choestelus/neo-go-sdk/utility"
	"github.com/stretchr/testify/assert"
)

func BenchmarkAddressToHash160(b *testing.B) {
	addr := "APyEx5f4Zm4oCHwFWiSTaph1fPBxZacYVR"
	for i := 0; i < b.N; i++ {
		utility.AddressToHash(addr, byte(23))
	}
}

func TestAddressToHash160(t *testing.T) {
	actual, err := utility.AddressToHash("AJShjraX4iMJjwVt8WYYzZyGvDMxw6Xfbe", byte(23))
	assert.NoError(t, err)
	assert.Equal(t, "bfc469dd56932409677278f6b7422f3e1f34481d", actual)
}
