package utils

import "net/http"

func Revalidate(path string) {
	http.Get("http://localhost:3000/api/revalidate?path=" + path)
}
