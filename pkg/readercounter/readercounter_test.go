package readercounter

import (
	"bufio"
	"os"
	"testing"
)

func TestEmpty(t *testing.T) {
	fileEmpty, err := os.Open("..\\names\\empty")
	defer func() {
		_ = fileEmpty.Close()
	}()
	if err != nil {
		t.Fatalf("Cannot open empty file.\nGot err = %v of %T.", err, err)
	}
	names := ReadCountNames(bufio.NewScanner(fileEmpty))
	if len(names) == 0 {
		t.Logf("Test with empty file PASSED.\nGot err = %v of %T, names = %v.", err, err, names)
	} else {
		t.Errorf("Test with empty file FAILED.\nGot err = %v of %T, names = %v.", err, err, names)
	}
}

func TestOneName(t *testing.T) {
	fileOne, err := os.Open("..\\names\\one")
	defer func() {
		_ = fileOne.Close()
	}()
	if err != nil {
		t.Fatalf("Cannot open file.\nGot err = %v of %T.", err, err)
	}
	names := ReadCountNames(bufio.NewScanner(fileOne))
	if len(names) == 1 && names["Екатерина"] == 1 {
		t.Logf("Test 1-name-in-a-file PASSED.\nGot err = %v of %T, names = %v.", err, err, names)
	} else {
		t.Errorf("Test 1-name-in-a-file FAILED.\nGot err = %v of %T, names = %v.", err, err, names)
	}
}

func TestTwoSame(t *testing.T) {
	fileSame, err := os.Open("..\\names\\twosame")
	defer func() {
		_ = fileSame.Close()
	}()
	if err != nil {
		t.Fatalf("Cannot open file.\nGot err = %v of %T.", err, err)
	}
	names := ReadCountNames(bufio.NewScanner(fileSame))
	if len(names) == 1 && names["Kate"] == 2 {
		t.Logf("Test 2-same-names-in-a-file PASSED.\nGot err = %v of %T, names = %v.", err, err, names)
	} else {
		t.Errorf("Test 2-same-names-in-a-file FAILED.\nGot err = %v of %T, names = %v.", err, err, names)
	}
}

func TestTwoDif(t *testing.T) {
	fileDif, err := os.Open("..\\names\\twodif")
	defer func() {
		_ = fileDif.Close()
	}()
	if err != nil {
		t.Fatalf("Cannot open file.\nGot err = %v of %T.", err, err)
	}
	names := ReadCountNames(bufio.NewScanner(fileDif))
	if len(names) == 2 && names["Екатерина"] == names["Алиса"] {
		t.Logf("Test 2-different-names-in-a-file PASSED.\nGot err = %v of %T, names = %v.", err, err, names)
	} else {
		t.Errorf("Test 2-different-names-in-a-file FAILED.\nGot err = %v of %T, names = %v.", err, err, names)
	}
}

func TestThree(t *testing.T) {
	fileThree, err := os.Open("..\\names\\three")
	defer func() {
		_ = fileThree.Close()
	}()
	if err != nil {
		t.Fatalf("Cannot open file.\nGot err = %v of %T.", err, err)
	}
	names := ReadCountNames(bufio.NewScanner(fileThree))
	if len(names) == 2 && names["Михаил"] == 2 && names["Александра"] == 1 {
		t.Logf("Test 3-names-in-a-file PASSED.\nGot err = %v of %T, names = %v.", err, err, names)
	} else {
		t.Errorf("Test 3-names-in-a-file FAILED.\nGot err = %v of %T, names = %v.", err, err, names)
	}
}
