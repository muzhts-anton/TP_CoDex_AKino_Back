package usrdelivery

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/log"

	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"firebase.google.com/go"
	"github.com/mailru/easyjson"
	"google.golang.org/api/option"
)

func (handler *UserHandler) AddUserToNotificationTopic(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tokenForm := new(domain.UserNotificationToken)
	err = easyjson.Unmarshal(b, tokenForm)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
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
	out, err := easyjson.Marshal(us)
	w.Write(out)
}
