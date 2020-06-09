package dice

type BuffDieRollBonus struct {
	BaseDiePoolBuff
	BaseDiePoolBuffAmount
	BaseDiePoolBuffContext
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

type BuffParamsDieRollBonus struct {
	BuffParams
	BuffParamsAmount
	BuffParamsDieRollContext
}
