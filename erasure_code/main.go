package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/klauspost/reedsolomon"
)

func fillRandom(data []byte) {
	lorem_ipsum := `Lorem ipsum dolor sit amet, feugiat abhorreant et sed, mei no feugait neglegentur. Justo tritani molestiae mea at. Ius no iracundia omittantur deterruisset. Augue feugiat disputationi vim cu, ex vis labore intellegebat.

Mel cu vero qualisque euripidis, vix ut posse mucius, pri ad menandri tacimates interpretaris. Mei alienum ocurreret eu. Ex pro praesent temporibus, cu mea vide eius possim. Ex animal regione disputando sea, vero saperet theophrastus an vis.

Per porro fuisset contentiones te. Vis case partem feugiat et. Ne pro dolorum fastidii reprehendunt, alia impedit sadipscing eu mei, duo dicit tamquam commune id. Odio laoreet fuisset cu mei.

At quod persius eum. Id detraxit honestatis comprehensam vel, amet mollis an pri, vis ad dicit docendi praesent. Malis consetetur reprimique usu ea, pro te bonorum dissentiet theophrastus. Possim antiopam est an, odio decore concludaturque ne sed, conceptam abhorreant an nam. Scripta apeirian est id. Solet utroque incorrupte eum at, ne duo erat lorem discere.

Assum possim mel in, mei persius eripuit an. Eos propriae legendos te, fugit meliore suscipiantur vel an. Cum te quas appareat, nam et ancillae suscipit, alii semper consequuntur vim no. Ea per minim nostrum sententiae, has ne dicat iudico, libris consectetuer qui ne. Mel ei iudico maluisset, ei dicunt invenire philosophia usu.`

	copy(data[:], lorem_ipsum)
}

func main() {
	// Create some sample data
	var data = make([]byte, 10000)
	fillRandom(data)
	fmt.Println("First 100 bytes:", string(data[:100]))

	// Create an encoder with 17 data and 3 parity slices.
	enc, _ := reedsolomon.New(17, 3)

	// Split the data into shards
	shards, _ := enc.Split(data)
	fmt.Printf("# of shards (Type: %T): %v\n", shards, len(shards))

	// Encode the parity set
	_ = enc.Encode(shards)

	// Verify the parity set
	ok, _ := enc.Verify(shards)
	if ok {
		fmt.Println("Shards are verified")
	}

	// Delete two shards
	shards[0], shards[11] = nil, nil

	// Reconstruct the shards
	_ = enc.Reconstruct(shards)

	// Verify the data set
	ok, _ = enc.Verify(shards)
	if ok {
		fmt.Println("Reconstruction successful")
	}

	buf := &bytes.Buffer{}
	err := enc.Join(buf, shards, 10000)
	if err != nil {
		log.Fatalln("Error during .Join():", err)
	}

	fmt.Println("First 100 bytes:", string(buf.Bytes()[:100]))
}
