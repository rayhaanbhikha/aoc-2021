package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/rayhaanbhikha/aoc-2021/aoc-cli/language"
	"github.com/urfave/cli/v2"
)

func main() {
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
					return setToken(c.String("token"))
				},
			},
			{
				Name: "new",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "day",
						Required: true,
						Aliases:  []string{"d"},
						Usage:    "Advent of code calenday day",
					},
					&cli.StringFlag{
						Name:    "language",
						Value:   "go",
						Aliases: []string{"l"},
						Usage:   "language to generate AOC code template",
					},
				},
				Usage: "Create new code template",
				Action: func(c *cli.Context) error {
					return runApp(c.Int("day"), c.String("language"))
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func runApp(day int, languageTemplate string) error {
	shouldCleanup := false

	lang, err := language.NewLanguage(languageTemplate)
	if err != nil {
		return err
	}

	rootDirName := fmt.Sprintf("day%d-%s", day, lang)

	if _, err := os.Stat(rootDirName); err == nil {
		log.Fatal(fmt.Errorf("%s already exists", rootDirName))
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

	inputData, err := fetchInput(day)
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

func fetchInput(day int) ([]byte, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	sessionToken, err := readToken()
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

func setToken(token string) error {
	return os.WriteFile(".session", []byte(token), os.ModePerm)
}

func readToken() (string, error) {
	data, err := os.ReadFile(".session")
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.New("need to set session token by running: aoc session -t <token>")
		}

		return "", err
	}
	return strings.TrimSpace(string(data)), err
}
