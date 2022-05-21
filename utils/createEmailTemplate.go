package utils

import (
	"bytes"
	"html/template"
)

func CreateEmailTemplate(data CalculatedData) (string, error) {
	var tmpl = `
	<!DOCTYPE html>
  <html lang="es">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Transaction Notification</title>
  </head>
  <body style="font-family: Helvetica, Arial, sans-serif; color: #333333;">
    <div style="padding: 40px 20px; max-width: 600px; margin: auto; background-color: #fff;">
      <h2 style="font-weight: normal; font-size: 24px; margin-top: 0; color: #200944">
        Hi, this is your personal transaction notification
      </h2>

			<div style="font-size: 15px; margin: 5px 0;">
        <span style="font-weight: 600;">Total balance: </span><span style="font-weight: normal;">{{.Total}}</span>
      </div>

			<div style="font-size: 15px; margin: 5px 0;">
        <span style="font-weight: 600;">Average Credit: </span><span style="font-weight: normal;">{{.AvgCredit}}</span>
      </div>
			<div style="font-size: 15px; margin: 5px 0;">
        <span style="font-weight: 600;">Average Debit: </span><span style="font-weight: normal;">{{.AvgDebit}}</span>
      </div>

			{{range .NumTransactionsByMonth}}
				<div style="font-size: 15px; margin: 5px 0;">
					<span style="font-weight: 600;">Number of transactions in {{.Month}}: </span><span style="font-weight: normal;">{{.NumTransactions}}</span>	
				</div>
			{{end}}

      <footer style="text-align:center">
  			<span style="display:inline;">Made with love by  </span> <span><img src="https://comunidadblogger.net/wp-content/uploads/2022/03/stori-logo.jpg" height="50px" style="vertical-align:middle; padding-left:3px"></span>
			</footer>
    </div>
  </body>
  </html>
	`

	t, err := template.New("emailTemplate").Parse(tmpl)

	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	err = t.Execute(&buf, data)

	if err != nil {
		return "", err
	}

	return  buf.String(), nil
}
