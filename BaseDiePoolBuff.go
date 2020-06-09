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

	ModifyDuration(delta int)

	Buff(buffContext interface{})
}

/*END*/

type BaseDiePoolBuff struct {
	duration int
}

func makeBaseDiePoolBuff(duration int) *BaseDiePoolBuff {
	return &BaseDiePoolBuff{duration: duration}
}

func (baseBuff *BaseDiePoolBuff) Duration() int {
	return baseBuff.duration
}

func (baseBuff *BaseDiePoolBuff) ModifyDuration(delta int) {
	baseBuff.duration += delta
}

type BaseDiePoolBuffAmount struct {
	amount int
}

func makeBaseDiePoolBuffAmount(amount int) *BaseDiePoolBuffAmount {
	return &BaseDiePoolBuffAmount{amount: amount}
}

type BaseDiePoolBuffContext struct {
	context DieRollContext
}

func makeBaseDiePoolBuffContext(context DieRollContext) *BaseDiePoolBuffContext {
	return &BaseDiePoolBuffContext{context: context}
}

type BaseDiePoolBuffVsFate struct {
	fate   IDiePool
	vsFate IDiePool
}

func makeBaseDiePoolBuffVsFate(fate IDiePool, vsFate IDiePool) *BaseDiePoolBuffVsFate {
	return &BaseDiePoolBuffVsFate{fate: fate, vsFate: vsFate}
}
