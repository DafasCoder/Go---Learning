package main

type Dictionary map[string] string
type DictionaryErr string

const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExist = DictionaryErr("This word already exists")
	ErrWordDoesNotExist = DictionaryErr("This word does not exist")
	
)


func (d Dictionary) Search( key string) (string, error){
	definition, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key string, word string) error  {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = word
	case nil:
		return ErrWordExist
	default: 
		return err
	}

	return nil
}

func (d Dictionary) Update(word string, newDefinition string) error{
	_, err := d.Search(word)

	
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d,word)
}


func (e DictionaryErr) Error()  string {
	return string(e)
}