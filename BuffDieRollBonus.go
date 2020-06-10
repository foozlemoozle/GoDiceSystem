package dice

type buffDieRollBonus struct {
	*baseDiePoolBuff
	*baseDiePoolBuffAmount
	*baseDiePoolBuffContext
}

func makeBuffDieRollBonus(duration int, amount int, context DieRollContext) IDiePoolBuff {
	return &buffDieRollBonus{
		baseDiePoolBuff:        makeBaseDiePoolBuff(duration),
		baseDiePoolBuffAmount:  makeBaseDiePoolBuffAmount(amount),
		baseDiePoolBuffContext: makeBaseDiePoolBuffContext(context),
	}
}

func (buff *buffDieRollBonus) Buff(buffContext interface{}) {

	switch casted := buffContext.(type) {
	case *BuffContextRollDice:
		if (buff.context & casted.context) != 0 {
			casted.bonus += buff.amount
		}
		break
	}
}
