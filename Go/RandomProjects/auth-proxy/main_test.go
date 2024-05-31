package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

var permissions = []Permission{
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

func TestGetUserPermissionsAuthV2(t *testing.T) {
	var tests = []struct {
		ldapGroup           string
		expectedPermissions AuthorizationV2
	}{
		{
			ldapGroup: "noPermissions",
			expectedPermissions: AuthorizationV2{
				Admin: false,
			},
		},
		{
			ldapGroup: "testRead",
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
			ldapGroup: "testWrite",
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
			ldapGroup: "testRW",
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
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.ldapGroup)
		t.Run(testName, func(t *testing.T) {
			actualPermissions := getAuthV1Permissions(tt.ldapGroup, &permissions)
			for _, item := range actualPermissions{
				sort.Strings(item.Authorizations)
			}
			if !reflect.DeepEqual(actualPermissions, tt.expectedPermissions.Applications) {
				t.Errorf("expected permissions don't match:\n%+v\n%+v\n", actualPermissions, tt.expectedPermissions.Applications)
			}
		})
	}
}

func TestFilterPermissions(t *testing.T) {
	var tests = []struct {
		ldapGroups               []string
		expectedPermissions []AuthorizationV1
	}{
		{
			ldapGroups: []string{"group1"},
			expectedPermissions: []AuthorizationV1{
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
		},
		{
			ldapGroups: []string{"group2"},
			expectedPermissions: []AuthorizationV1{
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
		},
		{
			ldapGroups: []string{"group1", "group2"},
			expectedPermissions: []AuthorizationV1{
				{
					Name:           "multiplePermissionsAll",
					Authorizations: []string{"READ", "WRITE" },
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
		},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.ldapGroups)
		t.Run(testName, func(t *testing.T) {
			actualPermissions := FilterPermissions(tt.ldapGroups, &permissions)
			for _, item := range actualPermissions{
				sort.Strings(item.Authorizations)
			}
			if !reflect.DeepEqual(actualPermissions, tt.expectedPermissions) {
				t.Errorf("expected permissions don't match:\n%+v\n%+v\n", actualPermissions, tt.expectedPermissions)
			}
		})
	}
}
