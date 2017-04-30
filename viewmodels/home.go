package viewmodels

// HomeViewModel holds the view information for the home.gohtml template
type HomeViewModel struct {
	Title    string
	Nav      NavViewModel
	Messages []string
}

// NavViewModel holds the information for the nav bar.
type NavViewModel struct {
	Active          string
	IsAuthenticated bool
	Messages        []string
	UserName        string
}

// GetHomeViewModel populates the items required by the home.gohtml view
func GetHomeViewModel(messages []string, isAuthenticated bool, userName string) HomeViewModel {
	nav := NavViewModel{
		Active:          "home",
		IsAuthenticated: isAuthenticated,
		UserName:        userName,
	}

	result := HomeViewModel{
		Title:    "Go Ping Sites - Home",
		Nav:      nav,
		Messages: messages,
	}

	return result
}
