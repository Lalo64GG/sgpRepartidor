package services

type Notification interface {
	SendPushNotification(driverToken, title, body string) (error)
}
