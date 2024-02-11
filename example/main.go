package main

import (
	"fmt"

	"github.com/ariefsn/terkirim"
)

func main() {
	client := terkirim.New("YOUR_API_KEY")

	res, err := client.Whatsapp(terkirim.WhatsappPayload{
		From: "REGISTERED_WHATSAPP_NUMBER_FROM_TERKIRIM",
		To:   "PHONE_NUMBER_TO_SEND_WHATSAPP_WITH_COUNTRY_CODE",
		Body: "Hello World",
	})
	if err != nil {
		fmt.Println("Whatsapp error:", err)
		return
	}

	fmt.Println("Whatsapp response", res)

	res, err = client.Email(terkirim.EmailPayload{
		From: terkirim.EmailFrom{
			Username: "REGISTERED_ACCOUNT_FROM_TERKIRIM",
			Name:     "ALIAS",
		},
		To: []terkirim.EmailAccount{
			{
				Email: "someone@mail.com",
				Name:  "Alias",
			},
		},
		Subject:  "Hello World",
		Category: "Terkirim",
		Body:     "<mjml>\n\t<mj-body>\n\t\t<mj-section>\n\t\t\t<mj-column>\n\t\t\t\t<mj-divider></mj-divider>\n\t\t\t\t<mj-text font-family=\"helvetica\" font-size=\"20px\" align=\"center\">Hello, <mj-text>{{ firstName }} {{ lastName }}</mj-text></mj-text>\n\t\t\t\t<mj-text align=\"center\">Thank you for joining Terkirim</mj-text>\n\t\t\t\t<mj-divider></mj-divider>\n\t\t\t</mj-column>\n\t\t</mj-section>\n\t</mj-body>\n</mjml>\n",
		Variables: terkirim.M{
			"firstName": "Captain",
			"lastName":  "Tsubasa",
		},
	})
	if err != nil {
		fmt.Println("Email error:", err)
		return
	}

	fmt.Println("Email response", res)
}
