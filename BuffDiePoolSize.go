/// Created by Kirk George
/// Copyright: Kirk George

package dice

type buffDiePoolSize struct {
	*baseDiePoolBuff
	*baseDiePoolBuffAmount
}

func makeBuffDiePoolSize(duration int, amount int) IDiePoolBuff {
	return &buffDiePoolSize{
		baseDiePoolBuff:       makeBaseDiePoolBuff(duration),
		baseDiePoolBuffAmount: makeBaseDiePoolBuffAmount(amount),
	}
}

func (buff *buffDiePoolSize) Buff(buffContext interface{}) {

	switch casted := buffContext.(type) {
	case *BuffContextAddDice:
		casted.maxSize += buff.amount
		break
	}
}
