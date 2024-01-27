package github

//func (r *Repo) SetUserVisitorsByID(ctx context.Context, id int64) (int64, error) {
//	var (
//		visitors int64
//		err      error
//	)
//
//	memKey := fmt.Sprintf(constant.InMemoryUserByID, id)
//	visitors, err = r.GetUserVisitorsByID(ctx, id)
//	if err != nil {
//		visitors = 0
//	}
//
//	visitors++
//	r.inMemory.Set(memKey, visitors)
//
//	return visitors, nil
//}
//
//func (r *Repo) GetUserVisitorsByID2(ctx context.Context, id int64) (int64, error) {
//	var result int64
//
//	memKey := fmt.Sprintf(constant.InMemoryUserByID, id)
//	rawData, ok := r.inMemory.Get(memKey)
//	if !ok {
//		return result, errors.New("failed to get visitors from InMemory")
//	}
//
//	result, ok = rawData.(int64)
//	if !ok {
//		return result, errors.New("failed to parse visitors result")
//	}
//
//	return result, nil
//}
