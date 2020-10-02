package main

import (
	"fmt"
	"log"
	"os"

	"github.com/getsentry/raven-go"
)

func main() {
	dsns := []string{
		// Default Sentry DSN; got from project settings page
		// Format: https://ba****************************54@o9**8.ingest.sentry.io/6***9
		os.Getenv("SENTRY_DSN_DEFAULT"),

		// Deprecated DSN, got from project settings page
		// Format: https://ba****************************54:a9****************************93@o9**8.ingest.sentry.io/6***9
		os.Getenv("SENTRY_DSN_DEPRECATED"),

		// Legacy DSN, which our Go application is using in production
		// Format: https://ba****************************54:a9****************************93@app.getsentry.com/6***9"
		os.Getenv("SENTRY_DSN_LEGACY"),
	}
	for i, dsn := range dsns {
		log.SetPrefix(fmt.Sprintf("[%d]: ", i))
		err := raven.SetDSN(dsn)
		if err != nil {
			log.Print(err)
		}
		a, s := raven.CapturePanic(func() {
			// do all of the scary things here
			panic("My first Sentry error!")
		}, nil)
		log.Printf("%#v %#v\n", a, s)
	}
}
