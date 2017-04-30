package viewmodels

// AboutViewModel holds the view information for the about.gohtml template
type AboutViewModel struct {
	Title    string
	Nav      NavViewModel
	Messages []string
	Version  string
}

// GetAboutViewModel populates the items required by the about.gohtml view
func GetAboutViewModel(messages []string, isAuthenticated bool, userName string, version string) AboutViewModel {
	nav := NavViewModel{
		Active:          "about",
		IsAuthenticated: isAuthenticated,
	}

	result := AboutViewModel{
		Title:    "Go Ping Sites - About",
		Nav:      nav,
		Messages: messages,
		Version:  version,
	}
	return result
}
