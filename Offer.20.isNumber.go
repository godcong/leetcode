package leetcode

/*
剑指 Offer 20. 表示数值的字符串
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"-1E-16"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"及"12e+5.4"都不是。
*/

const (
	StateInitial = iota + 1
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
	CharNumber = iota + 1
	CharExp
	CharPoint
	CharSign
	CharSpace
	CharIllegal
	CharMax
)

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
	case ' ':
		return CharSpace
	default:
		return CharIllegal
	}
}

var isNumberTransfer = [StateMax][CharMax]int{
	StateInitial: {
		CharSpace:  StateInitial,
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
		CharSpace:  StateEnd,
	},
	StatePoint: {
		CharNumber: StateFraction,
		CharExp:    StateExp,
		CharSpace:  StateEnd,
	},
	StatePointWithoutInt: {
		CharNumber: StateFraction,
	},
	StateFraction: {
		CharNumber: StateFraction,
		CharExp:    StateExp,
		CharSpace:  StateEnd,
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
		CharSpace:  StateEnd,
	},
	StateEnd: {
		CharSpace: StateEnd,
	},
}

func isNumber(s string) bool {
	state := StateInitial
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if s := isNumberTransfer[state][typ]; s == 0 {
			return false
		} else {
			state = isNumberTransfer[state][typ]
		}
	}
	return state == StateInteger || state == StatePoint || state == StateFraction || state == StateExpNumber || state == StateEnd
}
