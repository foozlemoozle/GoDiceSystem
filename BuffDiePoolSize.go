package dice

type BuffDiePoolSize struct {
	*BaseDiePoolBuff
	*BaseDiePoolBuffAmount
}

func MakeBuffDiePoolSize(duration int, amount int) IDiePoolBuff {
	return &BuffDiePoolSize{
		BaseDiePoolBuff:       makeBaseDiePoolBuff(duration),
		BaseDiePoolBuffAmount: makeBaseDiePoolBuffAmount(amount),
	}
}

func (buff *BuffDiePoolSize) Buff(buffContext interface{}) {

	switch casted := buffContext.(type) {
	case *BuffContextAddDice:
		casted.maxSize += buff.amount
		break
	}
}
