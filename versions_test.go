package versions_test

import (
	"fmt"
	"testing"

	. "github.com/Amqp-prtcl/versions"
)

func Test(t *testing.T) {
	a := Version("0.0.1")
	b := Version("1.0.32")
	c := Version("1.0.32")
	d := Version("1.1.32")
	e := Version("2.0.2")

	var vrs Versions = Versions{
		a, c, e, b, d,
	}

	fmt.Printf("%#v\n", vrs)
	fmt.Printf("last version %#v\n", vrs.GetLastVersion())

	fmt.Printf("%t\n", Version("4.0.3").IsNewerThan(Version("1.2.1")))
}

// 32.00.403
