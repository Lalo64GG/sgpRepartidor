package fcm

import (
	"context"
	"fmt"
	"sync"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	appInstance *firebase.App
	once        sync.Once
)

func GetFirebaseApp() (*firebase.App, error) {
	var err error
	once.Do(func() {
		opt := option.WithCredentialsFile("src/Repartidor.json")
		appInstance, err = firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Printf("error initializing Firebase App: %v\n", err)
		}
	})
	return appInstance, err
}