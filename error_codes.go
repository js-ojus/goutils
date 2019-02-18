package goutils

const (
	// ErrInternalDBError : an error with a database.
	ErrInternalDBError = 101
	// ErrInternalSystemError indiates an error at an unidentified point
	// in the program.
	ErrInternalSystemError = 102
	// ErrFileSystemError : an error in some file system operation.
	ErrFileSystemError = 103
	// ErrCryptError : an error in either encrypting or decrypting some
	// data.
	ErrCryptError = 104

	// ErrCorruptBody : inability to read request body.
	ErrCorruptBody = 1001
	// ErrCorruptEnvelope : inability to parse request envelope.
	ErrCorruptEnvelope = 1002
	// ErrInvalidJSON : inability to parse given JSON data.
	ErrInvalidJSON = 1003
	// ErrUnhandledMethod : unknown request method.
	ErrUnhandledMethod = 1004
	// ErrInvalidAPIKey : unknown API key.
	ErrInvalidAPIKey = 1005

	// ErrMissingArguments : one or more mandatory arguments not given.
	ErrMissingArguments = 1101
	// ErrUniquenessViolation : a uniqueness constraint violated.
	ErrUniquenessViolation = 1102
	// ErrCorruptCtokenFormat : corrupt continuation token format.
	ErrCorruptCtokenFormat = 1103
	// ErrInvalidCtoken : unkown continuation token.
	ErrInvalidCtoken = 1104
	// ErrIntegrityViolation : a referential integrity constraint was
	// violated.
	ErrIntegrityViolation = 1105
	// ErrEmptyResultSet : no records met the given query criteria.
	ErrEmptyResultSet = 1106
	// ErrMutuallyExclusiveOptions : a pair of mutually exclusive
	// options given.
	ErrMutuallyExclusiveOptions = 1107
)

// Errors is the global map of errors.  It serves as the central place
// to define error codes and their textual messages.
var Errors map[int]string

//

func init() {
	Errors = map[int]string{
		// Internal system errors.
		ErrInternalDBError:     "Internal database error.",
		ErrInternalSystemError: "Internal system error.",
		ErrFileSystemError:     "File system error.",
		ErrCryptError:          "Encryption system error.",

		// API / transport errors.
		ErrCorruptBody:     "Corrupt request body.",
		ErrCorruptEnvelope: "Corrupt request envelope.",
		ErrInvalidJSON:     "Given data contains invalid JSON.",
		ErrUnhandledMethod: "Unhandled request method.",
		ErrInvalidAPIKey:   "Invalid API Key.",

		// Generic application logic errors.
		ErrMissingArguments:         "Missing mandatory arguments.",
		ErrUniquenessViolation:      "Uniqueness constraint violation.",
		ErrCorruptCtokenFormat:      "Corrupt continuation token format.",
		ErrInvalidCtoken:            "Invalid continuation token.",
		ErrIntegrityViolation:       "Referential integrity violation.",
		ErrEmptyResultSet:           "Empty result set; record could not be found.",
		ErrMutuallyExclusiveOptions: "Mutually exclusive options specified.",
	}
}
