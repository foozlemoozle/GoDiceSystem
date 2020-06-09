package dice

type BuffDiePoolSize struct {
	BaseDiePoolBuff
	BaseDiePoolBuffAmount
}

func (buff *BuffDiePoolSize) Buff(buffContext interface{}) {

	switch casted := buffContext.(type) {
	case *BuffContextAddDice:
		casted.maxSize += buff.amount
		break
	}
}
