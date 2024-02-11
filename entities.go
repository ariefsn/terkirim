package terkirim

type M map[string]interface{}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type EmailFrom struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type EmailAccount struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type EmailAttachment struct {
	Content     string `json:"content"`
	Filename    string `json:"filename"`
	Type        string `json:"type"`
	Disposition string `json:"disposition"`
	ContentID   string `json:"contentId"`
}

type EmailPayload struct {
	From        EmailFrom         `json:"from"`
	To          []EmailAccount    `json:"to"`
	Cc          []EmailAccount    `json:"cc"`
	Bcc         []EmailAccount    `json:"bcc"`
	Subject     string            `json:"subject"`
	Body        string            `json:"body"`
	Headers     M                 `json:"headers"`
	Text        string            `json:"text"`
	Attachments []EmailAttachment `json:"attachments"`
	Category    string            `json:"category"`
	Variables   M                 `json:"variables"`
	Tags        M                 `json:"tags"`
}

type WhatsappPayload struct {
	From string `json:"accountName"`
	To   string `json:"to"`
	Body string `json:"body"`
}
