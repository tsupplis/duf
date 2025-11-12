//go:build solaris || illumos
// +build solaris illumos

package main

func isFuseFs(_ Mount) bool {
	// FUSE is not commonly used on Solaris
	return false
}

func isNetworkFs(m Mount) bool {
	fs := []string{"nfs", "nfs3", "nfs4", "smbfs", "cifs"}

	for _, v := range fs {
		if m.Fstype == v {
			return true
		}
	}

	return false
}

func isSpecialFs(m Mount) bool {
	fs := []string{"devfs", "ctfs", "fd", "lofs", "mntfs", "objfs", "proc", "tmpfs", "dev", "sharefs", "bootfs", "autofs"}

	for _, v := range fs {
		if m.Fstype == v {
			return true
		}
	}

	return false
}

func isHiddenFs(_ Mount) bool {
	return false
}
