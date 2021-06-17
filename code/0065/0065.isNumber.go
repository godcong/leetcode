package _0065

const (
	StateInitial = iota
	StateIntSign
	StateInteger
	StatePoint
	StatePointWithoutInt
	StateFraction
	StateExp
	StateExpSign
	StateExpNumber
	StateEnd
	StateMax
)

const (
	CharNumber = iota
	CharExp
	CharPoint
	CharSign
	CharIllegal
	CharMax
)

var transfer = [StateMax][CharMax]int{
	StateInitial: {
		CharNumber: StateInteger,
		CharPoint:  StatePointWithoutInt,
		CharSign:   StateIntSign,
	},
	StateIntSign: {
		CharNumber: StateInteger,
		CharPoint:  StatePointWithoutInt,
	},
	StateInteger: {
		CharNumber: StateInteger,
		CharExp:    StateExp,
		CharPoint:  StatePoint,
	},
	StatePoint: {
		CharNumber: StateFraction,
		CharExp:    StateExp,
	},
	StatePointWithoutInt: {
		CharNumber: StateFraction,
	},
	StateFraction: {
		CharNumber: StateFraction,
		CharExp:    StateExp,
	},
	StateExp: {
		CharNumber: StateExpNumber,
		CharSign:   StateExpSign,
	},
	StateExpSign: {
		CharNumber: StateExpNumber,
	},
	StateExpNumber: {
		CharNumber: StateExpNumber,
	},
}

func toCharType(ch byte) int {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CharNumber
	case 'e', 'E':
		return CharExp
	case '.':
		return CharPoint
	case '+', '-':
		return CharSign
	default:
		return CharIllegal
	}
}

func isNumber(s string) bool {
	state := StateInitial
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		state = transfer[state][typ]
		if state == StateInitial {
			return false
		}
	}
	return state == StateInteger || state == StatePoint || state == StateFraction || state == StateExpNumber || state == StateEnd
}
