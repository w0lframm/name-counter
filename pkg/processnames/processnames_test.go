package processnames

import (
	"testing"
)

func TestInvalidPath(t *testing.T) {
	names, err := ProcessNames("abcd")
	if err != nil && names == nil {
		t.Logf("TestInvalidPath \"abcd\" PASSED.\nExpected name = map[], err = *fs.PathError.\nGot name = %v, err = %T.", names, err)
	} else {
		t.Errorf("TestInvalidPath \"abcd\" FAILED.\nExpected name = map[], err = *fs.PathError.\nGot name = %v, err = %T.", names, err)
	}
}

func TestValidPath(t *testing.T) {
	names, err := ProcessNames("..\\names\\test.txt")

	if err == nil && len(names) == 3 && names["Алёна"] == 2 && names["Миша"] == 1 && names["Дима"] == 1 {
		t.Logf("Test ..\\names\\test.txt PASSED.\nGot err = %v of %T, names = %v.", err, err, names)
	} else {
		t.Errorf("Test ..\\names\\test.txt FAILED.\nGot err = %v of %T, names = %v.", err, err, names)
	}
}
