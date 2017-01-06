package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	var (
		surah string
		ayah  string
	)

	app := cli.NewApp()
	app.Name = "Al-Qur'an in CLI"
	app.Usage = "Data from https://alquran.cloud/"
	app.Version = "1.0"
	app.Authors = []cli.Author{
		{
			Name:  "Muhammad Zaki Al-Afrani",
			Email: "zaki.afrani@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "surah",
			Value:       "1",
			Usage:       "Surah number",
			Destination: &surah,
		},
		cli.StringFlag{
			Name:        "ayah",
			Value:       "1",
			Usage:       "Ayah number",
			Destination: &ayah,
		},
	}

	app.Action = func(c *cli.Context) error {
		printAyah(surah, ayah)
		return nil
	}

	app.Run(os.Args)
}
