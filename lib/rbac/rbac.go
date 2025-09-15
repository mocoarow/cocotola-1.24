package rbac

import (
	mbuserdomain "github.com/mocoarow/cocotola-1.24/moonbeam/user/domain"
)

// space
var CreateDeckAction = mbuserdomain.NewRBACAction("CreateDeck")
var ListDecksAction = mbuserdomain.NewRBACAction("ListDecks")

// deck
var ReadDeckAction = mbuserdomain.NewRBACAction("ReadDeck")

// card
