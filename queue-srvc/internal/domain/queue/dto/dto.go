package dto

type PutQueueInput struct {
	Category string
	Item     string
}

type PutQueueOutput struct {
	Status int64
	Error  string
}

type GetQueueInput struct {
	Timeout  int64
	Category string
}

type GetQueueOutput struct {
	Status int64
	Item   string
	Error  string
}

func NewPutQueueInput(category string, item string) PutQueueInput {
	return PutQueueInput{
		Category: category,
		Item:     item,
	}
}

func NewPutQueueOutput(status int64, err string) PutQueueOutput {
	return PutQueueOutput{
		Status: status,
		Error:  err,
	}
}

func NewGetQueueInput(timeout int64, category string) GetQueueInput {
	return GetQueueInput{
		Timeout:  timeout,
		Category: category,
	}
}

func NewGetQueueOutput(status int64, item string, err string) GetQueueOutput {
	return GetQueueOutput{
		Status: status,
		Item:   item,
		Error:  err,
	}
}
