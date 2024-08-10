package tagError

type TagError struct {
	HtmlStatus int
	Inner      error
}

func GetTagError(htmlStatus int, inner error) *TagError {
	return &TagError{HtmlStatus: htmlStatus, Inner: inner}
}
