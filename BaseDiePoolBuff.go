package dice

type BuffContextNumDiceRolls struct {
	roll        func() int
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

	Initialize(params *IBuffParams)

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

func (buffParams *BuffParams) Duration() int {
	return buffParams.duration
}

type BuffParamsAmount struct {
	amount int
}

func (buffParams *BuffParamsAmount) Amount() int {
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

func (baseBuff *BaseDiePoolBuff) Initialize(params IBuffParams) {
	baseBuff.duration = params.Duration()
}

func (baseBuff *BaseDiePoolBuff) ModifyDuration(delta int) {
	baseBuff.duration += delta
}

type BaseDiePoolBuffAmount struct {
	amount int
}

func (baseBuff *BaseDiePoolBuffAmount) Initialize(params IBuffParams) {
	baseBuff.amount = params.(IBuffParamsAmount).Amount()
}

type BaseDiePoolBuffContext struct {
	context int
}

type BaseDiePoolBuffVsFate struct {
	fate   *IDiePool
	vsFate *IDiePool
}
