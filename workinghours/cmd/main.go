package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	folderNameEnv      = "env"
	tokenFileName      = "token.json"
	credentialFileName = "credentials.json"
)

var (
	pathTokenFile      = fmt.Sprintf("%s%s%s", folderNameEnv, string(filepath.Separator), tokenFileName)
	pathCredentialFile = fmt.Sprintf("%s%s%s", folderNameEnv, string(filepath.Separator), credentialFileName)
)

func createFolderEnv() {
	if _, err := os.Stat(folderNameEnv); os.IsNotExist(err) {
		if err := os.Mkdir(folderNameEnv, fs.ModeDir); err != nil {
			panic(err)
		}
	}
}

func getClient(config *oauth2.Config) *http.Client {

	tok, err := tokenFromFile(pathTokenFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(pathTokenFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tok, err := config.Exchange(ctx, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	createFolderEnv()
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()
	if err := json.NewEncoder(f).Encode(token); err != nil {
		panic(err)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	b, err := ioutil.ReadFile(pathCredentialFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	loc, err := time.LoadLocation("Europe/Madrid")
	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now().In(loc)
	firstDayMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, loc)
	lastDayMonth := firstDayMonth.AddDate(0, 1, -1)

	events, err := srv.Events.List("primary").ShowDeleted(true).SingleEvents(true).Q("Work hours").ShowDeleted(false).TimeMin(firstDayMonth.Format(time.RFC3339)).TimeMax(lastDayMonth.Format(time.RFC3339)).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	log.Println("Upcoming events:")
	if len(events.Items) == 0 {
		log.Println("No upcoming events found.")
		return
	}

	totalHours := 0.0
	for _, item := range events.Items {
		tStart, err := time.Parse(time.RFC3339, item.Start.DateTime)
		if err != nil {
			panic(err)
		}
		tEnd, err := time.Parse(time.RFC3339, item.End.DateTime)
		if err != nil {
			panic(err)
		}
		totalHours += tEnd.Sub(tStart).Hours()
	}
	log.Printf("Total hours in %s month were %.2f\n", now.Month(), totalHours)
}
