package auth

import (
	"github.com/celestix/gotgproto"
)

type webAuth struct{}

var (
	phoneChan  = make(chan string)
	codeChan   = make(chan string)
	passwdChan = make(chan string)
)

func GetWebAuth() gotgproto.AuthConversator {
	return &webAuth{}
}

func (w *webAuth) AskPhoneNumber() (string, error) {
	code := <-phoneChan
	return code, nil
}

func (w *webAuth) AskCode() (string, error) {
	code := <-codeChan
	return code, nil
}

func (w *webAuth) AskPassword() (string, error) {
	code := <-passwdChan
	return code, nil
}

func (w *webAuth) RetryPassword(attemptsLeft int) (string, error) {
	code := <-passwdChan
	return code, nil
}

func ReceivePhone(phone string) {
	phoneChan <- phone
}

func ReceiveCode(code string) {
	codeChan <- code
}

func ReceivePasswd(passwd string) {
	passwdChan <- passwd
}
