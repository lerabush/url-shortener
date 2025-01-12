package download_p2p

import (
	torrent_file "download-p2p/torrent-file"
	"log"
	"os"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrent_file.Open(inPath)

	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
