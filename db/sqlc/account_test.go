package db

// // helper สร้าง user สุ่ม 1 คนใน DB แล้วคืนค่ากลับมา
// func randomUser(t *testing.T) User { // ← struct ชื่อ User (เอกพจน์)
// 	t.Helper() // บอกว่าเป็น helper function

// 	arg := CreateUsersParams{ // sqlc สร้าง struct นี้จากคิวรี CreateUsers
// 		Username: util.Randomusername(),
// 		Balance:  "0.00",
// 		AffiliateID: uuid.NullUUID{ // NULL
// 			Valid: false,
// 		},
// 	}

// 	user, err := testQueries.CreateUsers(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, user)

// 	require.Equal(t, arg.Username, user.Username)
// 	require.Equal(t, arg.Balance, user.Balance)
// 	require.Equal(t, arg.AffiliateID, user.AffiliateID)
// 	require.NotZero(t, user.ID)

// 	return user
// }

// func TestCreateUser(t *testing.T) {
// 	randomUser(t) // ถ้ helper ไม่ panic = ผ่าน
// }

// func TestGetUser(t *testing.T) {
// 	// 1) สร้าง user ทดสอบ
// 	u1 := randomUser(t)

// 	// 2) ดึงกลับมา
// 	ctx := context.Background()
// 	u2, err := testQueries.GetUser(ctx, u1.ID)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, u2)

// 	require.Equal(t, u1.ID, u2.ID)
// 	require.Equal(t, u1.Username, u2.Username)
// 	require.Equal(t, u1.Balance, u2.Balance)
// 	require.Equal(t, u1.AffiliateID, u2.AffiliateID)

// 	// 3) ดึง ID สุ่มที่ไม่มีอยู่ → ควรได้ ErrRecordNotFound
// 	_, err = testQueries.GetUser(ctx, uuid.New())
// 	require.ErrorIs(t, err, sql.ErrNoRows)
// }
