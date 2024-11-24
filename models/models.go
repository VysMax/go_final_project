package models

type Task struct {
	ID      string `json:"id,omitempty"`
	Date    string `json:"date,omitempty"`
	Title   string `json:"title,omitempty"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat,omitempty"`
	Error   string `json:"error,omitempty"`
}

type Response struct {
	ID    int64  `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

type TasksList struct {
	Tasks []Task `json:"tasks,omitempty"`
	Error string `json:"error,omitempty"`
}

type Auth struct {
	Password string `json:"password,omitempty"`
}

type AuthResponse struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}
