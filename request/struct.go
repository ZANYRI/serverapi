package request

type Text struct {
	Texts string `json:"text"`
}

type Part struct {
	Parts []Text `json:"parts"`
}

type Content struct{
	Contents []Part  `json:"contents"`
}