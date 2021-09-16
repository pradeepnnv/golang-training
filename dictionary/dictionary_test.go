package dictionary

import "testing"

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		word := "test"
		def := "definition"
		dict := Dictionary{word: def}
		err := dict.Delete(word)
		assertError(t, err, nil)
		_, err = dict.Search(word)
		assertError(t, err, ErrNotFound)
	})
	t.Run("delete non available word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}
		err := dict.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
		_, err = dict.Search(word)
		assertError(t, err, ErrNotFound)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dict := Dictionary{word: definition}
		newDefinition := "this is an updated definition"
		dict.Update(word, newDefinition)
		assertDefinition(t, dict, word, newDefinition)
	})
	t.Run("Update new word", func(t *testing.T) {
		word := "test"
		definition := "definition"
		dict := Dictionary{}
		err := dict.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)

	})
}

func TestAdd(t *testing.T) {
	d := Dictionary{}
	t.Run("Add a word", func(t *testing.T) {
		word := "adamantine"
		definition := "rigidly firm"
		err := d.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, d, word, definition)
	})
	t.Run("Add an existing word", func(t *testing.T) {
		word := "adamantine"
		definition := "rigidly firm"
		err := d.Add(word, "new definition")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, d, word, definition)
	})
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"jeopardize": "to expose to danger or risk"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("jeopardize")
		want := "to expose to danger or risk"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search(("unkown"))
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestError(t *testing.T) {
	got := ErrNotFound
	want := "could not find the word you are looking for"
	if got.Error() != want {
		t.Errorf("got %q but want %q", got.Error(), want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, want string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, want)

}
