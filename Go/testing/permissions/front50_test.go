package permissions

import (
"encoding/json"
"fmt"
"net/http"
"net/http/httptest"
"reflect"
"sort"
"testing"
)

var permissions = []AppPermission{
	{
		Name: "noPermissions",
		Permissions: Access{
			READ:  []string{},
			WRITE: []string{},
		},
	},
	{
		Name: "readPermissions",
		Permissions: Access{
			READ:  []string{"testRead"},
			WRITE: []string{},
		},
	},
	{
		Name: "writePermissions",
		Permissions: Access{
			READ:  []string{},
			WRITE: []string{"testWrite"},
		},
	},
	{
		Name: "readWritePermissions",
		Permissions: Access{
			READ:  []string{"testRW"},
			WRITE: []string{"testRW"},
		},
	},
	{
		Name: "multiplePermissionsAll",
		Permissions: Access{
			READ:  []string{"group1", "group2"},
			WRITE: []string{"group1", "group2"},
		},
	},
	{
		Name: "multiplePermissionsMixed",
		Permissions: Access{
			READ:  []string{"group1", "group2"},
			WRITE: []string{"group1"},
		},
	},
	{
		Name: "multiplePermissionsSingle",
		Permissions: Access{
			READ:  []string{"group1"},
			WRITE: []string{"group2"},
		},
	},
}

var tests = []struct {
	ldapGroups          []string
	expectedPermissions AuthorizationV2
}{
	{
		ldapGroups: []string{"noPermissions"},
		expectedPermissions: AuthorizationV2{
			Admin: false,
		},
	},
	{
		ldapGroups: []string{"testRead"},
		expectedPermissions: AuthorizationV2{
			Applications: []AuthorizationV1{
				{
					Name:           "readPermissions",
					Authorizations: []string{"READ"},
				},
			},
			Admin: false,
		},
	},
	{
		ldapGroups: []string{"testWrite"},
		expectedPermissions: AuthorizationV2{
			Applications: []AuthorizationV1{
				{
					Name:           "writePermissions",
					Authorizations: []string{"WRITE"},
				},
			},
			Admin: false,
		},
	},
	{
		ldapGroups: []string{"testRW"},
		expectedPermissions: AuthorizationV2{
			Applications: []AuthorizationV1{
				{
					Name:           "readWritePermissions",
					Authorizations: []string{"READ", "WRITE"},
				},
			},
			Admin: false,
		},
	},
	{
		ldapGroups: []string{"group1"},
		expectedPermissions: AuthorizationV2{
			Applications: []AuthorizationV1{
				{
					Name:           "multiplePermissionsAll",
					Authorizations: []string{"READ", "WRITE"},
				},
				{
					Name:           "multiplePermissionsMixed",
					Authorizations: []string{"READ", "WRITE"},
				},
				{
					Name:           "multiplePermissionsSingle",
					Authorizations: []string{"READ"},
				},
			},
			Admin: false,
		},
	},
	{
		ldapGroups: []string{"group2"},
		expectedPermissions: AuthorizationV2{
			Applications: []AuthorizationV1{
				{
					Name:           "multiplePermissionsAll",
					Authorizations: []string{"READ", "WRITE"},
				},
				{
					Name:           "multiplePermissionsMixed",
					Authorizations: []string{"READ"},
				},
				{
					Name:           "multiplePermissionsSingle",
					Authorizations: []string{"WRITE"},
				},
			},
			Admin: false,
		},
	},
	{
		ldapGroups: []string{"group1", "group2"},
		expectedPermissions: AuthorizationV2{
			Applications: []AuthorizationV1{
				{
					Name:           "multiplePermissionsAll",
					Authorizations: []string{"READ", "WRITE"},
				},
				{
					Name:           "multiplePermissionsMixed",
					Authorizations: []string{"READ", "WRITE"},
				},
				{
					Name:           "multiplePermissionsSingle",
					Authorizations: []string{"READ", "WRITE"},
				},
			},
			Admin: false,
		},
	},
}

var server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	marshal, _ := json.Marshal(permissions)

	w.Write(marshal)

}))

func TestGetAuthV1Permissions(t *testing.T) {
	for _, tt := range tests {
		//only testing ldap groups with 1 group - getAuthV1Permissions only accepts 1 group
		if len(tt.ldapGroups) > 1 {
			continue
		}
		testName := tt.ldapGroups[0]
		t.Run(testName, func(t *testing.T) {
			actualPermissions := getAuthV1Permissions(tt.ldapGroups[0], &permissions)
			for _, item := range actualPermissions {
				sort.Strings(item.Authorizations)
			}
			if !reflect.DeepEqual(actualPermissions, tt.expectedPermissions.Applications) {
				t.Errorf("expected permissions don't match:\n%+v\n%+v\n", actualPermissions, tt.expectedPermissions.Applications)
			}
		})
	}
}

func TestFilterPermissions(t *testing.T) {
	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.ldapGroups)
		t.Run(testName, func(t *testing.T) {
			actualPermissions := filterPermissions(tt.ldapGroups, &permissions)
			for _, item := range actualPermissions {
				sort.Strings(item.Authorizations)
			}
			if !reflect.DeepEqual(actualPermissions, tt.expectedPermissions.Applications) {
				t.Errorf("expected permissions don't match:\n%+v\n%+v\n", actualPermissions, tt.expectedPermissions.Applications)
			}
		})
	}
}

func TestIsAdmin(t *testing.T) {
	admin := "test"
	if !isAdmin([]string{admin}, admin) {
		t.Errorf("isAdmin failed")
	}

	if isAdmin([]string{admin}, "fail") {
		t.Errorf("isAdmin failed")
	}
}

func TestGetUserPermissions(t *testing.T) {
	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.ldapGroups)
		t.Run(testName, func(t *testing.T) {
			actualAuthV2Permissions, _ := GetUserPermissions(tt.ldapGroups, server.URL, "testAdmin")

			for _, item := range tt.expectedPermissions.Applications {
				sort.Strings(item.Authorizations)
			}
			for _, item := range actualAuthV2Permissions.Applications {
				sort.Strings(item.Authorizations)
			}

			if !reflect.DeepEqual(actualAuthV2Permissions, tt.expectedPermissions) {
				t.Errorf("expected permissions don't match:\n%+v\n%+v\n", actualAuthV2Permissions, tt.expectedPermissions)
			}
		})
	}
}

func TestFetchPermissions(t *testing.T) {
	actualPermissions, _ := fetchPermissions(server.URL)
	if !reflect.DeepEqual(*actualPermissions, permissions) {
		t.Errorf("expected permissions don't match:\n%+v\n%+v\n", actualPermissions, permissions)
	}
}
