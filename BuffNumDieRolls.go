package dice

type buffNumDieRolls struct {
	*baseDiePoolBuff
	*baseDiePoolBuffAmount
}

func makeBuffNumDieRolls(duration int, amount int) IDiePoolBuff {
	return &buffNumDieRolls{
		baseDiePoolBuff:       makeBaseDiePoolBuff(duration),
		baseDiePoolBuffAmount: makeBaseDiePoolBuffAmount(amount),
	}
}

func (buff *buffNumDieRolls) Buff(buffContext interface{}) {

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
