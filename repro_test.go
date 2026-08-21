package repro

import (
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestApproval(t *testing.T) {
	approvals.VerifyString(t, "hello\n")
}
