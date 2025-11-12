//go:build openbsd
// +build openbsd

package main

func isFuseFs(_ Mount) bool {
	//FIXME: implement
	return false
}

func isNetworkFs(_ Mount) bool {
	//FIXME: implement
	return false
}

func isSpecialFs(m Mount) bool {
	return m.Fstype == "devfs"
}

func isHiddenFs(_ Mount) bool {
	return false
}
