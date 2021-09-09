package dictionary

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound   = DictionaryErr("could not find the word you are looking for")
	ErrWordExists = DictionaryErr("cannot add word becuase it already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
