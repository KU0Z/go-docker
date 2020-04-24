sudo ip link add mac0 link enp0s3 type macvlan mode bridge
sudo ip addr add 172.20.4.9/28 dev mac0
sudo ifconfig enp0s3 down
sudo ifconfig enp0s3 up promisc
sudo ifconfig mac0 up