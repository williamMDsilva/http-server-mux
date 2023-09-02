package main

func FirstOrDefault[T any](slice []T, filter func(*T) bool) (element *T) {

	for i := 0; i < len(slice); i++ {

		if filter(&slice[i]) {
			return &slice[i]
		}
	}

	return nil
}

func getUserById(userById string) string {
	user := FirstOrDefault(users, func(u *User) bool { return u.Id == userById })

	if user == nil {
		return ""
	}

	return user.Name
}
