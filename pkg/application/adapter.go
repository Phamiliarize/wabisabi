package application

// datastore describes the expected interface of any future datastore adapters
type datastore interface {
	// CreateSession adds a session to the data store. Returns nil if no error
	CreateSession(req Session) error

	// ValidateSession will return (true,nil) for valid sessions. It returns (false,nil)
	// for invalid sessions and (false, error) for anything else
	ValidateSession(req Token) (bool, error)

	// DeleteSessionByUserId is used to kill a specific session.
	DeleteSessionByTokenId(req Token) error

	// DeleteSessionByUserId is used to kill all active tokens belonging to a specific
	// user. This should be used on any action that mutates a user or tactically for implementing a
	// "sign out everywhere" function.
	DeleteSessionByUserId(req User) error
}
