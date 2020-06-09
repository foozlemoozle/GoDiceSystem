package dice

type BuffRemoveDice struct {
	*BaseDiePoolBuff
	*BaseDiePoolBuffVsFate
}

func MakeBuffRemoveDice(duration int, fate IDiePool, vsFate IDiePool) IDiePoolBuff {
	return &BuffRemoveDice{
		BaseDiePoolBuff:       makeBaseDiePoolBuff(duration),
		BaseDiePoolBuffVsFate: makeBaseDiePoolBuffVsFate(fate, vsFate),
	}
}

func (buff *BuffRemoveDice) Buff(buffContext interface{}) {

	switch casted := buffContext.(type) {
	case *BuffContextRemoveDice:
		result := ContestPools(buff.vsFate, VersusFate, buff.fate, Fate)
		if result == Actor {
			casted.amountToRemove--
		}
		break
	}
}
