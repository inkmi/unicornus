package uni

type Tag struct {
	Validation   *string
	Optional     bool
	ErrorMessage *string
	Choices      []Choice
	InputType    *string // "date" or "datetime" for time.Time fields
}
