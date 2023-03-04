package videos

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// bks, err := AllBooks()
	// if err != nil {
	// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	// 	return
	// }

}
