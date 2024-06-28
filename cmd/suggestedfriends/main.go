package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	friends, ok := friendships[username]
	if !ok {
		return nil
	}

	friendSet := make(map[string]struct{})
	for _, friend := range friends {
		friendSet[friend] = struct{}{}
	}

	suggestedSet := make(map[string]struct{})

	for _, friend := range friends {
		friendsOfFriend, ok := friendships[friend]
		if !ok {
			continue
		}
		for _, fof := range friendsOfFriend {
			if fof != username {
				if _, isFriend := friendSet[fof]; !isFriend {
					suggestedSet[fof] = struct{}{}
				}
			}
		}
	}
	suggesteds := make([]string, 0, len(suggestedSet))
	for suggested := range suggestedSet {
		suggesteds = append(suggesteds, suggested)
	}
	if len(suggesteds) == 0 {
		return nil
	}
	return suggesteds
}
