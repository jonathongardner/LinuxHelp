package cli

import (
	"fmt"
	"net"

	"github.com/jonathongardner/linuxhelp/setup/wol"
	"github.com/urfave/cli/v2"
)

var wolCommand = &cli.Command{
	Name:  "wol",
	Usage: "Send Wake on lan packat",
	Action: func(c *cli.Context) error {
		macAddr := c.Args().First()

		var localAddr *net.UDPAddr

		bcastAddr := fmt.Sprintf("%s:%s", "255.255.255.255", "9")
		udpAddr, err := net.ResolveUDPAddr("udp", bcastAddr)
		if err != nil {
			return err
		}

		// Build the magic packet.
		mp, err := wol.New(macAddr)
		if err != nil {
			return err
		}

		// Grab a stream of bytes to send.
		bs, err := mp.Marshal()
		if err != nil {
			return err
		}

		// Grab a UDP connection to send our packet of bytes.
		conn, err := net.DialUDP("udp", localAddr, udpAddr)
		if err != nil {
			return err
		}
		defer conn.Close()

		fmt.Printf("Attempting to send a magic packet to MAC %s\n", macAddr)
		fmt.Printf("... Broadcasting to: %s\n", bcastAddr)
		n, err := conn.Write(bs)
		if err == nil && n != 102 {
			err = fmt.Errorf("magic packet sent was %d bytes (expected 102 bytes sent)", n)
		}
		if err != nil {
			return err
		}

		fmt.Printf("Magic packet sent successfully to %s\n", macAddr)
		return nil
	},
}
