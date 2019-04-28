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
  -table
        Enable or disable output in table format. Set to false by default.
  -to string
        The date and time to end looking up recordings. Example: 2019-04-26T19:51:10.661Z.The date range has to be within one month.
```

### Example

```bash
zoom-recordings -account-email example@example.com -from 2019-03-26T19:51:10.661Z -meeting-id 105230406 -table
```

Expected output

```bash
Meeting ID: 105230406
All cloud recordings from 2019-03-26T19:51:10.661Z to :
Output type "table".

+--------------------------+----------------------+-------------------------------------------------------------------------+
|           NAME           |    DATE AND TIME     |                           MEETING RECORDING URL                         |
+--------------------------+----------------------+-------------------------------------------------------------------------+
| Meeting1                 | 2019-04-20T15:30:01Z | https://zoom.us/recording/play/9518540cac6e68f3c1502f04790516b99545cd21 |
| Meeting2                 | 2019-04-15T15:27:55Z | https://zoom.us/recording/play/d3876ee7d773f1f808c99a24bb16197689032182 |
| Meeting3                 | 2019-04-10T15:30:01Z | https://zoom.us/recording/play/fb71deb67b59cf05afccf9134f5e0f6bf1928334 |
+--------------------------+----------------------+-------------------------------------------------------------------------+
```