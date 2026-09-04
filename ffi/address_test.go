package ffi

import "testing"

const (
	test_address  = "+49123456789"
	test_deviceID = 1
)

func TestAddress(t *testing.T) {
	addr, err := NewAddress(test_address, 1)
	if err != nil {
		t.Fatalf("NewAddress(): %v", err)
	}
	defer addr.Destroy()

	name, err := addr.Name()
	if err != nil {
		t.Fatalf("*Address.Name(): %v", err)
	}

	if name != test_address {
		t.Fatalf("'%s' != '%s'", test_address, name)
	}
}
