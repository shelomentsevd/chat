package db

func Get(model interface{}, opts ...Option) error {
	query := Pool

	for _, option := range opts {
		query = option.apply(query)
	}

	result := query.Find(model)
	if result.RecordNotFound() {
		return ErrRecordNotFound
	}

	return processError(result.Error)
}
