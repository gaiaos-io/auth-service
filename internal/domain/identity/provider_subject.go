package identity

import "errors"

type ProviderSubject struct {
	value string
}

func NewProviderSubject(subject string) (*ProviderSubject, error) {
	providerSubject := &ProviderSubject{value: subject}
	if err := providerSubject.IsValid(); err != nil {
		return nil, err
	}
	return providerSubject, nil
}

func (subject ProviderSubject) IsValid() error {
	if subject.value == "" {
		return errors.New("provider subject cannot be empty")
	}
	return nil
}

func (subject ProviderSubject) String() string {
	return subject.value
}

func (subject ProviderSubject) Equal(other ProviderSubject) bool {
	return subject.value == other.value
}
