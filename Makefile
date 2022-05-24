test_concurrency:
	wrk -t12 -c2000 -d15s "http://localhost/"

install_go:
	apt install -y build-essential
	snap install go --classic

enable_bbr:
	echo net.core.default_qdisc=fq >> /etc/sysctl.conf
	echo net.ipv4.tcp_congestion_control=bbr >> /etc/sysctl.conf
	sysctl -p
