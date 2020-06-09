package dice

type BuffDiePoolSize struct {
	BaseDiePoolBuff
	BaseDiePoolBuffAmount
}

func (buff *BuffDiePoolSize) Buff(buffContext interface{}) {
	calc := buff.cachedAction(buffContext)

	switch casted := buffContext.(type) {
	case *BuffContextAddDice:
		casted.maxSize += buff.amount
		break
	}
}
