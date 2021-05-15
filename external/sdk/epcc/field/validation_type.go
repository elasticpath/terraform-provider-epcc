package field

type ValidationType string

const (
	Enum      ValidationType = "enum"
	Email     ValidationType = "email"
	Slug      ValidationType = "slug"
	Uuid      ValidationType = "uuid"
	Between   ValidationType = "between"
	OneToMany ValidationType = "one-to-many"
	OneToOne  ValidationType = "one-to-one"
)

func (s ValidationType) AsString() string {
	return string(s)
}
