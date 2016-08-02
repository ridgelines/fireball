package test

func TestIndexShouldGetID(t *testing.T) {
	for _, id := range []string{"one", "two", "three"} {
		c := gomock.NewContext(ctrl)

		c.EXPECT().
			PathVar("id").
			Return(id)

		c.EXPECT().
			HTML(200, "index.html", id).
			Return(nil, nil)

		Index(c)
	}
}
