package request

type Text struct {
	Text string `json:"text"`
}

type Part struct {
	Parts []Text `json:"parts"`
}

type Content struct{
	Contents []Parts `json:"contents"`
}