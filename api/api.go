package api

func RegisterUser(username string) (usr *User, err error) {
	// <TODO>
	// 1. Check if user exists
	// 2. Construct new user
	// 2. Save new user to DB
	// 3. Handle errors

	return nil, err
}

func LookupUser(username string) (usr *User, err error) {
	// <TODO>
	// 1. Search DB for user with username equal to input string
	// 2. Return user if found, nil if not.
	// 3. Handle errors

	return nil, err
}

// User tweeting
func (usr *User) Tweet(input []byte) (tst *Tweet, err error) {

	// <TODO>
	// 1. Append to user's tweet list
	// 2. Save user back to DB
	// 3. Handle errors
	return nil, err
}

// User tweeting
func (usr *User) Comment(input []byte, twt *Tweet) (twt *Tweet, err error) {

	// <TODO>
	// 1. Append to user's tweet list
	// 2. Save user back to DB
	// 3. Handle errors
	return nil, err
}

// User follows another user
func (usr *User) Follow(inputUser *User) (err Error) {

	// <TODO>
	// 1. Append to user's following list
	// 2. Save user back to DB
	// 3. Handle errors
	return nil
}
