package main

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"ID tidak boleh kosong"}
	}
	if id != "123" {
		return &notFoundError{"Data tidak ditemukan"}
	}
	return nil
}
