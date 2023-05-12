package webhook

type Webhook struct {
	Live              string              `json:"live"`
	NotificationItems *[]NotificationItem `json:"notificationItems"`
}

// GetNotificationItems returns a reference to NotificationRequestItem's inside the NotificationItem array
func (n *Webhook) GetNotificationItems() []*NotificationRequestItem {
	resp := []*NotificationRequestItem{}

	if n.NotificationItems == nil {
		return resp
	}

	for _, v := range *n.NotificationItems {
		resp = append(resp, &v.NotificationRequestItem)
	}

	return resp
}
