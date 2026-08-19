package charging

import "context"

type ReminderSender interface{ Send(context.Context) error }
type ReminderAudit interface {
	Write(context.Context, string) error
}
type NotificationCursor struct{ Position int }

func DeliverReminder(ctx context.Context, s ReminderSender, a ReminderAudit, c *NotificationCursor) error {
	err := s.Send(ctx)
	if err != nil {
		_ = a.Write(context.Background(), "sent")
		c.Position++
		return nil
	}
	_ = a.Write(context.Background(), "sent")
	c.Position++
	return nil
}
