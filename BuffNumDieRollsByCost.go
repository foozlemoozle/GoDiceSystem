package dice

type BuffNumDieRollsByCost struct {
	*BaseDiePoolBuff
}

func MakeBuffNumDieRollsByCost(duration int) IDiePoolBuff {
	return &BuffNumDieRollsByCost{BaseDiePoolBuff: makeBaseDiePoolBuff(duration)}
}

func (buff *BuffNumDieRollsByCost) Buff(buffContext interface{}) {

	switch casted := buffContext.(type) {
	case *BuffContextNumDiceRolls:
		for i := 0; i < casted.diceCost; i++ {
			roll := casted.roll()
			if roll > casted.highestRoll {
				casted.highestRoll = roll
			}
		}
		break
	}
}
