package helpers

import "strconv"

func StringToUint32(s string) (uint32, error) {
	ui64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}

	id := uint32(ui64)

	return id, nil
}
