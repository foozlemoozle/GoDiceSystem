package dice

type buffRemoveDice struct {
	*baseDiePoolBuff
	*baseDiePoolBuffVsFate
}

func makeBuffRemoveDice(duration int, fate IDiePool, vsFate IDiePool) IDiePoolBuff {
	return &buffRemoveDice{
		baseDiePoolBuff:       makeBaseDiePoolBuff(duration),
		baseDiePoolBuffVsFate: makeBaseDiePoolBuffVsFate(fate, vsFate),
	}
}

func (buff *buffRemoveDice) Buff(buffContext interface{}) {

	switch casted := buffContext.(type) {
	case *BuffContextRemoveDice:
		result := contestPools(buff.vsFate, VersusFate, buff.fate, Fate)
		if result == Actor {
			casted.amountToRemove--
		}
		break
	}
}
