package main

import "24lab.net/testlib"

func music_app() {
	song := testlib.GetMusic("Alicia Keys")
	println(song)
}
