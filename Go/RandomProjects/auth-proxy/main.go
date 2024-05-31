package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)
type AuthorizationV1 struct {
	Name           string   `json:"name"`
	Authorizations []string `json:"authorizations"`
}

// AuthorizationV2 Authorization Object which contains list of authorized applications and whether ldapGroup is admin or not
type AuthorizationV2 struct { //app-tap-vessel - ldap group for admin
	Applications []AuthorizationV1 `json:"applications"`
	Admin        bool              `json:"admin"`
}
type Access struct{
	READ  []string `json:"READ"`
	WRITE []string `json:"WRITE"`
}

type Permission struct {
	Name           string `json:"name"`
	LastModified   int64  `json:"lastModified"`
	LastModifiedBy string `json:"lastModifiedBy"`
	Permissions    Access `json:"permissions"`
}

func main(){
	permissions, err := GetUserPermissions([]string{"app-tap-spin-price-execution-ro"})
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}
	fmt.Printf("%+v\n", permissions)
}

//func GetUserPermissions(ldapGroup string) (AuthorizationV2, error){
//	permissions, err := fetchPermissions()
//	if err != nil{
//		fmt.Printf("failed to feth permissions from front50: %s\n", err.Error())
//		return AuthorizationV2{}, err
//	}
//
//	//userPermissions:= getPermissionsForUser(ldapGroup, permissions)
//	//fmt.Printf("%+v\n", userPermissions)
//	//userAuthV2Permissions := convertPermissionsToAuthV2(ldapGroup, userPermissions)
//	//fmt.Printf("%+v\n", userAuthV2Permissions)
//
//	userAuthV2Permissions := getUserPermissionsAuthV2(ldapGroup, permissions)
//
//	return userAuthV2Permissions, nil
//}
//
//func fetchPermissions() ([]Permission, error){
//	var permissions []Permission
//	get, err := http.Select("https://storespinnaker-front50.prod.target.com/permissions/applications")
//	if err != nil {
//		fmt.Printf("%s\n", err.Error())
//		return nil, err
//	}
//
//	all, err := io.ReadAll(get.Body)
//	if err != nil {
//		fmt.Printf("%s\n", err.Error())
//		return nil, err
//	}
//	err = json.Unmarshal(all, &permissions)
//	if err != nil {
//		fmt.Printf("%s\n", err.Error())
//		return nil, err
//	}
//	return permissions, nil
//}
//
//func getUserPermissionsAuthV2(ldapGroup string, permissions []Permission) AuthorizationV2{
//	av2 :=  AuthorizationV2{
//		Applications: []AuthorizationV1{},
//	}
//	for _, v := range permissions{
//		if len(v.Permissions.WRITE) == 0 && len(v.Permissions.READ) == 0 {
//			continue
//		}
//
//		var rwPermissions []string
//		for _, r := range v.Permissions.WRITE {
//			if r == ldapGroup {
//				rwPermissions = append(rwPermissions, "WRITE")
//				break
//			}
//		}
//
//		for _, r := range v.Permissions.READ {
//			if r == ldapGroup {
//				rwPermissions = append(rwPermissions, "READ")
//				break
//			}
//		}
//		if len(rwPermissions) > 0{
//			av1 := AuthorizationV1{
//				Name: v.Name,
//				Authorizations: rwPermissions,
//			}
//			av2.Applications = append(av2.Applications, av1)
//		}
//
//	}
//	return av2
//
//}

//func getPermissionsForUser(ldapGroup string, permissions []Permission) []Permission{
//	var userPermission []Permission
//	var matched bool
//
//	for _,p := range permissions{
//		matched = false
//
//		for _, v := range p.Permissions.WRITE{
//			if v == ldapGroup {
//				userPermission = append(userPermission, p)
//				matched = true
//			}
//		}
//		if !matched{
//			for _, v := range p.Permissions.READ {
//				if v == ldapGroup {
//					userPermission = append(userPermission, p)
//				}
//			}
//		}
//
//	}
//	return userPermission
//}
//
//func convertPermissionsToAuthV2(ldapGroup string, permissions []Permission) AuthorizationV2{
//	av2 :=  AuthorizationV2{
//		Applications: []AuthorizationV1{},
//	}
//
//	for _, v := range permissions{
//		var rwPermissions []string
//		for _, r := range v.Permissions.WRITE {
//			if r == ldapGroup {
//				rwPermissions = append(rwPermissions, "WRITE")
//				break
//			}
//		}
//
//		for _, r := range v.Permissions.READ {
//			if r == ldapGroup {
//				rwPermissions = append(rwPermissions, "READ")
//				break
//			}
//		}
//
//		av1 := AuthorizationV1{
//			Name: v.Name,
//			Authorizations: rwPermissions,
//		}
//		av2.Applications = append(av2.Applications, av1)
//	}
//	return av2
//}

//TODO update logging

func GetUserPermissions(ldapGroups []string) (AuthorizationV2, error) {
	av2 := AuthorizationV2{
		Applications: []AuthorizationV1{},
	}

	permissions, err := fetchPermissions()
	if err != nil {
		fmt.Printf("failed to feth permissions from front50: %s\n", err.Error())
		return av2, err
	}

	av2.Admin = isAdmin(ldapGroups)
	if av2.Admin {
		return av2, nil
	}

	av2.Applications = FilterPermissions(ldapGroups, permissions)

	return av2, nil
}

func isAdmin(ldapGroups []string) bool {
	for _, ldapGroup := range ldapGroups {
		if os.Getenv("ADMIN_AD_GROUP") == ldapGroup {
			return true
		}
	}
	return false
}

func FilterPermissions(ldapGroups []string, permissions *[]Permission) (av1 []AuthorizationV1) {
	for _, ldapGroup := range ldapGroups {
		av1 = mergeAuthV1Lists(av1, getAuthV1Permissions(ldapGroup, permissions))
	}
	return av1
}

func fetchPermissions() (*[]Permission, error) {
	var permissions []Permission
	get, err := http.Get(os.Getenv("F50_URL"))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, err
	}

	all, err := io.ReadAll(get.Body)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, err
	}
	err = json.Unmarshal(all, &permissions)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, err
	}
	return &permissions, nil
}

// getAuthV1Permissions filters front50 Permissions matching the ldapGroup and transforms them to AuthorizationV1
func getAuthV1Permissions(ldapGroup string, appPermissions *[]Permission) []AuthorizationV1 {
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
		if !found{
			existing = append(existing, a)
		}
	}
	return existing
}
