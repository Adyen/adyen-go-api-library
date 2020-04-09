package notification

type Notification struct {
	Live             string              `json:"live"`
	NotificationItems []*NotificationItem `json:"notificationItem"`
}
