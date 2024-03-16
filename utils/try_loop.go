package utils

type Run func(int) (bool, error)

type TryLoop struct {
	maxCount int
}

func NewTryLoop(maxCount int) *TryLoop {
	return &TryLoop{
		maxCount: maxCount,
	}
}

func (l *TryLoop) Run(run Run) error {
	var err error
	var isContinue bool
	for count := 0; count < l.maxCount; count++ {
		isContinue, err = run(count)
		if !isContinue {
			return err
		}
	}
	return err
}
