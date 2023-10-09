package ui

import "strings"

func hasOptional(typ string) bool {
	return strings.HasPrefix(typ, "Option[")
}

func removeOptional(typ string) string {
	if strings.HasPrefix(typ, "Option[") {
		typ = typ[len("Option[") : len(typ)-1]
	}
	return typ
}
