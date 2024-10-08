package uausgosdk

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type SDK struct{}

func (sdk *SDK) CreateUser(email, password string) (User, error) {
	url := URL + "/crud/createUser/" + SERVICE_ID

	payload := strings.NewReader(fmt.Sprintf("{\n \"email\" : \"%s\",\n \"password\" : \"%s\"\n}",
		email, password))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Println(err.Error())
		return User{}, err
	}

	req.Header.Add("Authorization", "Bearer "+API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return User{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return User{}, fmt.Errorf("failed")
	}

	var reponse User

	err = unmarshalbodytostruct(&res.Body, &reponse)
	if err != nil {
		log.Println(err.Error())
		return User{}, err
	}

	return reponse, nil
}

func (sdk *SDK) ChangePassword(token, newPassword string) error {
	url := URL + "/crud/changePassword/" + SERVICE_ID

	payload := strings.NewReader(fmt.Sprintf(`{\n "resetPasswordToken" : "%s",\n "newPassword" : "%s"\n}`,
		token, newPassword))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	req.Header.Add("Authorization", "Bearer "+API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed")
	}

	return nil
}

func (sdk *SDK) AuthenticateUser(email, password string) (string, error) {
	url := URL + "/authentication/authenticateUser/" + SERVICE_ID

	payload := strings.NewReader(fmt.Sprintf(`{\n "email" : "%s",\n "password" : "%s"\n}`,
		email, password))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed")
	}

	var reponse Token

	err = unmarshalbodytostruct(&res.Body, &reponse)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return reponse.Token, nil
}

func (sdk *SDK) ValidateJWT(token string) error {
	url := URL + "/authentication/validateJWT/" + SERVICE_ID

	payload := strings.NewReader(fmt.Sprintf(`{\n "jwt" : "%s",\n}`,
		token))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	req.Header.Add("Authorization", "Bearer "+API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed")
	}

	return nil
}

func (sdk *SDK) GenerateResetPasswordToken(email string) (string, error) {
	url := URL + "/authentication/resetPassword/" + SERVICE_ID

	payload := strings.NewReader(fmt.Sprintf(`{\n "email" : "%s"}`,
		email))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+API_KEY)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed")
	}

	var reponse Token

	err = unmarshalbodytostruct(&res.Body, &reponse)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return reponse.Token, nil
}

func unmarshalbodytostruct(body *io.ReadCloser, targetstruct any) error {
	bodybytes, err := io.ReadAll(*body)
	if err != nil {
		log.Println("read body error: ", err.Error())
		return err
	}

	err = json.Unmarshal(bodybytes, targetstruct)
	if err != nil {
		log.Println("unmarshal body error: ", err.Error())
		return err
	}

	return nil
}
