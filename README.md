```
$ export SENTRY_DSN_DEFAULT=https://ba****************************54@o9**8.ingest.sentry.io/6***9
$ export SENTRY_DSN_DEPRECATED=https://ba****************************54:a9****************************93@o9**8.ingest.sentry.io/6***9
$ export SENTRY_DSN_LEGACY=https://ba****************************54:a9****************************93@app.getsentry.com/6***9"
$ go run main.go
[0]: 2020/10/03 04:37:50 raven: dsn missing private key
[0]: 2020/10/03 04:37:50 "My first Sentry error!" "817a5c382e98495697f3f1d784df494c"
[1]: 2020/10/03 04:37:50 "My first Sentry error!" "36ba540d865f477ebb6c3a44d9451f3c"
[2]: 2020/10/03 04:37:50 "My first Sentry error!" "27f0d5828b1745678c664c6165895e36"
```
