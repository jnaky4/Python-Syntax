package permissions

import (
"encoding/json"
"fmt"
"io"
"net/http"
)

// TODO update logging

func GetUserPermissions(ldapGroups []string, f50Url string, adminLdapGroup string) (AuthorizationV2, error) {
	av2 := AuthorizationV2{
		Applications: []AuthorizationV1{},
	}
	if len(ldapGroups) == 0 {
		return av2, fmt.Errorf("failed to check Permissions, no ldap group provided\n")
	}

	av2.Admin = isAdmin(ldapGroups, adminLdapGroup)
	if av2.Admin {
		return av2, nil
	}

	permissions, err := fetchPermissions(f50Url)
	if err != nil {
		return av2, err
	}

	av2.Applications = filterPermissions(ldapGroups, permissions)

	return av2, nil
}

func isAdmin(ldapGroups []string, admin string) bool {
	for _, ldapGroup := range ldapGroups {
		if admin == ldapGroup {
			return true
		}
	}
	return false
}

func filterPermissions(ldapGroups []string, permissions *[]AppPermission) (av1 []AuthorizationV1) {
	for _, ldapGroup := range ldapGroups {
		av1 = mergeAuthV1Lists(av1, getAuthV1Permissions(ldapGroup, permissions))
	}
	return av1
}

func fetchPermissions(f50Url string) (*[]AppPermission, error) {
	var permissions []AppPermission
	get, err := http.Get(f50Url)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch f50 permissions %w", err)
	}
	defer get.Body.Close()

	all, err := io.ReadAll(get.Body)
	if err != nil {
		return nil, fmt.Errorf("failed f50 get body ReadAll %w", err)
	}
	err = json.Unmarshal(all, &permissions)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal f50 json %w", err)
	}
	return &permissions, nil
}

// getAuthV1Permissions filters front50 Permissions matching the ldapGroup and transforms them to AuthorizationV1
func getAuthV1Permissions(ldapGroup string, appPermissions *[]AppPermission) []AuthorizationV1 {
	var av []AuthorizationV1
	for _, v := range *appPermissions {
		if len(v.Permissions.WRITE) == 0 && len(v.Permissions.READ) == 0 {
			continue
		}

		var rwPermissions []string
		for _, r := range v.Permissions.WRITE {
			if r == ldapGroup {
				rwPermissions = append(rwPermissions, "WRITE")
				break
			}
		}

		for _, r := range v.Permissions.READ {
			if r == ldapGroup {
				rwPermissions = append(rwPermissions, "READ")
				break
			}
		}

		if len(rwPermissions) > 0 {
			av1 := AuthorizationV1{
				Name:           v.Name,
				Authorizations: rwPermissions,
			}
			av = append(av, av1)
		}

	}
	return av
}

// mergeAuthV1Lists checks for duplicate permissions in both lists
func mergeAuthV1Lists(existing []AuthorizationV1, adding []AuthorizationV1) []AuthorizationV1 {
	found := true
	for i, a := range adding {
		found = false
		for j, e := range existing {
			if e.Name == a.Name {
				found = true
				//Len cannot be greater than 2, Permissions has Write and Read structs only
				if len(e.Authorizations) == 2 {
					break
				}
				if len(a.Authorizations) == 2 {
					existing[j].Authorizations = a.Authorizations
					break
				}
				if e.Authorizations[0] != a.Authorizations[0] {
					existing[j].Authorizations = append(existing[j].Authorizations, adding[i].Authorizations...)
					break
				}
			}
		}
		if !found {
			existing = append(existing, a)
		}
	}
	return existing
}

