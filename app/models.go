package app

type Link struct {
  Href string `json:"href"`
  Templated bool `json:"templated"`
}

type HAL struct {
  Links map[string]Link `json:"_links"`
}