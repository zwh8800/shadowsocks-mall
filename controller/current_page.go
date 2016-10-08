package controller

type CurrentPage int

const (
	CurrentPageInIndex CurrentPage = iota
	CurrentPageInConsole
	CurrentPageInIssue
)

func (c CurrentPage) InIndex() bool {
	return c == CurrentPageInIndex
}

func (c CurrentPage) InConsole() bool {
	return c == CurrentPageInConsole
}

func (c CurrentPage) InIssue() bool {
	return c == CurrentPageInIssue
}
