package util

import (
	"golang.org/x/sys/windows"
	"log"
)

func IsRunAsAdmin() bool {
	var sid *windows.SID
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid)
	if err != nil {
		log.Printf("Failed to get SID")
		return false
	}
	defer windows.FreeSid(sid)

	token := windows.Token(0)
	isAdminMember, err := token.IsMember(sid)
	if err != nil {
		log.Printf("Failed to get token membership")
		return false
	}

	return token.IsElevated() || isAdminMember
}

func GetAdminUserName() string {
	return "administrator"
}
