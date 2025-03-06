package helpers

import (
	"context"
	"fmt"

	"firebase.google.com/go/messaging"
	"github.com/lalo64/sgp/src/delivery/application/services"
	fcm "github.com/lalo64/sgp/src/shared/FCM"
)


type PushNotificationService struct {}

func NewPushNotificationService() (services.Notification, error) {
	return &PushNotificationService{}, nil
}



func (s *PushNotificationService) SendPushNotification(driverToken, title, body string) error {
	app, err := fcm.GetFirebaseApp()
	if err != nil {
		return fmt.Errorf("error obteniendo Firebase App: %v", err)
	}

	// Crear cliente de FCM
	client, err := app.Messaging(context.Background())
	if err != nil {
		return fmt.Errorf("error obteniendo cliente FCM: %v", err)
	}

	// Crear el mensaje
	message := &messaging.Message{
		Token: driverToken,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	// Enviar el mensaje
	_, err = client.Send(context.Background(), message)
	if err != nil {
		return fmt.Errorf("error enviando la notificación push: %v", err)
	}

	// Si todo va bien, devolver nil
	return nil
}