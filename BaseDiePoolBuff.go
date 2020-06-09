package dice

import (
	"strconv"
)

type BuffContextNumDiceRolls struct {
	roll        func()
	diceCost    int
	highestRoll int
}

type BuffContextRollDice struct {
	context DieRollContext
	bonus   int
}

type BuffContextRemoveDice struct {
	amountToRemove int
}

type BuffContextAddDice struct {
	maxSize int
}

/*INTERFACE DEFINITIONS*/

type IDiePoolBuff interface {
	Duration() int

	Initialize(params *IBuffParams) *IDiePoolBuff

	ModifyDuration(delta int)

	Buff(buffContext interface{})
}

type IBuffParams interface {
	Duration() int
}

type IBuffParamsAmount interface {
	Amount() int
}

type IBuffParamsDieRollContext interface {
	Context() int
}

type IBuffParamsVsFate interface {
	Fate() *IDiePool
	VsFate() *IDiePool
}

/*END*/

/*BUFF PARAMS DEFINTIONS*/

type BuffParams struct {
	duration int
}

func (buffParams *BaseBuffParams) Duration() int {
	return buffParams.duration
}

type BuffParamsAmount struct {
	amount int
}

func (buffParams *BaseBuffParamsAmount) Amount() int {
	return buffParams.amount
}

type BuffParamsDieRollContext struct {
	context int
}

func (buffParams *BuffParamsDieRollContext) Context() int {
	return buffParams.context
}

type BuffParamsVsFate struct {
	fate   *IDiePool
	vsFate *IDiePool
}

func (buffParams *BuffParamsVsFate) Fate() *IDiePool {
	return buffParams.fate
}

func (buffParams *BuffParamsVsFate) VsFate() *IDiePool {
	return buffParams.vsFate
}

/*END*/

type BaseDiePoolBuff struct {
	duration int
}

func (baseBuff *BaseDiePoolBuff) Duration() int {
	return baseBuff.duration
}

func (baseBuff *BaseDiePoolBuff) Initialize(params *IBuffParams) *IDiePoolBuff {
	baseBuff.duration = params.Duration()
	return baseBuff
}

func (baseBuff *BaseDiePoolBuff) ModifyDuration(delta int) {
	baseBuff.duration += delta
}

type BaseDiePoolBuffAmount struct {
	amount int
}

func (baseBuff *BaseDiePoolBuffAmount) Initialize(params *IBuffParams) *IDiePoolBuff {
	baseBuff.BaseDiePoolBuff.Initialize(params)
	baseBuff.amount = params.(IBuffParamsAmount).Amount()

	return baseBuff
}

type BaseDiePoolBuffContext struct {
	context int
}

type BaseDiePoolBuffVsFate struct {
	fate   *IDiePool
	vsFate *IDiePool
}
