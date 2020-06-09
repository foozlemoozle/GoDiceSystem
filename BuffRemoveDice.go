package dice

type BuffRemoveDice struct {
	BaseDiePoolBuff
	BaseDiePoolBuffVsFate
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

type BuffParamsRemoveDice struct {
	BuffParams
	BuffParamsVsFate
}
