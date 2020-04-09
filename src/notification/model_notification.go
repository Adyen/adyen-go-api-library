package notification

type Notification struct {
	Live             string              `json:"live"`
	NotificationItem []*NotificationItem `json:"notificationItem"`
}
