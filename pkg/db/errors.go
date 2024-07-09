package db

import "errors"

var ErrorNotFound = errors.New("record not found")

var ErrorMultipleRecords = errors.New("multiple records found in FindOne")
