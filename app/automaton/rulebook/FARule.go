package rulebook

type FARule struct {
	state     int
	character string
	nextState int
}

func NewFARule(state int, character string, nextState int) *FARule {
	return &FARule{state: state, character: character, nextState: nextState}
}

func (rule *FARule) AppliesTo(state int, character string) bool {
	return rule.state == state && rule.character == character
}

func (rule *FARule) Follow() int {
	return rule.nextState
}

func (rule *FARule) inspect() string {
	return "<FARule " + string(rule.state) + " --" + rule.character + "--> " + string(rule.nextState) + ">"
}
