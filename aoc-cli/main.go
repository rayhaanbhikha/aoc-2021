package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/rayhaanbhikha/aoc-2021/aoc-cli/language"
	"github.com/rayhaanbhikha/aoc-2021/aoc-cli/session"
	"github.com/urfave/cli/v2"
)

var questions = []*survey.Question{
	{
		Name:     "day",
		Prompt:   &survey.Input{Message: "What day?"},
		Validate: survey.Required, // TODO: can check it's a number.
	},
	{
		Name: "language",
		Prompt: &survey.Select{
			Message: "Choose a language",
			Options: []string{"Go", "Node"},
			Default: "Go",
		},
	},
}

func main() {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name: "AOC cli to generate ready to go code templates",
		Commands: []*cli.Command{
			{
				Name: "session",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "token",
						Required: true,
						Aliases:  []string{"t"},
						Usage:    "Sets session cookie as token",
					},
				},
				Usage: "Set Session token (stored as cookie in browser)",
				Action: func(c *cli.Context) error {
					return sess.SetToken(c.String("token"))
				},
			},
			{
				Name:  "new",
				Usage: "Create new code template",
				Action: func(c *cli.Context) error {
					if err := sess.Init(); err != nil {
						return err
					}

					answers := struct {
						Day      int
						Language string
					}{}

					if err := survey.Ask(questions, &answers); err != nil {
						return err
					}

					return runApp(answers.Day, answers.Language, sess)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func runApp(day int, languageTemplate string, sess *session.Session) error {
	shouldCleanup := false

	lang, err := language.NewLanguage(languageTemplate)
	if err != nil {
		return err
	}

	rootDirName := fmt.Sprintf("day%d-%s", day, lang)

	if _, err := os.Stat(rootDirName); err == nil {
		return fmt.Errorf("%s already exists", rootDirName)
	}

	if err = os.Mkdir(rootDirName, os.ModePerm); err != nil {
		return err
	}

	defer func() {
		if !shouldCleanup {
			return
		}

		if err := os.RemoveAll(rootDirName); err != nil {
			log.Println("failed to clean up")
			panic(err)
		}
	}()

	inputData, err := fetchInput(day, sess.GetToken())
	if err != nil {
		shouldCleanup = true
		return err
	}

	if err := lang.Create(rootDirName, inputData); err != nil {
		shouldCleanup = true
		return err
	}

	return nil
}

func fetchInput(day int, sessionToken string) ([]byte, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Cookie", fmt.Sprintf("session=%s", sessionToken))
	response, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("failed to retrieve input")
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
