package api

import (
	gohttp "net/http"

	"github.com/0xc0d/MailHog-Server/config"
	"github.com/gorilla/pat"
)

func CreateAPI(conf *config.Config, r gohttp.Handler) {
	apiv1 := createAPIv1(conf, r.(*pat.Router))
	apiv2 := createAPIv2(conf, r.(*pat.Router))

	go func() {
		for {
			select {
			case msg := <-conf.MessageChan:
				apiv1.messageChan <- msg
				apiv2.messageChan <- msg
			}
		}
	}()
}
