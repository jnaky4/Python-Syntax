package permissions

type Access struct {
	READ  []string `json:"READ"`
	WRITE []string `json:"WRITE"`
}

type AppPermission struct {
	Name           string `json:"name"`
	LastModified   int64  `json:"lastModified"`
	LastModifiedBy string `json:"lastModifiedBy"`
	Permissions    Access `json:"permissions"`
}

// AuthorizationV1 and V2 is a legacy struct from fiat
type AuthorizationV1 struct {
	Name           string   `json:"name"`
	Authorizations []string `json:"authorizations"`
}

//todo remove references to AuthorizationV2

// AuthorizationV2 Authorization Object which contains list of authorized applications and whether user is admin or not
type AuthorizationV2 struct { //app-tap-vessel - ldap group for admin
	Applications []AuthorizationV1 `json:"applications"`
	Admin        bool              `json:"admin"`
}
