# simple script to enumerate UPNP devices
import socket
import sys

def discover_upnp_devices(timeout=2):
    # M-Search message body
    ms = ( \
        'M-SEARCH * HTTP/1.1\r\n' \
        'HOST:192.168.0.1:49153\r\n' \
        'ST:upnp:rootdevice\r\n' \
        'MX:2\r\n' \
        'MAN:"ssdp:discover"\r\n' \
        '\r\n'
    )

    soc = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, socket.IPPROTO_UDP)
    multicast_address = '192.168.0.255'
    multicast_port = 49153

    try:
        soc.settimeout(timeout)
        # Added option to reuse address
        soc.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        # Added option to use broadcast
        soc.setsockopt(socket.SOL_SOCKET, socket.SO_BROADCAST, 1)

        # Bind to all interfaces (0.0.0.0) and an available port (0)
        soc.bind(('', 0))

        #send the request to multicast
        soc.sendto(ms.encode('utf-8'), (multicast_address, multicast_port))
        print(f"Sending M-SEARCH request to {multicast_address}:{multicast_port}")

        # listen and capture returned responses
        try:
            while True:
                data, addr = soc.recvfrom(8192)
                print(f"Response from {addr}:")
                print(data.decode('utf-8'))
                print("-" * 20)
        except socket.timeout:
            print("No more responses received.")
    except OSError as e:
        print("OS error", e)
    except Exception as e:
        print("Error:", e)
    finally:
        soc.close()
if __name__ == "__main__":
    discover_upnp_devices()
