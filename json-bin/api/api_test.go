package api_test

import (
	"errors"
	"json-bin/api"
	"json-bin/bins"
	"json-bin/config"
	"testing"
)

var _api *api.Api = api.New(config.New("../.env"))

func TestGet(t *testing.T) {
	// Arrange - подготовка, ожидаемый результат
	tCases := []struct {
		name     string
		id       string
		expected error
	}{
		{
			name:     "valid",
			id:       "68f8fe14ae596e708f244c78",
			expected: nil,
		},
		{
			name:     "valid expected error not found",
			id:       "68f8fe14ae596e708f244c98",
			expected: api.ErrNotFound,
		},
		{
			name:     "valid expected error empty id",
			id:       "",
			expected: api.ErrEmptyId,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act - выполняем функцию
			err := _api.Get(tc.id)

			// Assert - проверка результата с ожидаемым
			if err != nil {
				if tc.expected != nil {
					if !errors.Is(err, tc.expected) {
						t.Errorf("Ожидалось: %v, получено: %v", tc.expected, err)
					}
				} else {
					t.Errorf("Пришла ошибка: %v", err)
				}
			} else {
				if tc.expected != nil {
					t.Errorf("Ожидалось: %v, получено: nil", tc.expected)
				}
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tCases := []struct {
		name     string
		id       string
		expected error
	}{
		{
			name:     "valid",
			id:       "68f8fe14ae596e708f244c78",
			expected: nil,
		},
		{
			name:     "valid expected error not found",
			id:       "68f8fe14ae596e708f244c98",
			expected: api.ErrNotFound,
		},
		{
			name:     "valid expected error empty id",
			id:       "",
			expected: api.ErrEmptyId,
		},
	}
	binRecord := bins.GenerateRecord()
	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act - выполняем функцию
			_, err := _api.Update(tc.id, binRecord)

			// Assert - проверка результата с ожидаемым
			if err != nil {
				if tc.expected != nil {
					if !errors.Is(err, tc.expected) {
						t.Errorf("Ожидалось: %v, получено: %v", tc.expected, err)
					}
				} else {
					t.Errorf("Пришла ошибка: %v", err)
				}
			} else {
				if tc.expected != nil {
					t.Errorf("Ожидалось: %v, получено: nil", tc.expected)
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	// Arrange - подготовка, ожидаемый результат
	tCases := []struct {
		name     string
		id       string
		expected error
	}{
		{
			name:     "valid",
			id:       "68f8fe14ae596e708f244c78",
			expected: nil,
		},
		{
			name:     "valid expected error not found",
			id:       "68f8fe14ae596e708f244c98",
			expected: api.ErrNotFound,
		},
		{
			name:     "valid expected error empty id",
			id:       "",
			expected: api.ErrEmptyId,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act - выполняем функцию
			err := _api.Delete(tc.id)

			// Assert - проверка результата с ожидаемым
			if err != nil {
				if tc.expected != nil {
					if !errors.Is(err, tc.expected) {
						t.Errorf("Ожидалось: %v, получено: %v", tc.expected, err)
					}
				} else {
					t.Errorf("Пришла ошибка: %v", err)
				}
			} else {
				if tc.expected != nil {
					t.Errorf("Ожидалось: %v, получено: nil", tc.expected)
				}
			}
		})
	}
}

func TestCreate(t *testing.T) {
	// Arrange - подготовка, ожидаемый результат
	// Act - выполняем функцию
	binRecord := bins.GenerateRecord()
	bin, err := _api.Create(binRecord)

	// Assert - проверка результата с ожидаемым
	if err != nil {
		t.Errorf("Пришла ошибка: %v", err)
	} else {
		_api.Delete(bin.Id)
	}
}
