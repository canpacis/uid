package uid_test

import (
	"fmt"
	"testing"

	"github.com/CanPacis/uid"
)

func TestUID(t *testing.T) {
	id := uid.New()
	if len(id) != 12 {
		t.Errorf("expected uid byte array length to be 12 but got %d", len(id))
	}
	if len(id.String()) != 16 {
		t.Errorf("expected uid string length to be 12 but got %d", len(id))
	}
}

func TestParse(t *testing.T) {
	idStr := "abcdefghijklmnop"
	id, err := uid.Parse(idStr)
	if err != nil {
		t.Error(err)
	}
	if id.String() != idStr {
		t.Errorf("expected uid string to be %s but got %s", idStr, id.String())
	}
	uid.MustParse("0123456789abcdef")

	var invalid string

	invalid = "abcdefghijklmno"
	if _, err := uid.Parse(invalid); err == nil {
		t.Error("expected error")
	}

	invalid = "abcdefghijklmnş"
	if _, err := uid.Parse(invalid); err == nil {
		t.Error("expected error")
	}
}

func TestMarshal(t *testing.T) {
	id := uid.MustParse("0123456789abcdef")
	bin, err := id.MarshalBinary()
	if err != nil {
		t.Error(err)
	}

	expected := []byte{0, 16, 131, 16, 81, 135, 32, 146, 139, 48, 211, 143}
	if len(expected) != len(bin) {
		t.Errorf("expected length %d but got %d", len(expected), len(bin))
	}
	for i := 0; i < len(expected); i++ {
		exp := expected[i]
		found := bin[i]

		if exp != found {
			t.Errorf("expected byte %b (%d) at index %d but found %b (%d)", exp, exp, i, found, found)
		}
	}

	if err := id.UnmarshalBinary(expected); err != nil {
		t.Error(err)
	}

	if err := id.UnmarshalBinary([]byte{}); err == nil {
		t.Error("expected error")
	}

	if err := id.UnmarshalBinary([]byte{0, 16, 131, 16, 81, 135, 32, 146}); err == nil {
		t.Error("expected error")
	}

	bin, err = id.MarshalText()
	if err != nil {
		t.Error(err)
	}
	if string(bin) != id.String() {
		t.Errorf("expected bin string to be %s but got %s", id.String(), bin)
	}

	bin, err = id.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bin) != fmt.Sprintf(`"%s"`, id.String()) {
		t.Errorf("expected bin string to be %s but got %s", fmt.Sprintf(`"%s"`, id.String()), bin)
	}
}
