package charging

import (
	"context"
	"errors"
	"testing"
)

type cancelledSender struct{}

func (cancelledSender) Send(ctx context.Context) error { return ctx.Err() }

type countingAudit struct{ writes int }

func (a *countingAudit) Write(context.Context, string) error { a.writes++; return nil }
func TestTaskBehavior(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	audit := &countingAudit{}
	cursor := &NotificationCursor{}
	err := DeliverReminder(ctx, cancelledSender{}, audit, cursor)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err=%v", err)
	}
	if audit.writes != 0 || cursor.Position != 0 {
		t.Fatalf("writes=%d cursor=%d", audit.writes, cursor.Position)
	}
}
