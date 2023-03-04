package videos

// func AllBooks() ([]Book, error) {
// 	rows, err := config.DB.Query("SELECT * FROM books")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	bks := make([]Book, 0)
// 	for rows.Next() {
// 		bk := Book{}
// 		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price) // order matters
// 		if err != nil {
// 			return nil, err
// 		}
// 		bks = append(bks, bk)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return bks, nil
// }
