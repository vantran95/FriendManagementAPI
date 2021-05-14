package users

//func TestServiceImpl_CreateUser(t *testing.T) {
//	tcs := []struct {
//		scenario             string
//		input                string
//		mockGetUserOutput    models.User
//		mockCreateUserOutput bool
//		mockErr              error
//		expResult            bool
//		expErr               error
//	}{
//		{
//			scenario:             "success",
//			input:                "a@mail.com",
//			mockGetUserOutput:    models.User{},
//			mockCreateUserOutput: true,
//			expResult:            true,
//		},
//		{
//			scenario:          "email existed",
//			input:             "b@gmail.com",
//			mockGetUserOutput: models.User{Email: "b@gmail.com"},
//			expErr:            errors.New("email already exists"),
//		},
//	}
//
//	for _, tc := range tcs {
//		t.Run(tc.scenario, func(t *testing.T) {
//			mockRepo := mockCreateRepository{
//				TestF: t,
//				GetUserInput: struct {
//					Input  string
//					Output models.User
//					Err    error
//				}{Input: tc.input, Output: tc.mockGetUserOutput, Err: tc.mockErr},
//				CreateUserInput: struct {
//					Input  string
//					Output bool
//					Err    error
//				}{Input: tc.input, Output: tc.mockCreateUserOutput, Err: tc.mockErr},
//			}
//
//			service := ServiceImpl{
//				CreateRepo: mockRepo,
//			}
//
//			rs, err := service.CreateUser(tc.input)
//
//			assert.Equal(t, tc.expErr, err)
//			if tc.expErr == nil {
//				assert.Equal(t, tc.expResult, rs)
//			}
//			assert.Equal(t, tc.expErr, err)
//		})
//	}
//}
