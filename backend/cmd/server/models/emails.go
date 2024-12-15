package models

type ZincResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Source struct {
				MessageID string `json:"message_id"`
				Date      string `json:"date"`
				From      string `json:"from"`
				To        string `json:"to"`
				Subject   string `json:"subject"`
				Content   string `json:"content"`
				XFrom     string `json:"x_from"`
				XTo       string `json:"x_to"`
				XBCC      string `json:"x_bcc"`
				XCC       string `json:"x_cc"`
				XFilename string `json:"x_filename"`
				XFolder   string `json:"x_folder"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type CustomResponse struct {
	MessageID string `json:"message_id"`
	Date      string `json:"date"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
	XFrom     string `json:"x_from"`
	XTo       string `json:"x_to"`
	XBCC      string `json:"x_bcc"`
	XCC       string `json:"x_cc"`
	XFilename string `json:"x_filename"`
	XFolder   string `json:"x_folder"`
}

type Email struct {
	ID      string `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

type StatsResponse struct {
	TotalEmails      int    `json:"total_emails"`
	TopSender        string `json:"top_sender"`
	TopRecipient     string `json:"top_recipient"`
	LargestEmailSize int    `json:"largest_email_size"`
}
