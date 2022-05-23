package usrdelivery

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/log"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"net/http"
	"strconv"
	"encoding/json"
)

func (handler *UserHandler) AddUserToNotificationTopic(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	tokenForm := new(domain.UserNotificationToken)

	err := json.NewDecoder(r.Body).Decode(&tokenForm)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	opt := option.WithCredentialsFile("firebasePrivateKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Warn(fmt.Sprintf("error initializing app: %v\n", err))
	}

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Warn(fmt.Sprintf("error getting Messaging client: %v\n", err))
	}

	registrationTokens := []string{tokenForm.Token}

	response, err := client.SubscribeToTopic(ctx, registrationTokens, "all")
	if err != nil {
		log.Error(err)
	}
	log.Info(strconv.Itoa(response.SuccessCount) + " tokens were subscribed successfully")

	us := domain.UserNotificationToken{}
	x, err := json.Marshal(us)

	w.Write(x)
}
