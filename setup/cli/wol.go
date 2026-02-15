package cli

import (
	"fmt"
	"net"

	"github.com/jonathongardner/linuxhelp/setup/wol"
	"github.com/urfave/cli/v2"
)

func ipFromInterface(iface string) (*net.UDPAddr, error) {
	ief, err := net.InterfaceByName(iface)
	if err != nil {
		return nil, err
	}

	addrs, err := ief.Addrs()
	if err == nil && len(addrs) <= 0 {
		err = fmt.Errorf("no address associated with interface %s", iface)
	}
	if err != nil {
		return nil, err
	}

	// Validate that one of the addrs is a valid network IP address.
	for _, addr := range addrs {
		switch ip := addr.(type) {
		case *net.IPNet:
			if !ip.IP.IsLoopback() && ip.IP.To4() != nil {
				return &net.UDPAddr{
					IP: ip.IP,
				}, nil
			}
		}
	}
	return nil, fmt.Errorf("no address associated with interface %s", iface)
}

// sudo tcpdump -i eth0 -n -v -s0 udp port 9
var wolCommand = &cli.Command{
	Name:  "wol",
	Usage: "Send Wake on lan packat",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "address",
			Usage: "Address to brodcast on",
		},
		&cli.StringFlag{
			Name:  "interface",
			Usage: "Interface to brodcast on",
		},
	},
	Action: func(c *cli.Context) error {
		macAddr := c.Args().First()

		var localAddr *net.UDPAddr
		if add := c.String("address"); add != "" {
			var err error
			localAddr, err = net.ResolveUDPAddr("udp", add)
			if err != nil {
				return fmt.Errorf("error resolving address %v (%v)", add, err)
			}
		}
		if iface := c.String("interface"); iface != "" {
			var err error
			localAddr, err = ipFromInterface(iface)
			if err != nil {
				return fmt.Errorf("error resolving interface %v (%v)", iface, err)
			}
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

		for _, port := range []string{"7", "9"} {
			bcastAddr := fmt.Sprintf("%s:%s", "255.255.255.255", port)
			udpAddr, err := net.ResolveUDPAddr("udp", bcastAddr)
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
			ip := ""
			if localAddr != nil {
				ip = localAddr.String()
			}
			fmt.Printf("... Broadcasting to: %s on: %s\n", bcastAddr, ip)
			n, err := conn.Write(bs)
			if err == nil && n != 102 {
				err = fmt.Errorf("magic packet sent was %d bytes (expected 102 bytes sent)", n)
			}
			if err != nil {
				return err
			}

			fmt.Printf("Magic packet sent successfully to %s\n", macAddr)
		}
		return nil
	},
}
