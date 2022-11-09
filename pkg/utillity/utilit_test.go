package utillity

import (
	"github.com/plavc/tutorial-go-cli/test/testutil"
	"testing"
)

func TestSolution(t *testing.T) {

	t.Run("expect reversed string", func(t *testing.T) {
		testutil.ExpectString(t, "bbaa", Reverse("aabb"))
	})
}
