package exception

func verifyParamsSize(size int, params []string) error {
	if len(params) < size {
		return Invalid
	}

	return nil
}
