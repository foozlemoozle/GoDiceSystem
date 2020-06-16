/// Created by Kirk George
/// Copyright: Kirk George

package dice

type buffNumDieRollsByCost struct {
	*baseDiePoolBuff
}

func makeBuffNumDieRollsByCost(duration int) IDiePoolBuff {
	return &buffNumDieRollsByCost{baseDiePoolBuff: makeBaseDiePoolBuff(duration)}
}

func (buff *buffNumDieRollsByCost) Buff(buffContext interface{}) {

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
