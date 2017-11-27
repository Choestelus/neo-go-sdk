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
	actual, err = utility.AddressToHash("AeidcdeJPNRCevppL2ickxgGozGbvgfzzw", byte(23))
	assert.NoError(t, err)
	assert.Equal(t, "73d32628738060a19b6d04d78766366e82eeadfb", actual)
}

func TestHash160ToAddress(t *testing.T) {
	actual, err := utility.HashToAddress("73d32628738060a19b6d04d78766366e82eeadfb", byte(23))
	assert.NoError(t, err)
	assert.Equal(t, "AeidcdeJPNRCevppL2ickxgGozGbvgfzzw", actual)
	actual, err = utility.HashToAddress("bfc469dd56932409677278f6b7422f3e1f34481d", byte(23))
	assert.NoError(t, err)
	assert.Equal(t, "AJShjraX4iMJjwVt8WYYzZyGvDMxw6Xfbe", actual)
}
