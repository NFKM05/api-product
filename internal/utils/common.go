package utils

import "database/sql"

// Função que representa o valor nulo usado nos DTOs representando campos opcionais
// Gravando a informação como nula no banco de dados ao invés de uma string vazia
func ToNullString(value *string) sql.NullString {
	if value == nil {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: *value, Valid: true}
}

// Função que faz o caminho inverso convertendo o valor devolvido pelo sqlc de volta pra *string
func FromNullString(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}
	return &value.String
}
