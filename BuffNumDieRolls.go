package dice

type BuffNumDieRolls struct {
	BaseDiePoolBuff
	BaseDiePoolBuffAmount
}

func (buff *BuffNumDieRolls) Buff(buffContext interface{}) {

	switch casted := buffContext.(type) {
	case *BuffContextNumDiceRolls:
		for i := 0; i < buff.amount; i++ {
			roll := casted.roll()
			if roll > casted.highestRoll {
				casted.highestRoll = roll
			}
		}

		break
	}
}
