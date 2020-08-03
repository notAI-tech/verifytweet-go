package configs

// App represents application configuration
type (
	App struct {
		Search *Search `json:"search"`
	}

	Search struct {
		Host string `json:"host"`
		URL  string `json:"url"`
	}
)

// Init ...
func Init() *App {
	return &App{
		Search: &Search{
			Host: "http://localhost:8000",
			URL:  "/api/v1/search",
		},
	}
}
