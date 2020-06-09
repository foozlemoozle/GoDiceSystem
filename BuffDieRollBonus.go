package dice

type BuffDieRollBonus struct {
	*BaseDiePoolBuff
	*BaseDiePoolBuffAmount
	*BaseDiePoolBuffContext
}

func MakeBuffDieRollBonus(duration int, amount int, context DieRollContext) IDiePoolBuff {
	return &BuffDieRollBonus{
		BaseDiePoolBuff:        makeBaseDiePoolBuff(duration),
		BaseDiePoolBuffAmount:  makeBaseDiePoolBuffAmount(amount),
		BaseDiePoolBuffContext: makeBaseDiePoolBuffContext(context),
	}
}

func (buff *BuffDieRollBonus) Buff(buffContext interface{}) {

	switch casted := buffContext.(type) {
	case *BuffContextRollDice:
		if (buff.context & casted.context) != 0 {
			casted.bonus += buff.amount
		}
		break
	}
}
