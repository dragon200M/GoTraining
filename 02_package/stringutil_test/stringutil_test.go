package stringutil_test

import (
	"github.com/dragon200M/GoTraining/02_package/stringutil"
	"testing"
)

func TestReverse(t *testing.T) {
	if stringutil.MyName != "Maciej" {
		t.Error("expected Maciej")
	}

	if stringutil.Reverse("Hello") != "olleH" {
		t.Error("expected olleH")
	}

}
