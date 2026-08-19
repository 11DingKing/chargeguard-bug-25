package charging

import (
	"context"
	"fmt"
)

type ReminderSender interface{ Send(context.Context) error }
type ReminderAudit interface {
	Write(context.Context, string) error
}
type NotificationCursor struct{ Position int }

func DeliverReminder(ctx context.Context, s ReminderSender, a ReminderAudit, c *NotificationCursor) error {
	err := s.Send(ctx)
	if err != nil {
		return fmt.Errorf("send reminder: %w", err)
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("cancel reminder: %w", err)
	}
	if err := a.Write(ctx, "sent"); err != nil {
		return fmt.Errorf("audit reminder: %w", err)
	}
	c.Position++
	return nil
}
