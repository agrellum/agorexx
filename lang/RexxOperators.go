package lang

type RexxOperators interface {
	OpAdd(set *RexxSet, rhs *Rexx) *Rexx
	OpAnd(set *RexxSet, rhs *Rexx) bool
	OpCc(set *RexxSet, rhs *Rexx) *Rexx
	OpCcblank(set *RexxSet, rhs *Rexx) *Rexx
	OpDiv(set *RexxSet, rhs *Rexx) *Rexx
	OpDivI(set *RexxSet, rhs *Rexx) *Rexx
	OpEq(set *RexxSet, rhs *Rexx) bool
	OpEqS(set *RexxSet, rhs *Rexx) bool
	OpGt(set *RexxSet, rhs *Rexx) bool
	OpGtEq(set *RexxSet, rhs *Rexx) bool
	OpGtEqS(set *RexxSet, rhs *Rexx) bool
	OpGtS(set *RexxSet, rhs *Rexx) bool
	OpLt(set *RexxSet, rhs *Rexx) bool
	OpLtEq(set *RexxSet, rhs *Rexx) bool
	OpLtEqS(set *RexxSet, rhs *Rexx) bool
	OpLtS(set *RexxSet, rhs *Rexx) bool
	OpMinus(set *RexxSet) *Rexx
	OpMult(set *RexxSet, rhs *Rexx) *Rexx
	OpNot(set *RexxSet) bool
	OpNotEq(set *RexxSet, rhs *Rexx) bool
	OpNotEqS(set *RexxSet, rhs *Rexx) bool
	OpOr(set *RexxSet, rhs *Rexx) bool
	OpPlus(set *RexxSet) *Rexx
	OpPow(set *RexxSet, rhs *Rexx) *Rexx
	OpRem(set *RexxSet, rhs *Rexx) *Rexx
	OpSub(set *RexxSet, rhs *Rexx) *Rexx
	OpXor(set *RexxSet, rhs *Rexx) bool
}
