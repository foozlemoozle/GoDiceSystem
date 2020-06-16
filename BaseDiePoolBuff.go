/// Created by Kirk George
/// Copyright: Kirk George

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

type baseDiePoolBuff struct {
	duration int
}

func makeBaseDiePoolBuff(duration int) *baseDiePoolBuff {
	return &baseDiePoolBuff{duration: duration}
}

func (baseBuff *baseDiePoolBuff) Duration() int {
	return baseBuff.duration
}

func (baseBuff *baseDiePoolBuff) ModifyDuration(delta int) {
	baseBuff.duration += delta
}

type baseDiePoolBuffAmount struct {
	amount int
}

func makeBaseDiePoolBuffAmount(amount int) *baseDiePoolBuffAmount {
	return &baseDiePoolBuffAmount{amount: amount}
}

type baseDiePoolBuffContext struct {
	context DieRollContext
}

func makeBaseDiePoolBuffContext(context DieRollContext) *baseDiePoolBuffContext {
	return &baseDiePoolBuffContext{context: context}
}

type baseDiePoolBuffVsFate struct {
	fate   IDiePool
	vsFate IDiePool
}

func makeBaseDiePoolBuffVsFate(fate IDiePool, vsFate IDiePool) *baseDiePoolBuffVsFate {
	return &baseDiePoolBuffVsFate{fate: fate, vsFate: vsFate}
}
