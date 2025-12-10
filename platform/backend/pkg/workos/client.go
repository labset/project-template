package workos

type ClientConfig struct {
	ClientID    string
	APIKey      string
	RedirectURI string
}

type Client interface {
	UserManagement() UserManagement
	Organisations() Organisations
}

type client struct {
	userManagement UserManagement
	organisations  Organisations
}

func (c *client) Organisations() Organisations {
	return c.organisations
}

func (c *client) UserManagement() UserManagement {
	return c.userManagement
}

func NewClient(cfg ClientConfig) Client {
	return &client{
		userManagement: newUserManagement(cfg),
		organisations:  newOrganisations(cfg),
	}
}
