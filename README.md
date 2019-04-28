# zoom-recordings

A command line tool to fetch Zoom cloud recordings.

## Usage

Install `zoom-recordings`.

```bash
go install github.com/anubhavmishra/zoom-recordings
```

You can also build `zoom-recordings` from scratch using the following
commands below.

```bash
go get github.com/anubhavmishra/zoom-recordings
```

```bash
cd $GOPATH/src/github.com/anubhavmishra/zoom-recordings
```

```bash
go build .
```

Set JWT Zoom API credentials

*Note: This project doesn't support authentication using OAuth. Read more about*
*generating a JWT token [here](https://marketplace.zoom.us/docs/guides/authorization/jwt/generating-jwt).*

```bash
export ZOOM_API_KEY="YOUR_ZOOM_JWT_KEY"
export ZOOM_API_SECRET="YOUR_ZOOM_JWT_SECRET"
```

```bash
Usage of ./zoom-recordings:
  -account-email string
        Zoom account email. It can also be supplied by using the "ZOOM_ACCOUNT_EMAIL" environment variable.
  -debug
        Enable or disable debugging. Set to false by default.
  -from string
        The date and time to start looking up recordings. Example: 2019-03-26T19:51:10.661Z.The date range has to be within one month.
  -meeting-id int
        Zoom meeting id to filter.
  -to string
        The date and time to end looking up recordings. Example: 2019-04-26T19:51:10.661Z.The date range has to be within one month.
```