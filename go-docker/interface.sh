ip link add mac0 link enp0s3 type macvlan mode bridge
ip addr add 172.20.4.9/28 dev mac0
ifconfig enp0s3 down
ifconfig enp0s3 up promisc
ifconfig mac0 up