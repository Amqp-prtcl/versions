package versions_test

import (
	"testing"

	"github.com/Amqp-prtcl/versions"
)

func Test(t *testing.T) {
	t.Logf("%v\n", versions.IsVersion(""))
	t.Logf("%v\n", versions.IsVersion("1.0.3"))
	t.Logf("%v\n", versions.IsVersion("12.3"))
	t.Logf("%v\n", versions.IsVersion("12..3"))
	t.Logf("%v\n", versions.IsVersion("1d0a3"))
	t.Logf("%v\n", versions.IsVersion("1.0."))
	t.Logf("%v\n", versions.IsVersion(".0.3"))
	t.Logf("%v\n", versions.IsVersion("1.."))
	t.Logf("%v\n", versions.IsVersion("..3"))
	t.Logf("%v\n", versions.IsVersion("1.0.3d"))
}

// 32.00.403
