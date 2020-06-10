package dice

type buffInfiniteDice struct {
	*baseDiePoolBuff
}

func makeBuffInfiniteDice() IDiePoolBuff {
	return &buffRemoveDice{
		baseDiePoolBuff: makeBaseDiePoolBuff(-1),
	}
}

func (buff *buffInfiniteDice) Buff(buffContext interface{}) {

	switch casted := buffContext.(type) {
	case *BuffContextRemoveDice:
		casted.amountToRemove = 0
		break
	}
}
