package translations

import (
	"github.com/df-mc/dragonfly/server/player/chat"
	"golang.org/x/text/language"
)

var NoTargetMatch = chat.Translate(str("commands.generic.noTargetMatch"), 0, "No targets matched selector")
var ToManyTargets = chat.Translate(str("commands.generic.tooManyTargets"), 0, "Too many targets matched selector")
var PlayerNotFound = chat.Translate(str("commands.generic.player.notFound"), 0, "That player cannot be found")
var InvalidOrigin = chat.Translate(str("commands.generic.invalidOrigin"), 0, "Command origin was invalid at command execution time")
var TargetNotPlayer = chat.Translate(str("commands.generic.targetNotPlayer"), 0, "Selector must be player-type")

type str string

func (s str) Resolve(language.Tag) string { return string(s) }
